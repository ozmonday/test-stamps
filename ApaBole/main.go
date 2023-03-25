package main

import (
	"fmt"
	"strings"
)

func ApaBole(n int) string {
	var data []string
	for i := 1; i <= n; i++ {
		if i%5 == 0 && i%3 == 0 {
			data = append(data, "ApaBole")
		} else if i%3 == 0 {
			data = append(data, "Apa")
		} else if i%5 == 0 {
			data = append(data, "Bole")
		} else {
			data = append(data, fmt.Sprintf("%d", i))
		}
	}
	return strings.Join(data, ", ")
}

func main() {
	fmt.Println(ApaBole(100))
}
