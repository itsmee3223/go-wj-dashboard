package controller

import (
	"net/http"
	"wj-dashboard/model"
	"wj-dashboard/usecase"
	"wj-dashboard/utils"
	"wj-dashboard/validators"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type IAdminController interface {
	Reister(c echo.Context) error
	Login(c echo.Context) error
}

type adminController struct {
	au usecase.IAdminUsecase
	av validators.IAdminValidator
}

func NewAdminController(au usecase.IAdminUsecase, av validators.IAdminValidator) IAdminController {
	return &adminController{au, av}
}

func (ac *adminController) Reister(c echo.Context) error {
	admin := new(model.Admin)
	if err := c.Bind(&admin); err != nil {
		if utils.IsUUIDValid(admin.RoleID.String()) {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid UUID length"})
		}
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if !utils.IsUUIDValid(admin.RoleID.String()) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid UUID"})
	}

	if err := ac.av.ValidateRegister(admin); err != nil {
		var validationErrors []string
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, e := range errs {
				validationErrors = append(validationErrors, e.Field()+" is "+e.Tag())
			}
		}
		return utils.ErrorResponse(c, http.StatusBadRequest, validationErrors)
	}

	userRes, err := ac.au.RegisterAdmin(admin)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, []string{err.Error()})
	}

	return c.JSON(http.StatusCreated, userRes)
}

func (ac *adminController) Login(c echo.Context) error {
	admin := new(model.Admin)
	if err := c.Bind(admin); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, []string{err.Error()})
	}

	if err := ac.av.ValidateLogin(admin); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, []string{err.Error()})
	}

	var token string
	var err error

	if admin.Email != "" {
		token, err = ac.au.LoginEmail(admin)
	} else if admin.Username != "" {
		token, err = ac.au.LoginUsername(admin)
	} else {
		return utils.ErrorResponse(c, http.StatusBadRequest, []string{"Email or Username is required"})
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
