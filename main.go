package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func getFileSize(filePath string) (int64, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}

	return fileInfo.Size(), nil
}

func getFileLines(filepath string) (int64, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return 0, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineCount := 0

	for scanner.Scan() {
		lineCount++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return int64(lineCount), nil
}

func main() {

	fileSize := flag.String("c", "", "size of the file")
	fileLines := flag.String("l", "", "line count in the file")

	filePath := os.Args[2]

	if len(os.Args) > 2 {
		fmt.Println("Usage: gowc [options] filepath")
		flag.PrintDefaults()
		os.Exit(1)
	}

	flag.Parse()

	switch {
	case *fileSize != "":
		f, _ := getFileSize(filePath)
		fmt.Println(f, filePath)
	case *fileLines != "":
		s, _ := getFileLines(filePath)
		fmt.Println(s, filePath)
	default:
		fmt.Println("Usage: gowc [options] filepath")
		flag.PrintDefaults()
		os.Exit(1)

	}
}
