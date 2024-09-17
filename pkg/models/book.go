package models

import (
	"main/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	ISBN        string `json:"isbn" gorm:"column:isbn;"`
	Name        string `json:"name" gorm:"column:name;"`
	Author      string `json:"author" gorm:"column:author;"`
	Publication string `json:"publication" gorm:"column:publication;"`
}

func (Book) TableName() string { return "books" }

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(b)

	return b
}

func GetAllBooks() []Book {
	var books []Book

	db.Find(&books)
	return books
}

func GetBookById(id int64) *Book {
	var book Book

	db.Where("id=?", id).Find(&book)
	return &book
}

func UpdateBookById(b *Book, id int64) error {
	return db.Table((Book{}.TableName())).Where("id=?", id).Update(b).Error
}

func DeleteBookById(id int64) error {
	return db.Table((Book{}.TableName())).Where("id=?", id).Delete(nil).Error
}
