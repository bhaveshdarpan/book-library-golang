package service

import (
	"com/github/book-go/internal/model"
	repo "com/github/book-go/internal/repository"

	"github.com/google/uuid"
)

type BookService struct {
	BookRepository repo.BookRepository
}

func NewBookService(bookRepository repo.BookRepository) *BookService {
	return &BookService{BookRepository: bookRepository}
}

func (s *BookService) GetBook(id string) (*model.Book, error) {
	bookID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.BookRepository.FindByID(bookID)
}

func (s *BookService) GetAllBooks() ([]model.Book, error) {
	return s.BookRepository.GetAll()
}

func (s *BookService) AddBook(book *model.Book) (*model.Book, error) {
	book.ID = uuid.New()
	err := s.BookRepository.Add(book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (s *BookService) UpdateBook(id string, updatedBook *model.Book) error {
	bookID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	updatedBook.ID = bookID
	return s.BookRepository.Update(updatedBook)
}

func (s *BookService) DeleteBook(id string) error {
	bookID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return s.BookRepository.Delete(bookID)
}
