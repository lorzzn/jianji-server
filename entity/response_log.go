package entity

type ResponseLog struct {
	Universal
	StatusCode int    `gorm:"type:int"`
	RequestURL string `gorm:"type:text"`
	SessionId  string `gorm:"type:text"`
	TraceId    string `gorm:"type:text"`
	Method     string `gorm:"type:text"`
	Request    string `gorm:"type:text"`
	Response   string `gorm:"type:text"`
}
