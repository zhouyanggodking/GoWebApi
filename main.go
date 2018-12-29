package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"   // mysql dialect
)

func main() {
	db, err := gorm.Open("mysql", "root@tcp(10.10.24.108:3306)/godking?charset=utf8&parseTime&loc=Local")
	defer db.Close()

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print("db connected")
	}
}

func init() {
	fmt.Println("init func in main package")
}
