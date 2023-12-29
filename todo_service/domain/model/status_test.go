package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStatus_TableName(t *testing.T) {
	type fields struct {
		ID        TodoStatus
		Name      string
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				ID:        TodoStatusUnspecified,
				Name:      "test",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: time.Now(),
			},
			want: "status",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Status{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
			}
			assert.Equal(t, tt.want, s.TableName())
		})
	}
}
