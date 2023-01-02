package test

import (
	"github.com/nskforward/fix"
	"log"
)

func send(sess *fix.Session, msg *fix.Message) (*fix.Message, error) {
	err := fix.SendMessage(sess, msg)
	if err != nil {
		return nil, err
	}
	log.Println("[debug] -->", string(fix.Dump(msg.Marshal(), '|')))
	sess.FreeMessage(msg)

	msg, err = fix.ReadMessage(sess)
	if err != nil {
		return nil, err
	}
	log.Println("[debug] <--", string(fix.Dump(msg.Marshal(), '|')))
	return msg, nil
}
