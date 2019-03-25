package reader

import (
	"bytes"
	"fmt"
	"strings"
)

type path []byte

func (p path) TruncateAtFinalSlash() {
	i := bytes.LastIndex(p, []byte("/"))
	if i >= 0 {
		p = p[0:i]
	}
}

func (p *path) ToUpper() {
	for i, b := range *p {
		if 'a' <= b && b <= 'z' {
			(*p)[i] = b + 'A' - 'a'
		}
	}
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var s []string
	for _, v := range ip {
		s = append(s, string(v))
	}
	return strings.Join(s, ".")
}

func main() {
	// ip := IPAddr{127, 0, 0, 1}
	// fmt.Printf("%v", ip)

	// a := []byte{1, 2, 3, 4}
	// fmt.Println(string(a[:]))

	// const nihongo = "彂"
	// fmt.Println(len(nihongo))
	// for i, w := 0, 0; i < len(nihongo); i += w {
	// 	runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
	// 	fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
	// 	w = width
	// }

	// pathName := path("/usr/bin/tso") // Conversion from string to path.
	// pathName.TruncateAtFinalSlash()
	// fmt.Printf("%s\n", pathName)
	// pathName.ToUpper()
	// fmt.Printf("%s\n", pathName)

	// var iBuffer [10]int
	// slice := iBuffer[0:0]
	// for i := 0; i < 10; i++ {
	// 	slice = Extend(slice, i)
	// 	fmt.Println(slice)
	// }

	// fmt.Println("\n", slice, len(slice), cap(slice))

	// slice = make([]int, 10, 15)
	// fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
	// newSlice := make([]int, len(slice), 2*cap(slice))
	// // for i := range slice {
	// // 	newSlice[i] = slice[i]
	// // }
	// copy(newSlice, slice)
	// slice = newSlice
	// fmt.Printf("slice: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))

	// slice = append(slice, 2)
	// slice = append(slice, 3)
	// fmt.Println(slice)
	// anotherSlice := Insert(slice, 1, 1)
	// fmt.Println(slice, anotherSlice, "\n")

	// s := make([]int, 10, 20)
	// for i := range s {
	// 	s[i] = i
	// }
	// fmt.Println(s)
	// sb := Insert(s, 5, 99)
	// fmt.Println(s, sb)

	// slice := make([]int, 0, 5)
	// for i := 0; i < 10; i++ {
	// 	slice = Extend(slice, i)
	// 	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice), cap(slice), slice)
	// 	fmt.Println("address of 0th element:", &slice[0])
	// }

	// s := []int{1, 2, 3}
	// s2 := []int{7, 8, 9}
	// fmt.Println(s)
	// s = Append(s, 4, 5, 6)
	// fmt.Println(s)
	// s = Append(s, s2...)
	// fmt.Println(s)

	// Create a couple of starter slices.
	// slice := []int{1, 2, 3}
	// slice2 := []int{55, 66, 77}
	// fmt.Println("Start slice: ", slice)
	// fmt.Println("Start slice2:", slice2)

	// // Add an item to a slice.
	// slice = append(slice, 4)
	// fmt.Println("Add one item:", slice)

	// // Add one slice to another.
	// slice = append(slice, slice2...)
	// fmt.Println("Add one slice:", slice)

	// // Make a copy of a slice (of int).
	// slice3 := append([]int(nil), slice...)
	// fmt.Println("Copy a slice:", slice3)

	// // Copy a slice to the end of itself.
	// fmt.Println("Before append to self:", slice)
	// slice = append(slice, slice...)
	// fmt.Println("After append to self:", slice)

	// p := path(nil)
	// fmt.Println(p)
	// fmt.Println([]int(nil))

	// sample := "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	// for i := 0; i < len(sample); i++ {
	// 	fmt.Printf("%q\n", sample[i])
	// }

	// sample := "\xbd\xb2=\xbc ⌘"
	// fmt.Printf("%+q\n", sample)

	// const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	// fmt.Println("Println:")
	// fmt.Println(sample)

	// fmt.Println("Byte loop:")
	// for i := 0; i < len(sample); i++ {
	// 	fmt.Printf("%q ", sample[i])
	// }
	// fmt.Printf("\n")

	// fmt.Println("Printf with %x:")
	// fmt.Printf("%x\n", sample)

	// fmt.Println("Printf with % x:")
	// fmt.Printf("% x\n", sample)

	// fmt.Println("Printf with %q:")
	// fmt.Printf("%q\n", sample)

	// fmt.Println("Printf with %+q:")
	// fmt.Printf("%+q\n", sample)

	const nihongo = "日本語"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
}

func Append(s []int, items ...int) []int {
	n := len(s)
	total := n + len(items)
	if total > cap(s) {
		newSize := total*3/2 + 1
		newSlice := make([]int, total, newSize)
		copy(newSlice, s)
		s = newSlice
	}
	s = s[:total]
	copy(s[n:], items)

	return s
}

func Extend(slice []int, element int) []int {
	n := len(slice)
	if n == cap(slice) {
		// Slice is full; must grow.
		// We double its size and add 1, so if the size is zero we still grow.
		newSlice := make([]int, len(slice), 2*len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

func Insert(slice []int, index, value int) []int {
	slice = slice[0 : len(slice)+1]
	copy(slice[index+1:], slice[index:])
	slice[index] = value
	// fmt.Println(slice)
	return slice
}
