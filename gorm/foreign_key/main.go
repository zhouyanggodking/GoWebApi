package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"   // mysql dialect
	"go-tutorials/models"
)

func main() {
	db, err := gorm.Open("mysql", "root@tcp(10.10.24.108:3306)/godking?charset=utf8&parseTime&loc=Local")
	defer db.Close()

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print("db connected")
	}

	db.DropTableIfExists(&models.Product{})
	db.DropTableIfExists(&models.User{})

	// db.AutoMigrate(&User{})
	db.AutoMigrate(&models.Product{}, &models.User{})

	user := models.User {
		Name: "godking", 
		Age: 10,
	}
	user2 := models.User {Name: "oceansky", Age: 20}
	db.Create(&user)
	db.Create(&user2)

	// should add foreign key manually
	// this is neccesary to add foreign key in database
	db.Model(&models.Product{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	
	products := []models.Product{
		{
			UserID: user.ID,
			Number: "123456789",
		},
		{
			UserID: user2.ID,
			Number: "000000000",
		},
	}

	db.Create(&products[0])
	db.Create(&products[1])
}

func init() {
	fmt.Println("init func in main package")
}