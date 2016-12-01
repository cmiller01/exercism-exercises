package hello

import "fmt"

const testVersion = 2

// HelloWorld will return a greeting
func HelloWorld(s string) string {
	if s == "" {
		return "Hello, World!"
	}
	return fmt.Sprintf("Hello, %s!", s)
}
