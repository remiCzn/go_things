package cmd

import "time"

type Data struct {
	NextId int    `json:"next_id"`
	Tasks  []Task `json:"tasks"`
}

type Task struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created_at"`
	Done    time.Time `json:"done_at"`
}
