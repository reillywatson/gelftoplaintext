package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type gelfEntry struct {
	ShortMessage string `json:"short_message"`
	FullMessage  string `json:"full_message,omitempty"`
}

func main() {
	in := os.Stdin
	if len(os.Args) > 1 {
		var err error
		in, err = os.Open(os.Args[1])
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			os.Exit(1)
		}
	}
	reader := bufio.NewReaderSize(in, 1024*1024*32)
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Read error: %v\n", err)
			os.Exit(1)
		}
		var gelf gelfEntry
		err = json.Unmarshal(line, &gelf)
		if err != nil {
			fmt.Printf("Error unmarshaling GELF: %v\n", err)
			os.Exit(1)
		}
		if gelf.FullMessage != "" {
			fmt.Println(gelf.FullMessage)
		} else {
			fmt.Println(gelf.ShortMessage)
		}
	}
}
