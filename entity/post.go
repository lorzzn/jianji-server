package entity

type Post struct {
	Universal
	UserFK
	Title         string    `gorm:"type:varchar(64);comment:标题" json:"title"`
	Content       string    `gorm:"comment:内容" json:"content"`
	Description   string    `gorm:"type:varchar(300);comment:文章描述" json:"description"`
	CategoryValue *uint64   `gorm:"comment:分类" json:"categoryValue"`
	Category      *Category `gorm:"foreignKey:CategoryValue;references:Value;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"category"`
	Tags          *[]Tag    `gorm:"many2many:post_tags;foreignKey:UUID;joinForeignKey:PostUUID;references:Value;joinReferences:TagValue;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"tags"`
	Favoured      *bool     `gorm:"comment:标记为喜爱收藏;default:false" json:"favoured"`
	Archived      *bool     `gorm:"comment:标记为归档;default:false" json:"archived"`
	Public        *bool     `gorm:"comment:公开;default:false" json:"public"`
	Status        *uint64   `gorm:"comment:文章状态:0:待编辑,1: 草稿,2: 已发布,3: 隐藏;default:1" json:"status"`
}

func (p Post) TableName() string {
	return "post"
}
