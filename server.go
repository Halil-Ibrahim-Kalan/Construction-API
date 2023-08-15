package main

import (
	"log"
	"net/http"

	"Construction-API/databases"
	"Construction-API/graph"
	"Construction-API/middlewares"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func main() {
	db := databases.InitDb()

	e := echo.New()
	middlewares.Middleware(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to the GraphQL API")
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
