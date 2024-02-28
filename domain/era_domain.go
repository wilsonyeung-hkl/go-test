package domain

type Era struct {
	EraId int    `json:"eraId" db:"era_id" key:"primary"`
	Name  string `json:"name" db:"name"`
	Sortable
}

type EraRepository interface {
	GetList() ([]Era, error)
	GetById(id int) (*Era, error)
}
