package fix

type ArrayInt struct {
	items []int
}

func (a *ArrayInt) Append(value int) {
	if a.items == nil {
		a.items = []int{value}
		return
	}
	a.items = append(a.items, value)
}

func (a *ArrayInt) Remove(value int) {
	for i, k := range a.items {
		if k == value {
			a.items[i] = a.items[len(a.items)-1]
			a.items = a.items[:len(a.items)-1]
			break
		}
	}
}

func (a *ArrayInt) Items() []int {
	//sort.Ints(a.items)
	return a.items
}

func (a *ArrayInt) Reset() {
	a.items = a.items[:0]
}
