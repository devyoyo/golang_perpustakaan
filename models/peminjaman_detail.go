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
	JudulBuku     string `json:"judul_buku"`
	KategoriBuku  string `json:"kategori_buku"`
	TahunBuku     string `json:"tahun_buku"`
	PengarangBuku string `json:"pengarang_buku"`
	PenerbitBuku  string `json:"penerbit_buku"`
}

type RequestPeminjamanDetail struct {
	BukuID       uint `json:"buku_id"`
	PeminjamanID uint `json:"peminjaman_id"`
}

type PeminjamanDetailByBukuResponse struct {
	Peminjaman ResponsePeminjamanByBuku `json:"peminjaman"`
}
