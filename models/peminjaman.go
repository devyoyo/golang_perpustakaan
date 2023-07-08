package models

import "gorm.io/gorm"

type Peminjaman struct {
	gorm.Model
	TanggalPinjam  string  `json:"tanggal_pinjam"`
	TanggalKembali string  `json:"tanggal_kembali"`
	AnggotaID      uint    `json:"anggota_id"`
	PetugasID      uint    `json:"petugas_id"`
	Anggota        Anggota `json:"anggota"`
	Petugas        Petugas `json:"petugas"`
}
