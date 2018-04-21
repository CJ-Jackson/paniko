package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	b := make([]byte, 16)
	rand.Read(b)

	fmt.Printf("%x", b)
}
