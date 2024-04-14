package dao

import (
	"jianji-server/entity"
	"jianji-server/utils"

	"github.com/google/uuid"
)

type Post struct {
}

func (*Post) GetUserPostByPostUUID(userUUID uuid.UUID, postUUID uuid.UUID) (post *entity.Post, err error) {
	query := utils.DBQueryBegin()

	//找到数据库记录
	err = query.Model(&entity.Post{}).
		Where("uuid = ? AND user_uuid = ?", postUUID, userUUID).
		First(&post).
		Error
	return
}
