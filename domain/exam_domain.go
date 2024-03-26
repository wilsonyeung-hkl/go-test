package domain

type Exam struct {
	ExamId int    `json:"examId" db:"exam_id"`
	Name   string `json:"name" db:"name"`
	Grade  int    `json:"grade" db:"grade"`
	CreateActionLog
	UpdateActionLog
	Status
}
