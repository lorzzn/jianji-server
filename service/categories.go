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

type Categories struct {
}

func (*Categories) List(c *gin.Context) (code int, message string, data *[]response.Category) {
	userUUID, _ := c.Get("UserUUID")

	err := utils.DB.Model(&entity.Category{}).Where("user_uuid = ?", userUUID).Find(&data).Error
	if err != nil {
		code = r.ERROR_DB_OPE
		return
	}
	return
}

func (*Categories) Create(c *gin.Context) (code int, message string, data []*response.Category) {
	params, _ := utils.GetRequestParams[request.CreateCategories](c)
	userUUID, _ := c.Get("UserUUID")

	tx := utils.DBQueryBegin()
	var err error
	for _, datum := range params.Data {
		category := &entity.Category{
			UserFK:        entity.UserFK{UserUUID: userUUID.(uuid.UUID)},
			Label:         *datum.Label,
			ParentValue:   datum.ParentValue,
			OrdinalNumber: datum.OrdinalNumber,
		}
		err = utils.DB.Create(&category).Error
		if err != nil {
			break
		}
		data = append(data, &response.Category{
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
		//if !errors.Is(err, gorm.ErrForeignKeyViolated) {
		//	message = "父级分类不存在"
		//}
		return
	}

	tx.Commit()
	return
}

func (*Categories) Update(c *gin.Context) (code int, message string, data []*response.Category) {
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

	var categories []entity.Category
	err := utils.DB.Where("value IN (?) AND user_uuid = ?", paramsValues, userUUID).Find(&categories).Error
	if err != nil {
		code = r.ERROR_DB_OPE
		message = "你操作的分类不存在"
		return
	}

	tx := utils.DBQueryBegin()

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
		data = append(data, &response.Category{
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

	err := utils.DB.Where("value IN (?) AND user_uuid = ?", params.Value, userUUID).Delete(&entity.Category{}).Error
	if err != nil {
		code = r.ERROR_DB_OPE
		data = nil
		message = "删除失败"
		//if errors.Is(err, gorm.ErrForeignKeyViolated) {
		//	message = "该分类不为空，请先清空删除子分类"
		//}
		return
	}

	return
}

func (*Categories) CategoryStatistics(c *gin.Context) (code int, message string, data *response.CategoryStatistics) {
	params, _ := utils.GetRequestParams[request.CategoryStatistics](c)
	userUUID := c.MustGet("UserUUID").(uuid.UUID)

	query := utils.DBQueryBegin()
	var category entity.Category
	var totalPosts int64

	err := query.Model(&entity.Post{}).Where(&entity.Post{CategoryValue: &params.Value, UserFK: entity.UserFK{UserUUID: userUUID}}).Count(&totalPosts).Error
	if err == nil {
		err = query.Model(&entity.Category{}).Where(&entity.Category{Value: params.Value, UserFK: entity.UserFK{UserUUID: userUUID}}).Find(&category).Error
	}

	if err != nil {
		query.Rollback()
		data = nil
		code = r.ERROR_DB_OPE
		return
	}

	data = &response.CategoryStatistics{
		TotalPosts: totalPosts,
		CreateAt:   category.CreatedAt,
		UpdatedAt:  category.UpdatedAt,
	}

	query.Commit()
	return
}
