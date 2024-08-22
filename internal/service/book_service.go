package service

import (
	"com/github/book-go/internal/model"
	"com/github/book-go/internal/repo"
)

type BookService struct {
	Book           *model.Book
	BookRepository *repo.BookRepository
}

func NewBookService(bookRepository *repo.BookRepository, book *model.Book) *BookService {
	return &BookService{BookRepository: bookRepository}
}

func (s *BookService) GetBook(id string) (*Book, error) {
	return s.BookRepository.GetBook(id)
}

func (s *BookService) GetAllBooks() ([]Book, error) {
	return s.BookRepository.GetAllBooks()
}

func (s *BookService) AddBook(book Book) (*Book, error) {
	
	return s.BookRepository.AddBook(book)
}

func (s *BookService) UpdateBook(id string, updatedBook Book) error {
	return s.BookRepository.UpdateBook(id, updatedBook)
}

func (s *BookService) DeleteBook(id string) error {
	return s.BookRepository.DeleteBook(id)
}
