package routes

import (
	"fmt"
	"log"
	"net/http"
	"project_perpustakaan/config"
	"project_perpustakaan/models"

	"github.com/gin-gonic/gin"
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
			Nip:      Petugas.Nip,
			Role:     Petugas.Role,
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
		Nip:      Petugas.Nip,
		Role:     Petugas.Role,
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
	if RequestPetugas.Role <= 0 {
		err["role"] = "role is required"
	}

	if len(err) > 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"error":   err,
			"message": "failed created Petugas",
		})
	} else {
		var PetugasModel models.Petugas

		last_petugas := config.DB.Last(&PetugasModel)

		var Nip string

		if last_petugas.Error == nil {
			last_id := PetugasModel.ID
			last_id++

			Nip = fmt.Sprintf("PTGS%04d", last_id)
		} else {
			Nip = "PTGS0001"
		}

		hash, err := models.HashPassword(RequestPetugas.Password)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed Hash Password",
				"error":   err.Error(),
			})

			c.Abort()
			return
		}

		Petugas := models.Petugas{
			Name:     RequestPetugas.Name,
			Username: RequestPetugas.Username,
			Role:     RequestPetugas.Role,
			Password: hash,
			Nip:      Nip,
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

	var RequestPetugas models.RequestPetugas

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
		hash, err := models.HashPassword(RequestPetugas.Password)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed Hash Password",
				"error":   err.Error(),
			})

			c.Abort()
			return
		}

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
