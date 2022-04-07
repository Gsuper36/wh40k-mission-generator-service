package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Gsuper36/wh40k-mission-generator-service/endpoints"
	"github.com/Gsuper36/wh40k-mission-generator-service/pb"
	"github.com/Gsuper36/wh40k-mission-generator-service/service"
	"github.com/Gsuper36/wh40k-mission-generator-service/service/models/deployment"
	"github.com/Gsuper36/wh40k-mission-generator-service/service/models/mission"
	"github.com/Gsuper36/wh40k-mission-generator-service/service/models/objective"
	"github.com/Gsuper36/wh40k-mission-generator-service/service/models/twist"
	"github.com/Gsuper36/wh40k-mission-generator-service/transports"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpcServerPort = getEnv("APP_PORT", "50051")
	dbUrl = getEnv("DB_URL", "postgres://mission_generator:mission_generator@localhost:5432/mission_generator")
)


func main() {
	var logger log.Logger
	grpcServerEndpoint := ":" + grpcServerPort;
	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	dbCtx := context.Background()
	dbCtx, dbCancel := context.WithCancel(dbCtx)
	defer dbCancel()

	conn, err := pgx.Connect(dbCtx, dbUrl)

	if err != nil {
		logger.Log("during", "connect db", "err", err)
		os.Exit(1)
	}

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// repository, err := mission.NewPostgresRepo(context.Background(), "postgres://mission_generator:mission_generator@localhost:5432/mission_generator") //@todo repo instance from ENV
	
	mRepo := mission.NewInMemoryRepo()
	oRepo := objective.NewPostgresRepo(conn)
	tRepo := twist.NewPostgresRepo(conn)
	dRepo := deployment.NewPostgresRepo(conn)

	// if err != nil {
		// logger.Log("during", "connect db", "err", err)
		// os.Exit(1)
	// }

	listener, err := net.Listen("tcp", grpcServerEndpoint) //@todo port from ENV

	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	service := service.NewService(logger, mRepo, oRepo, tRepo, dRepo) 
	endpoints := endpoints.MakeEndpoints(service)
	server := transports.NewGRPCServer(endpoints, logger)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 3)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	//listener, err := net.Listen("tcp", ":50051") //@todo port from ENV

	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	go func() {
		baseServer := grpc.NewServer()
		pb.RegisterMissionGeneratorServer(baseServer, server)
		level.Info(logger).Log("msg", "gRPC server is starting")
		err := baseServer.Serve(listener)

		if err != nil {
			errs <- err
		}

	}()

	go func ()  {
		err := pb.RegisterMissionGeneratorHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
		
		if err != nil {
			errs <- err
		}
		level.Info(logger).Log("msg", "Proxy server is starting")
		err = http.ListenAndServe(":6011", mux)

		if err != nil {
			errs <- err
		}

	}() //@todo move proxy to another file and run in different container !!! this stuff here is only for testing purposes

	level.Error(logger).Log("exit", <-errs)
}

func getEnv(key, def string) string {
	v, exists := os.LookupEnv(key)

	if !exists {
		return def
	}

	return v
}