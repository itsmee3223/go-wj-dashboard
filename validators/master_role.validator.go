package validators

import (
	"wj-dashboard/model"

	"github.com/go-playground/validator/v10"
)

type IMasterRoleValidator interface {
	Validate(role *model.MasterRole) error
}

type masterRoleValidator struct {
	validate *validator.Validate
}

func NewMasterRoleValidator() IMasterRoleValidator {
	return &masterRoleValidator{
		validate: validator.New(),
	}
}

func (v *masterRoleValidator) Validate(role *model.MasterRole) error {
	return v.validate.Struct(role)
}
