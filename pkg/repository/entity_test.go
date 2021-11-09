package repository

import (
	"github.com/dgodyna/bazel-playground/pkg/entity"
	"gotest.tools/assert"
	"gotest.tools/assert/cmp"
	"sync"
	"testing"
)

func Test_entityRepository_Create(t *testing.T) {
	type fields struct {
		mux   sync.RWMutex
		cache map[string]*entity.Entity
	}
	type args struct {
		entity *entity.Entity
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected map[string]*entity.Entity
	}{
		{
			name: "Sunny Day",
			fields: fields{
				cache: map[string]*entity.Entity{},
			},
			args: args{
				entity: entity.NewEntity("Test"),
			},
			expected: map[string]*entity.Entity{
				"Test": {Name: "Test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			rep := entityRepository{
				cache: tt.fields.cache,
			}

			rep.Create(tt.args.entity)
			assert.Assert(t, cmp.DeepEqual(tt.expected, rep.cache))

		})
	}
}
