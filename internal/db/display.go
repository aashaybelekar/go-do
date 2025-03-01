package db

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/aashaybelekar/go-do/internal/util"
	"github.com/mergestat/timediff"
)

func format_header(row []string, a_tag bool) string {
	task_id := row[0]
	task_des := row[1]
	task_time := row[2]
	task_completion := row[3]

	till_time := task_time
	if a_tag {
		out_string := fmt.Sprintf("\t%s\t%s\t%s\t%s\t", task_id, task_des, till_time, task_completion)
		return out_string
	} else {
		out_string := fmt.Sprintf("\t%s\t%s\t%s\t", task_id, task_des, till_time)
		return out_string
	}
}

func format_row(row []string, a_tag bool) (string, bool) {
	task_id := row[0]
	task_des := row[1]
	task_time, err := time.Parse(time.RFC1123, row[2])
	if err != nil {
		log.Print(err)
	}

	task_completion := row[3]

	var complete bool
	if task_completion == "true" {
		complete = true
	} else {
		complete = false
	}

	till_time := timediff.TimeDiff(task_time)
	if a_tag {
		out_string := fmt.Sprintf("\t%s\t%s\t%s\t%s\t", task_id, task_des, till_time, task_completion)
		return out_string, complete
	} else {
		out_string := fmt.Sprintf("\t%s\t%s\t%s\t", task_id, task_des, till_time)
		return out_string, complete
	}
}

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

	out_string := format_header(records[0], a_tag)
	fmt.Fprintln(w, out_string)

	for _, row := range records[1:] {
		out_string, complete := format_row(row, a_tag)
		if !complete {
			fmt.Fprintln(w, out_string)
		}
	}

	w.Flush()

	return nil
}
