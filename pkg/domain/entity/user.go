package entity

type User struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}
