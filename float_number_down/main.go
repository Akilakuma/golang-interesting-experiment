package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	test1()
	test2()
}

func test1() {
	origin := "18.9"
	originToFloat, _ := strconv.ParseFloat(origin, 64)
	fmt.Println("原本的數字：")
	fmt.Println(originToFloat)
	fmt.Println()

	result1 := math.Trunc(originToFloat*1e2) * 1e-2
	fmt.Println("math.Trunc ==>")
	fmt.Println(result1)
	fmt.Println()

	s := fmt.Sprintf("%f", originToFloat)
	result2, _ := strconv.ParseFloat(s, 64)
	fmt.Println("sprintf to string than parse to float ==>")
	fmt.Println(result2)
	fmt.Println()

	fmt.Println("math.Floor with large number ==>")
	fmt.Println(math.Floor(originToFloat*100000000) / 100000000)
}

func test2() {
	fmt.Println()
	base := 10000000
	origin := "18.9"
	originToFloat, _ := strconv.ParseFloat(origin, 64)
	fmt.Println("原本的數字：")
	fmt.Println(originToFloat)

	b := originToFloat * float64(base)
	c := math.Trunc(b*1e2) * 1e-2 / float64(base)
	fmt.Println("right method ==>")
	fmt.Println(c)
	fmt.Println()

	fmt.Println("fail method ==>")
	g := math.Trunc(originToFloat*1e20) * 1e-20
	d := math.Trunc(originToFloat*1e10) * 1e-10
	e := math.Trunc(originToFloat*1e5) * 1e-5
	f := math.Trunc(originToFloat*1e2) * 1e-2

	fmt.Println(g)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
