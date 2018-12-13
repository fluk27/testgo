package services

import (
	"fmt"

	"github.com/fluk27/testgo/models"

	"github.com/fluk27/testgo/database"
	_ "github.com/go-sql-driver/mysql"
)

type UserService struct {
}

func (UserService) FindAll() ([]models.User, error) {

	DB, err := database.ConnectDatabase()
	if err != nil {
		return nil, err
	}
	users, err := database.QueryString("SELECT * FROM memberppe", DB)
	if err != nil {
		return nil, err
	}
	fieldName, err := database.ResultColumns(users)
	if err != nil {
		return nil, err
	}
	dataUsers := []models.User{}
	resultFromDatabase := make([]interface{}, len(fieldName))
	valueFromDatabase := make([]string, len(fieldName))
	for i := range valueFromDatabase {
		resultFromDatabase[i] = &valueFromDatabase[i]
	}
	for users.Next() {
		users.Scan(resultFromDatabase...)
		allUsers := models.User{
			ID:        valueFromDatabase[0],
			FirstName: valueFromDatabase[2],
			LastName:  valueFromDatabase[3],
		}
		fmt.Printf("this is type dataUsers %T\n", dataUsers)
		fmt.Println("this is value of dataUsers",dataUsers)
		dataUsers = append(dataUsers, allUsers)
	}
	return dataUsers, nil
}

func (UserService) FindByID(id string) (*models.User, error) {
	allUsers := models.User{}
	DB, err := database.ConnectDatabase()
	if err != nil {
		return nil, err
	}
	queryString := "SELECT * FROM memberppe WHERE Idmem ="+id
	users, err := database.QueryString(queryString, DB)
	if err != nil {
		return nil, err
	}
	fieldName, err := database.ResultColumns(users)
	if err != nil {
		return nil, err
	}
	//dataUsers := []models.User{}
	resultFromDatabase := make([]interface{}, len(fieldName))
	valueFromDatabase := make([]string, len(fieldName))
	for i := range valueFromDatabase {
		resultFromDatabase[i] = &valueFromDatabase[i]
	}
	for users.Next() {
		users.Scan(resultFromDatabase...)
		
			allUsers.ID=valueFromDatabase[0]
			allUsers.FirstName= valueFromDatabase[2]
			allUsers.LastName=  valueFromDatabase[3]
		
	//return allUsers,nil	
	}
	if len(allUsers.ID) == 0 {
		fmt.Println("this is value of allUsers.ID",len(allUsers.ID))
		return nil,err
	}
	return &allUsers,nil
}
