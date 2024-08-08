// Code generated by http://www.gotool.top
package model

import (
	"time"
)

type InRecords struct {
	Id       int32     `gorm:"column:id;primary_key;NOT NULL" json:"id"`
	Type     string    `gorm:"column:type;default:NULL;comment:'类型'" json:"type"`
	TypeId   int32     `gorm:"column:type_id;default:NULL;comment:'类型id'" json:"type_id"`
	Weight   float32   `gorm:"column:weight;default:NULL;comment:'重量'" json:"weight"`
	Price    float32   `gorm:"column:price;default:NULL;comment:'单价'" json:"price"`
	Sum      float32   `gorm:"column:sum;default:NULL;comment:'总价'" json:"sum"`
	Date     time.Time `gorm:"column:date;default:NULL;comment:'时间'" json:"date"`
	Settle   float32   `gorm:"column:settle;default:NULL;comment:'结账金额'" json:"settle"`
	IsSettle int8      `gorm:"column:is_settle;default:NULL;comment:'是否结账'" json:"is_settle"`
}

func (i *InRecords) TableName() string {
	return "in_records"
}
