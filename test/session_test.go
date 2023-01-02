package test

import (
	"fmt"
	"github.com/nskforward/fix"
	"os"
	"testing"
)

func TestSession(t *testing.T) {
	addr := os.Getenv("FIX_ADDR")
	if addr == "" {
		t.Fatal("env FIX_ADDR is empty")
	}

	sender := os.Getenv("FIX_SENDER")
	if sender == "" {
		t.Fatal("env FIX_SENDER is empty")
	}

	target := os.Getenv("FIX_TARGET")
	if target == "" {
		t.Fatal("env FIX_TARGET is empty")
	}

	pass := os.Getenv("FIX_PASS")
	if pass == "" {
		t.Fatal("env FIX_PASS is empty")
	}

	sess, err := fix.NewSession(addr, sender, target, pass)
	if err != nil {
		t.Fatal(err)
	}

	msgLogin := sess.BuildLoginMessage()
	err = sess.Send(msgLogin)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("[debug] -->", string(fix.Dump(msgLogin.Marshal(), '|')))
	sess.Free(msgLogin)

	msg, err := sess.Read()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("[debug] <--", string(fix.Dump(msg.Marshal(), '|')))
}
