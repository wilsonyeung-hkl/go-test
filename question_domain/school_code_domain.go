package domain

type SchoolCode struct {
	SchoolCodeId int    `json:"schoolCodeId" db:"school_code_id"`
	SchoolCode   string `json:"schoolCode" db:"school_code"`
	CreateActionLog
	UpdateActionLog
	Status
	Sortable
}
