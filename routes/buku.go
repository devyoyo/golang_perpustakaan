package routes

import (
	"log"
	"net/http"
	"project_perpustakaan/config"
	"project_perpustakaan/models"

	"github.com/gin-gonic/gin"
)

func GetBuku(c *gin.Context) {
	Bukus := []models.Buku{}

	config.DB.Find(&Bukus)

	ResponseBuku := []models.ResponseBuku{}

	for _, Buku := range Bukus {
		data := models.ResponseBuku{
			ID:        Buku.ID,
			Judul:     Buku.Judul,
			Kategori:  Buku.Kategori,
			Tahun:     Buku.Tahun,
			Stok:      Buku.Stok,
			Pengarang: Buku.Pengarang,
			Penerbit:  Buku.Penerbit,
		}

		ResponseBuku = append(ResponseBuku, data)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success get Buku",
		"data":    ResponseBuku,
	})

}

func GetBukuById(c *gin.Context) {
	id := c.Param("id")

	var Buku models.Buku

	data := config.DB.First(&Buku, "id = ?", id)

	if data.Error != nil {
		log.Printf(data.Error.Error())

		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "data not found",
		})

		return
	}

	ResponseBuku := models.ResponseBuku{
		ID:        Buku.ID,
		Judul:     Buku.Judul,
		Kategori:  Buku.Kategori,
		Tahun:     Buku.Tahun,
		Stok:      Buku.Stok,
		Pengarang: Buku.Pengarang,
		Penerbit:  Buku.Penerbit,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success get Buku",
		"data":    ResponseBuku,
	})
}

func PostBuku(c *gin.Context) {
	var RequestBuku models.Buku

	c.BindJSON(&RequestBuku)

	err := map[string]string{}

	if RequestBuku.Judul == "" {
		err["judul"] = "judul is required"
	}
	if RequestBuku.Kategori == "" {
		err["kategori"] = "kategori is required"
	}
	if RequestBuku.Tahun == "" {
		err["tahun"] = "tahun is required"
	}
	if RequestBuku.Stok <= 0 {
		err["stok"] = "stok must be greater than zero"
	}
	if RequestBuku.Pengarang == "" {
		err["pengarang"] = "pengarang is required"
	}
	if RequestBuku.Penerbit == "" {
		err["penerbit"] = "penerbit is required"
	}

	if len(err) > 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"error":   err,
			"message": "failed created Buku",
		})

		c.Abort()
		return
	} else {
		Buku := models.Buku{
			Judul:     RequestBuku.Judul,
			Kategori:  RequestBuku.Kategori,
			Tahun:     RequestBuku.Tahun,
			Stok:      RequestBuku.Stok,
			Pengarang: RequestBuku.Pengarang,
			Penerbit:  RequestBuku.Penerbit,
		}

		config.DB.Create(&Buku)

		c.JSON(http.StatusCreated, gin.H{
			"data":    Buku,
			"message": "success created Buku",
		})
	}

}

func PutBuku(c *gin.Context) {
	id := c.Param("id")

	var Buku models.Buku

	data := config.DB.First(&Buku, "id = ?", id)

	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})

		return
	}

	var RequestBuku models.Buku

	c.BindJSON(&RequestBuku)

	err := map[string]string{}

	if RequestBuku.Judul == "" {
		err["judul"] = "judul is required"
	}
	if RequestBuku.Kategori == "" {
		err["kategori"] = "kategori is required"
	}
	if RequestBuku.Tahun == "" {
		err["tahun"] = "tahun is required"
	}
	if RequestBuku.Stok <= 0 {
		err["stok"] = "stok must be greater than zero"
	}
	if RequestBuku.Pengarang == "" {
		err["pengarang"] = "pengarang is required"
	}
	if RequestBuku.Penerbit == "" {
		err["penerbit"] = "penerbit is required"
	}

	if len(err) > 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"error":   err,
			"message": "failed created Buku",
		})
	} else {
		update_Buku := models.Buku{
			Judul:     RequestBuku.Judul,
			Kategori:  RequestBuku.Kategori,
			Tahun:     RequestBuku.Tahun,
			Stok:      RequestBuku.Stok,
			Pengarang: RequestBuku.Pengarang,
			Penerbit:  RequestBuku.Penerbit,
		}

		update := config.DB.Model(&Buku).Where("id = ?", id).Updates(&update_Buku)

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
			"data":    RequestBuku,
		})
	}

}

func DeleteBuku(c *gin.Context) {
	id := c.Param("id")

	var Buku models.Buku

	data := config.DB.First(&Buku, "id = ?", id)

	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})

		return
	}

	config.DB.Delete(&Buku, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "delete success",
	})
}
