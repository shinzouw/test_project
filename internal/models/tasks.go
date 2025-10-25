package models

import "time"

type Task struct {
	Id          int       `json:"id"`
	UserId      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	Duedate     time.Time `json:"due_date,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

var Tasks = []Task{
	{Id: 1, UserId: 1, Title: "Изучить Go", Description: "Изучить основы языка Go", Status: "pending", Priority: "High", CreatedAt: time.Now()},
	{Id: 2, UserId: 2, Title: "Изучить PostgreSQL", Description: "Изучить основы PostgreSQL", Status: "pending", Priority: "Medium", CreatedAt: time.Now()},
}
