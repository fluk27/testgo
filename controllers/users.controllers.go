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
	user, errorFindByID := userService.FindByID(idUser)
	fmt.Println("this is errorFindByID",errorFindByID)
	msgError := fmt.Sprintf("%v", errorFindByID)
	//fmt.Println("this is user ID before =", user)
	if  msgError != "find not found" {
		fmt.Println("this is user =", user, err)
		return c.JSON(http.StatusNotFound, &map[string]interface{}{
			"code":    "ID_DUPLICATE",
			"message": "id duplicate ",
		})
	}
	/*user, err := userService.FindByID(idUser)
	if err != nil || user == nil {
		return c.JSON(http.StatusNotFound, &map[string]interface{}{
			"code":    "ID_DUPLICATE",
			"message": "id duplicate ",
		})
	}*/
	return c.JSON(http.StatusOK, user)
}
func (UserController) CheckEmpty(id string, firstName string, lastName string) ([]string, []interface{}) {
	err := make([]interface{}, 0)
	resultCheckEmpty := make([]string, 0)
//	resultCheckEmpty = append(resultCheckEmpty, id, firstName, lastName)
	if id == "" {
		errID := map[string]interface{}{
			"code":    "ID_IS_REQUEST",
			"message": "id is request ",
		}
		err = append(err, errID)

	}else {
		resultCheckEmpty = append(resultCheckEmpty, id)	
	}
	
	if firstName == "" {
		errfirstName := map[string]interface{}{
			"code":    "firstName_IS_REQUEST",
			"message": "firstName is request ",
		}
		err = append(err, errfirstName)
	}else {
		resultCheckEmpty = append(resultCheckEmpty,firstName)	
	}

	
	if lastName == "" {
		errlastName := map[string]interface{}{
			"code":    "lastName_IS_REQUEST",
			"message": "lastName is request ",
		}
		err = append(err, errlastName)

	}else {
		resultCheckEmpty = append(resultCheckEmpty,  lastName)	
	}
	
	
	//fmt.Println(len(resultCheckEmpty)," and ",len(err))
	if len(resultCheckEmpty) <=len(err) {
		fmt.Println(len(resultCheckEmpty)," and ",len(err))
		return nil,err
	}
	fmt.Println("this is resultCheckEmpty =",resultCheckEmpty)
	return resultCheckEmpty,nil

}
