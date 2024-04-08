package domain

type Question struct {
	QuestionId      int    `json:"questionId" db:"question_id"`
	PublisherCodeId int    `json:"publisherCodeId" db:"publisher_code_id"`
	PublishYear     int    `json:"publishYear" db:"publish_year"`
	SchoolCodeId    int    `json:"schoolCodeId" db:"school_code_id"`
	GradeId         int    `json:"gradeId" db:"grade_id"`
	LanguageId      int    `json:"languageId" db:"language_id"`
	SubjectId       int    `json:"subjectId" db:"subject_id"`
	Question        string `json:"question" db:"question"`
	Answer          string `json:"answer" db:"answer"`
	Media           string `json:"media" db:"media"`
	CreateActionLog
	UpdateActionLog
	Status
}
