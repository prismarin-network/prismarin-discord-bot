package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type CustomWriter struct {
	File *os.File
}

func (writer *CustomWriter) Write(data []byte) (int, error) {
	message := string(data)
	fmt.Printf(message)
	writer.File.WriteString(message)
	return len(data), nil
}

func ModifyLogger() {
	os.Mkdir("logs/", 0755)
	file, _ := os.OpenFile(fmt.Sprintf("logs/%s.log", time.Now().Format("2006-01-02")), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	w := CustomWriter{
		File: file,
	}
	defer file.Close()
	log.SetOutput(&w)
}
