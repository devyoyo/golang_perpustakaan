package config

import (
	"log"
	"net/http"
	"project_perpustakaan/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

const DbConnect = "root:@tcp(127.0.0.1:3306)/golang_perpustakaan?charset=utf8mb4&parseTime=True&loc=Local"

func InitDB() {
	var err error

	DB, err = gorm.Open(mysql.Open(DbConnect), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Println("error")
	}

}

func Migrate(c *gin.Context) {
	DB.AutoMigrate(
		&models.Anggota{},
		&models.AnggotaDetail{},
		&models.Peminjaman{},
		&models.PeminjamanDetail{},
		&models.Buku{},
		&models.Petugas{})

	var Buku = []models.Buku{
		{
			Judul:     "Akuntansi Pengantar 1",
			Kategori:  "Pendidikan",
			Tahun:     "2009",
			Stok:      2,
			Pengarang: "Supardi",
			Penerbit:  "Gava Media",
		},
		{
			Judul:     "Aplikasi Klinis Induk Ovulasi & Stimulasi Ovariu",
			Kategori:  "Pendidikan",
			Tahun:     "2013",
			Stok:      2,
			Pengarang: "Samsulhadi",
			Penerbit:  "Sagung Seto",
		},
		{
			Judul:     "Aplikasi Praktis Asuhan Keperawatan Keluarga",
			Kategori:  "Pendidikan",
			Tahun:     "2012",
			Stok:      2,
			Pengarang: "Komang Ayu Heni",
			Penerbit:  "Sagung Seto",
		},
		{
			Judul:     "A-Z Psikologi : Berbagai kumpulan topik Psikologi",
			Kategori:  "Pendidikan",
			Tahun:     "2009",
			Stok:      2,
			Pengarang: "Zainul Anwar",
			Penerbit:  "Andi Offset",
		},
		{
			Judul:     "Bangsa gagal ; Mencari identitas kebangsaan",
			Kategori:  "Politik",
			Tahun:     "2008",
			Stok:      2,
			Pengarang: "Nasruddin Anshoriy",
			Penerbit:  "LKiS",
		},
		{
			Judul:     "Biografi Gus Dur ; The Authorized Biography of KH. Abdurrahman Wahid (Soft Cover)",
			Kategori:  "Politik",
			Tahun:     "2011",
			Stok:      2,
			Pengarang: "Greg Barton",
			Penerbit:  "LKiS",
		},
	}

	hash_1, _ := models.HashPassword("iksan_password")
	hash_2, _ := models.HashPassword("rudi_password")

	var Petugas = []models.Petugas{
		{
			Name:     "iksan",
			Username: "user_iksan",
			Password: hash_1,
			Nip:      "PTGS0001",
			Role:     1,
		},
		{
			Name:     "Rudi",
			Username: "user_rudi",
			Password: hash_2,
			Nip:      "PTGS0002",
			Role:     2,
		},
	}

	var Anggota = []models.Anggota{
		{
			Name:      "Suryono",
			Email:     "suryono@gmail.com",
			NoAnggota: "M0001",
			AnggotaDetail: models.AnggotaDetail{
				NoTelp: "081201462110",
				Alamat: "jl komando raya",
			},
		},
		{
			Name:      "Jono",
			Email:     "jono@gmail.com",
			NoAnggota: "M0002",
			AnggotaDetail: models.AnggotaDetail{
				NoTelp: "085697282128",
				Alamat: "jl raya mauk",
			},
		},
	}

	DB.Create(Buku)
	DB.Create(Petugas)
	DB.Create(Anggota)

	c.JSON(http.StatusOK, gin.H{
		"message": "success migrate database",
	})
}
