package model

type Role struct {
	VModel
	Name        string `gorm:"size:255" json:"name"`
	Description string `gorm:"size:255" json:"description"`
	ResourceID  *uint  `json:"resource_id"`
	Users       []User `gorm:"many2many:user_role;joinForeignKey:RoleID;joinReferences:UserID" json:"users"`
}
