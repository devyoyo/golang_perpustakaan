package main

import (
	"project_perpustakaan/config"
	"project_perpustakaan/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	r := gin.Default()

	api := r.Group("api")
	{
		anggota := api.Group("anggota")
		{
			anggota.GET("/", routes.GetAnggota)
			anggota.GET("/:id", routes.GetAnggotaById)
			anggota.POST("/", routes.PostAnggota)
			anggota.PUT("/:id", routes.PutAnggota)
			anggota.DELETE("/:id", routes.DeleteAnggota)
		}

		Buku := api.Group("buku")
		{
			Buku.GET("/", routes.GetBuku)
			Buku.GET("/:id", routes.GetBukuById)
			Buku.POST("/", routes.PostBuku)
			Buku.PUT("/:id", routes.PutBuku)
			Buku.DELETE("/:id", routes.DeleteBuku)
		}

		Petugas := api.Group("petugas")
		{
			Petugas.GET("/", routes.GetPetugas)
			Petugas.GET("/:id", routes.GetPetugasById)
			Petugas.POST("/", routes.PostPetugas)
			Petugas.PUT("/:id", routes.PutPetugas)
			Petugas.DELETE("/:id", routes.DeletePetugas)
		}
	}

	r.GET("/migrate", config.Migrate)

	r.Run(":8081")
}
