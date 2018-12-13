package models

type User struct {
	ID        string `json:"id" mapstructure:"Idmem"`
	FirstName string `json:"first_name" mapstructure:"Namemem"`
	LastName  string `json:"last_name" mapstructure:"Lastnamemem"`
}
