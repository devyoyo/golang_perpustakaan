package routes

import (
	"net/http"
	"project_perpustakaan/auth"
	"project_perpustakaan/config"
	"project_perpustakaan/models"

	"github.com/gin-gonic/gin"
)

func GenerateToken(c *gin.Context) {
	reqToken := models.TokenRequest{}
	petugas := models.Petugas{}

	if err := c.ShouldBindJSON(&reqToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mesage": "bad request",
			"error":  err.Error(),
		})

		c.Abort()
		return
	}

	// check email
	checkNip := config.DB.Where("email = ?", reqToken.Nip).First(&petugas)
	if checkNip.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "email not found",
			"error":   checkNip.Error.Error(),
		})

		c.Abort()
		return
	}

	// check password
	credentialError := petugas.CheckPassword(reqToken.Password)
	if credentialError != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "password not match",
			"error":   credentialError.Error(),
		})

		c.Abort()
		return
	}

	// generate token
	tokenString, err := auth.GenerateJWT(petugas.Nip, petugas.Username, petugas.Role)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed generate token",
			"error":   err.Error(),
		})

		c.Abort()
		return
	}

	// response token

	c.JSON(http.StatusCreated, gin.H{
		"token": tokenString,
	})
}
