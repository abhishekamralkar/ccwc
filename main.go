package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func getFileSize(filePath string) (int64, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}

	return fileInfo.Size(), nil
}

func main() {
	filePath := flag.String("c", "", "Path of the file")

	flag.Parse()

	if *filePath == "" {
		fmt.Println("Usage: ./ccwc -c filepath")
		flag.PrintDefaults()
		os.Exit(1)
	}

	filesize, err := getFileSize(*filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(filesize, *filePath)
}
