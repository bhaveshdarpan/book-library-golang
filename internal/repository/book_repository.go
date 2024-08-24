package repository

import (
	"com/github/book-go/internal/model"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type BookRepository interface {
	FindByID(id uuid.UUID) (*model.Book, error)
	GetAll() ([]model.Book, error)
	Add(book *model.Book) error
	Update(book *model.Book) error
	Delete(id uuid.UUID) error
}

type SqlxBookRepository struct {
	db *sqlx.DB
}

func NewBookRepository(db *sqlx.DB) *SqlxBookRepository {
	return &SqlxBookRepository{db: db}
}

func (repo *SqlxBookRepository) FindByID(id uuid.UUID) (*model.Book, error) {
	var book model.Book
	query := `SELECT * FROM BookBank WHERE id = $1`
	err := repo.db.Get(&book, query, id)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (repo *SqlxBookRepository) GetAll() ([]model.Book, error) {
	var books []model.Book
	query := `SELECT * FROM BookBank`
	err := repo.db.Select(&books, query)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (repo *SqlxBookRepository) Add(book *model.Book) error {
	query := `INSERT INTO BookBank (id, Book_Title, Author, Genre) VALUES ($1, $2, $3, $4)`
	_, err := repo.db.Exec(query, book.ID, book.Title, book.Author, book.Genre)
	return err
}

func (repo *SqlxBookRepository) Update(book *model.Book) error {
	query := `UPDATE BookBank SET Book_Title = $2, Author = $3, Genre = $4 WHERE id = $1`
	_, err := repo.db.Exec(query, book.ID, book.Title, book.Author, book.Genre)
	return err
}

func (repo *SqlxBookRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM BookBank WHERE id = $1`
	_, err := repo.db.Exec(query, id)
	return err
}
