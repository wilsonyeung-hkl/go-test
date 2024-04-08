package domain

type Subject struct {
	SubjectId int    `json:"subjectId" db:"subject_id"`
	Subject   string `json:"subject" db:"subject"`
	Status
	Sortable
}
