package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/xtareq/go_api/config"
	"github.com/xtareq/go_api/entity"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
}

func NewAuthController() AuthController {
	return &authController{}
}

type LoginRequest struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type EmptyObj interface{}

func (c *authController) Login(ctx *gin.Context) {
	db := config.DbConnection()

	var newUser LoginRequest
	user := entity.User{}

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	error := db.Where(entity.User{Email: newUser.Email, Password: newUser.Password}).First(&user).Error
	if error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Not Registerd!"})
		return
	}

	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: 15000,
		Issuer:    "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, tokenErr := token.SignedString(mySigningKey)
	if tokenErr != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Token Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "login successfully",
		"data":    ss,
	})
}

type RegisterRequest struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (c *authController) Register(ctx *gin.Context) {
	db := config.DbConnection()

	var newUser RegisterRequest
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := entity.User{Name: newUser.Name, Email: newUser.Email, Password: newUser.Password}
	result := db.Create(&user)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "register successfully",
		"data":    result.RowsAffected,
	})
}
