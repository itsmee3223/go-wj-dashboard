package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type MasterRole struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;"`
	Name      string         `json:"name" gorm:"type:varchar(255)" validate:"required"`
	Access    datatypes.JSON `json:"access" gorm:"type:json" validate:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type MasterRoleResponse struct {
	ID        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	Access    datatypes.JSON `json:"access"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

func (masterRole *MasterRole) BeforeCreate(tx *gorm.DB) (err error) {
	masterRole.ID = uuid.New()
	return
}

func (m MasterRole) ToResponse() MasterRoleResponse {
	return MasterRoleResponse{
		ID:        m.ID,
		Name:      m.Name,
		Access:    m.Access,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
