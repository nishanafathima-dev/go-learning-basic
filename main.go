// package main

// import "fmt"

// func main() {
//     fmt.Println("Hello Nishana!")
// }


// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// func main() {
// 	fmt.Println("My favorite number is", rand.Intn(5))
// }

// package main 
// import "fmt"

// func add(x int,y int,w int,n int)int{
//     return x+y+w+n
// }
// func main(){
//     fmt.Println(add(1,2,3,4))
// }
// package main 
// import "fmt"
// func add(x ,y,w int)int{
//     return x+y+w
// }
// func main(){
//     fmt.Println(add(12,23,34))
// }

// package main 
// import "fmt"

// func swap(x,y,z string)(string,string,string){
//     return y,x,z
// }
// func main(){
//     a,b,c:=swap("nishana","love","niyas")
//     fmt.Println(a,b,c)
// }
// package main 
// import "fmt"

// func split(sum int)(x,y int){
//     x= sum*10/4
//     y=sum-x
//     return 
// }
// func main(){
//     fmt.Println(split(17))
// }

// package main 
// import "fmt"
// var c,python,java bool
// func main(){
//     var i int
//     fmt.Println(i,c,python,java)
// }

// package main 
// import "fmt"

// func main(){
//     var i,j =1,2
//     var c,python,java = true,false,"no!"
//     fmt.Println(i,j,c,python,java)
// }

// "ignoreDeprecations": "6.0",
// package main 
// import "fmt"
// var name string = "Nishana" 
// func main(){
//     user:= "Nishana"   //:= is used for define type of variable and assign value to it and it can only be used inside a function
//     fmt.Println("Hello",user)
//     fmt.Println("Hello",name)
// }
// package main 
// import (
//     "fmt"
//     "math/cmplx"
// )
// var(
//     tobe bool = false
//     maxInt uint64 = 1<<64 -1 //uint64 is an unsigned integer type that can hold values from 0 to 2^64-1
//     z complex128 = cmplx.Sqrt(-5 + 12i) //complex128 is a complex number type that can hold values with a real part and an imaginary part, each of which is a float64. The cmplx.Sqrt function calculates the square root of a complex number, in this case, -5 + 12i.
// )
// func main(){
//     fmt.Printf("Type:%T Value:%v\n",tobe,tobe)
//     fmt.Printf("Type:%T Value:%v\n",maxInt,maxInt)
//     fmt.Printf("Type:%T Value:%v\n",z,z)
// }
// int8 //the int8 type is an integer type that can hold values from -128 to 127. It is a signed 8-bit integer, which means it can represent both positive and negative numbers. The int8 type is often used when you need to save memory and you know that the values you will be working with will fit within this range.

// int      → Whole numbers
// float64  → Decimal numbers
// string   → Text
// bool     → true/false
// byte     → File/binary data
// rune     → Characters

// package main 
// import "fmt"

// func main(){
//     var i int
//     var f float64
//     var b bool
//     var s string
//     fmt.Printf("Default value of int: %d\n", i)
//     fmt.Printf("Default value of float64: %f\n", f)
//     fmt.Printf("Default value of bool: %t\n", b)
//     fmt.Printf("Default value of string: %q\n", s)
// }

// package main 
// import "fmt"

// var age int = 30
// func main(){
//     fmt.Printf("My age is %f\n", float64(age)) //float64 is a floating-point number type that can hold values with a decimal point. By converting the integer variable age to float64, we can print it as a floating-point number with a decimal point, even though it will still represent the same value of 30.
//     fmt.Printf("My age is %d\n", age)
// }
package main 
import "fmt"

func main(){
    var name="nishana "
    fmt.Println("Hello",name)
}