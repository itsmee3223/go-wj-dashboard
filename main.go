package main

import (
	"wj-dashboard/controller"
	"wj-dashboard/db"
	"wj-dashboard/repository"
	"wj-dashboard/router"
	"wj-dashboard/usecase"
	"wj-dashboard/validators"
	"wj-dashboard/model"
)

func main() {
	db := db.NewDB()
	db.AutoMigrate(&model.MasterRole{}, &model.Admin{})

	masterRoleValidator := validators.NewMasterRoleValidator()
	masterRoleRepository := repository.NewMasterRoleRepository(db)
	masterRoleUsecase := usecase.NewMasterRoleUsecase(masterRoleRepository)
	masterRoleController := controller.NewMasterRoleController(masterRoleUsecase, masterRoleValidator)

	adminValidator := validators.NewAdminValidator()
	adminRepository := repository.NewAdminRepository(db)
	adminUsecase := usecase.NewAdminUsecase(adminRepository)
	adminController := controller.NewAdminController(adminUsecase, adminValidator)

	e := router.NewRouter(masterRoleController, adminController)

	e.Logger.Fatal(e.Start(":8080"))
}
