package main

import (
	"fmt"

	"github.com/eltonlai2014/my-pkg/pkg/mylib"
	mylib2 "github.com/eltonlai2014/my-pkg/pkg/mylibex"
)

func main() {
	// test code for mylib
	result := mylib.Add(2, 3)
	fmt.Println(result) // 5

	fmt.Println(mylib2.Hello("aaa"))
}
