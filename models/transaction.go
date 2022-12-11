package models

import "time"

type Transaction struct{
	ID uint64 `gorm:"primary_key:auto_increment" json:"id"`
	UserID uint64 `gorm:"not null" json:"-"`
	User User `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"users"`
	MenuID uint64 `gorm:"not null" json:"-"`
	Menu Menu `gorm:"foreignKey:MenuID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"menu"`
	Tanggal time.Time
	DurasaRental uint64 `gorm:"type:int" json:"durasi_pinjam"`
	TotalPembayaran uint64 `gorm:"type:int" json:"total_pembayaran"`
}