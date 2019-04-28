package main

import (
	"fmt"
	"go-plugins/gate"
	"plugin"
)

func main() {
	var io = "ff"
	var mod string
	switch io {
	case "oo":
		mod = "/rus/rus.so"
	case "ff":
		mod = "/timus/timus.go"
	}
	plus, err := plugin.Open(mod)
	if err != nil {
		fmt.Println(err)
	}

	hi, err := plus.Lookup("Hi")
	if err != nil {
		fmt.Println("PAPAPAPA", err)
	}
	var hi1 gate.Hi
	hi1, ok := hi.(gate.Hi)
	if !ok {
		fmt.Println("ERROR TYPE")
	}
	hi1.Hello()
}
