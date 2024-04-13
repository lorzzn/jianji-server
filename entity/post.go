package entity

type Post struct {
	Universal
	UserFK
	Title          string    `gorm:"type:varchar(64);comment:标题" json:"title"`
	Content        string    `gorm:"comment:内容" json:"content"`
	Category       *uint64   `gorm:"comment:分类" json:"category"`
	CategoryStruct *Category `gorm:"foreignKey:Category;references:Value;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	Tags           *[]Tag    `gorm:"many2many:post_tags;foreignKey:UUID;joinForeignKey:PostUUID;references:Value;joinReferences:TagValue;" json:"tags"`
	Favoured       *bool     `gorm:"comment:标记为喜爱收藏;default:false" json:"favoured"`
	Public         *bool     `gorm:"comment:公开;default:false" json:"public"`
	Status         *uint64   `gorm:"comment:文章状态: 1: 草稿; 2: 已发布; 3: 隐藏;default:1" json:"status"`
}
