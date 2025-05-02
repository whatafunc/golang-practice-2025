package interfaces

import (
	"fmt"
)

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("Unknown type %T!\n", v)
	}
}

func describe(i interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", i, i)
}
