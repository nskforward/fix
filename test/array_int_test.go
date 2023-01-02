package test

import (
	"github.com/nskforward/fix"
	"testing"
)

func TestArrayInt(t *testing.T) {
	var a fix.ArrayInt
	a.Append(1)
	a.Append(3)
	a.Append(5)
	a.Append(7)
	a.Append(9)
	a.Append(0)
	a.Append(2)
	a.Append(4)
	a.Append(6)
	a.Append(8)

	a.Remove(9)

	for i, n := range a.Items() {
		if i != n {
			t.Fatal("want", i, "got", n)
		}
	}
}
