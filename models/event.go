package models

import (
	"log"
	"time"

	"github.com/salimmia/events-go/db"
)

type Event struct{
	ID 			int64		`json:"id"`
	Name 		string  	`json:"name"`
	Description string		`json:"description"`
	Location 	string		`json:"location"`
	DateTime 	time.Time	`json:"date_time"`
	UserId 		int64		`json:"user_id"`
}

var events = []Event{}

func (e *Event) Save() error{
	sql := `
		INSERT INTO events(name, description, location, date_time, user_id)
		VALUES($1, $2, $3, $4, $5)
		RETURNING id;
	`

	// stmt, err := db.DB.Prepare(sql)

	// if err != nil{
	// 	return err
	// }

	// result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)

	var id int64

	err := db.DB.QueryRow(sql, e.Name, e.Description, e.Location, e.DateTime, e.UserId).Scan(&id)

	if err != nil{
		return err
	}

	// id, err = result.LastInsertId()

	// if err != nil{
	// 	panic(err)
	// }

	log.Println(id)

	e.ID = id
	
	return err
}

func GetEvent() ([]Event, error){
	events := []Event{}

	query := `
		SELECT id, name, description, location, date_time, user_id FROM events;
	`

	rows, err := db.DB.Query(query)

	if err != nil{
		return events, err
	}

	defer rows.Close()

	for rows.Next(){
		event := Event{}

		err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.DateTime,
			&event.UserId,
		)

		if err != nil{
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}