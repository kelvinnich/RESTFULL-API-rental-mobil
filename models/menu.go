package models

type Menu struct{
	ID uint64 `gorm:"primary_key:auto_increment" json:"id"`
	NamaMobil string `gorm:"type:varchar(255)" json:"nama_mobil"`
	Harga uint64 `gorm:"type:int" json :"harga"`
	TipeMobil string `gorm:"type:varchar(255)" json:"tipe_mobil"`
	Status string `gorm:"type:varchar(20)" json:"status"`
}