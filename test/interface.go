package main

import "fmt"

func tes(value ...interface{}) {
	fmt.Println(value[1])
}
func main() {
	tes("hehe", "napa pak?")
}
