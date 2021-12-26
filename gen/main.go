package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

//go:generate go run main.go

func main() {
	stem := strconv.Itoa(int(time.Now().Unix()))
	ext := ".txt"
	baseName := stem + ext
	tmp, err := os.Create(filepath.Join("..", "generated", baseName))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	tmp.Write([]byte{98, 101, 97, 116, 115}) // beats
}
