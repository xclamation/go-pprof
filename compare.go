package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	// Подход 1: Использование bufio.Scanner

	scanner := bufio.NewScanner(file)
	//scanner := bufio.NewScannerSize(file, 1024)
	users := make([]map[string]interface{}, 0)
	start := time.Now()
	var line []byte

	for scanner.Scan() {
		line = scanner.Bytes()
		user := make(map[string]interface{})
		err := json.Unmarshal(line, &user)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to scan file: %s", err)
	}

	elapsed := time.Since(start)
	fmt.Printf("Подход 1 (bufio.Scanner): %s\n", elapsed)

	// Подход 2: Использование bufio.Reader
	file.Seek(0, 0) // Сброс позиции чтения файла

	reader := bufio.NewReader(file)
	users = make([]map[string]interface{}, 0)
	start = time.Now()
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("failed to read line: %s", err)
		}
		user := make(map[string]interface{})
		err = json.Unmarshal([]byte(line), &user)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	elapsed = time.Since(start)
	fmt.Printf("Подход 2 (bufio.Reader): %s\n", elapsed)
}
