package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Start a new process
	cmd := exec.Command("ls", "-al", "-tr") // list files in reverse order of modification time
	cmd.Stdin = strings.NewReader(".")
	var out strings.Builder // store the output in case we want to analyse it
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("💥💥 Error doing Syscall unix.SYS_WRITE : %v\n", err)
		return
	}
	log.Printf("✅ ✅ Successfull call to  cmd.Run() with ls -al -tr . returned:\n %s\n", out.String())
	// Get the current process ID
	pid := os.Getpid()
	fmt.Println("Current process ID:", pid)
}
