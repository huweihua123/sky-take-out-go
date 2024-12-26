/*
 * @Author: weihua hu
 * @Date: 2024-12-20 00:24:42
 * @LastEditTime: 2024-12-22 21:07:31
 * @LastEditors: weihua hu
 * @Description:
 */
package dish

import (
	"fmt"
	"sky-take-out-go/global"
	"sky-take-out-go/models/dto"
	"sky-take-out-go/models/entity"
)

func Create(dishDTO dto.DishDTO) error {
	// 开启事务
	tx := global.DB.Begin()
	if tx.Error != nil {
		return fmt.Errorf("开启事务失败: %s", tx.Error)
	}

	// 保存菜品基本信息
	dish := entity.Dish{
		Name:       dishDTO.Name,
		CategoryId: dishDTO.CategoryId,
		Price:      dishDTO.Price,
	}

	if dishDTO.Image != "" {
		dish.Image = dishDTO.Image
	}
	if dishDTO.Description != "" {
		dish.Description = dishDTO.Description
	}
	if dishDTO.Status != 0 {
		dish.Status = dishDTO.Status
	}

	if err := tx.Create(&dish).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("保存菜品失败: %s", err.Error())
	}

	// 保存口味信息
	if len(dishDTO.Flavors) > 0 {
		flavors := make([]entity.DishFlavor, 0)
		for _, flavorDTO := range dishDTO.Flavors {
			flavor := entity.DishFlavor{
				DishId: dish.ID,
				Name:   flavorDTO.Name,
				Value:  flavorDTO.Value,
			}
			flavors = append(flavors, flavor)
		}

		if err := tx.Create(&flavors).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("保存口味信息失败: %s", err.Error())
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("提交事务失败: %s", err.Error())
	}

	return nil
}

func Update(dishDTO dto.DishDTO) error {
	// 开启事务
	tx := global.DB.Begin()
	if tx.Error != nil {
		return fmt.Errorf("开启事务失败: %s", tx.Error)
	}

	// 更新菜品基本信息
	dish := entity.Dish{
		Name:       dishDTO.Name,
		CategoryId: dishDTO.CategoryId,
		Price:      dishDTO.Price,
	}

	if dishDTO.Image != "" {
		dish.Image = dishDTO.Image
	}
	if dishDTO.Description != "" {
		dish.Description = dishDTO.Description
	}
	if dishDTO.Status != 0 {
		dish.Status = dishDTO.Status
	}

	if err := tx.Model(&entity.Employee{BaseModel: entity.BaseModel{ID: dishDTO.ID}}).Updates(&dish).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("更新菜品失败: %s", err.Error())
	}

	// 删除原有口味信息
	if err := tx.Where("dish_id = ?", dish.ID).Delete(&entity.DishFlavor{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("删除原有口味信息失败: %s", err.Error())
	}

	// 保存口味信息
	if len(dishDTO.Flavors) > 0 {
		flavors := make([]entity.DishFlavor, 0)
		for _, flavorDTO := range dishDTO.Flavors {
			flavor := entity.DishFlavor{
				DishId: dishDTO.ID,
				Name:   flavorDTO.Name,
				Value:  flavorDTO.Value,
			}
			flavors = append(flavors, flavor)
		}

		if err := tx.Create(&flavors).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("保存口味信息失败: %s", err.Error())
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("提交事务失败: %s", err.Error())
	}

	return nil
}

func StartOrStop(status int, id int64) error {
	if err := global.DB.Model(&entity.Dish{BaseModel: entity.BaseModel{ID: id}}).Update("status", status).Error; err != nil {
		return fmt.Errorf("更新菜品状态失败: %s", err.Error())
	}

	return nil
}

func GetById(id int64) (dto.DishDTO, error) {
	// 1. 查询菜品基本信息
	var dish entity.Dish
	if err := global.DB.First(&dish, id).Error; err != nil {
		return dto.DishDTO{}, fmt.Errorf("查询菜品失败: %s", err.Error())
	}

	// 2. 查询口味信息
	var flavors []entity.DishFlavor
	if err := global.DB.Where("dish_id = ?", id).Find(&flavors).Error; err != nil {
		return dto.DishDTO{}, fmt.Errorf("查询口味失败: %s", err.Error())
	}

	// 3. 组装 DishDTO
	dishDTO := dto.DishDTO{
		ID:          dish.ID,
		Name:        dish.Name,
		CategoryId:  dish.CategoryId,
		Price:       dish.Price,
		Image:       dish.Image,
		Description: dish.Description,
		Status:      dish.Status,
		Flavors:     flavors,
	}

	return dishDTO, nil
}

func BatchDelete(ids []int64) error {
	// 开启事务
	tx := global.DB.Begin()
	if tx.Error != nil {
		return fmt.Errorf("开启事务失败: %s", tx.Error)
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 检查菜品状态
	var count int64
	if err := tx.Model(&entity.Dish{}).Where("id IN ? AND status = 1", ids).Count(&count).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("查询菜品状态失败: %s", err.Error())
	}

	if count > 0 {
		tx.Rollback()
		return fmt.Errorf("无法删除在售菜品")
	}

	// 删除菜品
	if err := tx.Where("id IN ?", ids).Delete(&entity.Dish{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("删除菜品失败: %s", err.Error())
	}

	// 删除关联的口味
	if err := tx.Where("dish_id IN ?", ids).Delete(&entity.DishFlavor{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("删除口味失败: %s", err.Error())
	}

	return tx.Commit().Error
}
