package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
)

func main() {
	hashDestination := "07811dc6c422334ce36a09ff5cd6fe71"
	a := 521
	b := 757
	fmt.Println("Start brute forcing...")

	for i := 1; i <= 10_000; i++ {
		h := md5.New()
		salt := a + b*i
		stoi := strconv.Itoa(salt)
		io.WriteString(h, stoi)
		hash := fmt.Sprintf("%x", h.Sum(nil))
		if hash == hashDestination {
			fmt.Println("done")
		}
	}
}
