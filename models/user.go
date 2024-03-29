package models

import (
	"errors"

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

func GetAllUsers() (*[]User, error){
	var users []User
	sql := `
		SELECT id, email FROM users;
	`

	rows, err := db.DB.Query(sql)
	if err != nil{
		return nil, err
	}

	for rows.Next(){
		user := User{}
		err = rows.Scan(
			&user.ID,
			&user.Email,
		)
		if err != nil{
			return nil, err
		}

		users = append(users, user)
	}

	return &users, nil
}

func GetUserByID(id int64) (*User, error){
	var user User
	sql := `
		SELECT id, email FROM users WHERE id = $1;
	`

	row := db.DB.QueryRow(sql, id)
	
	err := row.Scan(
		&user.ID,
		&user.Email,
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

func (user *User) ValidateCredentials()  error{
	email := user.Email

	findUser, err := GetUserByEmail(email)

	if err != nil{
		return errors.New("Invalid email")
	}

	if !utils.CheckPasswordHash(user.Password, findUser.Password){
		return errors.New("Invalid password")
	}
	user.ID = findUser.ID
	
	return nil
}