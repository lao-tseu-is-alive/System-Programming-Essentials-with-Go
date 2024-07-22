package main

import (
	"golang.org/x/sys/unix"
	"log"
)

func main() {
	write, err := unix.Write(1, []byte("\nHello, World! done with unix.Write\n"))
	if err != nil {
		log.Fatalf("💥💥 Error writing to stdout with unix.Write: %v", err)
	}
	log.Printf("✅✅ Successfull write of %d Bytes to stdout\n", write)
}
