/*
 * @Author: weihua hu
 * @Date: 2024-12-15 20:46:59
 * @LastEditTime: 2024-12-19 23:07:22
 * @LastEditors: weihua hu
 * @Description:
 */
package service

import (
	"errors"
	"sky-take-out-go/common/constant"
	"sky-take-out-go/common/result"
	"sky-take-out-go/global"
	"sky-take-out-go/models/dto"
	"sky-take-out-go/models/entity"
	"sky-take-out-go/utils"
)

func Login(EmployeeLoginDTO dto.EmployeeLoginDTO) (entity.Employee, error) {
	username := EmployeeLoginDTO.Username
	password := EmployeeLoginDTO.Password

	employee := entity.Employee{}

	result := global.DB.Where("username = ?", username).Find(&employee)

	if result.RowsAffected == 0 {
		return entity.Employee{}, errors.New(constant.ACCOUNT_NOT_FOUND)
	}

	encryptedPassword := utils.MD5Encrypt(password)

	if encryptedPassword != employee.Password {
		return entity.Employee{}, errors.New(constant.PASSWORD_ERROR)
	}

	if employee.Status == constant.DISABLE {
		// 账号被锁定
		return entity.Employee{}, errors.New(constant.ACCOUNT_LOCKED)
	}
	return employee, nil

}

func Save(employeeDTO dto.EmployeeDTO) error {
	employee := entity.Employee{
		Username: employeeDTO.Username,
		Name:     employeeDTO.Name,
		Phone:    employeeDTO.Phone,
		Sex:      employeeDTO.Sex,
		IDNumber: employeeDTO.IDNumber,
	}

	employee.Password = utils.MD5Encrypt(constant.DEFAULT_PASSWORD)
	employee.Status = constant.ENABLE

	// 保存到数据库
	if err := global.DB.Create(&employee).Error; err != nil {
		return errors.New("创建员工失败: ")
	}

	return nil
}

func StartOrStop(status int, id int64) error {
	var employee entity.Employee

	// 查找员工记录
	if err := global.DB.First(&employee, id).Error; err != nil {
		return errors.New("员工不存在")
	}

	// 更新员工状态
	if err := global.DB.Model(&employee).Update("status", status).Error; err != nil {
		return errors.New("更新员工状态失败")
	}

	return nil
}

func Update(employeeDTO dto.EmployeeDTO) error {
	// 直接更新员工信息
	if err := global.DB.Model(&entity.Employee{BaseModel: entity.BaseModel{ID: employeeDTO.ID}}).Updates(entity.Employee{
		Username: employeeDTO.Username,
		Name:     employeeDTO.Name,
		Phone:    employeeDTO.Phone,
		Sex:      employeeDTO.Sex,
		IDNumber: employeeDTO.IDNumber,
	}).Error; err != nil {
		return errors.New("更新员工信息失败: " + err.Error())
	}
	return nil
}

func GetById(id int64) (entity.Employee, error) {
	var employee entity.Employee
	if err := global.DB.First(&employee, id).Error; err != nil {
		return entity.Employee{}, errors.New("员工不存在")
	}
	return employee, nil
}

func PageQuery(employeePageQueryDTO dto.EmployeePageQueryDTO) (Pageresult result.PageResult, err error) {
	var employees []entity.Employee
	// 计算偏移量
	offset := (employeePageQueryDTO.Page - 1) * employeePageQueryDTO.PageSize

	// 分页查询员工记录
	query := global.DB.Offset(offset).Limit(employeePageQueryDTO.PageSize)

	// 添加过滤条件
	if employeePageQueryDTO.Name != "" {
		query = query.Where("name LIKE ?", "%"+employeePageQueryDTO.Name+"%")
	}

	// 执行查询
	if err := query.Find(&employees).Error; err != nil {
		return result.PageResult{}, err
	}
	// 获取总记录数
	var total int64
	global.DB.Model(&entity.Employee{}).Count(&total)

	results := result.PageResult{
		Total:   int(total),
		Records: employees,
	}
	return results, nil
}
