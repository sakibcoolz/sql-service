package database

import (
	"fmt"
	"log"
	"sql-service/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBInstance struct {
	DB *gorm.DB
}

func Connection(configSetup *config.Config) *DBInstance {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configSetup.Creden.User,
		configSetup.Creden.Password,
		configSetup.Creden.Host,
		configSetup.Creden.Port,
		configSetup.Creden.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return &DBInstance{DB: db}
}
