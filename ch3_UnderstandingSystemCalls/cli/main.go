package main

import (
	"fmt"
	"log"
	"os"
)

const APP = "MyGOCliApp"
const Version = "0.0.1"

func main() {
	words := os.Args[1:]
	l := log.New(os.Stdout, fmt.Sprintf("%s ", APP), log.Ldate|log.Ltime|log.Lshortfile)
	l.Printf("🚀🚀 Starting App %s, v%s :", APP, Version)
	if len(words) == 0 {
		nBytesWritten, err := fmt.Fprintln(os.Stderr, "No words provided.")
		if err != nil {
			l.Fatalf("💥💥 Error writing to stderr: %v", err)
		}
		l.Printf("## Successfull write of %d Bytes to stderr\n", nBytesWritten)
		os.Exit(1)
	}

	for i, w := range words {
		if len(w)%2 == 0 {
			nBytesWritten, err := fmt.Fprintf(os.Stdout, "[%4d] word '%s' is even\n", i+1, w)
			if err != nil {
				l.Fatalf("💥💥 Error writing to Stdout: %v", err)
			}
			l.Printf("✅ ✅ Successfull  write of %d Bytes to Stdout\n", nBytesWritten)
		} else {
			nBytesWritten, err := fmt.Fprintf(os.Stderr, "[%4d] word '%s' is odd\n", i+1, w)
			if err != nil {
				l.Fatalf("💥💥 Error writing to Stderr: %v", err)
			}
			l.Printf("✅ ✅ Successfull  write of %d Bytes to Stderr\n", nBytesWritten)
		}
	}
}
