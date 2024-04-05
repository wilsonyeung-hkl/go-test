package domain

type Subtopic struct {
	SubtopicId int    `json:"subtopicId" db:"subtopic_id"`
	Subtopic   string `json:"subtopic" db:"subtopic"`
	CreateActionLog
	UpdateActionLog
	Status
	Sortable
}
