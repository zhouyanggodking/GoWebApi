package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"   // mysql dialect
)

// demo for ont2many relationship

type User struct {
	gorm.Model  // ID, CreatedAt, UpdatedAt, DeletedAt
	Name string
	Age uint
	Products []Product    
}

type Product struct {
	gorm.Model
	Number string
	UserID uint
}

func main() {
	db, err := gorm.Open("mysql", "root@tcp(10.10.24.108:3306)/godking?charset=utf8&parseTime&loc=Local")
	defer db.Close()

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print("db connected")
	}

	db.DropTableIfExists(&Product{})
	db.DropTableIfExists(&User{})

	// db.AutoMigrate(&User{})
	db.AutoMigrate(&Product{}, &User{})

	user := User {
		Name: "godking", 
		Age: 10,
	}
	user2 := User {Name: "oceansky", Age: 20}
	db.Create(&user)
	db.Create(&user2)

	// should add foreign key manually
	// this is neccesary to add foreign key in database
	db.Model(&Product{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	
	products := []Product{
		{
			UserID: user.ID,
			Number: "123456789",
		},
		{
			UserID: user2.ID,
			Number: "000000000",
		},
	}

	for _, product := range(products) {
		db.Create(&product)
	}
}

func init() {
	fmt.Println("init func in main package")
}