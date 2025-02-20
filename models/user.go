package models

type User struct {
	Id        string `json:id`
	Name      string `json:name`
	UserName  string `json:username`
	Email     string `json:email`
	Password  string `json:password`
	CreatedAt string `json:created_at`
}
