package service

import (
	"jianji-server/dao"
	"jianji-server/entity"
	"jianji-server/model/request"
	"jianji-server/model/response"
	"jianji-server/utils"
	"jianji-server/utils/r"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type Posts struct {
}

func (*Posts) List(c *gin.Context) (code int, message string, data *[]response.Post) {
	userUUID, _ := c.Get("UserUUID")
	query := utils.DBQueryBegin()

	err := query.Model(&entity.Post{}).Preload(clause.Associations).Where("user_uuid = ?", userUUID).Find(&data).Error
	if err != nil {
		code = r.ERROR_DB_OPE
		data = nil
		return
	}
	return
}

func (*Posts) Create(c *gin.Context) (code int, message string, data *response.Post) {
	params, _ := utils.GetRequestParams[request.CreatePost](c)
	userUUID, _ := c.Get("UserUUID")
	query := utils.DBQueryBegin()

	var tags []entity.Tag
	if params.TagValues != nil {
		for _, tag := range *params.TagValues {
			tags = append(tags, entity.Tag{Value: tag})
		}
	} else {
		tags = nil
	}

	//创建
	post := &entity.Post{
		UserFK:        entity.UserFK{UserUUID: userUUID.(uuid.UUID)},
		Title:         params.Title,
		Content:       params.Content,
		CategoryValue: params.CategoryValue,
		Tags:          &tags,
		Favoured:      params.Favoured,
		Public:        params.Favoured,
		Status:        params.Status,
	}
	err := query.Create(post).Error
	if err != nil {
		query.Rollback()
		data = nil
		code = r.ERROR_DB_OPE
		message = "创建文章失败"
		return
	}
	//读取
	err = query.Preload(clause.Associations).Where(post).Find(&data).Error
	if err != nil {
		query.Rollback()
		data = nil
		code = r.ERROR_DB_OPE
		return
	}
	query.Commit()
	return
}

func (*Posts) Update(c *gin.Context) (code int, message string, data *response.Post) {
	params, _ := utils.GetRequestParams[request.UpdatePost](c)
	userUUID, _ := c.Get("UserUUID")
	query := utils.DBQueryBegin()

	//找到数据库记录
	post, err := dao.PostDao.GetUserPostByPostUUID(userUUID.(uuid.UUID), params.UUID)
	if err != nil {
		code = r.ERROR_DB_OPE
		data = nil
		message = "文章不存在"
		return
	}

	//更新标签关联
	if params.TagValues != nil {
		err = query.Model(&post).Association("Tags").Clear()
		if err == nil {
			var tagArray []entity.Tag
			for _, tag := range *params.TagValues {
				tagArray = append(tagArray, entity.Tag{Value: tag})
			}
			post.Tags = &tagArray
			err = query.Save(&post).Error
		}
	}
	if err != nil {
		utils.DBQueryRollback(query)
		code = r.ERROR_DB_OPE
		data = nil
		message = "保存失败"
		return
	}

	//更新其他数据
	updated := entity.Post{
		Title:         params.Title,
		Content:       params.Content,
		CategoryValue: params.CategoryValue,
		Favoured:      params.Favoured,
		Public:        params.Public,
		Status:        params.Status,
	}

	err = query.Model(&post).Updates(&updated).Preload(clause.Associations).First(&data).Error
	if err != nil {
		utils.DBQueryRollback(query)
		code = r.ERROR_DB_OPE
		data = nil
		message = "保存失败"
		return
	}
	utils.DBQueryCommit(query)

	return
}

func (*Posts) Delete(c *gin.Context) (code int, message string, data any) {
	params, _ := utils.GetRequestParams[request.DeletePost](c)
	userUUID, _ := c.Get("UserUUID")
	query := utils.DBQueryBegin()

	//找到数据库记录
	post, err := dao.PostDao.GetUserPostByPostUUID(userUUID.(uuid.UUID), params.UUID)
	if err != nil {
		code = r.ERROR_DB_OPE
		data = nil
		message = "文章不存在"
		return
	}

	//删除标签关联和文章
	err = query.Model(&post).Association("Tags").Clear()
	if err == nil {
		err = query.Where("user_uuid = ? AND uuid = ?", userUUID, params.UUID).Delete(&entity.Post{}).Error
	}
	if err != nil {
		query.Rollback()
		code = r.ERROR_DB_OPE
		data = nil
		message = "删除失败"
		return
	}

	query.Commit()
	return
}
