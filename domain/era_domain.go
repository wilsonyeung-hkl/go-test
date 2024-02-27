package domain

type Era struct {
	EraId int    `json:"eraId" col:"era_id" pk:"true"`
	Name  string `json:"name" col:"name"`
	Sortable
}

type EraRepository interface {
	GetList() ([]Era, error)
	GetById(id int) (*Era, error)
}
