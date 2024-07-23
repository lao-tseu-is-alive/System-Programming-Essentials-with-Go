package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const APP = "MyGOCliApp"
const Version = "0.0.2"

var l *log.Logger

type CliConfig struct {
	ErrStream, OutStream io.Writer
}

func app(words []string, cfg CliConfig, l *log.Logger) {
	for i, w := range words {
		if len(w)%2 == 0 {
			nBytesWritten, err := fmt.Fprintf(cfg.OutStream, "[%4d] word '%s' is even\n", i+1, w)
			if err != nil {
				l.Fatalf("💥💥 Error writing to cfg.OutStream: %v", err)
			}
			l.Printf("✅ ✅ Successfull write of %d Bytes to cfg.OutStream\n", nBytesWritten)
		} else {
			nBytesWritten, err := fmt.Fprintf(cfg.ErrStream, "[%4d] word '%s' is odd\n", i+1, w)
			if err != nil {
				l.Fatalf("💥💥 Error writing to cfg.ErrStream: %v", err)
			}
			l.Printf("✅ ✅ Successfull write of %d Bytes to cfg.ErrStream\n", nBytesWritten)
		}
	}
}

func NewCliConfig(opts ...Option) (CliConfig, error) {
	c := CliConfig{
		ErrStream: os.Stderr,
		OutStream: os.Stdout,
	}

	for _, opt := range opts {
		if err := opt(&c); err != nil {
			return CliConfig{}, err
		}
	}
	return c, nil
}

type Option func(*CliConfig) error

func WithErrStream(errStream io.Writer) Option {
	return func(c *CliConfig) error {
		c.ErrStream = errStream
		return nil
	}
}

func WithOutStream(outStream io.Writer) Option {
	return func(c *CliConfig) error {
		c.OutStream = outStream
		return nil
	}
}

func main() {
	words := os.Args[1:]
	l = log.New(os.Stdout, fmt.Sprintf("%s ", APP), log.Ldate|log.Ltime|log.Lshortfile)
	l.Printf("🚀🚀 Starting App %s, v%s :", APP, Version)
	if len(words) == 0 {
		nBytesWritten, err := fmt.Fprintln(os.Stderr, "No words provided.")
		if err != nil {
			l.Fatalf("💥💥 Error writing to stderr: %v", err)
		}
		l.Printf("## Successfull write of %d Bytes to stderr\n", nBytesWritten)
		os.Exit(1)
	}

	cfg, err := NewCliConfig()
	if err != nil {
		nBytesWritten, err := fmt.Fprintf(os.Stderr, "Error creating config: %v\n", err)
		if err != nil {
			l.Fatalf("💥💥 Error writing to stderr: %v", err)
		}
		l.Printf("✅ ✅ Successfull write of %d Bytes to stderr\n", nBytesWritten)
		os.Exit(1)
	}

	app(words, cfg, l)
}
