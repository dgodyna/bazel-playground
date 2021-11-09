package entity

import (
	"reflect"
	"testing"
)

func TestNewEntity(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *Entity
	}{
		{
			name: "Simple Test",
			args: args{
				name: "Test",
			},
			want: &Entity{Name: "Test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEntity(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEntity() = %v, want %v", got, tt.want)
			}
		})
	}
}
