package main

import (
	"Construction-API/databases"
	"Construction-API/graph"
	"Construction-API/middleware"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

func main() {
	db := databases.InitDb()

	e := echo.New()
	middleware.Middleware(e)

	resolver := graph.NewResolver(db)

	graphqlHandler := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))
	graphqlHandler.AddTransport(transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	playgroundHandler := playground.Handler("GraphQL playground", "/query")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to the GraphQL API")
	})

	e.POST("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	println("http://localhost:8080/playground")
	log.Fatal(e.Start(":" + "8080"))
}
