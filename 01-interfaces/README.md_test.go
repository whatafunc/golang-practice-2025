package interfaces

import "testing"

func Test_readmes(t *testing.T) {
	describe(42)      // Type: int, Value: 42
	describe("hello") // Type: string, Value: hello
	do(42)            // Type: int, Value: 42
	do("hello")       // Type: string, Value: hello
}
