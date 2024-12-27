/*
 * @Author: weihua hu
 * @Date: 2024-12-25 14:02:50
 * @LastEditTime: 2024-12-25 14:05:18
 * @LastEditors: weihua hu
 * @Description:
 */
package entity

type AddressBook struct {
	BaseModel
	UserId       int64  `json:"userId" gorm:"column:user_id;type:bigint;not null"`         // 用户id
	Consignee    string `json:"consignee" gorm:"column:consignee;size:50;not null"`        // 收货人
	Phone        string `json:"phone" gorm:"column:phone;size:20;not null"`                // 手机号
	Sex          string `json:"sex" gorm:"column:sex;size:2"`                              // 性别 0 女 1 男
	ProvinceCode string `json:"provinceCode" gorm:"column:province_code;size:20"`          // 省级区划编号
	ProvinceName string `json:"provinceName" gorm:"column:province_name;size:50"`          // 省级名称
	CityCode     string `json:"cityCode" gorm:"column:city_code;size:20"`                  // 市级区划编号
	CityName     string `json:"cityName" gorm:"column:city_name;size:50"`                  // 市级名称
	DistrictCode string `json:"districtCode" gorm:"column:district_code;size:20"`          // 区级区划编号
	DistrictName string `json:"districtName" gorm:"column:district_name;size:50"`          // 区级名称
	Detail       string `json:"detail" gorm:"column:detail;size:200"`                      // 详细地址
	Label        string `json:"label" gorm:"column:label;size:50"`                         // 标签
	IsDefault    int    `json:"isDefault" gorm:"column:is_default;type:tinyint;default:0"` // 是否默认 0否 1是
}
