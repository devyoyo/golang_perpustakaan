package models

import "gorm.io/gorm"

type Anggota struct {
	gorm.Model
	ID            uint
	Name          string        `json:"name"`
	Email         string        `json:"email"`
	NoAnggota     string        `json:"no_anggota"`
	AnggotaDetail AnggotaDetail `json:"detail"`
}

type ResponseAnggota struct {
	ID                    uint                  `json:"id"`
	Name                  string                `json:"name"`
	Email                 string                `json:"email"`
	NoAnggota             string                `json:"no_anggota"`
	AnggotaDetailResponse AnggotaDetailResponse `json:"detail"`
}

type ResponseAnggotaPeminjaman struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	NoAnggota string `json:"no_anggota"`
}

type RequestAnggota struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	NoTelp string `json:"no_telp"`
	Alamat string `json:"alamat"`
}
