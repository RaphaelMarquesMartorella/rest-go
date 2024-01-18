package repository

import (
	"github.com/RaphaelMarquesMartorella/go-project/domain/model"
	"gorm.io/gorm"
)

type BookRepository struct {
	Db *gorm.DB
}

func (b BookRepository) AddBook(book *model.Book) error {
	err := b.Db.Create(book).Error
	if err != nil {
		return err
	}
	return nil
}

func (b BookRepository) FindBookById(id string) (*model.Book, error) {
	var book model.Book
	result := b.Db.First(&book, "id = ?", id)
	return &book, result.Error
}
