package fix

import (
	"sync"
)

type Pool struct {
	messages sync.Pool
}

func NewPool() *Pool {
	return &Pool{messages: sync.Pool{
		New: func() any {
			return new(Message)
		},
	}}
}

func (p *Pool) Get() *Message {
	return p.messages.Get().(*Message)
}

func (p *Pool) Put(m *Message) {
	p.messages.Put(m)
}
