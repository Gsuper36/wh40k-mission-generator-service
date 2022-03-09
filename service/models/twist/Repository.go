package twist

import (
	"context"
	"errors"
	"math/rand"
)

type Repository interface {
	Random(ctx context.Context, max int) ([]*Twist, error)
}

type InMemoryRepo struct {
	twists []*Twist
}

func (r *InMemoryRepo) Random(ctx context.Context, max int) ([]*Twist, error) {
	if max < 1 {
		return make([]*Twist, 0), errors.New("max can't be lower than 1")
	}

	length := len(r.twists)
	maxBound := length - max

	if maxBound < 0 {
		maxBound = 1
		max -= length
	}

	from := rand.Intn(maxBound)
	max += from

	return r.twists[from:max], nil
}

func NewInMemoryRepo() *InMemoryRepo {
	t1 := &Twist{}
	t1.SetTitle("Twist 1")
	t1.SetDescription("Description")
	t1.SetRules("Rules")

	t2 := &Twist{}
	t2.SetTitle("Twist 2")
	t2.SetDescription("Description")
	t2.SetRules("Rules")

	return &InMemoryRepo{
		twists: []*Twist{t1, t2},
	}
}