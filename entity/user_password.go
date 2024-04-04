package entity

type UserPassword struct {
	Universal
	UniqueUserFK
	Password string `gorm:"type:varchar(256);comment:密码" json:"-"`
}
