package domain

type Question struct {
	QuestionId   int    `json:"questionId" db:"question_id"`
	Publisher    string `json:"publisher" db:"publisher"`
	PublishYear  int    `json:"publishYear" db:"publish_year"`
	School       string `json:"school" db:"school"`
	Grade        int    `json:"grade" db:"grade"`
	QuestionType int    `json:"questionType" db:"question_type"`
	Language     string `json:"language" db:"language"`
	Question     string `json:"question" db:"question"`
	Answer       string `json:"answer" db:"answer"`
	CreateActionLog
	UpdateActionLog
	Status
}
