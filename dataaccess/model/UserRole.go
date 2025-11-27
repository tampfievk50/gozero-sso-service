package model

type UserRole struct {
	UserID     uint `gorm:"primaryKey" json:"user_id"`
	RoleID     uint `gorm:"primaryKey" json:"role_id"`
	ResourceID uint `gorm:"primaryKey" json:"resource_id"`

	User     User     `gorm:"foreignKey:UserID;references:ID" json:"user"`
	Role     Role     `gorm:"foreignKey:RoleID;references:ID" json:"role"`
	Resource Resource `gorm:"foreignKey:ResourceID;references:ID" json:"resource"`
}
