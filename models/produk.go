package models

import "time"

type Produk struct {
	ID            int       `gorm:"primaryKey" json:"id"`
	IDToko        int       `gorm:"foreignKey:IDToko,references:Toko" json:"id_toko"`
	IDCategory    int       `gorm:"foreignKey:IDCategory,references:Category" json:"id_category"`
	NamaProduk    string    `json:"nama_produk"`
	Slug          string    `json:"slug"`
	HargaReseller string    `json:"harga_reseller"`
	HargaKonsumen string    `json:"harga_konsumen"`
	Stok          int       `json:"stok"`
	Deskripsi     string    `json:"deskripsi"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	LogProduk 	  LogProduk	`gorm:"foreignKey:IDProduk"`
	FotoProduk []FotoProduk `gorm:"foreignKey:IDProduk"`
}