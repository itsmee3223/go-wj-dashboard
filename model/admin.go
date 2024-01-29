package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;"`
	Name      string         `gorm:"type:varchar(255)"`
	Username  string         `gorm:"type:varchar(255);unique"`
	Email     string         `gorm:"type:varchar(255);unique"`
	Password  string         `gorm:"type:varchar(255)"`
	RoleID    uuid.UUID      `gorm:"type:uuid;index;"`
	Role      MasterRole     `gorm:"foreignkey:RoleID"`
	CreatedAt time.Time      `gorm:"type:datetime"`
	UpdatedAt time.Time      `gorm:"type:datetime"`
	DeletedAt gorm.DeletedAt `gorm:"type:datetime;index"`
}

type AdminResponse struct {
	ID        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	RoleID    uuid.UUID      `json:"role_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (admin *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	admin.ID = uuid.New()
	return
}
