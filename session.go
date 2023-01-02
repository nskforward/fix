package fix

import (
	"bufio"
	"time"
)

type Session struct {
	conn     *Conn
	seq      Sequence
	sender   string
	target   string
	password string
	pool     *Pool
	r        *bufio.Reader
}

func NewSession(addr, sender, target, password string) (*Session, error) {
	conn, err := NewConn("tcp", addr)
	return &Session{
		conn:     conn,
		sender:   sender,
		target:   target,
		password: password,
		pool:     NewPool(),
		r:        bufio.NewReader(conn),
	}, err
}

func NewSessionTLS(addr, sender, target, password string, cert []byte) (*Session, error) {
	conn, err := NewConnTLS("tcp", addr, cert)
	return &Session{
		conn:     conn,
		sender:   sender,
		target:   target,
		password: password,
		pool:     NewPool(),
		r:        bufio.NewReader(conn),
	}, err
}

func (s *Session) SetHeader(m *Message, msgType string) {
	m.SetString(35, msgType)
	m.SetInt(34, s.seq.Next())
	m.SetString(49, s.sender)
	m.SetString(56, s.target)
	m.SetTime(52, time.Now())
}

func (s *Session) FreeMessage(m *Message) {
	s.pool.Put(m)
}

func (s *Session) Close() {
	s.conn.Close()
}
