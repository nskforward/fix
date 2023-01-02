package fix

import "sync/atomic"

type Sequence struct {
	value int64
}

func (s *Sequence) Next() int64 {
	return atomic.AddInt64(&s.value, 1)
}
