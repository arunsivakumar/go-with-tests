package main

import "fmt"

//Hello this is a test
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("worlds"))
}
