package models

type User struct {
	NIK uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`
	Email string `gorm:"type:varchar(255);not null" json:"email"`
	Password string `gorm:"->;<-;not null" json:"-"`
	Token string `gorm:"-" json:"token, omitempty"`
	Transactions *[]Transaction `json:"transaction, omitempty"`
	NoWa uint64 `gorm:"uniqueIndex;type:int" json:"no_wa"`
}