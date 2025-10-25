package models

import "time"

type User struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

var Users = []User{
	{Id: 1, Username: "user1", Email: "user1@gmail.com", CreatedAt: time.Now()},
	{Id: 2, Username: "user2", Email: "user2@gmail.com", CreatedAt: time.Now()},
}
