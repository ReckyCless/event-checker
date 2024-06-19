package app

type Role struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" binding:"required" db:"name"`
}

type UpdateRoleInput struct {
	Name *string `json:"name" binding:"required" db:"name"`
}
