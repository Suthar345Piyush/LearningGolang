// struct in go

// package main

// import "fmt"

// type struct1 struct {
// 	x int
// 	y int
// }

// func main() {
// 	fmt.Println(struct1{1, 2})
// }

package main

import "fmt"

type vertex1 struct {
	x int
	y int
}

var (
	v1 = vertex1{1, 2}
	v2 = vertex1{x: 1}
	v3 = vertex1{}
	p  = &vertex1{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
