package fix

import "time"

func (s *Session) BuildLoginMessage() *Message {
	m := s.pool.Get()
	m.Reset()

	m.SetString(35, "A")
	m.SetInt(34, s.seq.Next())
	m.SetString(49, s.sender)
	m.SetString(56, s.target)
	m.SetTime(52, time.Now())

	m.SetString(98, "0")
	m.SetString(108, "60")
	m.SetString(141, "Y")
	m.SetString(554, s.password)

	return m
}
