package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Petugas struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username"`
	Nip      string `json:"nip"`
	Password string `json:"password"`
	Role     uint   `json:"role"`
}

type ResponsePetugas struct {
	ID       uint
	Name     string `json:"name"`
	Username string `json:"username"`
	Nip      string `json:"nip"`
	Role     uint   `json:"role"`
}

type RequestPetugas struct {
	ID       uint
	Name     string `json:"name"`
	Username string `json:"username"`
	Nip      string `json:"nip"`
	Password string `json:"password"`
	Role     uint   `json:"role"`
}

type TokenRequest struct {
	Nip      string `json:"nip"`
	Password string `json:"password"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (petugas *Petugas) SetHashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return err
	}

	petugas.Password = string(bytes)

	return nil
}

func (petugas *Petugas) CheckPassword(reqPass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(petugas.Password), []byte(reqPass))

	if err != nil {
		return err
	}

	return nil
}
