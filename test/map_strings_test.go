package test

import (
	"bytes"
	"github.com/nskforward/fix"
	"testing"
)

func TestMapStrings(t *testing.T) {
	var m fix.MapStrings
	m.Set(9, "9")
	m.Set(8, "8")
	m.Set(7, "7")
	m.Set(6, "6")
	m.Set(5, "5")
	m.Set(4, "4")
	m.Set(3, "3")
	m.Set(2, "2")
	m.Set(1, "1")

	var buf bytes.Buffer
	buf.WriteString(m.GetAndRemove(3)[0])
	buf.WriteString(m.GetAndRemove(4)[0])
	buf.WriteString(m.GetAndRemove(5)[0])

	m.Range(func(field int, value []string) bool {
		buf.WriteString(value[0])
		return true
	})

	if !bytes.Equal(buf.Bytes(), []byte("345126789")) {
		t.Fatal("want 345126789, got", buf.String())
	}
}
