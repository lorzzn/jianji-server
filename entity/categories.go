package entity

import "github.com/google/uuid"

type Categories struct {
	Universal
	UserUUID         uuid.UUID   `gorm:"not null" json:"userUUID"`
	User             User        `gorm:"foreignKey:UserUUID;references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Label            string      `gorm:"type:varchar(32);comment:名称" json:"label"`
	Value            uint64      `gorm:"auto_increment;unique;not null;autoIncrement:100;comment:值" json:"value"`
	ParentValue      *uint64     `gorm:"comment:父级值" json:"parentValue"`
	ParentCategories *Categories `gorm:"foreignKey:ParentValue;references:Value;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	OrdinalNumber    *uint64     `gorm:"comment:序数" json:"ordinalNumber"`
}
