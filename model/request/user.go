package request

type Login struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RefreshToken struct {
	Token        string `json:"token" form:"token"`
	RefreshToken string `json:"refreshToken" form:"refreshToken"`
}

type EditProfile struct {
	Name string `json:"name" form:"name"`
}
