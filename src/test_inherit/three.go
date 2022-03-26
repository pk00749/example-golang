package three

import "fmt"

type Base interface {
	protocol()
}

//--------------------------------
type A struct {
	Base
}

func (a A) protocol() {

}

//---------------------------------
type B struct {
	Base
}

func (b B) protocol() {

}

//---------------------------------
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
