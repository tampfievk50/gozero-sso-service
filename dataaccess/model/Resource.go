package model

type Resource struct {
	VModel
	Name        string `gorm:"size:255" json:"name"`
	Description string `gorm:"size:255" json:"description"`
}
