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
	}

	host := os.Getenv("DB_HOST")
   	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
   	user := os.Getenv("DB_USER")
   	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	log.Println(user, password, host, port, dbname)
	dsn := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, dbname)

   	db, errSql := sql.Open("postgres", dsn)

   	if errSql != nil {
    	fmt.Println("There is an error while connecting to the database ", err)
      	panic(err)
   	}else {
      	DB = db
      	fmt.Println("Successfully connected to database!")
	}

	createTables()
}

func createTables(){
	createEventTables := `
		CREATE TABLE IF NOT EXISTS events(
			id BIGSERIAL PRIMARY KEY,
			name varchar(100) NOT NULL,
			description varchar(1000) NOT NULL,
			location varchar(50) NOT NULL,
			user_id BIGINT
		);
	`

	_, err := DB.Exec(createEventTables)

	if err != nil{
		panic("Could not create event table.")
	}
}