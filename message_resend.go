package fix

func NewMessageResend(s *Session, resendMsgSeqNum string) *Message {
	m := s.pool.Get()
	m.Reset()
	s.SetHeader(m, "2")
	m.SetString(7, resendMsgSeqNum)  // begin range
	m.SetString(16, resendMsgSeqNum) // end range
	return m
}
