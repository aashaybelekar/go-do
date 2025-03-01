package db

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/aashaybelekar/go-do/internal/util"
)

func Listtask(a_tag bool) error {
	config, err := util.LoadConfig()
	if err != nil {
		log.Print(err)
		return err
	}

	filePath := config.Location

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDONLY, 0644)
	if err != nil {
		log.Print(err)
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		log.Print(err)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)

	for _, row := range records {
		task_id := row[0]
		task_des := row[1]
		task_time := row[2]
		task_completion := row[3]
		if a_tag {
			out_string := fmt.Sprintf("\t%s\t%s\t%s\t%s\t", task_id, task_des, task_time, task_completion)
			fmt.Fprintln(w, out_string)
		} else {
			out_string := fmt.Sprintf("\t%s\t%s\t%s\t", task_id, task_des, task_time)
			fmt.Fprintln(w, out_string)
		}
	}

	w.Flush()

	return nil
}
