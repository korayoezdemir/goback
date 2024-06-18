package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/koray.oezdemir/go-back-end/internal/api"
	"github.com/koray.oezdemir/go-back-end/internal/db"
	"github.com/koray.oezdemir/go-back-end/internal/model"
	"github.com/koray.oezdemir/go-back-end/internal/repo"
	"github.com/koray.oezdemir/go-back-end/internal/service"
)

func main() {
    fmt.Println("Hello, World!")

    database := db.GetDB()
    
    //! Migrate the schema
    database.AutoMigrate(&model.User{})
    database.AutoMigrate(&model.Recipe{})
    database.AutoMigrate(&model.Ingredient{})
    database.AutoMigrate(&model.RecipeIngredient{})

    userRepo := repo.NewUserRepository(database)
	userService := service.NewUserService(userRepo)
	
    server := api.NewServer(userService)

	router := gin.Default()
	server.SetupRoutes(router)

	router.Run(":8080") // Startet den Server auf Port 8080

}
