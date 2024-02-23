package domain

type CardGrade struct {
	CardGradeId int    `json:"cardGradeId"`
	Code        string `json:"code"`
	Sortable
}

type CardGradeRepository interface {
	GetList() ([]CardGrade, error)
}
