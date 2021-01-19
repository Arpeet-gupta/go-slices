package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := a[2:4]
	fmt.Printf("before -> s=%v\n", s)
	fmt.Printf("before -> a=%v\n", a)
	// fmt.Printf("before ")

	s = append(s, 50, 60, 70)
	fmt.Printf("after  -> s=%v\n", s)
	fmt.Printf("after  -> a=%v\n", a)
	fmt.Println("&a[2]= == $s[0] is ", &a[2] == &s[0])

	s = append(s, 80, 90, 100, 110)
	fmt.Printf("after  -> s=%v\n", s)
	fmt.Printf("after  -> a=%v\n", a)
	fmt.Println("&a[2]= == $s[0] is ", &a[2] == &s[0])
	// newS := append(s, 55, 66)
	// fmt.Printf("s=%v, newS=%v\n", s, newS)
	// fmt.Printf("len=%d, cap=%d\n", len(newS), cap(newS))
	// fmt.Printf("a=%v", a)
}
