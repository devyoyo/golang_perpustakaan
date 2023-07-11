package models

import "gorm.io/gorm"

type PeminjamanDetail struct {
	gorm.Model
	BukuID       uint       `json:"buku_id"`
	PeminjamanID uint       `json:"peminjaman_id"`
	Buku         Buku       `json:"buku"`
	Peminjaman   Peminjaman `json:"peminjaman"`
}

type ResponsePeminjamanDetail struct {
	ID            uint
	JudulBuku     string `json:"judul_buku"`
	KategoriBuku  string `json:"kategori_buku"`
	TahunBuku     string `json:"tahun_buku"`
	StokBuku      uint   `json:"stok_buku"`
	PengarangBuku string `json:"pengarang_buku"`
	PenerbitBuku  string `json:"penerbit_buku"`
}

type RequestPeminjamanDetail struct {
	BukuID uint `json:"buku_id"`
}
