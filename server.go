package main

import (
	"fmt"
	"log"
	"net/http"

	"Construction-API/graph"
	"Construction-API/graph/model"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var db *gorm.DB
var e *echo.Echo

func initDB() {
	var err error
	dataSourceName := "root:@tcp(localhost:3306)/test_db?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	db.LogMode(true)

	db.AutoMigrate(&model.Location{}, &model.Project{}, &model.Staff{}, &model.Task{}, &model.Department{})
}

func Middleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
	}))
}

func main() {
	initDB()

	e = echo.New()
	Middleware(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to the GraphQL playground")
	})

	graphqlHandler := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: db}}))
	playgroundHandler := playground.Handler("GraphQL playground", "/query")

	e.POST("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	log.Fatal(e.Start(":" + "8080"))
}
