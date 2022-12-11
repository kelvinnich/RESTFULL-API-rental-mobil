package dto

type UpdateUserDTO struct{
	ID int `json:"id" form:"id"`
	Name string `json:"name" form:"name" binding:"required, omitempty"`
	Email string `json:email" form:"email" binding:"required,  omitempty"`
	Password string `json:"password,omitempty" form:"password, omitempty"`
}

type RegisterDTO struct{
	NoWa uint64 `json:"no_wa" form:"no_wa"`
	Name string `json:"name" form:"name" binding:"required"`
	Email string `json:email" form:"email" binding:"required, omitempty"`
	Password string `json:"password,omitempty" form:"password, omitempty"` 
}

type LoginDTO struct{
	Email string `json:email" form:"email" binding:"required, omitempty"`
	Password string `json:"password,omitempty" form:"password, omitempty"`
}