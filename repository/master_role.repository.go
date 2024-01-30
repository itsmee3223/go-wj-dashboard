package repository

import (
	"errors"
	"wj-dashboard/model"

	"gorm.io/gorm"
)

type IMasterRoleRepository interface {
	GetAllMasterRoles(roles *[]model.MasterRole) error
	GetMasterRoleById(role *model.MasterRole, roleId string) error
	CreateMasterRole(role *model.MasterRole) error
	UpdateMasterRole(role *model.MasterRole, roleId string) error
	DeleteMasterRole(roleId string) error
}

type masterRoleRepository struct {
	db *gorm.DB
}

func NewMasterRoleRepository(db *gorm.DB) IMasterRoleRepository {
	return &masterRoleRepository{db}
}

func (r *masterRoleRepository) GetAllMasterRoles(roles *[]model.MasterRole) error {
	if err := r.db.Find(roles).Error; err != nil {
		return err
	}
	return nil
}

func (r *masterRoleRepository) GetMasterRoleById(role *model.MasterRole, roleId string) error {
	if err := r.db.Where("id = ?", roleId).First(role).Error; err != nil {
		return err
	}
	return nil
}

func (r *masterRoleRepository) CreateMasterRole(role *model.MasterRole) error {
	if err := r.db.Create(role).Error; err != nil {
		return err
	}
	return nil
}

func (r *masterRoleRepository) UpdateMasterRole(role *model.MasterRole, roleId string) error {
	if err := r.db.Where("id = ?", roleId).Updates(role).Error; err != nil {
		return err
	}
	return nil
}

func (r *masterRoleRepository) DeleteMasterRole(roleId string) error {
	var role model.MasterRole
	if err := r.db.Where("id = ?", roleId).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}

	if err := r.db.Where("id = ?", roleId).Delete(&model.MasterRole{}).Error; err != nil {
		return err
	}
	return nil
}
