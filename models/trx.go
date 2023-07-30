package models

import "time"

type Trx struct {
	ID               int       `gorm:"primaryKey" json:"id"`
	IDUser           int       `gorm:"foreignKey:IDUser,references:User" json:"id_user"`
	AlamatPengiriman int       `gorm:"foreignKey:AlamatPengiriman,references:Alamat" json:"alamat_pengiriman"`
	HargaTotal       int       `json:"harga_total"`
	KodeInvoice      string    `json:"kode_invoice"`
	MethodBayar      string    `json:"method_bayar"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	DetailTrx DetailTrx 		`gorm:"foreignKey:IDTrx"`
}