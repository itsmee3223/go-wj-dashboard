package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type MasterRole struct {
	ID        uuid.UUID      `json:"id" gorm:"type:char(36);primary_key"`
	Name      string         `json:"name" gorm:"type:varchar(255);not null" validate:"required"`
	Access    datatypes.JSON `json:"access" gorm:"type:json;not null" validate:"required"`
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

func (m MasterRole) ConvertToResponse() MasterRoleResponse {
	response := MasterRoleResponse(m)
	return response
}

func (m MasterRole) ToResponse() MasterRoleResponse {
	return m.ConvertToResponse()
}
