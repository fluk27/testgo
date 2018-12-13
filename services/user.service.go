package services

import (
	"github.com/fluk27/testgo/database"
	"github.com/fluk27/testgo/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mitchellh/mapstructure"
)

type UserService struct {
}

func (UserService) FindAll() ([]models.User, error) {
	DB, err := database.ConnectDatabase()
	if err != nil {
		return nil, err
	}
	query, err := database.QueryString("SELECT * FROM memberppe", DB)
	if err != nil {
		return nil, err
	}
	var dataUsers []models.User
	users, err := database.ResultValue(query)
	if err != nil {
		return nil, err
	}
	err = mapstructure.Decode(users, &dataUsers)
	if err != nil {
		return nil, err
	}

	return dataUsers, nil
}

func (UserService) FindByID(id string) (*models.User, error) {
	DB, err := database.ConnectDatabase()
	if err != nil {
		return nil, err
	}
	query, err := database.QueryString("SELECT * FROM memberppe WHERE Idmem ="+id, DB)
	if err != nil {
		return nil, err
	}

	var user models.User
	result, err := database.ResultValue(query)
	if err != nil {
		return nil, err
	}
	err = mapstructure.Decode(result[0], &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
