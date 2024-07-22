package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type CliConfig struct {
	ErrStream, OutStream io.Writer
}

func app(words []string, cfg CliConfig) {
	for _, w := range words {
		if len(w)%2 == 0 {
			nBytesWritten, err := fmt.Fprintf(cfg.OutStream, "word %s is even\n", w)
			if err != nil {
				log.Fatalf("💥💥 Error writing to cfg.OutStream: %v", err)
			}
			log.Printf("Successfull write of %d Bytes to cfg.OutStream\n", nBytesWritten)
		} else {
			nBytesWritten, err := fmt.Fprintf(cfg.ErrStream, "word %s is odd\n", w)
			if err != nil {
				log.Fatalf("💥💥 Error writing to cfg.ErrStream: %v", err)
			}
			log.Printf("Successfull write of %d Bytes to cfg.ErrStream\n", nBytesWritten)
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
	if len(words) == 0 {
		nBytesWritten, err := fmt.Fprintln(os.Stderr, "No words provided.")
		if err != nil {
			log.Fatalf("💥💥 Error writing to stderr: %v", err)
		}
		log.Printf("Successfull write of %d Bytes to stderr\n", nBytesWritten)
		os.Exit(1)
	}

	cfg, err := NewCliConfig()
	if err != nil {
		nBytesWritten, err := fmt.Fprintf(os.Stderr, "Error creating config: %v\n", err)
		if err != nil {
			log.Fatalf("💥💥 Error writing to stderr: %v", err)
		}
		log.Printf("Successfull write of %d Bytes to stderr\n", nBytesWritten)
		os.Exit(1)
	}

	app(words, cfg)
}
