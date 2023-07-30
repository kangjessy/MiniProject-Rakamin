package models

import "time"

type Alamat struct {
	ID           int       `gorm:"primaryKey" json:"id"`
	IDUser       int       `gorm:"foreignKey:IDUser,references:User" json:"id_user"`
	JudulAlamat  string    `json:"judul_alamat"`
	NamaPenerima string    `json:"nama_penerima"`
	NoTelp       string    `json:"notelp"`
	DetailAlamat string    `json:"detail_alamat"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Trx Trx 				`gorm:"foreignKey:AlamatPengiriman"`
}