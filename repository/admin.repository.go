package repository

import (
	"wj-dashboard/model"

	"gorm.io/gorm"
)

type IAdminRepository interface {
	GetAdminByUsername(admin *model.Admin, username string) error
	GetAdminByEmail(admin *model.Admin, email string) error
	CreateAdmin(admin *model.Admin) error
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) IAdminRepository {
	return &adminRepository{db}
}

func (ar *adminRepository) GetAdminByUsername(admin *model.Admin, username string) error {
	if err := ar.db.Where("username=?", username).First(admin).Error; err != nil {
		return err
	}

	return nil
}

func (ar *adminRepository) GetAdminByEmail(admin *model.Admin, email string) error {
	if err := ar.db.Where("email=?", email).First(admin).Error; err != nil {
		return err
	}

	return nil
}

func (ar *adminRepository) CreateAdmin(admin *model.Admin) error {
	if err := ar.db.Create(admin).Error; err != nil {
		return err
	}

	return nil
}
