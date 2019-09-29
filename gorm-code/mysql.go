package gorm_code

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

var db *gorm.DB

func init() {
	dsn := fmt.Sprintf("root:@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	var err error
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal("连接数据库失败", err)
	}

	// 默认 model 的小写形式为表名
	// 若不设置此参数，则 model 需实现 TableName() 方法
	// 对于已经实现 TableName() 方法的 model, 此配置无影响
	db.SingularTable(true)
}

func Insert() error {
	product := &Product{}
	err := db.Model(product).Create(
		&Product{
			CodeId:  1,
			Price: 1.5,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		},
	).Error

	return err
}

func SingleQuery() (product *Product, err error) {
	product = &Product{}

	err = db.Model(product).Where(&Product{CodeId:1}).Find(product).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}

func MultiQuery() (products []*Product, err error) {
	products = make([]*Product, 0)

	err = db.Model(&Product{}).Where(&Product{CodeId:1}).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}