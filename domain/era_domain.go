package domain

type Era struct {
	EraId int    `json:"eraId"`
	Name  string `json:"name"`
	Sortable
}

type EraRepository interface {
	GetList() ([]Era, error)
	GetById(id int) (*Era, error)
}
