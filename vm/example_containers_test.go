package vm_test

import (
	"fmt"
	"log"

	"github.com/mattn/anko/vm"
)

func Example_vmArrays() {
	env := vm.NewEnv()

	err := env.Define("println", fmt.Println)
	if err != nil {
		log.Fatalf("Define error: %v\n", err)
	}

	script := `
a = [1, 2]
println(a)

a += 3
println(a)

a = []
// this automaticly appends to array
a[0] = 1
println(a)

println("")

a = []
// this would give an index out of range error
// a[1] = 1

a = [1, 2]
b = [3, 4]
c = a + b
println(c)

c = [1, 2] + [3, 4]
println(c)

println("")

c = [a] + b
println(c)

c = [a] + [b]
println(c)

c = [[1, 2]] + [[3, 4]]
println(c)
`

	_, err = env.Execute(script)
	if err != nil {
		log.Fatalf("execute error: %v\n", err)
	}

	// output:
	// [1 2]
	// [1 2 3]
	// [1]
	//
	// [1 2 3 4]
	// [1 2 3 4]
	//
	// [[1 2] 3 4]
	// [[1 2] [3 4]]
	// [[1 2] [3 4]]
}

func Example_vmModule() {
	env := vm.NewEnv()

	err := env.Define("println", fmt.Println)
	if err != nil {
		log.Fatalf("Define error: %v\n", err)
	}

	script := `
module rectangle {
	_length = 1
	_width = 1
	
	func setLength (length) {
		if length <= 0 {
			return
		}
		_length = length
	}
	
	func setWidth (width) {
		if width <= 0 {
			return
		}
		_width = width
	}
	
	func area () {
		return _length * _width
	}
	
	func perimeter () {
		return 2 * (_length + _width)
	}
 }

rectangle1 = rectangle

rectangle1.setLength(4)
rectangle1.setWidth(5)

println(rectangle1.area())
println(rectangle1.perimeter())

rectangle2 = rectangle

rectangle2.setLength(2)
rectangle2.setWidth(4)

println(rectangle2.area())
println(rectangle2.perimeter())
`

	_, err = env.Execute(script)
	if err != nil {
		log.Fatalf("execute error: %v\n", err)
	}

	// output:
	// 20
	// 18
	// 8
	// 12
}
