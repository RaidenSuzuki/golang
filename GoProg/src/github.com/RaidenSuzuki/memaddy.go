package main

import "fmt"

var (
    i int = 9
    hello string = "Hello world"
    pi float32 = 3.14
    c complex64 = 3+5i
)

func main() {
    fmt.Println("Hexadecimal address of i is: ", &i)
    fmt.Println("Hexadecimal address of hello is: ", &hello)
    fmt.Println("Hexadecimal address of pi is: ", &pi)
    fmt.Println("Hexadecimal address of c is: ", &c)
}