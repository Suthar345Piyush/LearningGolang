// slices in go , they are reference type

package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 4, 8, 2}

	var s []int = primes[1:4] // start:end
	fmt.Println(s)
}
