package domain

import (
	"gopkg.in/guregu/null.v3"
)

type Card struct {
	CardId         int    `json:"cardId"`
	Name           string `json:"name"`
	Image          string `json:"image"`
	EnergyTypeId   int    `json:"energyTypeId"`
	WeaknessTypeId int    `json:"weaknessTypeId"`
	RarityId       int    `json:"rarityId"`
	ExtensionId    int    `json:"extensionId"`
	Code           string `json:"code"`
	Sortable
}

type CardRepository interface {
	Search(energyTypeId null.Int, weaknessTypeId null.Int, rarityId null.Int, extensionId null.Int, pageSize int, offset int) ([]Card, error)
	GetById(id int) (*Card, error)
}
