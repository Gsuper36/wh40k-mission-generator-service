package objective

import "errors"

type Objective struct {
	title       string
	description string
	rules       string
}

func (o *Objective) Title() string {
	return o.title
}

func (o *Objective) SetTitle(title string) error {
	if len(title) == 0 {
		return errors.New("Title can't be null")
	}

	o.title = title

	return nil
}

func (o *Objective) Description() string {
	return o.description
}

func (o *Objective) SetDescription(description string) {
	o.description = description
}

func (o *Objective) Rules() string {
	return o.rules
}

func (o *Objective) SetRules(rules string) error {
	if len(rules) == 0 {
		return errors.New("Rules can't be empty")
	}

	o.rules = rules

	return nil
}

func NewObjective(title, rules, description string) (*Objective, error) {
	o := &Objective{}

	err := o.SetTitle(title)

	if err != nil {
		return &Objective{}, err
	}

	err = o.SetRules(rules)

	if err != nil {
		return &Objective{}, err
	}

	o.SetDescription(description)

	return o, nil
}