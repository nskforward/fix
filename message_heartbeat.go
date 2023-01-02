package fix

func NewMessageHeartbeat(s *Session, testReqID string) *Message {
	m := s.pool.Get()
	m.Reset()
	s.SetHeader(m, "0")
	m.SetString(112, testReqID) // random text
	return m
}
