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

func getWordCount(filepath string) (int64, error) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	wordCount := 0

	for scanner.Scan() {
		words := scanner.Text()
		wordCount += len(words)
	}

	return int64(wordCount), err
}

func main() {

	fileSize := flag.String("c", "", "size of the file")
	fileLines := flag.String("l", "", "line count in the file")
	wordCount := flag.String("w", "", "word count in a file")

	filePath := os.Args[2]

	/*
		if len(os.Args) > 2 {
			fmt.Println("Usage: gowc [options] filepath")
			flag.PrintDefaults()
			os.Exit(1)
		}
	*/

	flag.Parse()

	switch {
	case *fileSize != "":
		f, _ := getFileSize(filePath)
		fmt.Println(f, filePath)
	case *fileLines != "":
		s, _ := getFileLines(filePath)
		fmt.Println(s, filePath)
	case *wordCount != "":
		w, _ := getWordCount(filePath)
		fmt.Println(w, filePath)
	default:
		fmt.Println("Usage: gowc [options] filepath")
		flag.PrintDefaults()
		os.Exit(1)

	}
}
