package routes

import (
	"log"
	"net/http"
	"project_perpustakaan/config"
	"project_perpustakaan/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetPetugas(c *gin.Context) {
	Petugass := []models.Petugas{}

	config.DB.Find(&Petugass)

	ResponsePetugas := []models.ResponsePetugas{}

	for _, Petugas := range Petugass {
		data := models.ResponsePetugas{
			ID:       Petugas.ID,
			Name:     Petugas.Name,
			Username: Petugas.Username,
			Password: Petugas.Password,
		}

		ResponsePetugas = append(ResponsePetugas, data)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success get Petugas",
		"data":    ResponsePetugas,
	})

}

func GetPetugasById(c *gin.Context) {
	id := c.Param("id")

	var Petugas models.Petugas

	data := config.DB.First(&Petugas, "id = ?", id)

	if data.Error != nil {
		log.Printf(data.Error.Error())

		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "data not found",
		})

		return
	}

	ResponsePetugas := models.ResponsePetugas{
		ID:       Petugas.ID,
		Name:     Petugas.Name,
		Username: Petugas.Username,
		Password: Petugas.Password,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success get Petugas",
		"data":    ResponsePetugas,
	})
}

func PostPetugas(c *gin.Context) {
	var RequestPetugas models.Petugas

	c.BindJSON(&RequestPetugas)

	err := map[string]string{}

	if RequestPetugas.Name == "" {
		err["name"] = "nama is required"
	}
	if RequestPetugas.Username == "" {
		err["username"] = "username is required"
	}
	if RequestPetugas.Password == "" {
		err["password"] = "password is required"
	}

	if len(err) > 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"error":   err,
			"message": "failed created Petugas",
		})
	} else {
		hash, _ := HashPassword(RequestPetugas.Password)

		Petugas := models.Petugas{
			Name:     RequestPetugas.Name,
			Username: RequestPetugas.Username,
			Password: hash,
		}

		config.DB.Create(&Petugas)

		c.JSON(http.StatusCreated, gin.H{
			"data":    RequestPetugas,
			"message": "success created Petugas",
		})
	}

}

func PutPetugas(c *gin.Context) {
	id := c.Param("id")

	var Petugas models.Petugas

	data := config.DB.First(&Petugas, "id = ?", id)

	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})

		return
	}

	var RequestPetugas models.Petugas

	c.BindJSON(&RequestPetugas)

	err := map[string]string{}

	if RequestPetugas.Name == "" {
		err["name"] = "nama is required"
	}
	if RequestPetugas.Username == "" {
		err["username"] = "username is required"
	}
	if RequestPetugas.Password == "" {
		err["password"] = "password is required"
	}

	if len(err) > 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"error":   err,
			"message": "failed created Petugas",
		})
	} else {
		hash, _ := HashPassword(RequestPetugas.Password)

		update_Petugas := models.Petugas{
			Name:     RequestPetugas.Name,
			Username: RequestPetugas.Username,
			Password: hash,
		}

		update := config.DB.Model(&Petugas).Where("id = ?", id).Updates(&update_Petugas)

		if update.Error != nil {
			log.Printf(update.Error.Error())

			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": update.Error.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "update success",
			"data":    RequestPetugas,
		})
	}

}

func DeletePetugas(c *gin.Context) {
	id := c.Param("id")

	var Petugas models.Petugas

	data := config.DB.First(&Petugas, "id = ?", id)

	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})

		return
	}

	config.DB.Delete(&Petugas, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "delete success",
	})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
