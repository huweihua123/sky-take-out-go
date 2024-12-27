/*
 * @Author: weihua hu
 * @Date: 2024-12-26 20:26:39
 * @LastEditTime: 2024-12-27 23:52:51
 * @LastEditors: weihua hu
 * @Description:
 */
package address_book

import (
	"sky-take-out-go/global"
	"sky-take-out-go/models/entity"
)

func Create(addressBook entity.AddressBook) error {
	addressBook.IsDefault = 0
	if err := global.DB.Create(&addressBook); err.Error != nil {
		return err.Error
	}
	return nil
}

func SetDefault(id int64, userId int64) error {
	// 开启事务
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 先将该用户的所有地址设为非默认
	if err := tx.Model(&entity.AddressBook{}).
		Where("user_id = ?", userId).
		Update("is_default", 0).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 将指定地址设为默认
	if err := tx.Model(&entity.AddressBook{}).
		Where("id = ? AND user_id = ?", id, userId).
		Update("is_default", 1).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func GetDefault(userId int64) (entity.AddressBook, error) {
	var addressBook entity.AddressBook
	if err := global.DB.Where("user_id = ? and is_default = 1", userId).First(&addressBook).Error; err != nil {
		return addressBook, err
	}
	return addressBook, nil
}

func Update(addressBook entity.AddressBook) error {
	updateFields := map[string]interface{}{
		"city_code":     addressBook.CityCode,
		"city_name":     addressBook.CityName,
		"consignee":     addressBook.Consignee,
		"detail":        addressBook.Detail,
		"district_code": addressBook.DistrictCode,
		"district_name": addressBook.DistrictName,
		"is_default":    addressBook.IsDefault,
		"label":         addressBook.Label,
		"phone":         addressBook.Phone,
		"province_code": addressBook.ProvinceCode,
		"province_name": addressBook.ProvinceName,
		"sex":           addressBook.Sex,
		"user_id":       addressBook.UserId,
	}

	return global.DB.Model(&entity.AddressBook{}).
		Where("id = ?", addressBook.ID).
		Updates(updateFields).Error
}

func List(userId int64) ([]entity.AddressBook, error) {
	var addressBooks []entity.AddressBook
	if err := global.DB.Where("user_id = ?", userId).Find(&addressBooks).Error; err != nil {
		return nil, err
	}
	return addressBooks, nil
}

func DeleteById(id int64) error {
	return global.DB.Where("id = ?", id).Delete(&entity.AddressBook{}).Error
}

func GetById(id int64) (entity.AddressBook, error) {
	var addressBook entity.AddressBook
	if err := global.DB.Where("id = ?", id).First(&addressBook).Error; err != nil {
		return addressBook, err
	}
	return addressBook, nil
}
