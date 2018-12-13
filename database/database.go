package database

 import (
	"fmt"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectDatabase is result database
func ConnectDatabase() (*sql.DB, error) {

	db, err := sql.Open("mysql", "root:@/s.a.d")
	if err != nil {
		fmt.Println("error from connect",err)
		return nil, err
	}
	
	return db, nil
}
// QueryString is function result database
func QueryString(queryString string, db *sql.DB) (*sql.Rows, error) {

	rows, err := db.Query(queryString)
	if err != nil {
		fmt.Println("error from querystring",err)
		return nil, err
	}
	return rows, nil
}
func ResultColumns(db *sql.Rows)  ([]string,error){
	columns,err := db.Columns()
	if err != nil {
		fmt.Println("error from resultcolumns",err)
		return nil,err 
	} 
	

	return columns,err
}
func ResultValue(db *sql.Rows,name []string)  ([]string,error){
	resultValue :=make([]interface{},len(name))
	valueToFunction :=make([]string,len(name))
	ToFunction :=make([]string,0)
	fmt.Println("this is value of resultValue",len(resultValue))
	for i := range valueToFunction {
		resultValue[i] = &valueToFunction[i]
	}
	
	for db.Next(){
		err :=db.Scan(resultValue...)
		if err != nil {
			fmt.Println("error from next.scan",err)
			return nil,err 
		}
			for _, value := range valueToFunction {
				if value == "" {
					value ="555"
				}
				fmt.Println("this is  value of valueToFunction",valueToFunction)
			//	fmt.Println("this is value of ToFunction in array",ToFunction[15])
				fmt.Println("this is col =",value)
				ToFunction = append(ToFunction,value)	
				//fmt.Println("this is value in range =",ToFunction)
			}
			
			fmt.Println("this is resultValue",valueToFunction[0])
			//var value string
	//
	
		//return valueToFunction,nil
	}
	//fmt.Println("this is value of valueToFunction",valueToFunction)
	if err := db.Err();err != nil{
		return nil,err
	}

	fmt.Println("this is value of ToFunction",ToFunction,len(ToFunction))
	return ToFunction,nil
}