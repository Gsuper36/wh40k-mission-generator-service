package endpoints

import (
	"context"

	"github.com/Gsuper36/wh40k-mission-generator-service/pb"
	"github.com/Gsuper36/wh40k-mission-generator-service/service"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	Generate endpoint.Endpoint
}

type GenerateReq struct {
	Format *pb.MissionFormat
}

func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		Generate: makeGenerateEndpoint(s),
	}
}

func makeGenerateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GenerateReq)

		result, _ := s.Generate(ctx, *req.Format)

		return result, nil
	}
}