package main

import (
	"fmt"
	"strconv"

	"./valuechangepackage"
)

func main() {
	var a int = 1
	var b int32 = 2
	var c int64 = 3
	var d string = "999"
	var e float32 = 88.8
	var f float64 = 99.9
	var x string = "I Love Golang_"

	//a+b (int)
	z := a + int(b)
	fmt.Printf("a+b = %d\n", z)

	//a+b+c (int)
	fmt.Printf("a+b+c = %d\n", a+int(b)+int(c))

	//f/e (float)
	fmt.Printf("f/e = %f\n", float32(f)/e)

	//a+d (int)
	fmt.Printf("a+d = %d\n", a+strToInt(d))

	//x+a (string)
	fmt.Printf("x+a = %s\n", x+valuechangepackage.IntToString(a))
}

func strToInt(str string) int {
	value, _ := strconv.Atoi(str)

	return value
}
