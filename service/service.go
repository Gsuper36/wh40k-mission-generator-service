package service

import (
	"context"
	"errors"
	"time"

	"github.com/Gsuper36/wh40k-mission-generator-service/pb"
	"github.com/Gsuper36/wh40k-mission-generator-service/service/models/mission"
	"github.com/go-kit/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type service struct {
	logger log.Logger
	missionRepository mission.Repository
}

type Service interface {
	Generate(context.Context, pb.MissionFormat) (*mission.Mission, error)
	List(*pb.ListMissionRequest, pb.MissionGenerator_ListServer) error
}

func NewService(logger log.Logger, repo mission.Repository) Service {
	return &service{
		logger: logger,
		missionRepository: repo,
	}
}

func (s service) Generate(ctx context.Context, format pb.MissionFormat) (*mission.Mission, error) {
	m := &mission.Mission{}
	m.SetFormat(format)

	c, cancel := context.WithTimeout(ctx, 5 * time.Second);
	defer cancel();

	_, err := s.missionRepository.Save(c, m)

	if err != nil {
		s.logger.Log(err)

		return &mission.Mission{}, status.Error(codes.Internal, err.Error())
	}

	return m, nil //todo
}

func (s service) List(req *pb.ListMissionRequest, server pb.MissionGenerator_ListServer) error {
	return errors.New("unimplemented")
}