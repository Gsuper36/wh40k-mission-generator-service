package objective

import (
	"context"
	"errors"
	"math/rand"

	"github.com/jackc/pgx/v4"
)


type Repository interface {
	Random(ctx context.Context, max int) ([]*Objective, error)
}

type PostgresRepo struct {
	conn *pgx.Conn
}

func (r *PostgresRepo) Random(ctx context.Context, max int) ([]*Objective, error) {
	var (
		title 		string
		description string
		rules 		string 
		objs 		[]*Objective
	)

	rows, err := r.conn.Query(ctx, "select title, description, rules from objective order by random() limit $1", max)

	if err != nil {
		return []*Objective{}, err
	}

	for rows.Next() {
		rows.Scan(&title, &description, &rules)
		o, err := NewObjective(title, rules, description)
		
		if err != nil {
			return []*Objective{}, err
		}

		objs = append(objs, o)
	}

	return objs, nil
}

func NewPostgresRepo(conn *pgx.Conn) *PostgresRepo {
	r := &PostgresRepo{conn: conn}

	return r
}

type InMemoryRepo struct{
	objectives []*Objective
}

func (r *InMemoryRepo) Random(ctx context.Context, max int) ([]*Objective, error){
	if max < 1 {
		return make([]*Objective, 0), errors.New("max can't be lower than 1")
	}

	length := len(r.objectives)
	maxBound := length - max

	if maxBound < 0 {
		maxBound = 1
		max -= length
	}

	from := rand.Intn(maxBound)
	max += from

	return r.objectives[from:max], nil
}

func NewInMemoryRepo() *InMemoryRepo {
	obj1 := &Objective{}
	obj1.SetTitle("Objective 1")
	obj1.SetDescription("Description")
	obj1.SetRules("Rule")

	obj2 := &Objective{}
	obj2.SetTitle("Objective 2")
	obj2.SetDescription("Description")
	obj2.SetRules("Rule")

	obj3 := &Objective{}
	obj3.SetTitle("Objective 3")
	obj3.SetDescription("Description")
	obj3.SetRules("Rules")

	return &InMemoryRepo{
		objectives: []*Objective{obj1, obj2, obj3},
	}
}
