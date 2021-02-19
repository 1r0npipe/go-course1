package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Actions with pointers")
	a := 1
	b := 2
	c := unsafe.Pointer(&a)
	d := unsafe.Pointer(&b)
	fmt.Printf("This is a = %d, b = %d, c = %d, d = %d, &a = %p,  &b = %p, *c = %v, *d = %v\n", 
				a, b, c, d, &a, &b, *(*int)(c), *(*int)(d))
	temp := uintptr(d) - uintptr(c)
	fmt.Printf("Size of a = %d and b = %d\n", unsafe.Sizeof(a), unsafe.Sizeof(b))
	if (temp > 0) {
		fmt.Printf("Getting b via pointer to a = %v\n",*(*int)(unsafe.Pointer(uintptr(c) + temp)))
		fmt.Printf("Getting a via pointer to b = %v",*(*int)(unsafe.Pointer(uintptr(d) - temp)))
	}
}