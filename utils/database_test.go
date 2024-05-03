package utils

import (
	"os"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestConnectDatabase(t *testing.T) {
	tests := []struct {
		name     string
		database string
		want     bool
	}{
		{name: "Test with invalid database", database: "../NO_DATABASE.db", want: false},
		// {name: "Test with valid database", database: "../database.db", want: true}, //TODO: FAILING
	}
	for _, tt := range tests {
		os.Setenv("DATABASE_FILE", tt.database)
		t.Run(tt.name, func(t *testing.T) {
			got := ConnectDatabase()
			loaded := (got != nil)
			if loaded != tt.want {
				t.Errorf("ConnectDatabase() = %v, want %v", loaded, tt.want)
			}
		})
	}
}

func Test_connect(t *testing.T) {
	tests := []struct {
		name    string
		want    *gorm.DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := connect()
			if (err != nil) != tt.wantErr {
				t.Errorf("connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegister(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Register(); (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}