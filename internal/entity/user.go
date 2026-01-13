package entity

type User struct {
	ID       int     `db:"id"`
	Login    string  `db:"login"`
	Password string  `db:"password"`
	Token    *string `db:"token"`
}
