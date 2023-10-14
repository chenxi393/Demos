package main

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Product struct {
	ID   uint   `gorm:"primarykey"`
	Code string `gorm:"column: code"`
	//Deleted gorm.DeletedAt  //实现软删
	Price int `gorm:"default: 18"`
}

func (p Product) TableName() string {
	return "products"
}

// GORM性能提高 对于创建 更新 删除 为了保证数据的完整性
// GORM会默认封装在事务运行 但会降低性能 可以使用
// SkipDefaultTansaction 关闭默认事务
// 使用 PrepareStmt 缓存预编译语句提高后续调用速度 性能大概提高35%

// GORM 提供了CURD Hook的能力
// Hook 是在创建之前 创建之后自动调用的函数
// 如果任何Hook返回错误 GORM将停止后续的操作并回滚事务
// 使用hook操作 默认会开启事务 保证内部的一致性 原子性
func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	if p.Price < 0 {
		return errors.New("can't save invaid data") //错误不要大写
	}
	return
}

func Transaction(db *gorm.DB) {
	tx := db.Begin() //开启事务
	// 事务一旦开启 就得用tx 而不是db
	if err := tx.Create(&Product{Code: "This is Transaction", Price: 1000}).Error; err != nil {
		tx.Rollback() //遇到错误 回滚
		return
	}
	if err := tx.Create(&Product{Code: "This is Transaction Test", Price: 1000}).Error; err != nil {
		tx.Rollback() //遇到错误 回滚
		return
	}

	tx.Commit() //容易漏写 Commit Rollback
	// 提交事务

	// 自动提交事务
	if err := db.Transaction(func(tx *gorm.DB) error { //推荐Tansaction
		if err := tx.Create(&Product{Code: "AutoCommit Test"}).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return
	}
}

func main() {
	db, err := gorm.Open(mysql.Open(
		"root:123456@tcp(localhost:3306)/mydb?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Create 一条
	res := db.Create(&Product{Code: "nihapo"})
	if res.Error != nil {
		fmt.Println(res.Error)
	}
	// Create 多条记录
	products := []*Product{{Code: "nihfsao", Price: 201}, {Code: "gfdsadfgd", Price: 401}, {Code: "sfdfd", Price: 601}, {Code: "gdsfgds", Price: 801}}

	//这里其实可以验证 为什么是指针 因为下面Create还要写回products
	// for _, p := range products {
	// 	fmt.Println(p.ID)
	// }

	res = db.Create(products)
	if res.Error != nil {
		fmt.Println(res.Error)
	}
	// 验证主键
	for _, p := range products {
		fmt.Println(p.ID)
	}
	// 最终的操作一定是在最后（一调用就执行） 条件一定在前
	// 处理冲突
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&Product{ID: 1, Code: "fsdf"})

	// First 查询数据 查询不到数据会返回错误
	var product Product
	//fmt.Println(product.ID)
	db.First(&product, 1) // 这里也会写入1到product
	//fmt.Println(product.ID)     打印1    // 根据整形主键查找
	//如果找到会把找到的值写入
	fmt.Println(product.Code)
	fmt.Println(product.ID)
	fmt.Println(product.Price)
	// 查找 price 字段值为 18 的记录 并且主键为1
	//SELECT * FROM `product` WHERE price = '18' AND `product`.`id` = 1 ORDER BY `product`.`id` LIMIT 1
	db.First(&product, "price = ?", "18")

	// Find 查询多条数据 查询不到 不会返回错误
	pppp := make([]*Product, 0)
	result := db.Where("price=100").Find(&pppp) // 这里为什么要用指针 上面不用
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)

	//IN
	result = db.Where("price IN ?", []uint{100, 200}).Find(&pppp)
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)
	// LIKE
	result = db.Where("price LIKE ?", "%20%").Find(&pppp)
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)

	// where里还可以放结构体 但是忽略零值 map不忽略0值 可以看例子 这里不写了
	// SELECT也有API   db.Select()

	// Update - 将 product 的 price 更新为 200 并且ID=111
	db.Model(&Product{}).Where("price > ? ", 100).Update("price", 500)
	// // Update - 更新多个字段
	db.Model(&Product{}).Updates(Product{Price: 200, Code: "更新多个选项"}) // 仅更新非零值字段
	// 这里不用Model（因为后面的Product也有表名）也可以 没用where会用model里的条件
	// 且通过Model获取表名（这是原则？若后面没有提供表名？？ 不会算上MOdel里的条件么）
	// TO: 这里还需要测试 到底有条件会不会算上前面的

	db.Model(&Product{}).Select("code").Updates(map[string]interface{}{"Price": 200, "Code": "SELECT只会更新code"})

	// 可以再在 update里放表达式
	db.Model(&Product{}).Where(1).Update("price", gorm.Expr("price*?+?", 1.5, 100))
	// // Delete - 删除 product
	db.Delete(&Product{}, "price >?", "100") // 物理删除

	// 软删 不真的删 而是标记DeletedAt 使用Unscoped
	//db.Unscoped() 查找被软删的数据
}
