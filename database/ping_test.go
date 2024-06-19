package database

import (
	"context"
	"log"
	"os"
	"testing"
)

func TestPing(t *testing.T) {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("DSN not set")
	}

	err := Ping(context.Background(), dsn)
	if err != nil {
		t.Fatal(err)
	}
}
