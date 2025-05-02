package interfaces

import (
	//"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Greetings(t *testing.T) {
	var p *Person
	t.Run("check hello", func(t *testing.T) {
		expected := "Hello"
		require.Equal(t, expected, p.Greet())

	})

	t.Run("check greeter", func(t *testing.T) {
		var greeter Greeter
		greeter = p

		_, ok := greeter.(Greeter)
		require.True(t, ok, "var type is incorrect")
		//	fmt.Println("greeter is Greeter")
		/*	gr, ok := greeter.(Greeter)
				if ok {
			        fmt.Println("greeter is Greeter")
			        fmt.Println(	gr.Greet())
				}*/
	})

}
