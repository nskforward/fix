package fix

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
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

func (s *Session) Send(m *Message) error {
	_, err := s.conn.Write(m.Marshal())
	return err
}

func (s *Session) Free(m *Message) {
	s.pool.Put(m)
}

func (s *Session) readField() (line, field, value []byte, err error) {
	line, err = s.r.ReadBytes(SOH)
	if err != nil {
		return
	}
	if len(line) == 0 {
		err = fmt.Errorf("empty response")
		return
	}
	pair := bytes.Split(line[:len(line)-1], []byte{'='})
	if len(pair) != 2 {
		err = fmt.Errorf("cannot decode fild-value pair: %s", Dump(line, '|'))
		return
	}
	field = pair[0]
	value = pair[1]
	return
}

func (s *Session) readAll(buf *bytes.Buffer, size int) error {
	for i := 0; i < size; i++ {
		char, err := s.r.ReadByte()
		if err != nil {
			return err
		}
		buf.WriteByte(char)
	}
	return nil
}

func (s *Session) Read() (*Message, error) {
	var buf bytes.Buffer

	line, f, v, err := s.readField()
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	buf.Write(line)
	if !bytes.Equal(f, []byte{'8'}) {
		return nil, fmt.Errorf("unexpected field: want 8, got %s", string(f))
	}

	line, f, v, err = s.readField()
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	buf.Write(line)
	if !bytes.Equal(f, []byte{'9'}) {
		return nil, fmt.Errorf("unexpected field: want 9, got %s", string(f))
	}

	length, err := strconv.Atoi(string(v))
	if err != nil {
		return nil, err
	}

	err = s.readAll(&buf, length)
	if err != nil {
		return nil, err
	}

	m := s.pool.Get()
	m.Reset()

	err = m.Unmarshal(buf.Bytes(), SOH)
	return m, err
}
