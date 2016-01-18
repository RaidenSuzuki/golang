package main

import "fmt"

//& is called the address operator.
//* is called the dereferencing operator.
// If Ptr is a pointer to Var then Ptr == &Var. And *Ptr == Var.
//* is a pointer

func main () {
	hello:= "Hello, minna-sama!"

	//declare a ptr variable to strings
	var hello_ptr *string
	
	//make it point to hello varriable
	//assigns address to it
	hello_ptr = &hello
	
	//an int var and pointer to it

	i := 6
	i_ptr := &i //i_ptr is of type *int

	fmt.Println("The string hello is:", hello)
	fmt.Println("The string pointed to by hello_ptr is: ", *hello_ptr)
	fmt.Println("The value of i is: ", i)
	fmt.Println("The value pointed to by i_ptr is: ", *i_ptr)
}