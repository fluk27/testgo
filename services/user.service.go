package services

import (
	"database/sql"
	"errors"
	"fmt"

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
	if err != nil || len(result) == 0 {
		errors.New("c")
		fmt.Println("error from ResultValue", err)
		return nil, err
	}

	err = mapstructure.Decode(result[0], &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (UserService) CheckUser(id string) (*sql.DB, error) {
	DB, err := database.ConnectDatabase()
	if err != nil {
		return nil, err
	}
	query, err := database.QueryString("SELECT Idmem FROM memberppe WHERE Idmem ="+id, DB)
	if err != nil {
		return nil, err
	}

	result, err := database.ResultValue(query)
	//fmt.Println("value of result   =",result)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		fmt.Println("value of result == 0  =", result)
		return DB, nil
	}

	return nil, err

}
func (U UserService) AddUsers(id string, firstName string, lastName string) (interface{}, error) {
	DB, err := database.ConnectDatabase()
	if err != nil {
		return nil, err
	}
	user := models.User{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
	}
	//aa,err := DB.Prepare(" INSERT INTO  memberppe (Idmem,Namemem,Lastnamemem)	VALUES('000232',sdfd,dsfdfd)")
	//bb,err :=aa.Query()
	userID, err := U.FindByID(id)
	fmt.Println("this is user ID before =", userID)
	if err != nil || userID == nil {
		fmt.Println("this is userID =", userID, err)
		return nil, err
	}

	query, err := database.QueryString(" INSERT INTO  memberppe (Idmem,Namemem,Lastnamemem)	VALUES('"+user.ID+"','"+user.FirstName+"','"+user.LastName+"')", DB)
	if err != nil {
		return nil, err
	}

	defer query.Close()

	return id, nil

}
