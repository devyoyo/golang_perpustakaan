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
	config.DB.Preload(clause.Associations).Find(&Peminjaman)

	for _, p := range Peminjaman {
		ResponsePeminjamanDetail := []models.ResponsePeminjamanDetail{}

		for _, respPeminjamanDetail := range p.PeminjamanDetail {
			ResponsePeminjamanDetail = append(ResponsePeminjamanDetail, models.ResponsePeminjamanDetail{
				JudulBuku:     respPeminjamanDetail.Buku.Judul,
				KategoriBuku:  respPeminjamanDetail.Buku.Kategori,
				TahunBuku:     respPeminjamanDetail.Buku.Tahun,
				StokBuku:      respPeminjamanDetail.Buku.Stok,
				PengarangBuku: respPeminjamanDetail.Buku.Pengarang,
				PenerbitBuku:  respPeminjamanDetail.Buku.Penerbit,
			})
		}

		data_pinjam := models.ResponsePeminjaman{
			ID:             p.ID,
			TanggalPinjam:  p.TanggalPinjam,
			TanggalKembali: p.TanggalKembali,
			Anggota: models.Anggota{
				Name:      p.Anggota.Name,
				Email:     p.Anggota.Email,
				NoAnggota: p.Anggota.NoAnggota,
			},
			Petugas: models.Petugas{
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

		return
	}

	for _, p := range Peminjaman {
		ResponsePeminjamanDetail := []models.ResponsePeminjamanDetail{}

		for _, respPeminjamanDetail := range p.PeminjamanDetail {
			ResponsePeminjamanDetail = append(ResponsePeminjamanDetail, models.ResponsePeminjamanDetail{
				JudulBuku:     respPeminjamanDetail.Buku.Judul,
				KategoriBuku:  respPeminjamanDetail.Buku.Kategori,
				TahunBuku:     respPeminjamanDetail.Buku.Tahun,
				StokBuku:      respPeminjamanDetail.Buku.Stok,
				PengarangBuku: respPeminjamanDetail.Buku.Pengarang,
				PenerbitBuku:  respPeminjamanDetail.Buku.Penerbit,
			})
		}

		data_pinjam := models.ResponsePeminjaman{
			ID:             p.ID,
			TanggalPinjam:  p.TanggalPinjam,
			TanggalKembali: p.TanggalKembali,
			Anggota: models.Anggota{
				Name:      p.Anggota.Name,
				Email:     p.Anggota.Email,
				NoAnggota: p.Anggota.NoAnggota,
			},
			Petugas: models.Petugas{
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

		return
	}

	for _, p := range Peminjaman {
		ResponsePeminjamanDetail := []models.ResponsePeminjamanDetail{}

		for _, respPeminjamanDetail := range p.PeminjamanDetail {
			ResponsePeminjamanDetail = append(ResponsePeminjamanDetail, models.ResponsePeminjamanDetail{
				JudulBuku:     respPeminjamanDetail.Buku.Judul,
				KategoriBuku:  respPeminjamanDetail.Buku.Kategori,
				TahunBuku:     respPeminjamanDetail.Buku.Tahun,
				StokBuku:      respPeminjamanDetail.Buku.Stok,
				PengarangBuku: respPeminjamanDetail.Buku.Pengarang,
				PenerbitBuku:  respPeminjamanDetail.Buku.Penerbit,
			})
		}

		data_pinjam := models.ResponsePeminjaman{
			ID:             p.ID,
			TanggalPinjam:  p.TanggalPinjam,
			TanggalKembali: p.TanggalKembali,
			Anggota: models.Anggota{
				Name:      p.Anggota.Name,
				Email:     p.Anggota.Email,
				NoAnggota: p.Anggota.NoAnggota,
			},
			Petugas: models.Petugas{
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

		return
	}

	// siapin struct penampungnya
	ResponsePeminjamanDetail := []models.ResponsePeminjamanDetail{}

	for _, respPeminjamanDetail := range Peminjaman.PeminjamanDetail {
		ResponsePeminjamanDetail = append(ResponsePeminjamanDetail, models.ResponsePeminjamanDetail{
			JudulBuku:     respPeminjamanDetail.Buku.Judul,
			KategoriBuku:  respPeminjamanDetail.Buku.Kategori,
			TahunBuku:     respPeminjamanDetail.Buku.Tahun,
			StokBuku:      respPeminjamanDetail.Buku.Stok,
			PengarangBuku: respPeminjamanDetail.Buku.Pengarang,
			PenerbitBuku:  respPeminjamanDetail.Buku.Penerbit,
		})
	}

	data_pinjam := models.ResponsePeminjaman{
		ID:             Peminjaman.ID,
		TanggalPinjam:  Peminjaman.TanggalPinjam,
		TanggalKembali: Peminjaman.TanggalKembali,
		Anggota: models.Anggota{
			Name:      Peminjaman.Anggota.Name,
			Email:     Peminjaman.Anggota.Email,
			NoAnggota: Peminjaman.Anggota.NoAnggota,
		},
		Petugas: models.Petugas{
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
	var RequestPeminjaman models.RequestPeminjaman

	c.BindJSON(&RequestPeminjaman)

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

		pinjam := models.RequestPeminjaman{
			TanggalPinjam:    t.String(),
			AnggotaID:        RequestPeminjaman.AnggotaID,
			PetugasID:        RequestPeminjaman.PetugasID,
			CodeOrder:        code_order,
			PeminjamanDetail: RequestPeminjamanDetail,
		}

		insert := config.DB.Create(pinjam)

		if insert.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": insert.Error.Error(),
			})

			c.Abort()
			return
		} else {
			CallBuku := []models.Buku{}

			config.DB.Where("id IN ?", buku_id).Find(&CallBuku)

			for _, data_bukus := range CallBuku {
				total_stok := data_bukus.Stok - 1

				var UpdateBuku models.Buku

				update := models.Buku{
					Stok: total_stok,
				}

				config.DB.Model(&UpdateBuku).Where("id = ?", data_bukus.ID).Updates(&update)
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

	for _, data := range Peminjaman.PeminjamanDetail {
		total_stok := data.Buku.Stok + 1

		var UpdateBuku models.Buku

		update := models.Buku{
			Stok: total_stok,
		}

		config.DB.Model(&UpdateBuku).Where("id = ?", data.BukuID).Updates(&update)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success pengembalian buku",
		"data":    Peminjaman,
	})
}
