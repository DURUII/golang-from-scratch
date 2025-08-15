package main

import (
	"book_management/model"
	"fmt"
)

func main() {
	lib := model.Library{
		Name: "图书馆",
		Books: []*model.Book{
			{ID: 1, Title: "Go语言基础", Author: "张三", IsAvailable: true},
			{ID: 2, Title: "Python编程", Author: "李四", IsAvailable: false},
			{ID: 3, Title: "Java编程", Author: "王五", IsAvailable: true},
		},
		Magazines: []*model.Magazine{
			{ID: 1, Title: "科技日报", Issue: 1, IsAvailable: true},
			{ID: 2, Title: "生活周刊", Issue: 2, IsAvailable: false},
		},
	}

	lib.PrintAllAvailableItems()
	fmt.Println("-----------")
	lib.GetBookByID(1).Borrow()
	lib.PrintAllAvailableItems()
	fmt.Println("-----------")

	lib.GetMagazineByID(1).Borrow()
	lib.PrintAllAvailableItems()
	fmt.Println("-----------")

	lib.GetMagazineByID(1).Return()
	lib.PrintAllAvailableItems()
	fmt.Println("-----------")
}
