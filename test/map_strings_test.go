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
	m.Set(6, "6")

	var buf bytes.Buffer
	m.Range(func(field int, values []string) bool {
		for _, v := range values {
			buf.WriteString(v)
		}
		return true
	})

	if !bytes.Equal(buf.Bytes(), []byte("9876654321")) {
		t.Fatalf("unexpected result:\nwnt: 9876654321\ngot: %s", buf.String())
	}
}
