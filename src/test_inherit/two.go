package two

import "fmt"

type Base interface {
}

//-----------------
type A struct {
	Base
}

//-----------------
type B struct {
	Base
}

//-----------------
func test(a Base) {
	fmt.Println(a)
}

func main() {
	fmt.Println("Hello, playground")
	var a = A{}
	var b = B{}
	test(a)
	test(b)
}
