package main

import (
	"fmt"
)

func main() {
	i := 42         //var i int = 42
	f := float64(i) //var f float64 = float64(i)
	u := uint(f)    //var u uint = uint(f)

	fmt.Println(i, f, u)
}
