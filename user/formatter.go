package user

type UserFornatter struct {
	IDUser string `json:"id_user"`
	Username string `json:"username"`
	Password string `json:"password"`
	NamaAdmin string `json:"nama_admin"`
	IDLevel string `json:"id_level"`
}

func FormatUser(user User) UserFornatter {
	formatter := UserFornatter {
		IDUser: user.IDUser,
		Username: user.Username,
		Password: user.Password,
		NamaAdmin: user.NamaAdmin,
		IDLevel: user.IDLevel,		
	}

	return formatter
}