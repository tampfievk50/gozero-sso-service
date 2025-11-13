package model

import "time"

type User struct {
	VModel
	Username    string     `gorm:"size:255" json:"username"`
	Email       string     `gorm:"size:255" json:"email"`
	Password    string     `gorm:"size:255" json:"password"`
	ResourceID  *uint      `json:"resource_id"`
	IsSupper    bool       `gorm:"default:false" json:"is_supper"`
	LastIP      string     `gorm:"size:255" json:"last_ip"`
	LastLogin   *time.Time `json:"last_login"`
	Avatar      string     `gorm:"size:255" json:"avatar"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	Age         *int32     `json:"age"`
	Gender      int32      `gorm:"default:0" json:"gender"` // 1=Male, 2=Female, 3=Other
	FirstName   string     `gorm:"size:255" json:"first_name"`
	LastName    string     `gorm:"size:255" json:"last_name"`
	IsActive    bool       `gorm:"default:true" json:"is_active"`

	Roles []Role `gorm:"many2many:user_role;joinForeignKey:UserID;joinReferences:RoleID" json:"roles"`
}
