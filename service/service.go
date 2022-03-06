package service

import (
	"context"
	"errors"
	"time"

	"github.com/Gsuper36/wh40k-mission-generator-service/pb"
	"github.com/Gsuper36/wh40k-mission-generator-service/service/models/deployment"
	"github.com/Gsuper36/wh40k-mission-generator-service/service/models/mission"
	"github.com/Gsuper36/wh40k-mission-generator-service/service/models/objective"
	"github.com/Gsuper36/wh40k-mission-generator-service/service/models/twist"
	"github.com/go-kit/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type service struct {
	logger            log.Logger
	missionRepository mission.Repository
}

type Service interface {
	Generate(context.Context, pb.MissionFormat) (*mission.Mission, error) //@todo remove dependencies on transports
	List(*pb.ListMissionRequest, pb.MissionGenerator_ListServer) error    //@todo remove dependencies on transports (any pb.*)
}

func NewService(logger log.Logger, repo mission.Repository) Service {
	return &service{
		logger:            logger,
		missionRepository: repo,
	}
}

func (s service) Generate(ctx context.Context, format pb.MissionFormat) (*mission.Mission, error) {
	// TEST CODE
	d := &deployment.Deployment{}
	err := d.SetImageUrl("https://i.picsum.photos/id/840/200/300.jpg?hmac=Z8Mc1xk7GaQHQ1hkPTK4cY0dYIxDKGBCHrgyaDqE0u0")

	if err != nil {
		s.logger.Log(err)
		return &mission.Mission{}, err
	}

	o := &objective.Objective{}
	o.SetTitle("Dummy objective")
	o.SetDescription("Dummy description")
	o.SetRules("Dummy rules")

	t := &twist.Twist{}
	t.SetTitle("Dummy twist")
	t.SetDescription("Dummy description")
	t.SetRules("Dummy rules")

	m := &mission.Mission{}
	m.SetFormat(format)
	m.SetDeployment(d)
	m.SetTitle("Dummy mission")
	m.SetDescription("Dummy description")
	m.SetRules("Dummy rules")
	m.AddObjective(o)
	m.AddTwist(t)

	//
	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err = s.missionRepository.Save(c, m)

	if err != nil {
		s.logger.Log(err)

		return &mission.Mission{}, status.Error(codes.Internal, err.Error())
	}

	return m, nil
}

func (s service) List(req *pb.ListMissionRequest, server pb.MissionGenerator_ListServer) error {
	return errors.New("unimplemented")
}
