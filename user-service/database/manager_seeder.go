package database

import (
	"log"
	"micro-warehouse/user-service/models"
	"micro-warehouse/user-service/pkg/conv"

	"gorm.io/gorm"
)

func SeedManager(db *gorm.DB) {
	bytes, err := conv.HashPassword("manager123")
	if err != nil {
		log.Fatalf("%s:%v", err.Error(), err)
	}

	modelRole := models.Role{}
	err = db.Where("name = ?", "Manager").First(&modelRole).Error
	if err != nil {
		log.Fatalf("%s:%v", err.Error(), err)
	}

	admin := models.User{
		Name:     "manager",
		Email:    "manager@mail.com",
		Password: bytes,
		Roles:    []models.Role{modelRole},
	}

	if err := db.FirstOrCreate(&admin, models.User{Email: "manager@mail.com"}).Error; err != nil {
		log.Fatalf("%s:%v", err.Error(), err)
	} else {
		log.Printf("Admin %s created", admin.Name)
	}

}
