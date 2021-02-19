package main

import (
	"fmt"
	"time"
)

func main() {
	x := time.Now()
	y := x.String()
	//z, _ := time.Parse(time.ANSIC, y)
	d := time.Now().Unix()
	fmt.Printf("%s : %s : %d", x, y, d)
}
