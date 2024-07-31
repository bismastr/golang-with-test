package helloworld

import "fmt"

func main() {
	fmt.Println(Hello("Chris", "Spanish"))
}

func Hello(q string, lang string) string {
	spanish := "Spanish"
	english := "English"
	french := "French"
	prefix := "Hello"

	switch lang {
	case spanish:
		prefix = "Hola"
	case english:
		prefix = "Hello"
	case french:
		prefix = "Bonjour"
	}

	if q == "" {
		q = "World"
	}

	return fmt.Sprintf("%s, %s", prefix, q)
}
