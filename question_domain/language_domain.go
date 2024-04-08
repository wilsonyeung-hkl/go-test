package domain

type Language struct {
	LanguageId int    `json:"languageId" db:"language_id"`
	Language   string `json:"language" db:"language"`
	Status
	Sortable
}
