/*
 * @Author: weihua hu
 * @Date: 2024-12-25 13:59:24
 * @LastEditTime: 2024-12-25 14:01:01
 * @LastEditors: weihua hu
 * @Description:
 */
package entity

import (
	"time"
)

// 订单状态常量
const (
	// 订单状态 1待付款 2待接单 3已接单 4派送中 5已完成 6已取消
	PendingPayment     = 1
	ToBeConfirmed      = 2
	Confirmed          = 3
	DeliveryInProgress = 4
	Completed          = 5
	Cancelled          = 6
)

// 支付状态常量
const (
	// 支付状态 0未支付 1已支付 2退款
	UnPaid = 0
	Paid   = 1
	Refund = 2
)

// Order 订单实体
type Order struct {
	BaseModel
	Number                string    `json:"number" gorm:"column:number;type:varchar(50);not null"`            // 订单号
	Status                int       `json:"status" gorm:"column:status;type:int;not null"`                    // 订单状态
	UserId                int64     `json:"userId" gorm:"column:user_id;type:bigint;not null"`                // 下单用户id
	AddressBookId         int64     `json:"addressBookId" gorm:"column:address_book_id;type:bigint;not null"` // 地址id
	OrderTime             time.Time `json:"orderTime" gorm:"column:order_time"`                               // 下单时间
	CheckoutTime          time.Time `json:"checkoutTime" gorm:"column:checkout_time"`                         // 结账时间
	PayMethod             int       `json:"payMethod" gorm:"column:pay_method;type:int"`                      // 支付方式 1微信 2支付宝
	PayStatus             int       `json:"payStatus" gorm:"column:pay_status;type:int"`                      // 支付状态
	Amount                float64   `json:"amount" gorm:"column:amount;type:decimal(10,2)"`                   // 实收金额
	Remark                string    `json:"remark" gorm:"column:remark;type:varchar(100)"`                    // 备注
	UserName              string    `json:"userName" gorm:"column:user_name;type:varchar(50)"`                // 用户名
	Phone                 string    `json:"phone" gorm:"column:phone;type:varchar(20)"`                       // 手机号
	Address               string    `json:"address" gorm:"column:address;type:varchar(255)"`                  // 地址
	Consignee             string    `json:"consignee" gorm:"column:consignee;type:varchar(50)"`               // 收货人
	CancelReason          string    `json:"cancelReason" gorm:"column:cancel_reason;type:varchar(255)"`       // 取消原因
	RejectionReason       string    `json:"rejectionReason" gorm:"column:rejection_reason;type:varchar(255)"` // 拒绝原因
	CancelTime            time.Time `json:"cancelTime" gorm:"column:cancel_time"`                             // 取消时间
	EstimatedDeliveryTime time.Time `json:"estimatedDeliveryTime" gorm:"column:estimated_delivery_time"`      // 预计送达时间
	DeliveryStatus        int       `json:"deliveryStatus" gorm:"column:delivery_status;type:int"`            // 配送状态 1立即送出 0选择具体时间
	DeliveryTime          time.Time `json:"deliveryTime" gorm:"column:delivery_time"`                         // 送达时间
	PackAmount            int       `json:"packAmount" gorm:"column:pack_amount;type:int"`                    // 打包费
	TablewareNumber       int       `json:"tablewareNumber" gorm:"column:tableware_number;type:int"`          // 餐具数量
	TablewareStatus       int       `json:"tablewareStatus" gorm:"column:tableware_status;type:int"`          // 餐具数量状态 1按餐量提供 0选择具体数量
}
