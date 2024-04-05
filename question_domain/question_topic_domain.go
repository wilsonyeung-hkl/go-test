package domain

type QuestionTopic struct {
	QuestionId `json:"questionId" db:"question_id"`
	TopicId    `json:"topicId" db:"topic_id"`
}
