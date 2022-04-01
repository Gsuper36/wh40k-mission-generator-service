package deployment

import (
	"context"
	"errors"
	"math/rand"
)

type Repository interface {
	Random(ctx context.Context) (*Deployment, error)
}

type InMemoryRepo struct {
	deployments []*Deployment
}

func (r *InMemoryRepo) Random(ctx context.Context) (*Deployment, error) {
	len := len(r.deployments)

	if len == 0 {
		return &Deployment{}, errors.New("empty repo")
	}

	return r.deployments[rand.Intn(len)], nil
}

func NewInMemoryRepo() *InMemoryRepo {
	d := &Deployment{}
	d.SetImageUrl("https://i.picsum.photos/id/840/200/300.jpg?hmac=Z8Mc1xk7GaQHQ1hkPTK4cY0dYIxDKGBCHrgyaDqE0u0")

	return &InMemoryRepo{
		deployments: []*Deployment{
			d,
		},
	}
}