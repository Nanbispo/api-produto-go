package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	dbConnection, err := db.ConncetDB()
	if err != nil {
		panic(err)
	}
	//camada de repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	//camada usecase
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)
	//Camada de controllers
	ProductController := controller.NewProductController(ProductUsecase)

	server := gin.Default()

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreatedProduct)

	server.Run(":8000")
}
