package services

import (
	"fmt"
	"errors"

	"github.com/fluk27/testgo/models"
)

//BookService is Bookservice
type BookService struct {
	Name string
}

var bookname string
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
// FindByID is for Controller
func (BookService) FindByID(id string) (*models.Book, error) {

	if id == "0" {
		return nil, errors.New("not found")
	}

	var b1,b2 *models.Book
	b1 = &models.Book {
		ID: "1",
		Name: "AAAA",
	}



	fmt.Println(b1.Name)
	fmt.Println(b2.Name)

	//b2.Name = "CCCC"

	fmt.Println(b1.Name)
	fmt.Println(b2)

	b2 = b1
	//b2.Name = "DDDD"
	fmt.Println(b1.Name)
	fmt.Println(b2.Name)

	 


	book := models.Book{
		ID:   id,
		Name: "php",
	}
	return &book, nil
}
