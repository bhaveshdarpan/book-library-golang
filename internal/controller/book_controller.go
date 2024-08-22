package controller

import (
	"com/github/book-go/internal/model"
	"com/github/book-go/internal/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type BookController struct {
	BookService *service.BookService
}

func NewBookController(bookService *service.BookService) *BookController {
	return &BookController{BookService: bookService}
}

func (bc *BookController) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, error := bc.BookService.GetAllBooks()
	if error != nil {
		http.Error(w, error.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(books)
}

func (bc *BookController) AddBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	json.NewDecoder(r.Body).Decode(&book)
	addedBook, error := bc.BookService.AddBook(book)
	if error != nil {
		http.Error(w, error.Error(), http.StatusExpectationFailed)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(addedBook)
}

func (bc *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var book model.Book
	json.NewDecoder(r.Body).Decode(&book)

	err := bc.BookService.UpdateBook(id, book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (bc *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := bc.BookService.DeleteBook(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (bc *BookController) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	book, exists := bc.BookService.GetBook(id)
	if exists != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(book)
}
