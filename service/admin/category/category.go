/*
 * @Author: weihua hu
 * @Date: 2024-12-20 00:18:06
 * @LastEditTime: 2024-12-20 18:30:19
 * @LastEditors: weihua hu
 * @Description:
 */
package category

import (
	"fmt"
	"sky-take-out-go/common/result"
	"sky-take-out-go/global"
	"sky-take-out-go/models/dto"
	"sky-take-out-go/models/entity"
)

func Create(categortDTO dto.CategoryDTO) (err error) {
	var category entity.Category = entity.Category{
		Name: categortDTO.Name,
		Sort: categortDTO.Sort,
		Type: categortDTO.Type,
	}

	if err := global.DB.Create(&category).Error; err != nil {
		return fmt.Errorf("创建category失败: %s", err.Error())
	}

	return nil
}

func PageQuery(pageDTO dto.CategoryPageQueryDTO) (pageResult result.PageResult, err error) {
	var categorys []entity.Category
	var total int64

	query := global.DB.Offset((pageDTO.Page - 1) * pageDTO.PageSize).Limit(pageDTO.PageSize)

	if pageDTO.Name != "" {
		query = query.Where("name like ?", "%"+pageDTO.Name+"%")
	}

	if pageDTO.Type != 0 {
		query = query.Where("type = ?", pageDTO.Type)
	}

	if err := query.Find(&categorys).Error; err != nil {
		return result.PageResult{}, fmt.Errorf("查询category失败: %s", err.Error())
	}

	global.DB.Model(&entity.Category{}).Count(&total)

	result := result.PageResult{
		Total:   int(total),
		Records: categorys,
	}

	return result, nil
}

func StartOrStop(status int, id int64) error {

	if err := global.DB.Model(&entity.Category{BaseModel: entity.BaseModel{ID: id}}).Update("status", status).Error; err != nil {
		return fmt.Errorf("更新category状态失败: %s", err.Error())
	}

	return nil
}

func Update(categoryDTO dto.CategoryDTO) error {
	if err := global.DB.Model(&entity.Category{BaseModel: entity.BaseModel{ID: categoryDTO.ID}}).Updates(entity.Category{
		Name: categoryDTO.Name,
		Sort: categoryDTO.Sort,
		Type: categoryDTO.Type,
	}).Error; err != nil {
		return fmt.Errorf("更新category信息失败: %s", err.Error())
	}

	return nil
}

func Delete(id int64) error {
	// 创建更新对象
	updates := map[string]interface{}{
		"is_deleted": true,
	}

	// 执行软删除并更新 is_deleted 字段
	if err := global.DB.Model(&entity.Category{}).
		Where("id = ?", id).
		Updates(updates).
		Delete(&entity.Category{}).Error; err != nil {
		return fmt.Errorf("删除category失败: %s", err.Error())
	}

	return nil
}

func List(typet *int) ([]entity.Category, error) {
	var categorys []entity.Category
	query := global.DB

	if typet != nil {
		query = query.Where("type = ?", *typet)
	}

	if err := query.Find(&categorys).Error; err != nil {
		return nil, fmt.Errorf("查询category失败: %s", err.Error())
	}
	return categorys, nil
}
