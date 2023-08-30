package models

import (
	"github.com/diiplexus/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

type Book struct {
	gorm.Model
	Name      string `gorm:"" json:"name"`
	Author    string `json:"" json:"author"`
	Publisher string `json:"" json:"publisher"`
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(id int64) (Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", id).Find(&book)
	return book, db
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(&book)
	return book
}
