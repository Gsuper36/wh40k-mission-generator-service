package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Gsuper36/wh40k-mission-generator-service/endpoints"
	"github.com/Gsuper36/wh40k-mission-generator-service/pb"
	"github.com/Gsuper36/wh40k-mission-generator-service/service"
	"github.com/Gsuper36/wh40k-mission-generator-service/transports"
	"github.com/go-kit/log/level"
	"github.com/go-kit/log"
	"google.golang.org/grpc"
)


func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	service := service.NewService(logger, nil) //@todo repo instance from ENV
	endpoints := endpoints.MakeEndpoints(service)
	server := transports.NewGRPCServer(endpoints, logger)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 3)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	listener, err := net.Listen("tcp", ":50051") //@todo port from ENV

	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	go func ()  {
		baseServer := grpc.NewServer()
		pb.RegisterMissionGeneratorServer(baseServer, server)
		level.Info(logger).Log("msg", "Server started succesfully")
		baseServer.Serve(listener)
	}()

	level.Error(logger).Log("exit", <-errs)
}