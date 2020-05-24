package tonic

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbSource = "root:root@/tonic?charset=utf8&parseTime=True&loc=Local"
)

func TestOpenConnection(t *testing.T) {
	_, err := Open("mysql", dbSource)
	if err != nil {
		t.Error(err)
	}
}

func TestCloseConnection(t *testing.T) {
	db, err := Open("mysql", dbSource)
	if err != nil {
		t.Error(err)
	}
	if err = db.Close(); err != nil {
		t.Errorf("expected db connection to close without error, got %#v", err)
	}
}
