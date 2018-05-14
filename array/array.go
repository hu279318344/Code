package main

import (
	"fmt"
)

func array01() {
	//var a [2]int
	//var b [1]int
	a := [20]int{1, 2: 6}

	fmt.Println(a)

}

func array02() {

	a := [...]int{99: 1}
	//var p *[100]int = &a
	var p = &a
	fmt.Println(p)

}

func array03() {

	x, y := 1, 2
	a := [...]*int{&x, &y}
	fmt.Println(a)

}

func main() {
	array01()
	array02()
	array03()
}
