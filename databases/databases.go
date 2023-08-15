package databases

import (
	"Construction-API/graph/model"
	"Construction-API/utils"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	Db := connectPostgresDB()
	Db.AutoMigrate(&model.Location{}, &model.Project{}, &model.Staff{}, &model.Task{}, &model.Department{})
	return Db
}

func connectMysqlDB() *gorm.DB {
	config := utils.LoadConfig()
	dataSourceName := config.DBUser + ":" + config.DBPassword + "@tcp" + "(" + config.DBHost + ":" + config.DBPort + ")/" + config.DBName + "?" + "parseTime=true"
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		panic("failed to connect database! Error: " + err.Error())
	}
	return db
}

func connectPostgresDB() *gorm.DB {
	config := utils.LoadConfig()
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName, config.DBSSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database! Error: " + err.Error())
	}

	return db
}
