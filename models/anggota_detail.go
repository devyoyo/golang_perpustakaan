package models

import "gorm.io/gorm"

type AnggotaDetail struct {
	gorm.Model
	ID        uint
	NoTelp    string `json:"no_telp"`
	Alamat    string `json:"alamat"`
	AnggotaID uint   `json:"anggota_id"`
}

type AnggotaDetailResponse struct {
	NoTelp string `json:"no_telp"`
	Alamat string `json:"alamat"`
}
