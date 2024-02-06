package router

import (
	"net/http"
	"wj-dashboard/controller"
	"wj-dashboard/middleware"

	"github.com/labstack/echo/v4"
)

func NewRouter(rc controller.IMasterRoleController, ac controller.IAdminController) *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/login", ac.Login)
	e.POST("/register", ac.Reister)

	r := e.Group("/roles", middleware.JWTMiddleware)
	r.GET("", rc.GetAllMasterRoles)
	r.GET("/:roleId", rc.GetMasterRoleById)
	r.POST("", rc.CreateMasterRole)
	r.PUT("/:roleId", rc.UpdateMasterRole)
	r.DELETE("/:roleId", rc.DeleteMasterRole)

	return e
}
