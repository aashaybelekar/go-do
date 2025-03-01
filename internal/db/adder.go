package db

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/aashaybelekar/go-do/internal/db"
	"github.com/aashaybelekar/go-do/internal/util"
)

func Addtask(p string, t string) error {
	filePath := db.LoadConfig()

	file, err := os.OpenFile(filePath, os.O_APPEND, 0644)

	if err != nil {
		log.Print(err)
		return err
	}

	defer file.Close()

	id, err := util.GetNextId(filePath)
	if err != nil {
		log.Print(err)
		return err
	}

	cur_time := time.Now().Format(time.RFC1123)

	data := []string{strconv.Itoa(id), t, cur_time, "false"}

	writer := csv.NewWriter(file)

	if err := writer.Write(data); err != nil {
		log.Print(err)
		return err
	}
	defer writer.Flush()
	return nil
}

func main() {
	Addtask("", "Buy a Bikes")
}
