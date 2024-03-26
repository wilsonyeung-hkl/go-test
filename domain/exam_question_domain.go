package domain

type ExamQuestion struct {
	ExamQuestionId int    `json:"examQuestionId" db:"exam_question_id"`
	Question       string `json:"question" db:"question"`
	Answer         string `json:"answer" db:"answer"`
	CreateActionLog
	UpdateActionLog
	Status
	Sortable
}
