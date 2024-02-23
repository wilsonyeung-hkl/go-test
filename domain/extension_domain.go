package domain

type Extension struct {
	ExtensionId int    `json:"extensionId"`
	Code        string `json:"code"`
	Icon        string `json:"icon"`
	Name        string `json:"name"`
	EraId       int    `json:"eraId"`
	Sortable
}

type ExtensionRepository interface {
	GetList() ([]Extension, error)
	GetListByEra(eraId int) ([]Extension, error)
	GetById(id int) (*Extension, error)
	GetByCode(code string) (*Extension, error)
}
