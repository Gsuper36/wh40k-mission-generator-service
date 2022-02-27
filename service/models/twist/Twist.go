package twist

import "errors"

type Twist struct {
	title       string
	description string
	rules       string
}

func (t *Twist) Title() string {
	return t.title
}

func (t *Twist) SetTitle(title string) error {
	if len(title) == 0 {
		return errors.New("Title mustn't be empty")
	}

	t.title = title

	return nil
}

func (t *Twist) Description() string {
	return t.description
}

func (t *Twist) SetDescription(descr string) {
	t.description = descr
}

func (t *Twist) Rules() string {
	return t.rules
}

func (t *Twist) SetRules(rules string) error {
	if len(rules) == 0 {
		return errors.New("Rules mustn't be empty")
	}

	t.rules = rules

	return nil
} 

func NewTwist(title, description, rules string) (*Twist, error) {
	t := &Twist{
		description: description,
	};

	err := t.SetTitle(title);

	if err != nil {
		return &Twist{}, err
	}

	err = t.SetRules(rules);

	if err != nil {
		return &Twist{}, err
	}

	return t, nil
}