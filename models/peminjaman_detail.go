package models

import "gorm.io/gorm"

type PeminjamanDetail struct {
	gorm.Model
	BukuID       uint       `json:"buku_id"`
	PeminjamanID uint       `json:"peminjaman_id"`
	Buku         Buku       `json:"buku"`
	Peminjaman   Peminjaman `json:"peminjaman"`
}
