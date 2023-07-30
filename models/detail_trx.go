package models

import "time"

type DetailTrx struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	IDTrx       int       `gorm:"foreignKey:IDTrx,references:Trx" json:"id_trx"`
	IDToko      int       `gorm:"foreignKey:IDUser,references:Toko" json:"id_toko"`
	Kuantitas   int       `json:"kuantitas"`
	HargaTotal  int       `json:"harga_total"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	LogProduk []LogProduk  `gorm:"foreignKey:IDDetailTrx"`
}