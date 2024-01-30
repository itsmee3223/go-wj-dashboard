package router

import (
	"net/http"
	"wj-dashboard/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(rc controller.IMasterRoleController) *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	r := e.Group("/roles")
	r.GET("", rc.GetAllMasterRoles)
	r.GET("/:roleId", rc.GetMasterRoleById)
	r.POST("", rc.CreateMasterRole)
	r.PUT("/:roleId", rc.UpdateMasterRole)
	r.DELETE("/:roleId", rc.DeleteMasterRole)

	return e
}
