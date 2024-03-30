package response

type Categories struct {
	Label       string  `gorm:"type:varchar(32);comment:名称" json:"label"`
	Value       uint64  `gorm:"auto_increment;unique;not null;autoIncrement:100;comment:值" json:"value"`
	ParentValue *uint64 `gorm:"foreignKey:Value;comment:父级值" json:"parentValue"`
	Path        string  `gorm:"type:varchar(32);comment:路径" json:"path"`
}
