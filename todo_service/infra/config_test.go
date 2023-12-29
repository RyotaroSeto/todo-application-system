package infra

import (
	"context"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	port := os.Getenv("PORT")
	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	DBName := os.Getenv("DB_NAME")
	DBSSLMode := os.Getenv("DB_SSM_MODE")
	DBTimeZone := os.Getenv("DB_TIME_ZONE")
	DBMaxConn := os.Getenv("DB_MAX_CONN")
	DBMaxIdle := os.Getenv("DB_MAX_IDLE")

	defer func() {
		os.Setenv("PORT", port)
		os.Setenv("DB_USER", DBUser)
		os.Setenv("DB_PASSWORD", DBPassword)
		os.Setenv("DB_HOST", DBHost)
		os.Setenv("DB_PORT", DBPort)
		os.Setenv("DB_NAME", DBName)
		os.Setenv("DB_SSM_MODE", DBSSLMode)
		os.Setenv("DB_TIME_ZONE", DBTimeZone)
		os.Setenv("DB_MAX_CONN", DBMaxConn)
		os.Setenv("DB_MAX_IDLE", DBMaxIdle)
	}()
	tests := []struct {
		name       string
		port       string
		DBUser     string
		DBPassword string
		DBHost     string
		DBPort     string
		DBName     string
		DBSSLMode  string
		DBTimeZone string
		DBMaxConn  int
		DBMaxIdle  int
		want       *Config
		assertion  assert.ErrorAssertionFunc
	}{
		{
			name:       "success",
			port:       "8080",
			DBUser:     "user",
			DBPassword: "password",
			DBHost:     "host",
			DBPort:     "5432",
			DBName:     "name",
			DBSSLMode:  "disable",
			DBTimeZone: "Asia/Tokyo",
			DBMaxConn:  0,
			DBMaxIdle:  0,
			want: &Config{
				Port:       "8080",
				DBUser:     "user",
				DBPassword: "password",
				DBHost:     "host",
				DBPort:     "5432",
				DBName:     "name",
				DBSSLMode:  "disable",
				DBTimeZone: "Asia/Tokyo",
				DBMaxConn:  0,
				DBMaxIdle:  0,
			},
			assertion: assert.NoError,
		},
		{
			name:       "failed",
			port:       "8080",
			DBUser:     "user",
			DBPassword: "password",
			DBHost:     "",
			DBPort:     "5432",
			DBName:     "name",
			DBSSLMode:  "disable",
			DBTimeZone: "Asia/Tokyo",
			DBMaxConn:  0,
			DBMaxIdle:  0,
			want:       nil,
			assertion:  assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Unsetenv("PORT")
			os.Unsetenv("DB_USER")
			os.Unsetenv("DB_PASSWORD")
			os.Unsetenv("DB_HOST")
			os.Unsetenv("DB_PORT")
			os.Unsetenv("DB_NAME")
			os.Unsetenv("DB_SSM_MODE")
			os.Unsetenv("DB_TIME_ZONE")
			os.Unsetenv("DB_MAX_CONN")
			os.Unsetenv("DB_MAX_IDLE")

			if tt.port != "" {
				os.Setenv("PORT", tt.port)
			}
			if tt.DBUser != "" {
				os.Setenv("DB_USER", tt.DBUser)
			}
			if tt.DBPassword != "" {
				os.Setenv("DB_PASSWORD", tt.DBPassword)
			}
			if tt.DBHost != "" {
				os.Setenv("DB_HOST", tt.DBHost)
			}
			if tt.DBPort != "" {
				os.Setenv("DB_PORT", tt.DBPort)
			}
			if tt.DBName != "" {
				os.Setenv("DB_NAME", tt.DBName)
			}
			if tt.DBSSLMode != "" {
				os.Setenv("DB_SSM_MODE", tt.DBSSLMode)
			}
			if tt.DBTimeZone != "" {
				os.Setenv("DB_TIME_ZONE", tt.DBTimeZone)
			}
			if tt.DBMaxConn != 0 {
				os.Setenv("DB_MAX_CONN", strconv.Itoa(tt.DBMaxConn))
			}
			if tt.DBMaxIdle != 0 {
				os.Setenv("DB_MAX_IDLE", strconv.Itoa(tt.DBMaxIdle))
			}
			got, err := LoadConfig(context.Background())
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetConfig(t *testing.T) {
	tests := []struct {
		name      string
		globalVar *Config
		want      *Config
	}{
		{
			name: "success",
			globalVar: &Config{
				Port:       "8080",
				DBUser:     "user",
				DBPassword: "password",
				DBHost:     "host",
				DBPort:     "5432",
				DBName:     "name",
				DBSSLMode:  "disable",
				DBTimeZone: "Asia/Tokyo",
				DBMaxConn:  0,
				DBMaxIdle:  0,
			},
			want: &Config{
				Port:       "8080",
				DBUser:     "user",
				DBPassword: "password",
				DBHost:     "host",
				DBPort:     "5432",
				DBName:     "name",
				DBSSLMode:  "disable",
				DBTimeZone: "Asia/Tokyo",
				DBMaxConn:  0,
				DBMaxIdle:  0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c = tt.globalVar
			assert.Equal(t, tt.want, GetConfig())
		})
	}
}

func TestConfig_DNS(t *testing.T) {
	type fields struct {
		Port       string
		DBUser     string
		DBPassword string
		DBHost     string
		DBPort     string
		DBName     string
		DBSSLMode  string
		DBTimeZone string
		DBMaxConn  int
		DBMaxIdle  int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				Port:       "8080",
				DBUser:     "user",
				DBPassword: "password",
				DBHost:     "host",
				DBPort:     "5432",
				DBName:     "name",
				DBSSLMode:  "disable",
				DBTimeZone: "Asia/Tokyo",
			},
			want: "host=host user=user password=password dbname=name port=5432 sslmode=disable TimeZone=Asia/Tokyo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Config{
				Port:       tt.fields.Port,
				DBUser:     tt.fields.DBUser,
				DBPassword: tt.fields.DBPassword,
				DBHost:     tt.fields.DBHost,
				DBPort:     tt.fields.DBPort,
				DBName:     tt.fields.DBName,
				DBSSLMode:  tt.fields.DBSSLMode,
				DBTimeZone: tt.fields.DBTimeZone,
				DBMaxConn:  tt.fields.DBMaxConn,
				DBMaxIdle:  tt.fields.DBMaxIdle,
			}
			assert.Equal(t, tt.want, r.DNS())
		})
	}
}
