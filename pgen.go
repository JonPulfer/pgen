// Package pgen provides a password generator
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var pwdlen int

func init() {
	flag.IntVar(&pwdlen, "pwdlen", 12, "Desired length of password "+
		"to generate")
	flag.Parse()
}

// Generate a random string.
func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
	REDO:
		c := byte(randInt(33, 126))
		if !checkRepeat(bytes, c) {
			bytes[i] = c
		} else {
			goto REDO
		}
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func checkRepeat(b []byte, c byte) bool {
	for _, n := range b {
		if n == c {
			return true
		}
	}
	return false
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(randomString(pwdlen))
}
