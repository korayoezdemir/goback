// internal/api/user_controller.go

package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/koray.oezdemir/go-back-end/internal/model"
	"github.com/koray.oezdemir/go-back-end/internal/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}


func (uc *UserController) GetUser(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Ungültige Benutzer-ID"})
        return
    }

    user, err := uc.userService.GetUser(uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Fehler beim Abrufen des Benutzers"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Benutzer gefunden", "user": user})
}

func (uc *UserController) CreateUser(c *gin.Context) {
    var newUser model.User
    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := uc.userService.CreateUser(&newUser); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Fehler beim Erstellen des Benutzers"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Benutzer erstellt", "user": newUser})
}


func (uc *UserController) RegisterRoutes(router *gin.Engine) {
	router.GET("/users/:id", uc.GetUser)
	router.POST("/users", uc.CreateUser)
}