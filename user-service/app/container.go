package app

import (
	"micro-warehouse/user-service/config"
	"micro-warehouse/user-service/controllers"
	"micro-warehouse/user-service/database"
	"micro-warehouse/user-service/repositories"
	"micro-warehouse/user-service/usecase"

	"github.com/gofiber/fiber/v2/log"
)

type Container struct {
	RoleController controllers.RoleControllerInterface
}

func BuildContainer() *Container {
	config := config.NewConfig()
	db, err := database.ConnectionPostgres(*config)
	if err != nil {
		log.Fatalf("Failed connect to database: %v", err)
	}

	roleRepo := repositories.NewRoleRepository(db.DB)
	roleUsecase := usecase.NewRoleUsecase(roleRepo)
	roleController := controllers.NewRoleController(roleUsecase)

	return &Container{
		RoleController: roleController,
	}
}