package domain

type EnergyType struct {
	EnergyTypeId int    `json:"energyTypeId"`
	Name         string `json:"name"`
	Icon         string `json:"icon"`
	Sortable
}

type EnergyTypeRepository interface {
	GetList() ([]EnergyType, error)
	GetById(id int) (*EnergyType, error)
}
