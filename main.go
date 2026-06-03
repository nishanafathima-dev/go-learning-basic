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

package main 
import "fmt"

func add(x int,y int,w int,n int)int{
    return x+y+w+n
}
func main(){
    fmt.Println(add(1,2,3,4))
}