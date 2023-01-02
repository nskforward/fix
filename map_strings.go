package fix

type MapStrings struct {
	fields map[int][]string
	keys   ArrayInt
}

func (m *MapStrings) Get(field int) []string {
	if m.fields == nil {
		return nil
	}
	return m.fields[field]
}

func (m *MapStrings) GetAndRemove(field int) []string {
	if m.fields == nil {
		return nil
	}
	v, ok := m.fields[field]
	if !ok {
		return nil
	}
	delete(m.fields, field)
	m.keys.Remove(field)
	return v
}

func (m *MapStrings) Set(field int, value string) {
	if m.fields == nil {
		m.fields = map[int][]string{field: {value}}
		m.keys.Append(field)
		return
	}
	arr, ok := m.fields[field]
	if !ok {
		m.fields[field] = []string{value}
		m.keys.Append(field)
		return
	}
	m.fields[field] = append(arr, value)
}

func (m *MapStrings) Range(f func(field int, values []string) bool) {
	for _, k := range m.keys.Items() {
		if !f(k, m.fields[k]) {
			break
		}
	}
}

func (m *MapStrings) Reset() {
	for k := range m.fields {
		delete(m.fields, k)
	}
	m.keys.Reset()
}
