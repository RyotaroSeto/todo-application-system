package infra

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBConnect(t *testing.T) {
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
				dsn: "host=host user=user password=password dbname=name port=5432 sslmode=disable TimeZone=Asia/Tokyo",
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
