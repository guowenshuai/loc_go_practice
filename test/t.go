package main

import (
	"fmt"
	"time"
)

func main() {
	//t := time.Date(2017, time.November, 10, 23, 0, 0, 0, time.UTC)
	//fmt.Printf("Go launched at %s\n", t.Local())
	//x, y := 3, 4
	//fmt.Println(SumAndProduct(x, y))
	//r1 := Rectangle{12, 2}
	//fmt.Println("the area of r1 is: ", r1.area())
	tstart := time.Now()
	n := 13
	fmt.Printf("0-%d have %d nums 1\n", n, f(n))
	fmt.Println(" f(n) == n, n is ", testn())

	fmt.Println("execute time ", time.Now().Sub(tstart).String())
}

func SumAndProduct(A, B int) (int, int) {
	return A + B, A * B
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func testn() int {
	times := 0
	numbers := make(map[int]int)
	for n := 2; n < 200000000; n++ {
		s, ok := numbers[n-1]
		if ok {
			numbers[n] = s + checkOne(n)
		} else {
			numbers[n] = f(n)
		}

		if numbers[n] == n {
			fmt.Printf("f(%d) == %d\n", n, n)
			times = n
			break
		}
	}
	return times
}

func f(n int) int {
	r := 0

	for i := 0; i <= n; i++ {
		r += checkOne(i)
	}

	return r
}

func checkOne(n int) (r int) {
	for ; n >= 10; n /= 10 {
		l := n % 10
		if l == 1 {
			r++
		}
	}
	if n == 1 {
		r++
	}
	return
}
