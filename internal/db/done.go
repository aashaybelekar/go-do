package db

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/aashaybelekar/go-do/internal/util"
)

func Completetask(task_id int) error {
	config, err := util.LoadConfig()
	if err != nil {
		log.Print(err)
		return err
	}

	filePath := config.Location

	read_file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		log.Print(err)
	}

	reader := csv.NewReader(read_file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Print(err)
	}

	write_file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)

	writer := csv.NewWriter(write_file)

	for i, rec := range records {
		if task_id == i {
			rec[3] = "true"
		}
		writer.Write(rec)
	}
	writer.Flush()

	return nil
}
