package routes

import (
	"fmt"
	"log"
	"net/http"
	"project_perpustakaan/config"
	"project_perpustakaan/models"

	"github.com/gin-gonic/gin"
)

func GetAnggota(c *gin.Context) {
	anggotas := []models.Anggota{}

	config.DB.Preload("AnggotaDetail").Find(&anggotas)

	ResponseAnggota := []models.ResponseAnggota{}

	for _, anggota := range anggotas {
		detail := anggota.AnggotaDetail

		data := models.ResponseAnggota{
			ID:        anggota.ID,
			Name:      anggota.Name,
			Email:     anggota.Email,
			NoAnggota: anggota.NoAnggota,
			AnggotaDetailResponse: models.AnggotaDetailResponse{
				NoTelp: detail.NoTelp,
				Alamat: detail.Alamat,
			},
		}

		ResponseAnggota = append(ResponseAnggota, data)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success get anggota",
		"data":    ResponseAnggota,
	})

}

func GetAnggotaById(c *gin.Context) {
	id := c.Param("id")

	var anggota models.Anggota

	data := config.DB.Preload("AnggotaDetail").First(&anggota, "id = ?", id)

	if data.Error != nil {
		log.Printf(data.Error.Error())

		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "data not found",
		})

		return
	}

	ResponseAnggota := models.ResponseAnggota{
		ID:        anggota.ID,
		Name:      anggota.Name,
		Email:     anggota.Email,
		NoAnggota: anggota.NoAnggota,
		AnggotaDetailResponse: models.AnggotaDetailResponse{
			NoTelp: anggota.AnggotaDetail.NoTelp,
			Alamat: anggota.AnggotaDetail.Alamat,
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success get anggota",
		"data":    ResponseAnggota,
	})
}

func PostAnggota(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	no_telp := c.PostForm("no_telp")
	alamat := c.PostForm("alamat")

	err := map[string]string{}

	if name == "" {
		err["name"] = "name is required"
	}
	if email == "" {
		err["email"] = "email is required"
	}
	if no_telp == "" {
		err["no_telp"] = "no telepon is required"
	}
	if alamat == "" {
		err["alamat"] = "alamat is required"
	}

	if len(err) > 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"error":   err,
			"message": "failed created anggota",
		})
	} else {
		var AnggotModel models.Anggota

		last_anggota := config.DB.Last(&AnggotModel)

		var no_anggota string

		if last_anggota.Error == nil {
			last_id := AnggotModel.ID
			last_id++

			no_anggota = fmt.Sprintf("M%04d", last_id)
		} else {
			no_anggota = "M0001"
		}

		anggota := models.Anggota{
			Name:      name,
			Email:     email,
			NoAnggota: no_anggota,
			AnggotaDetail: models.AnggotaDetail{
				NoTelp: no_telp,
				Alamat: alamat,
			},
		}

		config.DB.Create(&anggota)

		c.JSON(http.StatusCreated, gin.H{
			"data":    anggota,
			"message": "success created anggota",
		})
	}

}

func PutAnggota(c *gin.Context) {
	id := c.Param("id")

	var anggota models.Anggota

	data := config.DB.First(&anggota, "id = ?", id)

	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})

		return
	}

	var RequestAnggota models.RequestAnggota

	c.BindJSON(&RequestAnggota)

	update_anggota := models.Anggota{
		ID:    anggota.ID,
		Name:  RequestAnggota.Name,
		Email: RequestAnggota.Email,
		AnggotaDetail: models.AnggotaDetail{
			NoTelp: RequestAnggota.NoTelp,
			Alamat: RequestAnggota.Alamat,
		},
	}

	var AnggotaDetail models.AnggotaDetail

	config.DB.First(&AnggotaDetail, "anggota_id = ?", id)

	update := config.DB.Save(&update_anggota)

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
		"data":    RequestAnggota,
	})
}

func DeleteAnggota(c *gin.Context) {
	id := c.Param("id")

	var anggota models.Anggota

	data := config.DB.First(&anggota, "id = ?", id)

	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})

		return
	}

	config.DB.Delete(&anggota, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "delete success",
	})
}
