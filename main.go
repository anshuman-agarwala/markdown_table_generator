package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func outputMarkdown(markdown string) error {
	err := ioutil.WriteFile("markdown_table.md", []byte(markdown), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func renderMarkdown(content [][]string) (string, error) {
	var markdown string
	if len(content) == 0 {
		return "", fmt.Errorf("no data found")
	}

	line := "| " + strings.Join(content[0], " | ") + " |"
	markdown = line + "\n"
	headingLength := len(content[0])
	headingSeparator := "|"
	headingSeparator = headingSeparator + strings.Repeat(" - |", headingLength) + "\n"
	markdown = markdown + headingSeparator
	if len(content) == 1 {
		return markdown, nil
	}
	for _, item := range content[1:] {
		markdown = markdown + "| " + strings.Join(item, " | ") + " |" + "\n"
	}
	return markdown, nil
}

func parseCSV(filename string) [][]string {
	var fileContents [][]string

	// Handle file operations
	fileObj, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer fileObj.Close()

	s := bufio.NewScanner(fileObj)
	for s.Scan() {
		temp := strings.Split(s.Text(), ",")
		fileContents = append(fileContents, temp)
	}
	return fileContents
}

func main() {
	fileName := flag.String("file_path", "", "Path of the .csv file.")
	flag.Parse()
	_, err := os.Stat(*fileName)
	if os.IsNotExist(err) {
		fmt.Println("Please enter a valid filename as an argument to the function.")
		fmt.Println("for more help, run the program with a '-h' flag.")
		return
	}
	data := parseCSV(*fileName)
	markdown, err := renderMarkdown(data)
	if err != nil {
		panic(err)
	}
	err = outputMarkdown(markdown)
	if err != nil {
		panic(err)
	}
	// fmt.Println(renderMarkdown(data))
}
