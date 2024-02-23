package domain

type Rarity struct {
	RarityId int    `json:"rarityId"`
	Name     string `json:"name"`
	Icon     string `json:"icon"`
	Sortable
}

type RarityRepository interface {
	GetList() ([]Rarity, error)
	GetById(id int) (*Rarity, error)
}
