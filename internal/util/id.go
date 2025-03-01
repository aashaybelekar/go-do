package util

import (
	"encoding/csv"
	"log"
	"os"
)

func GetNextId(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Print(err)
		return -1, nil
	}
	fileInfo, err := file.Stat()
	if err != nil {
		log.Print(err)
		return -1, nil
	}
	isEmpty := fileInfo.Size() == 0
	if isEmpty {
		return 0, nil
	} else {
		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		if err != nil {
			log.Print(err)
			return -1, nil
		}
		return len(records), nil
	}
}
