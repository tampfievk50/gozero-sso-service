package model

type UserRole struct {
	UserID     string `gorm:"type:uuid;primaryKey" json:"user_id"`
	RoleID     string `gorm:"type:uuid;primaryKey" json:"role_id"`
	ResourceID string `gorm:"type:uuid;primaryKey" json:"resource_id"`

	User     User     `gorm:"foreignKey:UserID;references:ID" json:"user"`
	Role     Role     `gorm:"foreignKey:RoleID;references:ID" json:"role"`
	Resource Resource `gorm:"foreignKey:ResourceID;references:ID" json:"resource"`
}
