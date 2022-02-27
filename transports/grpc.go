package transports

import (
	"context"

	"github.com/Gsuper36/wh40k-mission-generator-service/endpoints"
	"github.com/Gsuper36/wh40k-mission-generator-service/pb"
	"github.com/Gsuper36/wh40k-mission-generator-service/service/models/mission"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
)

type gRPCServer struct {
	pb.UnimplementedMissionGeneratorServer
	generate gt.Handler
	list gt.Handler
}

func NewGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) pb.MissionGeneratorServer {
	return &gRPCServer {
		generate: gt.NewServer(
			endpoints.Generate,
			decodeGenerateRequest,
			encodeGenerateResponse,
		),
	}
}

func (s *gRPCServer) Generate(ctx context.Context, req *pb.GenerateMissionRequest) (*pb.Mission, error) {
	_, resp, err := s.generate.ServeGRPC(ctx, req)

	if err != nil {
		return &pb.Mission{}, err
	}

	return resp.(*pb.Mission), nil
}

func decodeGenerateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GenerateMissionRequest)

	return endpoints.GenerateReq{Format: &req.MissionFormat}, nil
}

func encodeGenerateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*mission.Mission)

	return &pb.Mission{
		Title: resp.Title(),
		Description: resp.Description(),
		Rules: resp.Rules(),
		MissionFormat: resp.Format(),
		//@todo wrap Twists: resp.Twists(),
		//@todo wrap Objectives: resp.Objectives()
		Deployment: &pb.Mission_Deployment{ImageUrl: resp.Deployment().ImageUrl()},
	}, nil
} 