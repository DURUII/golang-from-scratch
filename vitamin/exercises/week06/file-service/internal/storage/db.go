package storage

import (
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var db *gorm.DB

func init() {
	db = InitDB()
}

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}
	if err := db.AutoMigrate(&Item{}); err != nil {
		log.Fatal("迁移表结构失败:", err)
	}
	return db
}

func InsertItem(item Item) error {
	item.UUID = uuid.New().String()
	result := db.Create(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func ListItem(page, pageSize int) ([]Item, error) {
	var items []Item
	offset := (page - 1) * pageSize
	result := db.Where("is_delete = ?", 0).Offset(offset).Limit(pageSize).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

func FindItemByUUID(uuid string) (*Item, error) {
	var item Item
	result := db.Where("uuid = ? AND is_delete = ?", uuid, 0).First(&item)
	if result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}
