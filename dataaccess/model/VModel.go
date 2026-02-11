package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type VModel struct {
	ID        string       `gorm:"type:uuid;default:gen_random_uuid();primarykey" json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"index" json:"-"`
	CreatedBy string       `gorm:"type:uuid" json:"created_by"`
	UpdatedBy string       `gorm:"type:uuid" json:"updated_by"`
	DeletedBy string       `gorm:"type:uuid" json:"deleted_by"`
	IsDeleted bool         `gorm:"index" json:"is_deleted"`
}

func (v *VModel) BeforeCreate(tx *gorm.DB) (err error) {
	v.IsDeleted = false
	atTime := time.Now()
	v.CreatedAt = atTime
	v.UpdatedAt = atTime
	return
}

func (v *VModel) BeforeUpdate(tx *gorm.DB) (err error) {
	v.UpdatedAt = time.Now()
	return
}

func (v *VModel) BeforeDelete(tx *gorm.DB) (err error) {
	v.DeletedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	v.IsDeleted = true
	return
}
