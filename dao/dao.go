package dao

import (
	"GoGinBlog/until"
	"GoGinBlog/until/log"
	"fmt"
	"github.com/jinzhu/gorm"
	_"github.com/go-sql-driver/mysql"
)

var db *gorm.DB

//初始化数据库
func init(){
	db, err := gorm.Open(until.DbData, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		until.DbUser,
		until.DbPassWord,
		until.DbHost,
		until.DbName))
	if err != nil {
		log.Fatal("数据库连接异常: %v", err)
	}

	// Enable Logger, true:show detailed log  false:not show detailed log
	db.LogMode(true)

	// Debug a single operation, show detailed log for this operation
	//db.Debug().Where("name = ?", "jinzhu").First(&User{})

	// 启用Logger，显示详细日志
	db.LogMode(true)

	fmt.Println("数据库连接正常！")
}

func Test(){
	fmt.Println(123)
}


// Shutdown - close database connection
func Shutdown() error {
	log.Info("Closing database's connections")
	return db.Close()
}

// GetDb - get a database connection
func DB() *gorm.DB {
	return db
}

// 事务环绕
func Tx(txFunc func(tx *gorm.DB) error) (err error) {
	tx := db.Begin()
	if tx.Error != nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()

	err = txFunc(tx)
	return err
}




