package fix

func NewMessageSeqReset(s *Session, newSeqNum string) *Message {
	m := s.pool.Get()
	m.Reset()
	s.SetHeader(m, "4")
	m.SetString(36, newSeqNum)
	return m
}
