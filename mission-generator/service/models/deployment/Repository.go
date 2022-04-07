package deployment

import (
	"context"
	"errors"
	"math/rand"

	"github.com/jackc/pgx/v4"
)

type Repository interface {
	Random(ctx context.Context) (*Deployment, error)
}

type PostgresRepo struct {
	conn *pgx.Conn;
}

func (r *PostgresRepo) Random(ctx context.Context) (*Deployment, error) {
	var imageUrl string
	err := r.conn.QueryRow(ctx, "select image_url from deployment order by random() limit 1").Scan(&imageUrl)

	if err != nil {
		return &Deployment{}, err
	}

	d, err := NewDeployment(imageUrl)

	if err != nil {
		return &Deployment{}, err
	}

	return d, nil
}

func NewPostgresRepo(conn *pgx.Conn) *PostgresRepo {
	r := &PostgresRepo{conn: conn}

	return r
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