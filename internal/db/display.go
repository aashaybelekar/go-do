package db

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

func Listtask(p string) error {
	if p == "" {
		p = "./artifacts/data.csv"
	}

	filePath := p
	tempPath := strings.Split(filePath, "/")
	folderPath := strings.Join(tempPath[:len(tempPath)-1], "/")

	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		log.Print("Error creating folder:", err)
		return err
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Print(err)
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Print(err)
		return err
	}
	isEmpty := fileInfo.Size() == 0

	writer := csv.NewWriter(file)
	if isEmpty {
		header := []string{"ID", "Task", "Created", "Complete"}
		if err := writer.Write(header); err != nil {
			log.Print(err)
			return err
		}
	}
	return nil
}
