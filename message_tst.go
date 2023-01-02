package fix

func NewMessageTest(s *Session, testReqID string) *Message {
	m := s.pool.Get()
	m.Reset()
	s.SetHeader(m, "1")
	m.SetString(112, testReqID) // random text
	return m
}
