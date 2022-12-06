package models

type Artist struct {
	Name         string `json:"name"`
	Arts         []*Art `json:"art"`
	AtExhibition bool   `json:"at_exhibition"`
}
