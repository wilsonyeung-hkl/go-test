package domain

type Role struct {
	RoleId int    `json:"roleId" db:"role_id"`
	Name   string `json:"name" db:"name"`
	Status
}
