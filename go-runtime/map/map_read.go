package main

import "fmt"

func main() {
	m := map[int]int{1: 2, 3: 4}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
