package domain

type PublisherCode struct {
	PublisherCodeId int    `json:"publisherCodeId" db:"publisher_code_id"`
	PublisherCode   string `json:"publisherCode" db:"publisher_code"`
	CreateActionLog
	UpdateActionLog
	Status
	Sortable
}
