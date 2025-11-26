package model

type Role struct {
	VModel
	Name        string `gorm:"size:255" json:"name"`
	Description string `gorm:"size:255" json:"description"`
}
