package interfaces

// import "fmt"
type Person struct {
	Name    string
	Age     int
	Address string
}

func (p *Person) Greet() string {
	//return "Hello, " + p.Name
	return "Hello"
}

type Greeter interface {
	Greet() string
}

type Employee struct {
	Person     //Embedded struct
	EmployeeID int
}

/*
func main(){
    //
    var p *Person
    //p.Name = "John"
    var greeter Greeter

    if greeter == nil {
        fmt.Println("g is nil")
	} else {
        fmt.Println("g is not nil")
	}


    if p == nil {
        fmt.Println("p is nil")
	} else {
        fmt.Println("p is not nil")
	}
    fmt.Println("result:", p.Greet()) // Hello, John

    greeter = p
    res := greeter.Greet()

    fmt.Println("-result:", res) // Hello, John
    fmt.Println("---------", p) // Hello, John
    //fmt.Println(p.Greet()) // Hello, John

    // Type assertion (interface).
	gr, ok := greeter.(Greeter)
	if ok {
        fmt.Println("greeter is Greeter")
        fmt.Println(	gr.Greet())
	}
}
*/
