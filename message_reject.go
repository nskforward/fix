package fix

func NewMessageReject(s *Session, rejectedMsgSeqNum, reason string) *Message {
	m := s.pool.Get()
	m.Reset()
	s.SetHeader(m, "3")
	m.SetString(45, rejectedMsgSeqNum)
	m.SetString(58, reason)
	return m
}
