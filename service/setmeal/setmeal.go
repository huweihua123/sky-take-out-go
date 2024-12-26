/*
 * @Author: weihua hu
 * @Date: 2024-12-22 21:21:58
 * @LastEditTime: 2024-12-25 13:48:38
 * @LastEditors: weihua hu
 * @Description:
 */
package setmeal

import (
	"fmt"
	"sky-take-out-go/common/result"
	"sky-take-out-go/global"
	"sky-take-out-go/models/dto"
	"sky-take-out-go/models/entity"
	"sky-take-out-go/models/vo"
)

func Create(setmealDTO dto.SetMealDTO) error {
	tx := global.DB.Begin()
	if tx.Error != nil {
		return fmt.Errorf("开启事务失败: %s", tx.Error)
	}
	var setmeal entity.Setmeal = entity.Setmeal{
		CategoryId: setmealDTO.CategoryId,
		Image:      setmealDTO.Image,
		Name:       setmealDTO.Name,
		Price:      setmealDTO.Price,
		Status:     setmealDTO.Status,
	}
	if setmealDTO.Description != "" {
		setmeal.Description = setmealDTO.Description
	}
	// 保存套餐基本信息
	if err := tx.Create(&setmeal).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("创建套餐失败: %s", err.Error())
	}

	// 获取新创建的套餐ID，设置关联记录的setmealId
	for i := range setmealDTO.SetmealDishes {
		setmealDTO.SetmealDishes[i].SetmealId = setmeal.ID
	}

	// 批量创建套餐-菜品关联记录
	if err := tx.Create(&setmealDTO.SetmealDishes).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("创建套餐菜品关联失败: %s", err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("提交事务失败: %s", err.Error())
	}
	return nil
}

func Update(setmealDTO dto.SetMealDTO) error {
	tx := global.DB.Begin()
	if tx.Error != nil {
		return fmt.Errorf("开启事务失败: %s", tx.Error)
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 更新套餐基本信息
	setmeal := entity.Setmeal{
		Name:       setmealDTO.Name,
		CategoryId: setmealDTO.CategoryId,
		Price:      setmealDTO.Price,
		Image:      setmealDTO.Image,
		Status:     setmealDTO.Status,
	}
	if setmealDTO.Description != "" {
		setmeal.Description = setmealDTO.Description
	}

	if err := tx.Model(&entity.Setmeal{}).Where("id = ?", setmealDTO.Id).Updates(&setmeal).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("更新套餐失败: %s", err.Error())
	}

	// 删除原有套餐菜品关联
	if err := tx.Where("setmeal_id = ?", setmealDTO.Id).Delete(&entity.SetmealDish{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("删除原有套餐菜品失败: %s", err.Error())
	}

	// 重新插入新的套餐菜品关联
	for i := range setmealDTO.SetmealDishes {
		setmealDTO.SetmealDishes[i].SetmealId = setmealDTO.Id
	}
	if err := tx.Create(&setmealDTO.SetmealDishes).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("插入新的套餐菜品失败: %s", err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("提交事务失败: %s", err.Error())
	}
	return nil
}

func Delete(ids []int64) error {
	tx := global.DB.Begin()
	if tx.Error != nil {
		return fmt.Errorf("开启事务失败: %s", tx.Error)
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 检查套餐状态
	var count int64
	if err := tx.Model(&entity.Setmeal{}).Where("id IN ? AND status = 1", ids).Count(&count).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("查询套餐状态失败: %s", err.Error())
	}

	if count > 0 {
		tx.Rollback()
		return fmt.Errorf("无法删除在售套餐")
	}

	// 删除套餐
	if err := tx.Where("id IN ?", ids).Delete(&entity.Setmeal{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("删除套餐失败: %s", err.Error())
	}

	// 删除关联的套餐菜品
	if err := tx.Where("setmeal_id IN ?", ids).Delete(&entity.SetmealDish{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("删除套餐菜品失败: %s", err.Error())
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("提交事务失败: %s", err.Error())
	}
	return nil

}

func StartOrStop(id int64, status int) error {
	if err := global.DB.Model(&entity.Setmeal{BaseModel: entity.BaseModel{ID: id}}).Update("status", status).Error; err != nil {
		return fmt.Errorf("更新菜品状态失败: %s", err.Error())
	}

	return nil
}

func GetById(id int64) (vo.SetMealVO, error) {
	var setmeal entity.Setmeal
	if err := global.DB.Where("id = ?", id).First(&setmeal).Error; err != nil {
		return vo.SetMealVO{}, fmt.Errorf("查询套餐失败: %s", err.Error())
	}
	var Category entity.Category
	if err := global.DB.Where("id = ?", setmeal.CategoryId).First(&Category).Error; err != nil {
		return vo.SetMealVO{}, fmt.Errorf("查询套餐分类失败: %s", err.Error())
	}

	var setmealDishes []entity.SetmealDish
	if err := global.DB.Where("setmeal_id = ?", id).Find(&setmealDishes).Error; err != nil {
		return vo.SetMealVO{}, fmt.Errorf("查询套餐菜品失败: %s", err.Error())
	}

	return vo.SetMealVO{
		Id:            setmeal.ID,
		CategoryId:    setmeal.CategoryId,
		Description:   setmeal.Description,
		CategoryName:  Category.Name,
		Image:         setmeal.Image,
		Name:          setmeal.Name,
		Price:         setmeal.Price,
		SetMealDishes: setmealDishes,
		Status:        setmeal.Status,
	}, nil
}

func PageQuery(setmealPageQueryDTO dto.SetmealPageQueryDTO) (result.PageResult, error) {
	var setmeals []entity.Setmeal
	var count int64
	query := global.DB.Model(&entity.Setmeal{}).
		Select("setmeal.*, category.name AS category_name").
		Joins("LEFT JOIN category ON setmeal.category_id = category.id")

		// 如果有按分类过滤的需求
	if setmealPageQueryDTO.CategoryId != 0 {
		query = query.Where("setmeal.category_id = ?", setmealPageQueryDTO.CategoryId)
	}
	if setmealPageQueryDTO.Name != "" {
		query = query.Where("setmeal.name LIKE ?", "%"+setmealPageQueryDTO.Name+"%")
	}

	if setmealPageQueryDTO.Status != nil {
		query = query.Where("setmeal.status = ?", setmealPageQueryDTO.Status)
	}

	if err := query.Count(&count).Error; err != nil {
		return result.PageResult{}, fmt.Errorf("查询套餐总数失败: %s", err.Error())
	}

	offset := (setmealPageQueryDTO.Page - 1) * setmealPageQueryDTO.PageSize
	query = query.Offset(offset).Limit(setmealPageQueryDTO.PageSize).Order("setmeal.id DESC")

	if err := query.Find(&setmeals).Error; err != nil {
		return result.PageResult{}, fmt.Errorf("查询套餐失败: %s", err.Error())
	}
	return result.PageResult{
		Total:   int(count),
		Records: setmeals,
	}, nil
}
