package domain

type Grade struct {
	GradeId int    `json:"grade_id" db:"grade_id"`
	Grade   string `json:"grade" db:"grade"`
	Status
	Sortable
}
