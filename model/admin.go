package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	ID        uuid.UUID      `json:"id" gorm:"type:char(36);primary_key"`
	Name      string         `json:"name" gorm:"type:varchar(255)"`
	Username  string         `json:"username" gorm:"type:varchar(255);unique"`
	Email     string         `json:"email" gorm:"type:varchar(255);unique"`
	Password  string         `json:"password" gorm:"type:varchar(255)"`
	RoleID    uuid.UUID      `json:"role_id" gorm:"type:char(36);index"`
	Role      MasterRole     `json:"role" gorm:"foreignkey:RoleID"`
	CreatedAt time.Time      `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"type:datetime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"type:datetime;index"`
}

type AdminResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	RoleID    uuid.UUID `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (admin *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	admin.ID = uuid.New()
	return
}
