package domain

type Topic struct {
	TopicId int    `json:"topicId" db:"topic_id"`
	Topic   string `json:"topic" db:"topic"`
	CreateActionLog
	UpdateActionLog
	Status
	Sortable
}
