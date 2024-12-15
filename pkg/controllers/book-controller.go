package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rohan03122001/bookmgm/pkg/models"
	"github.com/rohan03122001/bookmgm/pkg/utils"
)

// GetBook handles GET requests to retrieve all books
func GetBook(w http.ResponseWriter, r *http.Request) {
    // Fetch all books from database
    books := models.GetAllBooks()
    
    // Convert to JSON and send response
    res, _ := json.Marshal(books)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

// GetBookById handles GET requests to retrieve a specific book by ID
func GetBookById(w http.ResponseWriter, r *http.Request) {
    // Extract book ID from URL parameters
    vars := mux.Vars(r)
    bookId := vars["bookId"]
    ID, err := strconv.ParseInt(bookId, 0, 0)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // Fetch book details from database
    bookDetails, _ := models.GetBookById(ID)
    res, _ := json.Marshal(bookDetails)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

// CreateBook handles POST requests to create a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
    // Parse request body into Book struct
    newBook := &models.Book{}
    utils.ParseBody(r, newBook)
    
    // Create book in database
    book := newBook.CreateBook()
    res, _ := json.Marshal(book)
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write(res)
}

// UpdateBook handles PUT requests to update an existing book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
    // Parse request body and URL parameters
    updateBook := &models.Book{}
    utils.ParseBody(r, updateBook)
    vars := mux.Vars(r)
    
    // Convert book ID to int64
    bookId := vars["bookId"]
    id, err := strconv.ParseInt(bookId, 0, 0)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // Fetch existing book and update fields if provided
    bookDetails, db := models.GetBookById(id)
    if updateBook.Name != "" {
        bookDetails.Name = updateBook.Name
    }
    if updateBook.Author != "" {
        bookDetails.Author = updateBook.Author
    }
    if updateBook.Publication != "" {
        bookDetails.Publication = updateBook.Publication
    }

    // Save updated book to database
    db.Save(&bookDetails)
    res, _ := json.Marshal(bookDetails)
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

// DeleteBook handles DELETE requests to remove a book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
    // Extract and parse book ID
    vars := mux.Vars(r)
    bookId := vars["bookId"]
    id, err := strconv.ParseInt(bookId, 0, 0)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // Delete book from database
    book := models.DeleteBook(id)
    res, _ := json.Marshal(book)
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}
