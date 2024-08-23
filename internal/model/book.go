package model

import (
	"github.com/google/uuid"
)

type Book struct {
	ID     uuid.UUID `db:"id"`
	Title  string    `db:"book_title"`
	Author string    `db:"author"`
	Genre  string    `db:"genre"`
}

const TableName = "BookBank"
