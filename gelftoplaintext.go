package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type LogLevel int

const (
	EmergencyLevel = LogLevel(0)
	AlertLevel     = LogLevel(1)
	CriticalLevel  = LogLevel(2)
	ErrorLevel     = LogLevel(3)
	WarningLevel   = LogLevel(4)
	NoticeLevel    = LogLevel(5)
	InfoLevel      = LogLevel(6)
	DebugLevel     = LogLevel(7)
)

type gelf struct {
	Version          string                 `json:"version"`
	Host             string                 `json:"host"`
	ShortMessage     string                 `json:"short_message"`
	FullMessage      string                 `json:"full_message,omitempty"`
	Timestamp        int64                  `json:"timestamp"`
	Level            LogLevel               `json:"level"`
	AdditionalFields map[string]interface{} `json:"-"` // we custom-marshal this
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
		var g gelf
		err = json.Unmarshal(line, &g)
		if err != nil {
			fmt.Printf("Error unmarshaling GELF: %v\n", err)
			os.Exit(1)
		}
		if g.FullMessage != "" {
			fmt.Println(g.FullMessage)
		} else {
			fmt.Println(g.ShortMessage)
		}
	}
}
