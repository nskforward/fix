package test

import (
	"github.com/nskforward/fix"
	"testing"
	"time"
)

func TestSession(t *testing.T) {
	sess, err := session()
	if err != nil {
		t.Fatal(err)
	}

	// LOGON
	msg := fix.NewMessageLogon(sess, time.Minute, true)
	resp, err := send(sess, msg)
	if err != nil {
		t.Fatal(err)
	}
	if resp.GetMsgType() != "A" {
		t.Fatal("response message must be type A")
	}
	sess.FreeMessage(resp)

	// HEARTBEAT
	msg = fix.NewMessageTest(sess, "heartbeat")
	resp, err = send(sess, msg)
	if err != nil {
		t.Fatal(err)
	}
	if resp.GetMsgType() != "0" {
		t.Fatal("response message must be type 0")
	}
	if resp.GetFieldValues(112)[0] != "heartbeat" {
		t.Fatal("heartbeat response field 112 must contain value 'heartbeat', got:", resp.GetFieldValues(112)[0])
	}
	sess.FreeMessage(resp)

	// LOGOUT
	msg = fix.NewMessageLogout(sess)
	resp, err = send(sess, msg)
	if err != nil {
		t.Fatal(err)
	}
	if resp.GetMsgType() != "5" {
		t.Fatal("response message must be type 5")
	}
	sess.FreeMessage(resp)

	sess.Close()
}
