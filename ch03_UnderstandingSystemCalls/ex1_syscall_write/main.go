package main

import (
	"fmt"
	"log"
	"unsafe"

	"golang.org/x/sys/unix"
)

func main() {

	const msg = "👋👋 Hello World ! from %s\n"
	// The native way to print "Hello, World!" to stdout
	fmt.Printf(msg, "fmt.Printf")

	// The overly complicated way to print "Hello, World!" to stdout
	msgSysCall := fmt.Sprintf(msg, "unix.Syscall")
	r1, r2, err := unix.Syscall(unix.SYS_WRITE, 1,
		uintptr(unsafe.Pointer(&[]byte(msgSysCall)[0])),
		uintptr(len(msgSysCall)),
	)
	if err != 0 {
		log.Fatalf("💥💥 Error doing Syscall unix.SYS_WRITE : %v", err)
	}
	log.Printf("✅ ✅ Successfull call to Syscall unix.SYS_WRITE returned: r1=%d, r2=%d\n", r1, r2)
}
