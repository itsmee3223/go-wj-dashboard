package validators

import (
	"fmt"
	"strings"
	"wj-dashboard/model"

	"github.com/go-playground/validator/v10"
)

type IAdminValidator interface {
	ValidateRegister(admin *model.Admin) error
	ValidateLogin(admin *model.Admin) error
}

type adminValidator struct {
	validate *validator.Validate
}

func NewAdminValidator() IAdminValidator {
	v := validator.New()
	return &adminValidator{validate: v}
}

func (v *adminValidator) ValidateRegister(admin *model.Admin) error {
	return v.validate.StructExcept(admin, "Role")
}

func (v *adminValidator) ValidateLogin(admin *model.Admin) error {
	var validationErrors []string

	if admin.Password == "" {
		validationErrors = append(validationErrors, "Password is required")
	}

	if admin.Email == "" && admin.Username == "" {
		validationErrors = append(validationErrors, "Either Email or Username is required")
	}

	if len(validationErrors) > 0 {
		return fmt.Errorf(strings.Join(validationErrors, ", "))
	}

	return nil
}
