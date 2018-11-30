package services

import (
	"errors"

	"github.com/fluk27/testgo/models"
)

//BookService is Bookservice
type BookService struct {
}

//FindAll list all data
func (BookService) FindAll() []models.Book {
	books := []models.Book{
		models.Book{
			ID:   "001",
			Name: "golang",
		},
		models.Book{
			ID:   "002",
			Name: "vue.js",
		},
	}
	return books
}
func (BookService) FindByID(id string) (*models.Book, error) {

	if id == "0" {
		return nil, errors.New("not found")
	}
	book := models.Book{
		ID:   id,
		Name: "php",
	}
	return &book, nil
}
