package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

func checkOutputContains(t *testing.T, received string, expected string) {
	if !strings.Contains(received, expected) {
		t.Fatalf("received '%s', does not contain %s", received, expected)
	}
}

func TestMainProgram(t *testing.T) {
	var stdoutBuf, stderrBuf bytes.Buffer
	l = log.New(os.Stdout, fmt.Sprintf("Test_%s ", APP), log.Ldate|log.Ltime|log.Lshortfile)
	l.Printf("🚀🚀 Starting App %s, v%s :", APP, Version)
	config, err := NewCliConfig(WithOutStream(&stdoutBuf), WithErrStream(&stderrBuf))
	if err != nil {
		t.Fatal("Error creating config:", err)
	}
	app([]string{"main", "alex", "golang", "error"}, config, l)
	output := stdoutBuf.String()
	if len(output) == 0 {
		t.Fatal("Expected output, got nothing")
	}
	checkOutputContains(t, output, "[   1] word 'main' is even")
	checkOutputContains(t, output, "[   2] word 'alex' is even")
	checkOutputContains(t, output, "[   3] word 'golang' is even")

	errors := stderrBuf.String()
	if len(errors) == 0 {
		t.Fatal("Expected errors, got nothing")
	}
	checkOutputContains(t, errors, "[   4] word 'error' is odd")

}
