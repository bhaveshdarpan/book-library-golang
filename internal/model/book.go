package model

import (
<<<<<<< Updated upstream
	"github.com/google/uuid"
)

<<<<<<< Updated upstream
=======
import (
	"github.com/google/uuid"
)

>>>>>>> Stashed changes
type Book struct {
	ID       uuid.UUID `db:"id"`
	Title    string    `db:"Book_Title"`
	Author   string    `db:"Author"`
	Genre    string    `db:"Genre"`
	Quantity int       `db:"quantity"`
}

const TableName = "BooksLibrary"
=======
	"github.com/pborman/uuid"
)

type Book struct {
	ID       uuid.UUID `gorm:"primaryKey;autoincrement" json:"id"`
	Title    string    `gorm:"column:Book_Title" json:"title"`
	Author   string    `gorm:"column:Author" json:"author"`
	Genre    string    `gorm:"column:Genre" json:"genre"`
	Quantity int       `gorm:"column:quantity" json:"quantity"`
}

func (Book) TableName() string {
	return "BooksLibrary"
}
>>>>>>> Stashed changes
