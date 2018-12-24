package controllers

import (
	"fmt"
	"net/http"

	"github.com/fluk27/testgo/services"
	"github.com/labstack/echo"
)

//UserController is struct
type UserController struct {
}
type errFilede struct {
	Code string                                 `json:code`
	Message string								`json:message`
	Fields map[string]map[string]interface{}    `json:fields`
}

// List function
func (UserController) ListUserAll(c echo.Context) error {
	userService := &services.UserService{}
	users, err := userService.FindAll()
	if err != nil {
		return c.JSON(http.StatusOK, "I don't have value from database")
	}
	return c.JSON(http.StatusOK, users)
}

// Find function
func (UserController) FindByID(c echo.Context) error {
	userService := &services.UserService{}
	valueByFindID := c.Param("id")
	user, err := userService.FindByID(valueByFindID)
	if err != nil || user == nil {
		return c.JSON(http.StatusNotFound, &map[string]interface{}{
			"code":    "FIND_NOT_FOUND",
			"message": "find not found ",
		})
	}
	return c.JSON(http.StatusOK, user)
}
func (Uc UserController) Add(c echo.Context) error {
	userService := &services.UserService{}
	//valueIDUser := make([]string, 0)
	id :=c.FormValue("id")
	firstName :=c.FormValue("first_name")
	lastName :=c.FormValue("last_name")
	resultdata, empty := Uc.CheckEmpty(id, firstName, lastName)
	if empty != nil  {
		fmt.Println("this is empty = ",empty)
		return c.JSON(http.StatusOK, empty)
	}
	fmt.Println("this is resultdata = ",resultdata)
	
	resultID := resultdata[0]
	resultFirstName := resultdata[1]
	resultLastName := resultdata[2]

	/*	resultCheck, err := userService.CheckUser(resultID)
		if err != nil || resultCheck == nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    "ID_DUPLICATE",
				"message": "id duplicate ",
			})
		}*/
	dataUser, err := userService.AddUsers(resultID, resultFirstName, resultLastName)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "ID_DUPLICATE",
			"message": "id duplicate ",
		})

	}
	fmt.Println("this is dataUser =",dataUser)
	idUser := fmt.Sprintf("%v", dataUser)

	//msg := string(user)
	user, err := userService.FindByID(idUser)
	//fmt.Println("this is errorFindByID",errorFindByID)
	//msgError := fmt.Sprintf("%v", errorFindByID)
	//fmt.Println("this is user ID before =", user)
	if  err != nil {
		fmt.Println("this is user =", user, err)
		return c.JSON(http.StatusNotFound, &map[string]interface{}{
			"code":    "ID_DUPLICATE",
			"message": "id duplicate ",
		})
	}

	return c.JSON(http.StatusOK, user)
}
func (UserController) CheckEmpty(id string, firstName string, lastName string) ([]string, interface{}) {
	
	ED:=&errFilede{}
	ED.Code="INVALID_PARAMS"
		ED.Message="Invalid parameters"
		var dataID = map[string]map[string]interface{}{}
	resultCheckEmpty := make([]string, 0)

	
	if id == "" {
		
		dataID["ID"]=map[string]interface{}{
				"code": "REQUIRED",
				"message": "non zero value required",
			}
	
		
		
		
		ED.Fields=dataID
	
	//fmt.Println(dataID)
	}else {
		resultCheckEmpty = append(resultCheckEmpty, id)	
	}
	
	fmt.Println("afer id empty =",ED)
	if firstName == "" {

		dataID["fristName"]=map[string]interface{}{
			"code": "REQUIRED",
			"message": "non zero value required",
		}

	
	
	
	ED.Fields=dataID
		
	}else {
		resultCheckEmpty = append(resultCheckEmpty,firstName)	
	}
	
	
	if lastName == "" {
		dataID["lastName"]=map[string]interface{}{
			"code": "REQUIRED",
			"message": "non zero value required",
		}

	
	
	
	ED.Fields=dataID
	}else {
		resultCheckEmpty = append(resultCheckEmpty,  lastName)	
	}
	
	fmt.Println("this all stuct =",ED)
	//fmt.Println(len(resultCheckEmpty)," and ",len(err))
	if len(ED.Fields) != 0 {
		fmt.Println(len(resultCheckEmpty)," and ",len(ED.Fields))
		return nil,ED
	}
	//fmt.Println("this is resultCheckEmpty =",resultCheckEmpty)
	return resultCheckEmpty,nil

}
