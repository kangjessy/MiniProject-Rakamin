package models

import "time"

type LogProduk struct {
	ID            int       `gorm:"primaryKey" json:"id"`
	IDProduk      int       `gorm:"foreignKey:IDProduk,references:Produk" json:"id_produk"`
	IDToko        int       `gorm:"foreignKey:IDToko,references:Toko" json:"id_toko"`
	IDDetailTrx   int       `gorm:"foreignKey:IDDetailTrx,references:LogProduk" json:"id_log_produk"`
	IDCategory    int       `gorm:"foreignKey:IDDetailTrx,references:IDCategory" json:"id_category"`
	NamaProduk    string    `json:"nama_produk"`
	Slug          string    `json:"slug"`
	HargaReseller string    `json:"harga_reseller"`
	HargaKonsumen string    `json:"harga_konsumen"`
	Deskripsi     string    `json:"deskripsi"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	FotoProduk []FotoProduk `gorm:"foreignKey:IDFotoProduk"`
}