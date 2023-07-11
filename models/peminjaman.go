package models

import "gorm.io/gorm"

type Peminjaman struct {
	gorm.Model
	TanggalPinjam    string             `json:"tanggal_pinjam"`
	TanggalKembali   string             `json:"tanggal_kembali"`
	CodeOrder        string             `json:"code_order"`
	AnggotaID        uint               `json:"anggota_id"`
	PetugasID        uint               `json:"petugas_id"`
	Anggota          Anggota            `json:"anggota"`
	Petugas          Petugas            `json:"petugas"`
	PeminjamanDetail []PeminjamanDetail `json:"detail"`
}

type ResponsePeminjaman struct {
	ID               uint                       `json:"id"`
	CodeOrder        string                     `json:"code_order"`
	TanggalPinjam    string                     `json:"tanggal_pinjam"`
	TanggalKembali   string                     `json:"tanggal_kembali"`
	AnggotaID        uint                       `json:"anggota_id"`
	PetugasID        uint                       `json:"petugas_id"`
	Anggota          Anggota                    `json:"anggota"`
	Petugas          Petugas                    `json:"petugas"`
	PeminjamanDetail []ResponsePeminjamanDetail `json:"detail"`
}

type RequestPeminjaman struct {
	CodeOrder        string                    `json:"code_order"`
	TanggalPinjam    string                    `json:"tanggal_pinjam"`
	AnggotaID        uint                      `json:"anggota_id"`
	PetugasID        uint                      `json:"petugas_id"`
	PeminjamanDetail []RequestPeminjamanDetail `json:"buku_id"`
}
