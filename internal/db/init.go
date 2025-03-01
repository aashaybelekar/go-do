package db

import (
	"encoding/csv"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type DataBaseConfig struct {
	location string
}

func LoadConfig() (*DataBaseConfig, error) {
	filename := "./config/db.yaml"
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config DataBaseConfig
	err = yaml.Unmarshal(file, &config)
	return &config, err
}

func init() {
	config, err := LoadConfig()
	if err != nil {
		log.Print("Error loading config:", err)
		return
	}
	log.Print("Loaded config sucessfully.")

	filePath := config.location

	if !strings.HasSuffix(filePath, ".csv") {
		panic("file path has to end with '.csv'")
	}

	tempPath := strings.Split(filePath, "/")
	folderPath := strings.Join(tempPath[:len(tempPath)-1], "/")

	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		log.Print("Error creating folder:", err)
	}

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Print(err)
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	fileInfo, err := file.Stat()
	if err != nil {
		log.Print(err)
		panic(err)
	}
	isEmpty := fileInfo.Size() == 0

	if isEmpty {
		header := []string{"ID", "Task", "Created", "Complete"}
		if err := writer.Write(header); err != nil {
			log.Print(err)
			panic(err)
		}
	}

}
