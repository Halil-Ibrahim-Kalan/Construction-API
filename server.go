package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Halil-Ibrahim-Kalan/Construction-API/graph"
	"github.com/Halil-Ibrahim-Kalan/Construction-API/graph/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func initDB() {
	var err error
	dataSourceName := "root:@tcp(localhost:3306)/?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	db.LogMode(true)

	// Create the database. This is a one-time step.
	// Comment out if running multiple times - You may see an error otherwise
	db.Exec("CREATE DATABASE test_db")
	db.Exec("USE test_db")

	// Migration to create tables for Order and Item schema
	db.AutoMigrate(&model.Location{}, &model.Project{}, &model.Staff{}, &model.Task{})
}

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	initDB()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
