package db

import (
	"testing"
)

func TestConnectMySQL(t *testing.T) {
	ConnectMySQL("sqlite3", "testing.db")

	err := MySQL.DB().Ping()
	if err != nil {
		t.Error(err)
	}
}

func TestAutoMigration(t *testing.T) {
	ConnectMySQL("sqlite3", "testing.db")

	type user struct {
		Name string
	}
	AutoMigration(user{})

	_, tableCheck := MySQL.CommonDB().Query("select * from user;")
	if tableCheck == nil {
		t.Error("Table is not created")
	}
}
