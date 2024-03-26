package domain

type RolePermission struct {
	RolePermissionId int `json:"rolePermissionId" db:"role_permission_id"`
	RoleId           int `json:"roleId" db:"role_id"`
	PermissionId     int `json:"permissionId" db:"permission_id"`
	CreateActionLog
	UpdateActionLog
	Status
}
