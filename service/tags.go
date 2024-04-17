package service

import (
	"jianji-server/entity"
	"jianji-server/model/request"
	"jianji-server/model/response"
	"jianji-server/utils"
	"jianji-server/utils/r"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Tags struct {
}

func (*Tags) List(c *gin.Context) (code int, message string, data *[]response.Tag) {
	userUUID, _ := c.Get("UserUUID")

	err := utils.DB.Model(&entity.Tag{}).Where("user_uuid = ?", userUUID).Find(&data).Error
	if err != nil {
		code = r.ERROR_DB_OPE
		return
	}
	return
}

func (*Tags) Create(c *gin.Context) (code int, message string, data []*response.Tag) {
	params, _ := utils.GetRequestParams[request.CreateTags](c)
	userUUID, _ := c.Get("UserUUID")

	tx := utils.DBQueryBegin()
	var err error
	for _, datum := range params.Data {
		tag := &entity.Tag{
			UserFK: entity.UserFK{UserUUID: userUUID.(uuid.UUID)},
			Label:  *datum.Label,
		}
		err = utils.DB.Create(&tag).Error
		if err != nil {
			break
		}
		data = append(data, &response.Tag{
			Label: tag.Label,
			Value: tag.Value,
		})
	}

	if err != nil {
		tx.Rollback()
		code = r.ERROR_DB_OPE
		data = nil
		message = "创建标签失败"
		return
	}

	tx.Commit()
	return
}

func (*Tags) Update(c *gin.Context) (code int, message string, data []*response.Tag) {
	params, _ := utils.GetRequestParams[request.UpdateTags](c)
	userUUID, _ := c.Get("UserUUID")

	var paramsValues []uint64
	paramsValueMap := make(map[uint64]request.UpdateTagDatum)
	for _, datum := range params.Data {
		if datum.Value != nil {
			paramsValues = append(paramsValues, *datum.Value)
			paramsValueMap[*datum.Value] = datum
		}
	}

	var tags []entity.Tag
	err := utils.DB.Where("value IN (?) AND user_uuid = ?", paramsValues, userUUID).Find(&tags).Error
	if err != nil {
		code = r.ERROR_DB_OPE
		message = "你操作的分类不存在"
		return
	}

	tx := utils.DBQueryBegin()

	for _, tag := range tags {
		paramValue := paramsValueMap[tag.Value]
		if paramValue.Label != nil {
			tag.Label = *paramValue.Label
		}
		err = tx.Save(&tag).Error
		if err != nil {
			break
		}
		data = append(data, &response.Tag{
			Label: tag.Label,
			Value: tag.Value,
		})
	}
	if err != nil {
		tx.Rollback()
		code = r.ERROR_DB_OPE
		message = "保存失败"
		data = nil
		return
	}

	tx.Commit()

	return
}

func (*Tags) Delete(c *gin.Context) (code int, message string, data any) {
	params, _ := utils.GetRequestParams[request.DeleteTags](c)
	userUUID, _ := c.Get("UserUUID")

	err := utils.DB.Where("value IN (?) AND user_uuid = ?", params.Value, userUUID).Delete(&entity.Tag{}).Error
	if err != nil {
		code = r.ERROR_DB_OPE
		data = nil
		message = "删除失败"
		return
	}

	return
}

func (*Tags) TagStatistics(c *gin.Context) (code int, message string, data *response.TagStatistics) {
	params, _ := utils.GetRequestParams[request.TagStatistics](c)
	userUUID := c.MustGet("UserUUID").(uuid.UUID)

	query := utils.DBQueryBegin()
	var tag entity.Tag
	var totalPosts int64

	err := query.
		Model(&entity.Post{}).
		Select("post.uuid,post_tags.post_uuid").
		Joins("INNER JOIN post_tags ON post.uuid = post_tags.post_uuid").
		Where("post_tags.tag_value = ?", params.Value).
		Count(&totalPosts).
		Error
	if err == nil {
		err = query.Model(&entity.Tag{}).Where(&entity.Tag{Value: params.Value, UserFK: entity.UserFK{UserUUID: userUUID}}).Find(&tag).Error
	}

	if err != nil {
		query.Rollback()
		data = nil
		code = r.ERROR_DB_OPE
		return
	}

	data = &response.TagStatistics{
		TotalPosts: totalPosts,
		CreateAt:   tag.CreatedAt,
		UpdatedAt:  tag.UpdatedAt,
	}

	query.Commit()
	return
}
