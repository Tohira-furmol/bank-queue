package models

import "time"

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Queue struct {
	ID         uint      `json:"id"`
	UserID     uint      `json:"user_id"`
	Branch     string    `json:"branch"`
	Service    string    `json:"service"`
	TimeSlot   time.Time `json:"time_slot"`
	TicketNo   string    `json:"ticket_no"`
	CreatedAt  time.Time `json:"created_at"`
}
