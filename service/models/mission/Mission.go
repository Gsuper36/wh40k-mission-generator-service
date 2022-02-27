package mission

import (
	"errors"

	"github.com/Gsuper36/wh40k-mission-generator-service/pb"
	"github.com/Gsuper36/wh40k-mission-generator-service/service/models/deployment"
	"github.com/Gsuper36/wh40k-mission-generator-service/service/models/objective"
	"github.com/Gsuper36/wh40k-mission-generator-service/service/models/twist"
)

type Mission struct {
	title       string
	description string
	rules       string
	format      pb.MissionFormat
	twists	    []*twist.Twist
	objectives  []*objective.Objective
	deployment  *deployment.Deployment
}

func (m *Mission) Title() string {
	return m.title
}

func (m *Mission) SetTitle(title string) error {
	if len(title) == 0 {
		return errors.New("title can't be null")
	}

	m.title = title

	return nil
}

func (m *Mission) Description() string {
	return m.description
}

func (m *Mission) SetDescription(description string) {
	m.description = description
}

func (m *Mission) Rules() string {
	return m.rules
}

func (m *Mission) SetRules(rules string) error {
	if len(rules) == 0 {
		return errors.New("rules can't be empty")
	}

	m.rules = rules

	return nil
}

func (m *Mission) Format() pb.MissionFormat {
	return m.format
}

func (m *Mission) SetFormat(format pb.MissionFormat) {
	m.format = format
} 

func (m *Mission) Twists() []*twist.Twist {
	return m.twists
}

func (m *Mission) AddTwist(t *twist.Twist) {
	m.twists = append(m.twists, t)
}

func (m *Mission) Objectives() []*objective.Objective {
	return m.objectives
}

func (m *Mission) AddObjective(o *objective.Objective) {
	m.objectives = append(m.objectives, o)
}

func (m *Mission) Deployment() *deployment.Deployment {
	return m.deployment
}

func (m *Mission) SetDeployemnt(d *deployment.Deployment) {
	m.deployment = d
}

func NewMission(title, description, rules string, format pb.MissionFormat, twists []*twist.Twist, deployment *deployment.Deployment, objectives []*objective.Objective) (*Mission, error) {
	m := &Mission{
		objectives: objectives,
		twists: twists,
		description: description,
		format: format,
		deployment: deployment,
	}

	err := m.SetTitle(title)

	if err != nil {
		return &Mission{}, err
	}

	err = m.SetRules(rules)

	if err != nil {
		return &Mission{}, err
	}

	return m, nil
}