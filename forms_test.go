package main

import (
	"fmt"
	"testing"
)

func Test_AddAccount(t *testing.T) {
	ld := LoginData{
		Fmem: "dave",
	}
	tt := FormAddAccount(ld)
	fmt.Println(tt)

}

func Test_Chpass(t *testing.T) {
	for i := 1; i < 1000; i++ {
		tt := FormPassword()
		fmt.Println("--PASSWORD--")
		fmt.Println(tt)
	}
}
