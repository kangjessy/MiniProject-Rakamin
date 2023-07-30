package models

import "time"

type Toko struct {
	ID        int 			`gorm:"primaryKey" json:"id"`
	IDUser    int 			`gorm:"foreignKey:IDUser,references:User" json:"id_user"`
	NamaToko  string 		`json:"nama_toko" binding:"required"`
	UrlFoto   string 		`json:"url_foto" binding:"required"`
	CreatedAt time.Time 	`json:"created_at"`
	UpdatedAt time.Time 	`json:"updated_at"`
	DetailTrx []DetailTrx 	`gorm:"foreignKey:IDToko"`
	Produk []Produk 		`gorm:"foreignKey:IDToko"`
	LogProduk []LogProduk 	`gorm:"foreignKey:IDToko"`
}