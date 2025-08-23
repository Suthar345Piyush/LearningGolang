// package main

// import "fmt"

// func main() {
// 	sum := 0
// 	for i := 0; i < 10; i++ {
// 		sum += i
// 	}
// 	fmt.Println(sum)
// }

// go has one loop statement , but we can implement all the other loops as well

// while and do while both

package main

import "fmt"

func main() {
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)
}

// we can run infinte loops in this and break them according to our need
