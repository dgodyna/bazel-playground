package sub

import (
	"github.com/dgodyna/bazel-playground/pkg/entity"
	"reflect"
	"testing"
	"time"
)

func TestNewSubEntity(t *testing.T) {
	type args struct {
		name   string
		parent *entity.Entity
	}

	parent := entity.NewEntity("Test")

	tests := []struct {
		name string
		args args
		want *SubEntity
	}{
		{
			name: "Simple Test",
			args: args{
				name:   "Test",
				parent: parent,
			},
			want: &SubEntity{
				Entity: entity.Entity{Name: "Test"},
				parent: parent,
			},
		},
	}
	for _, tt := range tests {
		tt := tt

		// sleep to check how bazel will solve deps update
		time.Sleep(5 * time.Second)

		t.Run(tt.name, func(t *testing.T) {
			if got := NewSubEntity(tt.args.name, tt.args.parent); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSubEntity() = %v, want %v", got, tt.want)
			}
		})
	}
}
