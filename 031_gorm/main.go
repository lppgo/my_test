package main

import (
	"gormdemo/db"

	"gorm.io/gorm"
)

func main() {

}

// Product 产品
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// Test 测试
func (Product) Test() {
	db := db.GetDB()
	// 自动迁移模式
	db.AutoMigrate(&Product{})

	// 创建
	db.Create(&Product{Code: "200", Price: 1000})

}
