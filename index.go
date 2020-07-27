/*
package main         // package declaration

import "fmt"         // importing the required packages

func main() {        // starts execution of program
	var x int= 12

	fmt.Printf("%T", x)
}

 we can declare variables in 3 ways:
var x int = 15
   OR
var x = 15
   OR
x := 15


// Mixed variable declaration
package main

import "fmt"

func main(){
	var a,b,c= 12, 15.6, "foo"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("Type of a is %T\n",a)
	fmt.Printf("Type of b is %T\n",b)
	fmt.Printf("Type of c is %T\n",c)
}

package main

import "fmt"

func main(){
    fmt.Println("Hello\tworld")
}        // Hello world


// constants declaration
package main

import "fmt"

func main(){
	const LENGTH int = 10     // It's a good practise to write constants in CAPITAL letters
	const WIDTH int =20
	var area int

	area= LENGTH * WIDTH
	fmt.Printf("Area is %d",area)
}

//  Incremental operator
package main

import "fmt"

func main(){
	var a int= 10
	a++
	fmt.Printf("value of a is %d",a)
}

// for loop implement
package main
import "fmt"
func main() {
	for a := 0; a < 10; a++ {
       fmt.Printf("a value is %d\n",a)
	}
}

// function declaration
package main

import "fmt"
func main() {
	var a int = 10
	var b int = 20
	var result int
	result = max(a, b)
	fmt.Printf("result is %d",result)

}
func max(a,b int) int{
	var res int
	res= a+b
	return res
}

// create Arrays
package main

import "fmt"

func main() {
	var cars = [2][2] int {{1, 2}, {2, 3}}
	fmt.Println("Car is", cars[0][1])
}

// Pointers
package main

import "fmt"

func main() {
	var a int = 10     // --> actual variable declaration
	var ip *int       // ---> pointer variable declaration
	ip = &a
	fmt.Printf("address of variable is %x\n", &a)
	fmt.Printf("address of pointerr is %x\n", ip)
	fmt.Printf("Value at pointer is %d\n", *ip)
}

// Nil Pointer in GO
package main

import "fmt"

func main() {
	var ip *int
	fmt.Printf("Address of variable is %x\n",ip)          // 0
}
*/

//
package main

func main() {

}
