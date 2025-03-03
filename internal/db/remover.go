package db

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/aashaybelekar/go-do/internal/util"
)

func Delete(task_id int) error {
	config, err := util.LoadConfig()
	if err != nil {
		log.Print(err)
		return err
	}

	filePath := config.Location

	// Read

	read_file, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer read_file.Close()

	reader := csv.NewReader(read_file)

	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// Write

	write_file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer write_file.Close()

	writer := csv.NewWriter(write_file)

	for _, record := range records[:task_id] {
		writer.Write(record)
	}

	for _, record := range records[task_id+1:] {
		prev_id, err := strconv.Atoi(record[0])
		if err != nil {
			return err
		}
		new_id := prev_id - 1
		record[0] = strconv.Itoa(new_id)
		writer.Write(record)
	}
	writer.Flush()
	return nil
}
