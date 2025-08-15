package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mailru/easyjson"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	dd "voc/models/db"
	jd "voc/models/json"
)

const batchSize = 750 // 每批处理数量

func main() {
	tic := time.Now()
	db := initDB()
	files := []struct {
		Path    string
		SrcName string
	}{
		{"data/json/3-CET4-顺序.json", "CET4"},
		{"data/json/4-CET6-顺序.json", "CET6"},
	}
	for _, f := range files {
		processSingleFile(db, f.Path, f.SrcName)
	}
	fmt.Println("导入完成，总耗时:", time.Since(tic).Seconds(), "秒")
}

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}
	// SQLite 优化
	db.Exec("PRAGMA journal_mode = WAL;")

	// 建表
	if err := db.AutoMigrate(&dd.Word{}, &dd.Translation{}, &dd.Phrase{}, &dd.Source{}); err != nil {
		log.Fatal("迁移表结构失败:", err)
	}
	return db
}

func processSingleFile(db *gorm.DB, path string, srcName string) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("读取 JSON 文件失败:", err)
	}

	var vocItems jd.VocItemList
	if err := easyjson.Unmarshal(data, &vocItems); err != nil {
		log.Fatal("解析 JSON 数据失败:", err)
	}

	source := dd.Source{SrcName: srcName}
	if err := db.FirstOrCreate(&source, dd.Source{SrcName: srcName}).Error; err != nil {
		log.Fatal("插入来源失败:", err)
	}

	if err := insertBatch(db, vocItems, source.ID); err != nil {
		log.Fatal("写入失败:", err)
	}
}

// 分批插入词条
func insertBatch(db *gorm.DB, vocItems jd.VocItemList, srcID uint) error {
	return db.Transaction(func(tx *gorm.DB) error {
		for start := 0; start < len(vocItems); start += batchSize {
			end := start + batchSize
			if end > len(vocItems) {
				end = len(vocItems)
			}
			batch := vocItems[start:end]

			var words []dd.Word
			for _, item := range batch {
				words = append(words, dd.Word{
					Word:  item.Word,
					SrcID: srcID,
				})
			}
			if err := tx.Create(&words).Error; err != nil {
				return err
			}

			var translations []dd.Translation
			var phrases []dd.Phrase
			for i, item := range batch {
				wid := words[i].ID
				for _, def := range item.Translations {
					translations = append(translations, dd.Translation{
						WordID:      wid,
						Translation: def.Translation,
						Type:        def.Type,
					})
				}
				for _, ph := range item.Phrases {
					phrases = append(phrases, dd.Phrase{
						WordID:      wid,
						Phrase:      ph.Phrase,
						Translation: ph.Translation,
					})
				}
			}
			if len(translations) > 0 {
				if err := tx.Create(&translations).Error; err != nil {
					return err
				}
			}
			if len(phrases) > 0 {
				if err := tx.Create(&phrases).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})
}
