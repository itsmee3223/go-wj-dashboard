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

type IMasterRoleController interface {
	GetAllMasterRoles(c echo.Context) error
	GetMasterRoleById(c echo.Context) error
	CreateMasterRole(c echo.Context) error
	UpdateMasterRole(c echo.Context) error
	DeleteMasterRole(c echo.Context) error
}

type masterRoleController struct {
	mru usecase.IMasterRoleUsecase
	mrv validators.IMasterRoleValidator
}

func NewMasterRoleController(mru usecase.IMasterRoleUsecase, mrv validators.IMasterRoleValidator) IMasterRoleController {
	return &masterRoleController{mru, mrv}
}

func (mc *masterRoleController) GetAllMasterRoles(c echo.Context) error {
	roles, err := mc.mru.GetAllMasterRoles()
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, []string{err.Error()})
	}
	return utils.SuccessResponse(c, http.StatusOK, roles)
}

func (mc *masterRoleController) GetMasterRoleById(c echo.Context) error {
	roleId := c.Param("roleId")
	role, err := mc.mru.GetMasterRoleById(roleId)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, []string{err.Error()})
	}
	return utils.SuccessResponse(c, http.StatusOK, role)
}

func (mc *masterRoleController) CreateMasterRole(c echo.Context) error {
	role := new(model.MasterRole)
	if err := c.Bind(role); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, []string{err.Error()})
	}

	if err := mc.mrv.Validate(role); err != nil {
		var validationErrors []string
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, e := range errs {
				validationErrors = append(validationErrors, e.Field()+" is "+e.Tag())
			}
		}
		return utils.ErrorResponse(c, http.StatusBadRequest, validationErrors)
	}

	createdRole, err := mc.mru.CreateMasterRole(*role)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, []string{err.Error()})
	}
	return utils.SuccessResponse(c, http.StatusCreated, createdRole)
}

func (mc *masterRoleController) UpdateMasterRole(c echo.Context) error {
	role := new(model.MasterRole)
	if err := c.Bind(role); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, []string{err.Error()})
	}

	roleId := c.Param("roleId")
	updatedRole, err := mc.mru.UpdateMasterRole(*role, roleId)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, []string{err.Error()})
	}
	return utils.SuccessResponse(c, http.StatusCreated, updatedRole)
}

func (mc *masterRoleController) DeleteMasterRole(c echo.Context) error {
	roleId := c.Param("roleId")
	err := mc.mru.DeleteMasterRole(roleId)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, []string{err.Error()})
	}
	return utils.InfoResponse(c, http.StatusOK, "Role successfully deleted")
}
