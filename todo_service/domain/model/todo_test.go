package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTodoTitle_String(t *testing.T) {
	tests := []struct {
		name string
		tr   TodoTitle
		want string
	}{
		{
			name: "success",
			tr:   "test",
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.tr.String())
		})
	}
}
