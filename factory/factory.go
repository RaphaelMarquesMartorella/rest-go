package factory

import (
	"github.com/RaphaelMarquesMartorella/go-project/infrastructure/repository"
	"gorm.io/gorm"
)

func CreateBookRepository(db *gorm.DB) repository.BookRepository {
	factoryBookRepository := repository.BookRepository{Db: db}
	return factoryBookRepository
}
