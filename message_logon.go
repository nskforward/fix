package fix

import "time"

func NewMessageLogon(s *Session, heartbeatInterval time.Duration, resetSeq bool) *Message {
	m := s.pool.Get()
	m.Reset()

	resetSeqFlag := "N"
	if resetSeq {
		resetSeqFlag = "Y"
	}

	s.SetHeader(m, "A")
	m.SetString(98, "0") // encrypt method - always must be unencrypted
	m.SetInt(108, int64(heartbeatInterval.Seconds()))
	m.SetString(141, resetSeqFlag) // indicates both sides of a FIX session should reset sequence numbers
	m.SetString(554, s.password)

	return m
}
