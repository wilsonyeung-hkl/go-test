package domain

type QuestionSubtopic struct {
	QuestionId `json:"questionId" db:"question_id"`
	SubtopicId `json:"subtopicId" db:"subtopic_id"`
}
