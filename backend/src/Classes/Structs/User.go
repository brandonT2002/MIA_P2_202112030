package structs

import "fmt"

type User struct {
	Id       string
	Group    string
	Name     string
	Password string
}

func NewUser(id, group, name, password string) *User {
	return &User{id, group, name, password}
}

func (u *User) ToString() string {
	return fmt.Sprintf("User(%s, %s, %s, %s)", u.Id, u.Group, u.Name, u.Password)
}
