package main

import (
	"fmt"
	"reflect"
)

func exemplo1() {
	x := 2

	d := reflect.ValueOf(&x).Elem()

	px := d.Addr().Interface().(*int)

	*px = 3

	fmt.Println(x)
}

func exemplo2() {
	a := 10

	var newValue *int = &a

	*newValue = 50

	fmt.Println(a)
}

func main() {
	exemplo1()
	exemplo2()
}
