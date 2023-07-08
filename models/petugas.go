package models

import "gorm.io/gorm"

type Petugas struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponsePetugas struct {
	ID       uint
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type RequestPetugas struct {
	ID       uint
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
