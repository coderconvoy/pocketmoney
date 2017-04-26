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
