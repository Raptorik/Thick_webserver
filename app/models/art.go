package models

type Art struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

func (a *Art) IsNotAssigned() bool {
	return a.Owner == ""
}
