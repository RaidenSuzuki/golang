package main

import "fmt"

//& is called the address operator.
//* is called the dereferencing operator.
// If Ptr is a pointer to Var then Ptr == &Var. And *Ptr == Var.
//* is a pointer

func main () {
	sum := 0
	var double_sum *int //a pointer to int
	for i:=0; i<10; i++ {
		sum+= i
	}
	double_sum = new(int) //allocate mem for an int and point to it
	*double_sum = sum*2 //use allocated mem, by dereferencing double_sum
	fmt.Println("The sum of numbers from 0 to 10 is: ", sum)
	fmt.Println("The double of this sum is: ", *double_sum)
}