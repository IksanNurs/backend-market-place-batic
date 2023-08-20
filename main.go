package main

import (
	"os"
	"time"

	_ "embed"

	database "e-commerce/database"
	handler "e-commerce/handler"
	utils1 "e-commerce/utils"

	middlewares "e-commerce/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//go:embed .env
var env string

func main() {
	utils1.LoadEnv(env)
	database.StartDB()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Static("/img", os.Getenv("TEMPLATE_DIR")+"/img")

	api := r.Group("/api")
	{
		//user
		api.POST("/auth/login", handler.Login)
		api.POST("/auth/register", handler.Register)

		//category
		api.POST("/category", middlewares.Authentication(), handler.PostCategory)
		api.PUT("/category", middlewares.Authentication(), handler.PutCategory)
		api.GET("/category", middlewares.Authentication(), handler.GetCategory)
		api.GET("/category/:id", middlewares.Authentication(), handler.GetOneCategory)

		//motif
		api.POST("/motif", middlewares.Authentication(), handler.PostMotif)
		api.PUT("/motif", middlewares.Authentication(), handler.PutMotif)
		api.GET("/motif", middlewares.Authentication(), handler.GetMotif)
		api.GET("/motif/:id", middlewares.Authentication(), handler.GetOneMotif)

		//product
		api.POST("/product", middlewares.Authentication(), handler.PostProduct)
		api.PUT("/product", middlewares.Authentication(), handler.PutProduct)
		api.GET("/product", middlewares.Authentication(), handler.GetProduct)
		api.GET("/product/sales", middlewares.Authentication(), handler.GetProductSales)
		api.GET("/product/:id", middlewares.Authentication(), handler.GetOneProduct)

		//size
		api.POST("/size", middlewares.Authentication(), handler.PostSize)
		api.PUT("/size", middlewares.Authentication(), handler.PutSize)

		//shipping
		api.POST("/shipping", middlewares.Authentication(), handler.PostShipping)
		api.PUT("/shipping", middlewares.Authentication(), handler.PutShipping)

	}

	r.Run(":" + os.Getenv("PORT"))
}
