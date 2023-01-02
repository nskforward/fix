package fix

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

const SOH = 0x01

type Message struct {
	msgType       string
	msgBegin      string
	msgBodyLength string
	msgCheckSum   string
	fields        MapStrings
	body          bytes.Buffer
	wrapped       bytes.Buffer
}

func (m *Message) GetMsgType() string {
	return m.msgType
}

func (m *Message) GetFieldValues(field int) []string {
	return m.fields.Get(field)
}

func (m *Message) SetString(field int, value string) {
	if field == 35 {
		m.msgType = value
		return
	}
	if field == 8 {
		m.msgBegin = value
		return
	}
	if field == 9 {
		m.msgBodyLength = value
		return
	}
	if field == 10 {
		m.msgCheckSum = value
		return
	}
	m.fields.Set(field, value)
}

func (m *Message) SetInt(field int, value int64) {
	m.SetString(field, strconv.FormatInt(value, 10))
}

func (m *Message) SetTime(field int, value time.Time) {
	m.SetString(field, value.UTC().Format("20060102-15:04:05.000"))
}

func (m *Message) Reset() {
	m.fields.Reset()
}

func (m *Message) Marshal() []byte {
	m.body.Reset()
	m.body.WriteString("35=")
	m.body.WriteString(m.msgType)
	m.body.WriteByte(SOH)
	m.fields.Range(func(field int, values []string) bool {
		m.body.WriteString(strconv.Itoa(field))
		m.body.WriteRune('=')
		m.body.WriteString(values[0])
		m.body.WriteByte(SOH)
		return true
	})
	m.wrapped.Reset()
	m.wrapped.WriteString("8=")
	if m.msgBegin != "" {
		m.wrapped.WriteString(m.msgBegin)
	} else {
		m.wrapped.WriteString("FIX.4.4")
	}
	m.wrapped.WriteByte(SOH)
	m.wrapped.WriteString("9=")
	m.wrapped.WriteString(strconv.Itoa(m.body.Len()))
	m.wrapped.WriteByte(SOH)
	m.wrapped.Write(m.body.Bytes())
	sum := 0
	for _, c := range m.wrapped.Bytes() {
		sum += int(c)
	}
	mod := sum % 256
	m.wrapped.WriteString("10=")
	m.wrapped.WriteString(fmt.Sprintf("%03d", mod))
	m.wrapped.WriteByte(SOH)
	return m.wrapped.Bytes()
}

func (m *Message) Unmarshal(data []byte, separator byte) error {
	var k, v bytes.Buffer
	equalSign := false
	var char byte

	for _, char = range data {
		if char == '=' {
			equalSign = true
			continue
		}

		if char == separator {
			if !equalSign {
				return fmt.Errorf("got separator before =")
			}

			field, err := strconv.Atoi(k.String())
			if err != nil {
				return err
			}

			m.SetString(field, v.String())

			k.Reset()
			v.Reset()
			equalSign = false
			continue
		}

		if !equalSign {
			k.WriteByte(char)
		} else {
			v.WriteByte(char)
		}
	}
	if char != separator {
		return fmt.Errorf("the last sign must be soh")
	}
	return nil
}
