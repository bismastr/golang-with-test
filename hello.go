package main

import "fmt"

func main() {
	fmt.Println(Hello("Chris"))
}

func Hello(q string) string {
	if q == "" {
		return "Hello World"
	}

	return fmt.Sprintf("Hello, %s", q)
}
