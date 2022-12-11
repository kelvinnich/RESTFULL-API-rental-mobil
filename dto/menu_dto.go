package dto

type CreateMenuDTO struct{
	NamaMobil string `json:"nama_mobil" form:"nama_mobil" binding:"required"`
	Harga uint64 `json:"harga" form:"harga"`
	Status string `json:"status" form:"status" binding:"required"`
	TipeMobil string `json:"tipe_mobil" form:"tipe_mobil" binding:"required"`
}

type UpdateMenuDTO struct{
	ID uint64 `json:"id" form:"id"`
	NamaMobil string `json:"nama_mobil" form:"nama_mobil" binding:"required"`
	Harga uint64 `json:"harga" form:"harga"`
	Status string `json:"status" form:"status" binding:"required"`
	TipeMobil string `json:"tipe_mobil" form:"tipe_mobil" binding:"required"`
}

