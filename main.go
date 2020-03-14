package main

import "fmt"

//Declaration and Primitives
//Pointers
//Constants

/*
This
is
a
multi-line
comment
*/

func main() {
	//Hello World!!
	fmt.Println("Hello World!!")

	demoPrimitives()

}

func demoPrimitives() {
	//Declaration and Primitives
	//---------------------------
	var i int
	i = 42
	fmt.Println(i) //prints 42

	var f float32 = 3.14 //concise declaration
	fmt.Println(f)

	firstName := "John"
	fmt.Println(firstName) //implicit initialization

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)
	r, im := real(c), imag(c)
	fmt.Println(r, im)
}
