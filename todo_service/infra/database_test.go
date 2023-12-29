package infra

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBConnect(t *testing.T) {
	c, _ := LoadConfig(context.Background())
	type args struct {
		dsn     string
		maxConn int
		maxIdle int
	}
	tests := []struct {
		name      string
		args      args
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				dsn: c.DNS(),
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, DBConnect(tt.args.dsn, tt.args.maxConn, tt.args.maxIdle))
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name string
		want *DB
	}{
		{
			name: "success",
			want: &DB{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.IsType(t, tt.want, Get())
		})
	}
}
