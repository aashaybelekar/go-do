package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func Loadfile(p string) error {
	if p == "" {
		p = "./artifacts/data.csv"
	}
	filePath := p
	tempPath := strings.Split(filePath, "/")
	folderPath := strings.Join(tempPath[:len(tempPath)-1], "/")

	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		fmt.Println("Error creating folder:", err)
		return err
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	// Check if file is empty
	fileInfo, _ := file.Stat()
	isEmpty := fileInfo.Size() == 0

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write headers if file is empty
	if isEmpty {
		headers := []string{"ID", "Description", "CreatedAt", "IsComplete"}
		if err := writer.Write(headers); err != nil {
			fmt.Println("Error writing headers:", err)
			return err
		}
	} else {
		log.Printf("The folder already exists, loaded '%v' sucessfully.", filePath)
	}

	return nil
}

func main() {
	Loadfile("")
}
