package service

import (
	"errors"
	"jianji-server/entity"
	"jianji-server/model/request"
	"jianji-server/model/response"
	"jianji-server/utils"
	"jianji-server/utils/r"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Categories struct {
}

func (*Categories) List(c *gin.Context) (code int, message string, data *[]response.Categories) {
	userUUID, _ := c.Get("UserUUID")

	err := utils.DB.Model(&entity.Categories{}).Where("user_uuid = ?", userUUID).Find(&data).Error
	if err != nil {
		code = r.ERROR_DB_OPE
		return
	}
	return
}

func (*Categories) Create(c *gin.Context) (code int, message string, data []*response.Categories) {
	params, _ := utils.GetRequestParams[request.CreateCategories](c)
	userUUID, _ := c.Get("UserUUID")

	tx := utils.DB.Begin()
	var err error
	for _, datum := range params.Data {
		category := &entity.Categories{
			UserUUID:      userUUID.(uuid.UUID),
			Label:         *datum.Label,
			ParentValue:   datum.ParentValue,
			OrdinalNumber: datum.OrdinalNumber,
		}
		err = utils.DB.Create(&category).Error
		if err != nil {
			break
		}
		data = append(data, &response.Categories{
			Label:         category.Label,
			Value:         category.Value,
			ParentValue:   category.ParentValue,
			OrdinalNumber: category.OrdinalNumber,
		})
	}

	if err != nil {
		tx.Rollback()
		code = r.ERROR_DB_OPE
		data = nil
		message = "创建分类失败"
		if !errors.Is(err, gorm.ErrForeignKeyViolated) {
			message = "父级分类不存在"
		}
		return
	}

	return
}

func (*Categories) Update(c *gin.Context) (code int, message string, data []*response.Categories) {
	params, _ := utils.GetRequestParams[request.UpdateCategories](c)
	userUUID, _ := c.Get("UserUUID")

	var paramsValues []uint64
	paramsValueMap := make(map[uint64]request.UpdateCategoriesDatum)
	for _, datum := range params.Data {
		if datum.Value != nil {
			paramsValues = append(paramsValues, *datum.Value)
			paramsValueMap[*datum.Value] = datum
		}
	}

	var categories []entity.Categories
	err := utils.DB.Where("value IN (?) AND user_uuid = ?", paramsValues, userUUID).Find(&categories).Error
	if err != nil {
		code = r.ERROR_DB_OPE
		message = "你操作的分类不存在"
		return
	}

	tx := utils.DB.Begin()

	for _, category := range categories {
		paramValue := paramsValueMap[category.Value]
		category.ParentValue = paramValue.ParentValue
		if paramValue.Label != nil {
			category.Label = *paramValue.Label
		}
		if paramValue.OrdinalNumber != nil {
			category.OrdinalNumber = paramValue.OrdinalNumber
		}
		err = tx.Save(&category).Error
		if err != nil {
			break
		}
		data = append(data, &response.Categories{
			Label:         category.Label,
			Value:         category.Value,
			ParentValue:   category.ParentValue,
			OrdinalNumber: category.OrdinalNumber,
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

func (*Categories) Delete(c *gin.Context) (code int, message string, data any) {
	params, _ := utils.GetRequestParams[request.DeleteCategories](c)
	userUUID, _ := c.Get("UserUUID")

	err := utils.DB.Where(&entity.Categories{Value: params.Value, UserUUID: userUUID.(uuid.UUID)}).Select("Categories").Delete(&entity.Categories{}).Error
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
