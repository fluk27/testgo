package services

import (
	"errors"
	"fmt"

	"github.com/fluk27/testgo/models"
)

//BookService is Bookservice
type BookService struct {
	Name string
}

var bookname string

//FindAll list all data
func (BookService) FindAll() *models.Book {
	books := models.Book{}
	fmt.Printf("this is value book %x", books)
	return &books
}

// FindByID is for Controller
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
