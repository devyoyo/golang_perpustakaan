package routes

import (
	"fmt"
	"log"
	"net/http"
	"project_perpustakaan/config"
	"project_perpustakaan/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetLoan(c *gin.Context) {
	Peminjaman := []models.Peminjaman{}

	// siapin struct penampungnya
	ResponsePeminjaman := []models.ResponsePeminjaman{}

	// config.DB.Find(&departments)
	data := config.DB.Preload(clause.Associations).Find(&Peminjaman)

	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})

		c.Abort()
		return
	}

	for _, p := range Peminjaman {
		ResponsePeminjamanDetail := []models.ResponsePeminjamanDetail{}

		for _, respPeminjamanDetail := range p.PeminjamanDetail {
			var ModelBuku models.Buku

			config.DB.First(&ModelBuku, "id = ?", respPeminjamanDetail.BukuID)

			ResponsePeminjamanDetail = append(ResponsePeminjamanDetail, models.ResponsePeminjamanDetail{
				JudulBuku:     ModelBuku.Judul,
				KategoriBuku:  ModelBuku.Kategori,
				TahunBuku:     ModelBuku.Tahun,
				PengarangBuku: ModelBuku.Pengarang,
				PenerbitBuku:  ModelBuku.Penerbit,
			})
		}

		data_pinjam := models.ResponsePeminjaman{
			ID:             p.ID,
			TanggalPinjam:  p.TanggalPinjam,
			TanggalKembali: p.TanggalKembali,
			CodeOrder:      p.CodeOrder,
			AnggotaID:      p.AnggotaID,
			PetugasID:      p.PetugasID,
			Anggota: models.ResponseAnggotaPeminjaman{
				Name:      p.Anggota.Name,
				Email:     p.Anggota.Email,
				NoAnggota: p.Anggota.NoAnggota,
			},
			Petugas: models.ResponsePetugasPeminjaman{
				Name:     p.Petugas.Name,
				Username: p.Petugas.Username,
				Nip:      p.Petugas.Nip,
			},
			PeminjamanDetail: ResponsePeminjamanDetail,
		}

		ResponsePeminjaman = append(ResponsePeminjaman, data_pinjam)

	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success get department",
		"data":    ResponsePeminjaman,
	})
}

func GetLoanByAnggota(c *gin.Context) {
	id := c.Param("id")

	Peminjaman := []models.Peminjaman{}

	// siapin struct penampungnya
	ResponsePeminjaman := []models.ResponsePeminjaman{}

	// config.DB.Find(&departments)
	data := config.DB.Preload(clause.Associations).Where("anggota_id = ?", id).Find(&Peminjaman)

	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})

		c.Abort()
		return
	}

	for _, p := range Peminjaman {
		ResponsePeminjamanDetail := []models.ResponsePeminjamanDetail{}

		for _, respPeminjamanDetail := range p.PeminjamanDetail {
			var ModelBuku models.Buku

			config.DB.First(&ModelBuku, "id = ?", respPeminjamanDetail.BukuID)

			ResponsePeminjamanDetail = append(ResponsePeminjamanDetail, models.ResponsePeminjamanDetail{
				JudulBuku:     ModelBuku.Judul,
				KategoriBuku:  ModelBuku.Kategori,
				TahunBuku:     ModelBuku.Tahun,
				PengarangBuku: ModelBuku.Pengarang,
				PenerbitBuku:  ModelBuku.Penerbit,
			})
		}

		data_pinjam := models.ResponsePeminjaman{
			ID:             p.ID,
			TanggalPinjam:  p.TanggalPinjam,
			TanggalKembali: p.TanggalKembali,
			CodeOrder:      p.CodeOrder,
			AnggotaID:      p.AnggotaID,
			PetugasID:      p.PetugasID,
			Anggota: models.ResponseAnggotaPeminjaman{
				Name:      p.Anggota.Name,
				Email:     p.Anggota.Email,
				NoAnggota: p.Anggota.NoAnggota,
			},
			Petugas: models.ResponsePetugasPeminjaman{
				Name:     p.Petugas.Name,
				Username: p.Petugas.Username,
				Nip:      p.Petugas.Nip,
			},
			PeminjamanDetail: ResponsePeminjamanDetail,
		}

		ResponsePeminjaman = append(ResponsePeminjaman, data_pinjam)

	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success get department",
		"data":    ResponsePeminjaman,
	})
}

func GetLoanByPetugas(c *gin.Context) {
	id := c.Param("id")

	Peminjaman := []models.Peminjaman{}

	// siapin struct penampungnya
	ResponsePeminjaman := []models.ResponsePeminjaman{}

	// config.DB.Find(&departments)
	data := config.DB.Preload(clause.Associations).Where("petugas_id = ?", id).Find(&Peminjaman)

	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})

		c.Abort()
		return
	}

	for _, p := range Peminjaman {
		ResponsePeminjamanDetail := []models.ResponsePeminjamanDetail{}

		for _, respPeminjamanDetail := range p.PeminjamanDetail {
			var ModelBuku models.Buku

			config.DB.First(&ModelBuku, "id = ?", respPeminjamanDetail.BukuID)

			ResponsePeminjamanDetail = append(ResponsePeminjamanDetail, models.ResponsePeminjamanDetail{
				JudulBuku:     ModelBuku.Judul,
				KategoriBuku:  ModelBuku.Kategori,
				TahunBuku:     ModelBuku.Tahun,
				PengarangBuku: ModelBuku.Pengarang,
				PenerbitBuku:  ModelBuku.Penerbit,
			})
		}

		data_pinjam := models.ResponsePeminjaman{
			ID:             p.ID,
			TanggalPinjam:  p.TanggalPinjam,
			TanggalKembali: p.TanggalKembali,
			CodeOrder:      p.CodeOrder,
			AnggotaID:      p.AnggotaID,
			PetugasID:      p.PetugasID,
			Anggota: models.ResponseAnggotaPeminjaman{
				Name:      p.Anggota.Name,
				Email:     p.Anggota.Email,
				NoAnggota: p.Anggota.NoAnggota,
			},
			Petugas: models.ResponsePetugasPeminjaman{
				Name:     p.Petugas.Name,
				Username: p.Petugas.Username,
				Nip:      p.Petugas.Nip,
			},
			PeminjamanDetail: ResponsePeminjamanDetail,
		}

		ResponsePeminjaman = append(ResponsePeminjaman, data_pinjam)

	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success get peminjaman",
		"data":    ResponsePeminjaman,
	})
}

func GetLoanByID(c *gin.Context) {
	id := c.Param("id")

	var Peminjaman models.Peminjaman

	data := config.DB.Preload(clause.Associations).Where("id = ?", id).First(&Peminjaman)

	if data.Error != nil {
		log.Printf(data.Error.Error())

		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "data not found",
		})

		c.Abort()
		return
	}

	// siapin struct penampungnya
	ResponsePeminjamanDetail := []models.ResponsePeminjamanDetail{}

	for _, respPeminjamanDetail := range Peminjaman.PeminjamanDetail {
		var ModelBuku models.Buku

		config.DB.First(&ModelBuku, "id = ?", respPeminjamanDetail.BukuID)

		ResponsePeminjamanDetail = append(ResponsePeminjamanDetail, models.ResponsePeminjamanDetail{
			JudulBuku:     ModelBuku.Judul,
			KategoriBuku:  ModelBuku.Kategori,
			TahunBuku:     ModelBuku.Tahun,
			PengarangBuku: ModelBuku.Pengarang,
			PenerbitBuku:  ModelBuku.Penerbit,
		})
	}

	data_pinjam := models.ResponsePeminjaman{
		ID:             Peminjaman.ID,
		TanggalPinjam:  Peminjaman.TanggalPinjam,
		TanggalKembali: Peminjaman.TanggalKembali,
		Anggota: models.ResponseAnggotaPeminjaman{
			Name:      Peminjaman.Anggota.Name,
			Email:     Peminjaman.Anggota.Email,
			NoAnggota: Peminjaman.Anggota.NoAnggota,
		},
		Petugas: models.ResponsePetugasPeminjaman{
			Name:     Peminjaman.Petugas.Name,
			Username: Peminjaman.Petugas.Username,
			Nip:      Peminjaman.Petugas.Nip,
		},
		PeminjamanDetail: ResponsePeminjamanDetail,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success get peminjaman",
		"data":    data_pinjam,
	})
}

func PostLoanByPetugas(c *gin.Context) {
	nipMustGet := c.MustGet("x-nip")

	var Petugas models.Petugas

	data := config.DB.First(&Petugas, "nip = ?", nipMustGet)

	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data petugas not found",
		})

		c.Abort()
		return
	}

	var RequestPeminjaman models.RequestPeminjaman

	c.BindJSON(&RequestPeminjaman)

	RequestPeminjaman.PetugasID = Petugas.ID

	err := map[string]string{}

	if len(RequestPeminjaman.PeminjamanDetail) <= 0 {
		err["detail"] = "buku is required"
	}

	if len(err) > 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"error":   err,
			"message": "failed created pinjam buku",
		})

		c.Abort()
		return
	} else {
		t := time.Now().Local()

		var RequestPeminjamanDetail = []models.RequestPeminjamanDetail{}

		var buku_id []uint

		for _, data := range RequestPeminjaman.PeminjamanDetail {
			reqbuku := models.RequestPeminjamanDetail{
				BukuID: data.BukuID,
			}

			buku_id = append(buku_id, data.BukuID)

			var Buku models.Buku

			dataBuku := config.DB.Where("id = ?", data.BukuID).First(&Buku)

			if dataBuku.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "buku id -" + strconv.FormatInt(int64(data.BukuID), 10) + "- not found",
					"status":  http.StatusInternalServerError,
				})

				c.Abort()

				return
			} else {
				if Buku.Stok > 0 {
					RequestPeminjamanDetail = append(RequestPeminjamanDetail, reqbuku)
				} else {
					c.JSON(http.StatusInternalServerError, gin.H{
						"message": "stok buku id -" + strconv.FormatInt(int64(data.BukuID), 10) + "- are empty",
						"status":  http.StatusInternalServerError,
					})

					c.Abort()

					return
				}
			}

		}

		var Peminjaman models.Peminjaman

		last_peminjaman := config.DB.Last(&Peminjaman)

		var code_order string
		var last_id int

		if last_peminjaman.Error == nil {
			last_id := Peminjaman.ID
			last_id++
		} else {
			last_id = 1
		}

		code_order = fmt.Sprintf("LN%04d", last_id)

		pinjam := models.Peminjaman{
			TanggalPinjam: t.String(),
			AnggotaID:     RequestPeminjaman.AnggotaID,
			PetugasID:     RequestPeminjaman.PetugasID,
			CodeOrder:     code_order,
		}

		insert := config.DB.Create(&pinjam)

		if insert.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": insert.Error.Error(),
			})

			c.Abort()
			return
		} else {
			for _, data := range RequestPeminjaman.PeminjamanDetail {
				reqbuku := models.PeminjamanDetail{
					BukuID:       data.BukuID,
					PeminjamanID: pinjam.ID,
				}
				config.DB.Create(&reqbuku)
			}

			CallBuku := []models.Buku{}

			config.DB.Where("id IN ?", buku_id).Find(&CallBuku)

			var total_stok uint

			for _, data_bukus := range CallBuku {
				total_stok = data_bukus.Stok - 1

				updateData := config.DB.Model(&models.Buku{}).Where("id = ?", data_bukus.ID).Update("stok", total_stok)

				if updateData.Error != nil {
					log.Println(updateData.Error.Error())
				}
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "success create peminjaman",
			"data":    RequestPeminjaman,
		})
	}
}

func BackLoan(c *gin.Context) {
	id := c.Param("id")

	var Peminjaman models.Peminjaman

	data := config.DB.Preload(clause.Associations).Where("id = ?", id).First(&Peminjaman)

	if data.Error != nil {
		log.Printf(data.Error.Error())

		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "data not found",
		})

		c.Abort()

		return
	}

	updateData := models.Peminjaman{
		TanggalKembali: time.Now().Local().String(),
	}

	config.DB.Model(&Peminjaman).Where("id = ?", id).Updates(&updateData)

	var total_stok uint

	for _, data := range Peminjaman.PeminjamanDetail {
		var Buku models.Buku

		config.DB.First(&Buku, "id = ?", data.BukuID)

		total_stok = Buku.Stok + 1

		log.Println(data.BukuID, ":", total_stok, " : ", Buku.Stok)

		updateData := config.DB.Model(&models.Buku{}).Where("id = ?", data.BukuID).Update("stok", total_stok)

		if updateData.Error != nil {
			log.Println(updateData.Error.Error())
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success pengembalian buku",
		"data":    Peminjaman,
	})
}

func GetLoanByBuku(c *gin.Context) {
	id := c.Param("id")

	var Buku models.Buku

	dataBuku := config.DB.First(&Buku, "id = ?", id)

	if dataBuku.Error != nil {
		log.Printf(dataBuku.Error.Error())

		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "data not found",
		})

		c.Abort()

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

	PeminjamanDetail := []models.PeminjamanDetail{}

	config.DB.Preload(clause.Associations).Find(&PeminjamanDetail, "buku_id = ?", id)

	ResponsePeminjamanByBuku := []models.ResponsePeminjamanByBuku{}

	for _, data := range PeminjamanDetail {
		var Anggota models.Anggota
		var Petugas models.Petugas

		config.DB.First(&Anggota, "id = ?", data.Peminjaman.AnggotaID)
		config.DB.First(&Petugas, "id = ?", data.Peminjaman.PetugasID)

		resp := models.ResponsePeminjamanByBuku{
			CodeOrder:      data.Peminjaman.CodeOrder,
			TanggalPinjam:  data.Peminjaman.TanggalPinjam,
			TanggalKembali: data.Peminjaman.TanggalKembali,
			AnggotaID:      data.Peminjaman.AnggotaID,
			PetugasID:      data.Peminjaman.PetugasID,
			Anggota: models.ResponseAnggotaPeminjaman{
				Name:      Anggota.Name,
				Email:     Anggota.Email,
				NoAnggota: Anggota.NoAnggota,
			},
			Petugas: models.ResponsePetugasPeminjaman{
				Name:     Petugas.Name,
				Username: Petugas.Username,
				Nip:      Petugas.Nip,
			},
		}

		ResponsePeminjamanByBuku = append(ResponsePeminjamanByBuku, resp)
	}

	config.DB.Find(&PeminjamanDetail, "buku_id = ?", id)

	c.JSON(http.StatusOK, gin.H{
		"message":   "success pengembalian buku",
		"data_buku": ResponseBuku,
		"data":      ResponsePeminjamanByBuku,
	})
}
