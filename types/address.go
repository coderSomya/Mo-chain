package types

import "fmt"

type Address [20]uint8

func (a Address) ToSlice() []byte {
	b := make([]byte, 20)

	for i := 0; i < 20; i++ {
		b[i] = a[i]
	}

	return b
}

func AddressFromBytes(b []byte) Address {

	if len(b) != 20 {
		msg := fmt.Sprintf("give bytes should be of lengh 20")
		panic(msg)
	}

	var value [20]uint8

	for i:=0; i < 20; i++{
		value[i] = b[i]
	}

	return Address(value)
}