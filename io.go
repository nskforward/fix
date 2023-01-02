package fix

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
)

func SendMessage(s *Session, m *Message) error {
	_, err := s.conn.Write(m.Marshal())
	return err
}

func ReadMessage(s *Session) (*Message, error) {
	var buf bytes.Buffer
	line, v, err := readField(s.r, []byte{'8'})
	if err != nil {
		return nil, err
	}
	buf.Write(line)
	line, v, err = readField(s.r, []byte{'9'})
	if err != nil {
		return nil, err
	}
	buf.Write(line)
	length, err := strconv.Atoi(string(v))
	if err != nil {
		return nil, err
	}
	err = readAll(s.r, &buf, length)
	if err != nil {
		return nil, err
	}
	m := s.pool.Get()
	m.Reset()
	err = m.Unmarshal(buf.Bytes(), SOH)
	return m, err
}

func readField(r *bufio.Reader, key []byte) (line, value []byte, err error) {
	for {
		line, err = r.ReadBytes(SOH)
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
		field := pair[0]
		value = pair[1]

		if bytes.Equal(field, key) {
			return
		}
	}
}

func readAll(r *bufio.Reader, buf *bytes.Buffer, size int) error {
	for i := 0; i < size; i++ {
		char, err := r.ReadByte()
		if err != nil {
			return err
		}
		buf.WriteByte(char)
	}
	return nil
}
