package model

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Todo struct {
	ID     string `json:"id"`
	UserID *User  `json:"userId"`
	Name   string `json:"name"`
}
