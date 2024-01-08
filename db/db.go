package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(){
	log.Println("Connecting to Database...")
	err := godotenv.Load()

	if err != nil{
		log.Println("Error is occurred  on .env file please check")
		return
	}

	host := os.Getenv("DB_HOST")
   	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
   	user := os.Getenv("DB_USER")
   	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, dbname)

   	db, errSql := sql.Open("postgres", dsn)

   	if errSql != nil {
    	fmt.Println("There is an error while connecting to the database ", err)
      	panic(err)
   	}else {
      	DB = db
      	log.Println("Connected to the Database!!!")
	}

	createTables()
}

func createTables(){
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users(
			id BIGSERIAL NOT NULL PRIMARY KEY,
			email VARCHAR(100) NOT NULL UNIQUE,
			password VARCHAR(500) NOT NULL
		);
	`
	
	_, err := DB.Exec(createUsersTable)
	if err != nil{
		panic("Could not create users table.")
	}

	createEventTable := `
		CREATE TABLE IF NOT EXISTS events(
			id BIGSERIAL PRIMARY KEY,
			name varchar(100) NOT NULL,
			description varchar(1000) NOT NULL,
			location varchar(50) NOT NULL,
			date_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
			user_id BIGINT,
			FOREIGN KEY(user_id) REFERENCES users(id)
		);
	`

	_, err = DB.Exec(createEventTable)
	if err != nil{
		panic("Could not create event table.")
	}

	createRegistrationTable := `
		CREATE TABLE IF NOT EXISTS registrations(
			id BIGSERIAL PRIMARY KEY,
			event_id BIGINT NOT NULL,
			user_id BIGINT NOT NULL,
			FOREIGN KEY(event_id) REFERENCES events(id),
			FOREIGN KEY(user_id) REFERENCES users(id)
		);
	`

	_, err = DB.Exec(createRegistrationTable)
	if err != nil{
		panic("Could not create Regestration table.")
	}
}