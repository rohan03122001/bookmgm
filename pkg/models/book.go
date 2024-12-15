package models

import (
	"github.com/jinzhu/gorm"
	"github.com/rohan03122001/bookmgm/pkg/config"
)

var db *gorm.DB

// Book represents the book entity in the database
type Book struct {
    gorm.Model
    Name        string `json:"name"`
    Author      string `json:"author"`
    Publication string `json:"publication"`
}

// init is called automatically when the package is imported
func init() {
    // Connect to database and perform auto-migration
    config.Connect()
    db = config.GetDB()
    db.AutoMigrate(&Book{})
}

// CreateBook saves a new book to the database
func (b *Book) CreateBook() *Book {
    db.Create(&b)
    return b
}

// GetAllBooks retrieves all books from the database
func GetAllBooks() []Book {
    var Books []Book
    db.Find(&Books)
    return Books
}

// GetBookById retrieves a specific book by ID
func GetBookById(Id int64) (*Book, *gorm.DB) {
    var GetBook Book
    db := db.Where("ID=?", Id).Find(&GetBook)
    return &GetBook, db
}

// DeleteBook removes a book from the database
func DeleteBook(Id int64) Book {
    var book Book
    db.Where("ID=?", Id).Delete(book)
    return book
}
