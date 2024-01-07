package models

import (
	"time"

	"github.com/salimmia/events-go/db"
)

type Event struct{
	ID 			int64		`json:"id"`
	Name 		string  	`json:"name"`
	Description string		`json:"description"`
	Location 	string		`json:"location"`
	DateTime 	time.Time	`json:"date_time"`
	UserId 		int			`json:"user_id"`
}

var events = []Event{}

func (e Event) Save() error{
	insert := `
		INSERT INTO events(name, description, location, date_time)
		VALUES($1, $2, $3, $4);
	`

	stmt, err := db.DB.Prepare(insert)

	if err != nil{
		return err
	}

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime)

	if err != nil{
		return err
	}

	id, err := result.LastInsertId()

	e.ID = id
	
	return err
}

func GetEvent() []Event{
	return events
}