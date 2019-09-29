package gorm_code

import "time"

type Product struct {
	Id         int64     `gorm:"primary_key"`
	CodeId     int8      `gorm:"column:code"`
	Code       Code      `gorm:"ForeignKey:Id;AssociationForeignKey:CodeId"`
	Price      float64   `gorm:"column:price"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}

func (this *Product) TableName() string {
	return "product"
}

type Code struct {
	Id   int8   `gorm:"primary_key"`
	Name string `gorm:"column:name;type:varchar(100)"`
}
