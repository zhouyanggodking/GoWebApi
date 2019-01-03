package models

import (
	"github.com/jinzhu/gorm"
	// "go-tutorials/model"
)

type User struct {
	gorm.Model  // ID, CreatedAt, UpdatedAt, DeletedAt
	Name string
	Age uint
	Products []Product
}