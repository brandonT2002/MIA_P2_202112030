package structs

import "fmt"

type Group struct {
	Id   string
	Name string
}

func NewGroup(id, name string) *Group {
	return &Group{id, name}
}

func (g *Group) ToString() string {
	return fmt.Sprintf("Id: %s Name: %s", g.Id, g.Name)
}
