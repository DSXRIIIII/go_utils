package main

import (
	"fmt"
	"net"
)

func main() {
	_, err := net.Dial("tcp", "116.198.246.11:5003")
	if err == nil {
		fmt.Println("success")
	} else {
		fmt.Println("false")
	}
}
