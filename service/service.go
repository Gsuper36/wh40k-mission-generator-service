package service

import (
	"context"

	"github.com/Gsuper36/wh40k-mission-generator-service/pb"
	"github.com/Gsuper36/wh40k-mission-generator-service/service/models/mission"
	"github.com/go-kit/kit/log"
)

type service struct {
	logger log.Logger
	missionRepository mission.Repository
}

type Service interface {
	Generate(context.Context, *pb.GenerateMissionRequest) (*pb.Mission, error)
	List(*pb.ListMissionRequest, pb.MissionGenerator_ListServer) error
}

func NewService(logger log.Logger, repo mission.Repository) Service {
	return &service{
		logger: logger,
		missionRepository: repo,
	}
}

func (s service) Generate(ctx context.Context, req *pb.GenerateMissionRequest) (*pb.Mission, error) {
	m := &mission.Mission{}
	m.SetFormat(req.MissionFormat)

	s.missionRepository.Save(m)

	return &pb.Mission{MissionFormat: m.Format()}, nil //@todo
}

func (s service) List(req *pb.ListMissionRequest, server pb.MissionGenerator_ListServer) error {
	panic("umimplemented")
}