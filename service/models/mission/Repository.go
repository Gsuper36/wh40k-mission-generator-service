package mission

import "github.com/jackc/pgx/v4"

type Repository interface {
	FindById(string) *Mission //@todo replace id with VO
	Save(*Mission) (string, error)
}

type PgRepo struct {
	conn *pgx.Conn
}

func (r *PgRepo) FindById(id string) *Mission {
	return &Mission{} //@todo
}

func (r *PgRepo) Save(m *Mission) (string, error) {
	return "", nil //@todo
}


func NewPostgresRepo(conn *pgx.Conn) *PgRepo {
	return &PgRepo{conn: conn}
}