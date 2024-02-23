package domain

type CardType struct {
	CardTypeId int    `json:"cardTypeId"`
	Name       string `json:"name"`
	Sortable
}

type CardTypeRepository interface {
	GetList() ([]CardType, error)
}
