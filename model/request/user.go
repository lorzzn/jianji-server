package request

type Signup struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	Name     string `form:"name" json:"name"`
	Avatar   string `form:"avatar" json:"avatar"`
}

type Login struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
