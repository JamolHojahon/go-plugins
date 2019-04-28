package main

import "fmt"

type gate struct{}

var Hi gate

func (g gate) Hello() {
	fmt.Println("helo ", Hi)
}

func main() {

}
