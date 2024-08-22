package model

import (
	"github.com/google/uuid"
)

type Book struct {
	ID       uuid.UUID `db:"id"`
	Title    string    `db:"Book_Title"`
	Author   string    `db:"Author"`
	Genre    string    `db:"Genre"`
	Quantity int       `db:"quantity"`
}

const TableName = "BooksLibrary"
