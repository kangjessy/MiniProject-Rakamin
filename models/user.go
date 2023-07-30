package models

import (
	"time"
)

type User struct {
	ID           int    `gorm:"primaryKey" json:"id"`
	Nama         string `json:"nama" binding:"required"`
	Email        string `json:"email" binding:"required"`
	KataSandi    string `json:"katasandi" binding:"required"`
	Notelp       string `json:"notelp" binding:"required"`
	TanggalLahir string `json:"tanggal_lahir" binding:"required"`
	JenisKelamin string `json:"jenis_kelamin" binding:"required"`
	Tentang      string `json:"tentang" binding:"required"`
	Pekerjaan    string `json:"pekerjaan" binding:"required"`
	IDProvinsi   string `json:"id_provinsi" binding:"required"`
	IDKota       string `json:"id_kota" binding:"required"`
	IsAdmin      bool 
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Toko Toko `gorm:"foreignKey:IDUser"`
	Trx Trx `gorm:"foreignKey:IDUser"`
	Alamat []Alamat `gorm:"foreignKey:IDUser"`
}
