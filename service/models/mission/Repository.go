package mission

import (
	"context"
	"errors"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository interface {
	FindById(context.Context, string) (*Mission, error) //@todo replace id with VO
	Save(context.Context, *Mission) (string, error)
}

type InMemoryRepo struct {
	missions map[string]*Mission
}

func (r *InMemoryRepo) FindById(ctx context.Context, id string) (*Mission, error) {
	m, ok := r.missions[id]

	if !ok {
		return &Mission{}, errors.New("Mission not found")
	}

	return m, nil
}

func (r *InMemoryRepo) Save(ctx context.Context, m *Mission) (string, error) {
	id, err := uuid.NewV4()

	if err != nil {
		return "", err
	}

	r.missions[id.String()] = m

	return id.String(), nil
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		missions: make(map[string]*Mission, 10),
	}
}

type PgRepo struct {
	pool *pgxpool.Pool
}

func (r *PgRepo) FindById(ctx context.Context, id string) (*Mission, error) {
	m := &Mission{}

	err := r.pool.QueryRow(ctx, "select * from mission where id=$1", id).Scan(
		&m.id, 
		&m.title, 
		&m.description, 
		&m.rules, 
		&m.format, 
		&m.deployment, 
		&m.twists, 
		&m.objectives,
	) //@todo rewrite, move this code to infrastructure

	if err != nil {
		return &Mission{}, err
	}

	return m, nil
}

func (r *PgRepo) Save(ctx context.Context, m *Mission) (string, error) {
	return "", errors.New("unimplemented") //@todo
}


func NewPostgresRepo(ctx context.Context, url string) (*PgRepo, error) {
	pool, err := pgxpool.Connect(ctx, url)

	if err != nil {
		return &PgRepo{}, err
	}

	return &PgRepo{pool: pool}, nil
}