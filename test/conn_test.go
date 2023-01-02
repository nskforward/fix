package test

import (
	"github.com/nskforward/fix"
	"os"
	"testing"
)

func TestConn(t *testing.T) {
	addr := os.Getenv("FIX_ADDR")
	if addr == "" {
		t.Fatal("env FIX_ADDR is empty")
	}

	conn, err := fix.NewConn("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
}
