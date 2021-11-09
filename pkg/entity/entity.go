package entity

type Entity struct {
	Name string
}

func NewEntity(name string) *Entity {
	return &Entity{
		Name: name,
	}
}
