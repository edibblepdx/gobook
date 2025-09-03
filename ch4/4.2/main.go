// Exercise 4.2: Write a command line program that prints the SHA256 hash of its
// standard input by default, but supports command line flags to print the SHA384
// or SHA512 hash instead.

package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var hash string

func init() {
	flag.StringVar(&hash, "h", "sha256", "one of {sha256, sha384, sha512}")

	if flag.Parse(); f[hash] == nil {
		fmt.Fprintln(os.Stderr, "Invalid hash type")
		os.Exit(1)
	}
}

type hf = func([]byte) []byte

var f = map[string]hf{
	"sha256": func(s []byte) []byte {
		h := sha256.Sum256(s)
		return h[:]
	},
	"sha384": func(s []byte) []byte {
		h := sha512.Sum384(s)
		return h[:]
	},
	"sha512": func(s []byte) []byte {
		h := sha512.Sum512(s)
		return h[:]
	},
}

func main() {
	var reader = bufio.NewReader(os.Stdin)

	fmt.Println("\"quit\" to exit")

	for {
		msg, _ := reader.ReadString('\n')
		msg = msg[:len(msg)-1] // strip newline
		if msg == "quit" {
			break
		}

		fmt.Printf("%x\n", f[hash]([]byte(msg)))
	}
}
