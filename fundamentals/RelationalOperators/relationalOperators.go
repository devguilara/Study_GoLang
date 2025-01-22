package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings: ", "Banana" == "Banana") //true
	fmt.Println("!= ", 3 != 5)                     //true
	fmt.Println("> ", 3 > 5)                       // false
	fmt.Println("< ", 3 < 5)                       // true
	fmt.Println(">= ", 3 >= 5)                     // false
	fmt.Println("<= ", 3 <= 5)                     // true

	d1 := time.Unix(0, 0)
	d2 := time.Unix(1, 0)
	fmt.Println("Datas", d1 == d2)
	fmt.Println("Datas Equal?", d1.Equal(d2))

	type Person struct {
		Nome string
	}

	p1 := Person{"Joao"}
	p2 := Person{"Joao"}

	fmt.Println("Pessoas : ", p1 == p2)
}
