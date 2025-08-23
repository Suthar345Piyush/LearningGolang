// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var x, y int = 2, 3
// 	var f float64 = math.Sqrt(float64(x*x + y*y))
// 	var z uint = uint(f)

// 	fmt.Println(x, y, z)
// }

package main

import "fmt"

const Pi = 3.14

func main() {
	const world = "Next"
	fmt.Println("hello", world)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true

	fmt.Println("Go rules?", Truth)
}
