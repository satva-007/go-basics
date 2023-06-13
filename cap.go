package main

import (
	"fmt"
	"strings"
)

func main() {
	var txt string
	fmt.Println("Enter a text")
	fmt.Scanln(&txt)

	result := cap(txt)
	fmt.Println(result)

}
func cap(txt2 string) string {
	return strings.ToUpper(txt2)
}
