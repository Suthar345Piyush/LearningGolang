/* defer keyword in golang -> it runs just before the ending of the function execution , behind the sence it uses the defer stack(first in last out) to handle the execution of the defer and other programs things */

//defer named statement always runs

package main

import "fmt"

func main() {
	fmt.Println("Counting") // printted first
	defer fmt.Println("Done")

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
