// Code generated by http://www.gotool.top
package model

type InTypes struct {
	Id   int32  `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL" json:"id"`
	Type string `gorm:"column:type;NOT NULL;comment:'类型'" json:"type"`
}

func (i *InTypes) TableName() string {
	return "in_types"
}
