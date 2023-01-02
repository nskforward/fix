package fix

import "bytes"

func Dump(data []byte, separator byte) []byte {
	return bytes.Replace(data, []byte{SOH}, []byte{separator}, -1)
}
