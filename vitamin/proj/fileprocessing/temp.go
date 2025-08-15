package main

// import (
// 	"fmt"
// 	"github.com/mailru/easyjson"
// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"
// 	"log"
// 	"os"
// 	"time"
// 	. "voc/models/db"
// 	jd "voc/models/json"
// )
// func main2() {
// 	tic := time.Now()
// 	// 连接数据库，创建 app.db 文件
// 	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("连接数据库失败:", err)
// 	}
// 	db.Exec("PRAGMA journal_mode = WAL;")
// 	db.Exec("PRAGMA synchronous = NORMAL;")

// 	// 自动迁移：创建空表
// 	if err := db.AutoMigrate(&Word{}, &Translation{}, &Phrase{}); err != nil {
// 		log.Fatal("迁移表结构失败:", err)
// 	}

// 	// 读取数据
// 	data, err := os.ReadFile("data/json/4-CET6-顺序.json")
// 	if err != nil {
// 		log.Fatal("读取 JSON 文件失败:", err)
// 	}

// 	var vocItems jd.VocItemList
// 	if err := easyjson.Unmarshal(data, &vocItems); err != nil {
// 		log.Fatal("解析 JSON 数据失败:", err)
// 	}

// 	// 遍历，插入数据表
// 	for _, item := range vocItems {
// 		// 插入单词到 dict_word
// 		word := Word{
// 			Word: item.Word,
// 		}

// 		if err := db.Create(&word).Error; err != nil {
// 			log.Fatal("插入单词失败:", err)
// 		}

// 		// 插入翻译到 dict_translation
// 		for _, def := range item.Translations {
// 			translation := Translation{
// 				WordID:      word.ID,
// 				Translation: def.Translation,
// 				Type:        def.Type,
// 			}
// 			if err := db.Create(&translation).Error; err != nil {
// 				log.Fatal("创建翻译失败:", err)
// 			}
// 		}

// 		// 插入短语到 dict_phrase
// 		for _, ph := range item.Phrases {
// 			phrase := Phrase{
// 				WordID:      word.ID,
// 				Phrase:      ph.Phrase,
// 				Translation: ph.Translation,
// 			}
// 			if err := db.Create(&phrase).Error; err != nil {
// 				log.Fatal("创建短语失败:", err)
// 			}
// 		}

// 		//fmt.Println("已插入单词:", word.Word)
// 	}
// 	fmt.Println(time.Since(tic))
// }
