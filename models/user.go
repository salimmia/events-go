package models

import (
	"github.com/salimmia/events-go/db"
	"github.com/salimmia/events-go/utils"
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

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil{
		return err
	}

	var id int64
	err = db.DB.QueryRow(sql, u.Email, hashedPassword).Scan(&id)
	if err != nil{
		return err
	}
	u.ID = id

	return nil
}

func GetUserByID(id int64) (*User, error){
	var user User
	sql := `
		SELECT id, email, password FROM users WHERE email = $1;
	`

	row, err := db.DB.Query(sql, id)
	if err != nil{
		return nil, err
	}

	err = row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
	)
	if err != nil{
		return nil, err
	}

	return &user, nil
}

func GetUserByEmail(email string) (*User, error){
	var user User
	sql := `
		SELECT * FROM users WHERE email = $1;
	`

	row := db.DB.QueryRow(sql, email)
	

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
	)
	if err != nil{
		return nil, err
	}

	return &user, nil
}