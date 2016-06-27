package main

import "fmt"

func main() {
	fruits := "木下博隆aa"

	for i, s := range fruits {
		fmt.Printf("fruits[%d]=%d\n", i, s)
	}
}
