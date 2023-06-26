package cmlabs_be

import (
	"fmt"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	for i := 1; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizz Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
