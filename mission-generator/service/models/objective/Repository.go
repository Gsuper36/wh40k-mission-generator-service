package objective

import (
	"context"
	"errors"
	"math/rand"
)


type Repository interface {
	Random(ctx context.Context, max int) ([]*Objective, error)
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
