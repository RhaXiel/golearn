package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	fmt.Println("Hello World!2")
	var age int = 5
	fmt.Println(age)
	var x float64 = 20.0

	y := 42
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("y is of type %T\n", y)

	fmt.Printf("Hello\tWorld!")

	var a int = 20 /* actual variable declaration */
	var ip *int    /* pointer variable declaration */

	ip = &a /* store address of a in pointer variable*/

	fmt.Printf("Address of a variable: %x\n", &a)

	/* address stored in pointer variable */
	fmt.Printf("Address stored in ip variable: %x\n", ip)

	/* access the value using the pointer */
	fmt.Printf("Value of *ip variable: %d\n", *ip)
}
