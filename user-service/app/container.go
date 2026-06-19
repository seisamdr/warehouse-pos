package app

import (
	"micro-warehouse/user-service/configs"
	"micro-warehouse/user-service/controllers"
	"micro-warehouse/user-service/database"
	"micro-warehouse/user-service/repositories"
	"micro-warehouse/user-service/service"
	"micro-warehouse/user-service/usecase"

	"github.com/gofiber/fiber/v2/log"
)

type Container struct {
	RoleController controllers.RoleControllerInterface
	UserController controllers.UserControllerInterface
}

func BuildContainer() *Container {
	config := configs.NewConfig()
	db, err := database.ConnectionPostgres(*config)
	if err != nil {
		log.Fatalf("Failed connect to database: %v", err)
	}

	rabbitMQService, err := service.NewRabbitMQService(*configs.NewConfig())
	if err != nil {
		log.Fatalf("Failed connect to RabbitMQ: %v", err)
	}

	roleRepo := repositories.NewRoleRepository(db.DB)
	roleUsecase := usecase.NewRoleUsecase(roleRepo)
	roleController := controllers.NewRoleController(roleUsecase)

	userRepo := repositories.NewUserRepository(db.DB)
	userUsecase := usecase.NewUserUsecase(userRepo, rabbitMQService)
	userController := controllers.NewUserController(userUsecase)

	return &Container{
		RoleController: roleController,
		UserController: userController,
	}
}