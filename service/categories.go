package service

import (
	"errors"
	"fmt"
	"jianji-server/entity"
	"jianji-server/model/request"
	"jianji-server/model/response"
	"jianji-server/utils"
	"jianji-server/utils/r"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type Categories struct {
}

func (*Categories) Load(c *gin.Context) (code int, message string, data *[]response.Categories) {
	userUUID, exists := c.Get("UserUUID")
	if !exists {
		code = r.USER_NOT_LOGIN
		return
	}

	err := utils.DB.Model(&entity.Categories{}).Where("user_uuid = ?", userUUID).Find(&data).Error
	if err != nil {
		code = r.ERROR_DB_OPE
		return
	}
	return
}

func (*Categories) Create(c *gin.Context) (code int, message string, data *response.Categories) {
	params, _ := utils.GetRequestParams[request.CreateCategories](c)
	userUUID, exists := c.Get("UserUUID")

	if !exists {
		code = r.USER_NOT_LOGIN
		data = nil
		return
	}

	categories := entity.Categories{
		UserUUID:    userUUID.(uuid.UUID),
		Label:       params.Label,
		ParentValue: lo.Ternary(params.ParentValue == 0, nil, lo.ToPtr(params.ParentValue)),
	}
	err := utils.DB.Create(&categories).Error

	if err != nil {
		code = r.ERROR_DB_OPE
		data = nil
		if !errors.Is(err, gorm.ErrForeignKeyViolated) {
			message = "父级分类不存在"
		}
		return
	}

	//遍历父元素，获取path
	parent := categories
	var path string
	for parent.ParentValue != nil {
		utils.DB.Preload("ParentCategories").First(&parent, parent.Value)
		parent = *parent.ParentCategories
		path += fmt.Sprintf("%d,", parent.Value)
	}

	//更新path到数据库
	if path != "" {
		categories.Path = path[:len(path)-1]
	}
	err = utils.DB.Save(&categories).Error
	if err != nil {
		code = r.ERROR_DB_OPE
		data = nil
		return
	}

	data = &response.Categories{
		Label:       categories.Label,
		Value:       categories.Value,
		ParentValue: categories.ParentValue,
		Path:        categories.Path,
	}
	return
}

func (*Categories) Delete(c *gin.Context) (code int, message string, data any) {
	params, _ := utils.GetRequestParams[request.DeleteCategories](c)
	err := utils.DB.Where(&entity.Categories{Value: params.Value}).Select("Categories").Delete(&entity.Categories{}).Error
	if err != nil {
		code = r.ERROR_DB_OPE
		data = nil
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			message = "该分类不为空，请先清空删除子分类"
		}
		return
	}

	return
}
