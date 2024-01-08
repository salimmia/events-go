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
	UserId 		int64		`json:"user_id"`
}

var events = []Event{}

func (e *Event) Save() error{
	sql := `
		INSERT INTO events(name, description, location, date_time, user_id)
		VALUES($1, $2, $3, $4, $5)
		RETURNING id;
	`
	var id int64

	err := db.DB.QueryRow(sql, e.Name, e.Description, e.Location, e.DateTime, e.UserId).Scan(&id)

	if err != nil{
		return err
	}

	e.ID = id
	
	return err
}

func GetEvents() ([]Event, error){
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

func GetEventById(id int64) (*Event, error){
	var event Event

	query := `
		SELECT id, name, description, location, date_time, user_id FROM events
		WHERE id = $1;
	`

	row := db.DB.QueryRow(query, id)

	err := row.Scan(
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

	return &event, nil
}

func (event *Event) UpdateEvent() error{
	sql := `
		UPDATE events
		SET name = $1, description = $2, location = $3, date_time = $4
		WHERE id = $5;
	`

	stmt, err := db.DB.Prepare(sql)
	if err != nil{
		return err
	}

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

func (event *Event) DeleteEvent() error{
	sql := `
		DELETE FROM events
		WHERE id = $1;
	`

	stmt, err := db.DB.Prepare(sql)
	if err != nil{
		return err
	}

	_, err = stmt.Exec(event.ID)
	return err
}