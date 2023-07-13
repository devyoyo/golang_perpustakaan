package main

import (
	"project_perpustakaan/config"
	"project_perpustakaan/middleware"
	"project_perpustakaan/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	r := gin.Default()

	api := r.Group("api")
	{
		anggota := api.Group("anggota").Use(middleware.Auth())
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
		}

		BukuAuth := api.Group("buku").Use(middleware.Auth())
		{
			BukuAuth.POST("/", routes.PostBuku)
			BukuAuth.PUT("/:id", routes.PutBuku)
			BukuAuth.DELETE("/:id", routes.DeleteBuku)
		}

		Petugas := api.Group("petugas").Use(middleware.Auth())
		{
			Petugas.GET("/", routes.GetPetugas)
			Petugas.GET("/:id", routes.GetPetugasById)
			Petugas.POST("/", routes.PostPetugas)
			Petugas.PUT("/:id", routes.PutPetugas)
			Petugas.DELETE("/:id", routes.DeletePetugas)
		}

		loaning := api.Group("peminjaman").Use(middleware.Auth())
		{
			loaning.GET("/", routes.GetLoan)
			loaning.GET("/anggota/:id", routes.GetLoanByAnggota)
			loaning.GET("/petugas/:id", routes.GetLoanByPetugas)
			loaning.GET("/buku/:id", routes.GetLoanByBuku)
			loaning.GET("/:id", routes.GetLoanByID)
			loaning.POST("/", routes.PostLoanByPetugas)
			loaning.GET("/back/:id", routes.BackLoan)
		}

		api.POST("/generate_token", routes.GenerateToken)
		api.GET("/migrate", config.Migrate)
	}

	r.Run(":8081")
}
