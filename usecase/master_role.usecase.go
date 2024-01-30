package usecase

import (
	"wj-dashboard/model"
	"wj-dashboard/repository"
)

type IMasterRoleUsecase interface {
	GetAllMasterRoles() ([]model.MasterRoleResponse, error)
	GetMasterRoleById(roleId string) (model.MasterRoleResponse, error)
	CreateMasterRole(role model.MasterRole) (model.MasterRoleResponse, error)
	UpdateMasterRole(role model.MasterRole, roleId string) (model.MasterRoleResponse, error)
	DeleteMasterRole(roleId string) error
}

type masterRoleUsecase struct {
	mr repository.IMasterRoleRepository
}

func NewMasterRoleUsecase(mr repository.IMasterRoleRepository) IMasterRoleUsecase {
	return &masterRoleUsecase{mr}
}

func (uc *masterRoleUsecase) GetAllMasterRoles() ([]model.MasterRoleResponse, error) {
	roles := []model.MasterRole{}
	if err := uc.mr.GetAllMasterRoles(&roles); err != nil {
		return nil, err
	}
	resRoles := []model.MasterRoleResponse{}
	for _, role := range roles {
		r := role.ToResponse()
		resRoles = append(resRoles, r)
	}
	return resRoles, nil
}

func (uc *masterRoleUsecase) GetMasterRoleById(roleId string) (model.MasterRoleResponse, error) {
	role := model.MasterRole{}

	if err := uc.mr.GetMasterRoleById(&role, roleId); err != nil {
		return model.MasterRoleResponse{}, err
	}
	resRole := role.ToResponse()
	return resRole, nil
}

func (uc *masterRoleUsecase) CreateMasterRole(role model.MasterRole) (model.MasterRoleResponse, error) {
	if err := uc.mr.CreateMasterRole(&role); err != nil {
		return model.MasterRoleResponse{}, err
	}

	resRole := role.ToResponse()
	return resRole, nil
}

func (uc *masterRoleUsecase) UpdateMasterRole(role model.MasterRole, roleId string) (model.MasterRoleResponse, error) {
	if err := uc.mr.UpdateMasterRole(&role, roleId); err != nil {
		return model.MasterRoleResponse{}, err
	}
	resRole := role.ToResponse()
	return resRole, nil
}

func (uc *masterRoleUsecase) DeleteMasterRole(roleId string) error {
	if err := uc.mr.DeleteMasterRole(roleId); err != nil {
		return err
	}
	return nil
}
