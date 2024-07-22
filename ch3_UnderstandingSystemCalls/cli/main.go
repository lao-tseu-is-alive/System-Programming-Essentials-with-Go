package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	words := os.Args[1:]
	if len(words) == 0 {
		nBytesWritten, err := fmt.Fprintln(os.Stderr, "No words provided.")
		if err != nil {
			log.Fatalf("💥💥 Error writing to stderr: %v", err)
		}
		log.Printf("Successfull write of %d Bytes to stderr\n", nBytesWritten)
		os.Exit(1)
	}

	for _, w := range words {
		if len(w)%2 == 0 {
			nBytesWritten, err := fmt.Fprintf(os.Stdout, "word %s is even\n", w)
			if err != nil {
				log.Fatalf("💥💥 Error writing to stderr: %v", err)
			}
			log.Printf("Successfull write of %d Bytes to Stdout\n", nBytesWritten)
		} else {
			nBytesWritten, err := fmt.Fprintf(os.Stderr, "word %s is odd\n", w)
			if err != nil {
				log.Fatalf("💥💥 Error writing to Stdout: %v", err)
			}
			log.Printf("Successfull write of %d Bytes to stdout\n", nBytesWritten)
		}
	}
}
