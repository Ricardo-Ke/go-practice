package main

import (
	"errors"
	"fmt"
)

// var name = *flag.String("name", "everyone", "......")

// func init() {
// 	flag.Parse()
// }
// var ques = "why"

// String not
// func String() *string {
// 	name := "you"
// 	return &name
// }

type operate func(x, y int) int

func main() {
	// var name string
	// flag.StringVar(&name, "name", "everyone", "The greeting object")

	// var ques = "how"
	// fmt.Printf("%s\n", ques)
	// name := String("name", "everyone", "......")
	// Parse()
	// name := *flag.String("name", "everyone", "......")
	// flag.Parse()

	// var srcInt = int16(-255)
	// var desInt = int8(srcInt)
	// var num = 9
	// if _, ok := interface{}(num).(int); ok {
	// 	fmt.Print(num)
	// }
	// fmt.Print(srcInt)
	// fmt.Print(desInt)
	// fmt.Print(string([]byte{'\xe4', '\xbd', '\xa0', '\xe5', '\xa5', '\xbd'}))
	// fmt.Print(string([]rune{'\u4F60', '\u597D'}))

	// s1 := make([]int, 5)
	// s2 := make([]int, 5, 8)
	// fmt.Printf("%d %d %d %d %d %d\n", len(s1), cap(s1), s1, len(s2), cap(s2), s2)

	// s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// s4 := s3[3:6]
	// s4 = s4[0:cap(s4)]
	// fmt.Printf("len:%d cap:%d value:%d\n", len(s4), cap(s4), s4)

	// s6 := make([]int, 0)
	// for i := 1; i <= 5; i++ {
	// 	s6 = append(s6, i)
	// 	fmt.Printf("s6(%v) len:%d cap:%d\n", s6, len(s6), cap(s6))
	// }
	// s6e1 := append(s6, make([]int, 5)...)
	// fmt.Printf("s8a: len: %d, cap: %d\n", len(s6e1), cap(s6e1))

	// var l *list.List
	// l = list.New()
	// fmt.Print(l.Len())
	// fmt.Print(l.Front())
	// x := 1
	// for i := 1; i <= 10; i++ {
	// 	x = x * 1.06
	// }
	// fmt.Print(x)

	// m := map[string]int{
	// 	"one":   1,
	// 	"two":   2,
	// 	"three": 3,
	// }
	// k := "tw"
	// v, ok := m[k]
	// if ok {
	// 	fmt.Println(v)
	// } else {
	// 	fmt.Println(k)
	// }

	// m := map[interface{}]int{
	// 	"1":      1,
	// 	[]int{1}: 2,
	// 	3:        3,
	// }

	// fmt.Println(m["1"])

	// ch1 := make(chan int, 3)
	// ch1 <- 1
	// ch1 <- 2
	// ch1 <- 3
	// _ = <-ch1
	// _ = <-ch1
	// ele1 := <-ch1
	// fmt.Println(ele1)
	// chan1 := make(chan int, 2)
	// go func() {
	// 	for i := 0; i < 10000; i++ {
	// 		fmt.Printf("Sender: sending element %v...\n", i)
	// 		chan1 <- i
	// 	}
	// 	fmt.Println("Sender: close the channel...")
	// 	close(chan1)
	// }()

	// for {
	// 	elem1, ok := <-chan1
	// 	if !ok {
	// 		fmt.Println("Receiver: closed channel")
	// 		break
	// 	}
	// 	fmt.Printf("Receiver: received an element: %v\n", elem1)
	// }

	// fmt.Println("End.")

	fmt.Println(calculate(1, 2, plus))
}

func plus(x int, y int) int {
	return x + y
}

func calculate(x int, y int, op operate) (int, error) {
	if op != nil {
		return op(x, y), nil
	} else {
		return 0, errors.New("op is not operate type")
	}
}
