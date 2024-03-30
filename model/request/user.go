package request

type Login struct {
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	Fingerprint string `json:"fingerprint" form:"fingerprint"`
	From        string `json:"-" form:"-"`
}

type Signup = Login

type Active struct {
	Email string `json:"email" form:"email"`
	State string `json:"state" form:"state"`
}

type RefreshToken struct {
	Token        string `json:"token" form:"token"`
	RefreshToken string `json:"refreshToken" form:"refreshToken"`
}

type EditProfile struct {
	Name string `json:"name" form:"name"`
}
