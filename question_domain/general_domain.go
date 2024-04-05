package domain

type Sortable struct {
	Sorting int `json:"sorting" db:"sorting"`
}

type CreateActionLog struct {
	CreateBy   int `json:"createBy" db:"create_by"`
	CreateTime int `json:"createTime" db:"create_time"`
}

type UpdateActionLog struct {
	UpdateBy   int `json:"updateBy" db:"update_by"`
	UpdateTime int `json:"updateTime" db:"update_time"`
}

type Status struct {
	IsActive int `json:"isActive" db:"is_active"`
	IsDelete int `json:"isDelete" db:"is_delete"`
}
