package models

import (
	"github.com/DeeGrant/golang-bookstore-management/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBook(Id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("Id=?", Id).Find(&book)
	return &book, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("Id=?", Id).Delete(&book)
	return book
}
