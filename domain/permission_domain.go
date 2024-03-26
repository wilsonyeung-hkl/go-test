package domain

type Permission struct {
	PermissionId int    `json:"permissionId" db:"permission_id"`
	Name         string `json:"name" db:"name"`
	CreateActionLog
	UpdateActionLog
	Status
}
