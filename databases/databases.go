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
	Db.AutoMigrate(&model.Department{}, &model.Location{}, &model.Project{}, &model.StaffData{}, &model.TaskData{})

	return Db
}

func connectMysqlDB() *gorm.DB {
	config := utils.LoadConfig()
	dataSourceName := config.DBUser + ":" + config.DBPassword + "@tcp" + "(" + config.DBHost + ":" + config.DBPort + ")/" + config.DBName + "?" + "parseTime=true"
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		panic("failed to connect database! Error: " + err.Error())
	}
	return db
}

func connectPostgresDB() *gorm.DB {
	config := utils.LoadConfig()
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		// Logger:                                   logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database! Error: " + err.Error())
	}

	return db
}
