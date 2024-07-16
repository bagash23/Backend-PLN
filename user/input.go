package user

type RegisterUserInput struct {
	IDUser string `json:"id_user" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	NamaAdmin string `json:"nama_admin" binding:"required"`
	IDLevel string `json:"id_level" binding:"required"`
}

type LoginUserInput struct {	
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}