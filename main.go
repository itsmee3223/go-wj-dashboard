package main

import (
	"wj-dashboard/controller"
	"wj-dashboard/db"
	"wj-dashboard/repository"
	"wj-dashboard/router"
	"wj-dashboard/usecase"
	"wj-dashboard/validators"
)

func main() {
	db := db.NewDB()

	masterRoleValidator := validators.NewMasterRoleValidator()
	masterRoleRepository := repository.NewMasterRoleRepository(db)
	masterRoleUsecase := usecase.NewMasterRoleUsecase(masterRoleRepository)
	masterRoleController := controller.NewMasterRoleController(masterRoleUsecase, masterRoleValidator)

	e := router.NewRouter(masterRoleController)

	e.Logger.Fatal(e.Start(":8080"))
}
