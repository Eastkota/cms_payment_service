package main

import (
	"cms_payment_service/repositories"
	"cms_payment_service/resolvers"
	"cms_payment_service/services"
	"cms_payment_service/graph"
	"cms_payment_service/helpers"
	"log"

	"cms_payment_service/handlers"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
        log.Fatal("Error loading .env file")
    }
	
	db, err := helpers.GetGormDB()
    if err != nil {
        log.Fatal("Failed to connect to database: " + err.Error())
    }
    cmsPaymentRepository := repositories.NewCmsPaymentRepository(db)
    cmsPaymentService := services.NewCmsPaymentService(cmsPaymentRepository)
    resolver := resolvers.NewCmsPaymentResolver(cmsPaymentService)

	queryType := schema.NewQueryType(resolver)

	schema.InitSchema(queryType)
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:3000", 
			"https://www.educareskill.com/",
		},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))
	e.POST("/graphql", handlers.Handler)
	e.Logger.Fatal(e.Start(":8100"))
}



