package main

import "fmt"

// Printfln func
func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}
