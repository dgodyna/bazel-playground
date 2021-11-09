package sub

import "github.com/dgodyna/bazel-playground/pkg/entity"

type SubEntity struct {
	entity.Entity
	parent *entity.Entity
}

func NewSubEntity(name string, parent *entity.Entity) *SubEntity {
	return &SubEntity{
		Entity: entity.Entity{Name: name},
		parent: parent,
	}

}
