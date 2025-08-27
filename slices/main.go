// slices in go , they are reference type

// slices in go , they are reference type

// package main

// import "fmt"

// func main() {
// 	primes := [6]int{2, 3, 4, 8, 2}

// 	var s []int = primes[1:4] // start:end
// 	fmt.Println(s)
// }

// package main

// import "fmt"

// func main() {
// 	names := [4]string{
// 		"Piyush",
// 		"suthar",
// 		"magic",
// 		"code",
// 	}

// 	fmt.Println(names)

// 	a := names[0:2]
// 	b := names[1:3]
// 	fmt.Println(a, b)

// 	b[0] = "XXX"
// 	fmt.Println(a, b)
// 	fmt.Println(names)
// }

// slices in broader view

package main

import "fmt"

func main() {
	// direct slice initialization without using array value , if value is present then it is array otherwise it is the slice
	q := []int{2, 3, 5, 10, 11}
	fmt.Println(q)

	r := []bool{true, false, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{10, false},
		{11, true},
	}
	fmt.Println(s)
}

// zero value of refrence types (slices) are nil
