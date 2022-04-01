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
	logger            	 log.Logger
	missionRepository 	 mission.Repository
	objectiveRepository  objective.Repository
	twistRepository 	 twist.Repository
	deploymentRepository deployment.Repository
}

type Service interface {
	Generate(context.Context, pb.MissionFormat) (*mission.Mission, error) //@todo remove dependencies on transports
	List(*pb.ListMissionRequest, pb.MissionGenerator_ListServer) error    //@todo remove dependencies on transports (any pb.*)
}

func NewService(
	logger log.Logger, 
	mRepo mission.Repository,
	oRepo objective.Repository,
	tRepo twist.Repository,
	dRepo deployment.Repository,
) Service {
	return &service{
		logger:            logger,
		missionRepository: mRepo,
		deploymentRepository: dRepo,
		objectiveRepository: oRepo,
		twistRepository: tRepo,
	}
}

func (s service) Generate(ctx context.Context, format pb.MissionFormat) (*mission.Mission, error) {
	// TEST CODE
	d, err := s.deploymentRepository.Random(ctx)

	if err != nil {
		s.logger.Log(err) // @todo Remove Log from business logic
		return &mission.Mission{}, err
	}

	objs, err := s.objectiveRepository.Random(ctx, 1)

	if err != nil {
		s.logger.Log(err)
		return &mission.Mission{}, err
	}

	twst, err := s.twistRepository.Random(ctx, 1)

	if err != nil {
		s.logger.Log(err)
		return &mission.Mission{}, err
	}

	m := &mission.Mission{}
	m.SetFormat(format)
	m.SetDeployment(d)
	m.SetTitle("Dummy mission")
	m.SetDescription("Dummy description")
	m.SetRules("Dummy rules")
	m.SetTwists(twst)
	m.SetObjectives(objs)

	//
	c, cancel := context.WithTimeout(ctx, 5 * time.Second)
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
