package models

import "time"

type Event struct{
	ID 			int			`json:"id"`
	Name 		string  	`json:"name"`
	Description string		`json:"description"`
	Location 	string		`json:"location"`
	DateTime 	time.Time	`json:"date_time"`
	UserId 		int			`json:"user_id"`
}

var events = []Event{}

func (e Event) Save(){
	events = append(events, e)
}

func GetEvent() []Event{
	return events
}