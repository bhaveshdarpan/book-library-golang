package model


import (
	"github.com/pborman/uuid"
	"gorm.io/gorm"
)

type Book struct{
	ID uuid         `gorm:"primaryKey;autoincrement" json:"id"` 
	Title string    `gorm:"column:Book_Title" json:"title"`
	Author string   `gorm:"column:Author" json:"author"` 
	Genre string    `gorm:"column:Genre" json:"genre"`
	Quantity int    `gorm:"column:quantity" json:"quantity"` 
}

func (Book) TableName() string{
	return "BooksLibrary"

}
