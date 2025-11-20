package model

type Permission struct {
	VModel
	Name        string `gorm:"size:255" json:"name"`
	Description string `gorm:"size:255" json:"description"`
	Path        string `gorm:"size:255" json:"path"`
	Action      string `gorm:"size:255" json:"action"`
}
