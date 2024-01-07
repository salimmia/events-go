package models

import (
	"github.com/salimmia/events-go/db"
)

type User struct{
	ID			int64	`json:"id"`
	Email 		string	`json:"email"`
	Password 	string	`json:"password"`
}

var user User

func (u *User) Save() error{
	sql := `
		INSERT INTO users(email, password)
		VALUES($1, $2)
		RETURNING id;
	`

	var id int64
	err := db.DB.QueryRow(sql, u.Email, u.Password).Scan(&id)
	if err != nil{
		return err
	}
	u.ID = id

	return nil
}