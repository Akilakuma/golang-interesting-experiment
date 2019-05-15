package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	// test1()
	// test2()
	// a := test3("123455.9823736", 2)
	// fmt.Println(a)
	// test4()
	b := Floor(123.0049999, 2)
	fmt.Println(b)
}

func test1() {
	origin := "19.08"
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
	base := 1000000000000
	origin := "19.08"
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

func test3(s string, length int) float64 {

	// 有包含小數點才處理
	if strings.Contains(s, ".") && length > 0 {
		// 有小數點才切割
		temp := strings.Split(s, ".")

		rightOfDot := temp[1]
		// 小數點右邊數字是否有超過指定長度
		lenOfRight := len(rightOfDot)
		if lenOfRight > length {
			var downNumber string
			// 小數點右邊數字 轉成 []string
			floatNum := strings.Split(rightOfDot, "")
			// 指定長度內的數字做字串相加
			for k, v := range floatNum {
				if k < length {
					downNumber += v
					k++
				}
			}
			// 小數點左邊數字 + . + 小數點右邊指定長度內的數字
			finalNum := temp[0] + "." + downNumber
			f, _ := strconv.ParseFloat(finalNum, 64)
			return f
		}

	}

	i, _ := strconv.ParseFloat(s, 64)
	return i
}

// 最終方案
func test4() {
	para := 18.9
	a := strconv.FormatFloat(para, 'f', 2, 64)
	fmt.Println(a)
	b, _ := strconv.ParseFloat(a, 64)
	fmt.Println(b)
}

// Floor 無條件捨去小數第N位
func Floor(number float64, N int) (floorNumber float64) {
	pow := math.Pow10(N)
	floorNumber = math.Floor(Round(number*pow, 4)) / pow
	return
}

// Ceil 無條件進位小數第N位
func Ceil(number float64, N int) (floorNumber float64) {
	pow := math.Pow10(N)
	floorNumber = math.Ceil(Round(number*pow, 4)) / pow
	return
}

// Round 四捨五入小數第N位
func Round(number float64, N int) float64 {
	var pow float64 = 1
	for i := 0; i < N; i++ {
		pow *= 10
	}
	return float64(int((number*pow)+0.5)) / pow
}
