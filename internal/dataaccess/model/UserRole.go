package model

type UserRole struct {
	UserID uint `gorm:"primaryKey" json:"user_id"`
	RoleID uint `gorm:"primaryKey" json:"role_id"`

	User User `gorm:"foreignKey:UserID;references:ID" json:"user"`
	Role Role `gorm:"foreignKey:RoleID;references:ID" json:"role"`
}
