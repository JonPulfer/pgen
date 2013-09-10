// Package pgen provides a password generator
package main

import (
	"fmt"
	"math/rand"
	"time"
	"flag"
)

var pwdlen int
func init() {
	flag.IntVar(&pwdlen, "pwdlen", 12, "Desired length of password to generate")
	flag.Parse()
}

// Generate a random string.
func randomString (l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(33,126))
	}
	return string(bytes)
}

func randInt(min int , max int) int {
	    return min + rand.Intn(max-min)
}

func main () {
	rand.Seed( time.Now().UTC().UnixNano())
	fmt.Println(randomString(pwdlen))
}
