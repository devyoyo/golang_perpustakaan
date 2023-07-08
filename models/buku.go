package models

import "gorm.io/gorm"

type Buku struct {
	gorm.Model
	Judul     string `json:"judul"`
	Kategori  string `json:"kategori"`
	Tahun     string `json:"tahun"`
	Stok      uint   `json:"stok"`
	Pengarang string `json:"pengarang"`
	Penerbit  string `json:"penerbit"`
}

type ResponseBuku struct {
	ID        uint
	Judul     string `json:"judul"`
	Kategori  string `json:"kategori"`
	Tahun     string `json:"tahun"`
	Stok      uint   `json:"stok"`
	Pengarang string `json:"pengarang"`
	Penerbit  string `json:"penerbit"`
}

type RequestBuku struct {
	Judul     string `json:"judul"`
	Kategori  string `json:"kategori"`
	Tahun     string `json:"tahun"`
	Stok      uint   `json:"stok"`
	Pengarang string `json:"pengarang"`
	Penerbit  string `json:"penerbit"`
}
