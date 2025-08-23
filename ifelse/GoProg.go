// if else statement in golang

package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}

package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 2, 20),
	)
}

// switch case statement in Go

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's sunday?")
	today := time.Now().Weekday()
	switch time.Sunday {
	case today + 0:
		fmt.Println("today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In next two days")
	case today + 3:
		fmt.Println("In some days")
	default:
		fmt.Println("Too far away.")
	}
}

//second switch case implementation

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")

	case t.Hour() < 17:
		fmt.Println("Good evening")

	default:
		fmt.Println("Good night")
	}
}
