package fix

func NewMessageLogout(s *Session) *Message {
	m := s.pool.Get()
	m.Reset()

	s.SetHeader(m, "5")
	return m
}
