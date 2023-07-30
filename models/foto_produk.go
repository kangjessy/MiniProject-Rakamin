package models

import "time"

type FotoProduk struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	IDProduk  int       `gorm:"foreignKey:IDProduk,references:Produk" json:"id_produk"`
	IDFotoProduk  int   `gorm:"foreignKey:IDFotoProduk,references:Produk" json:"id_foto_produk"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}