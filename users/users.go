package users

type User struct {
	Id int `json:"id"`
	Name string `json:"name,omitempty"`
	SurName string `json:"surname,omitempty"`
	Age byte `json:"age,omitempty"`
}