package main

import "fmt"

func main() {
	var (
		sum int
		i   int = 1
	)

	for {
		if i > 10 {
			break
		}
		if i%2 != 0 {
			continue
		}
	}

	sum += i
	fmt.Println(i, "-->", sum)
	i++
}
