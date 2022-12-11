package dto

type CreateTransactionDTO struct{
	MenuID uint64 `json:"menu_id" form:"menu_id" binding:"required"`
	UserID uint64 `json:"user_id" form:"user_id" binding:"required"`
	DurasiPinjam int `json:"durasi_pinjam" form:"durasi_pinjam" binding:"required"`
	TotalPembayaran int `json:"total_pembayaran" form:"total_pembayaran"`
}

type UpdateTransactionDTO struct{
	ID uint64 `json:"id" form:"id"`
	MenuID uint64 `json:"menu_id" form:"menu_id" binding:"required, omitempty"`
	UserID uint64 `json:"user_id" form:"user_id" binding:"required, omitempty"`
	DurasiPinjam uint64 `json:"durasi_pinjam" form:"durasi_pinjam" binding:"required"`
	TotalPembayaran uint64 `json:"total_pembayaran" form:"total_pembayaran" binding:"required"`
}