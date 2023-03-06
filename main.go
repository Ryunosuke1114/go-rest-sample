package main

import (
	"main/handler"
	"main/infra"
	"log"
	"main/service"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := infra.DBInit() 
	factory := service.NewService(engine)

	defer func(){
		log.Println("engine closed")
		engine.Close()
	}()

	g := gin.Default()

	routes := g.Group("/v1")

	{
		routes.POST("/users" , handler.Create)
		routes.GET("/users" , handler.GetAll)
		routes.GET("/users/:user-id", handler.GetOne)
		routes.PUT("/users/:user-id" , handler.Update)
		routes.DELETE("/users/:user-id" , handler.Delete)
	}

	g.Run(":3000")
}
