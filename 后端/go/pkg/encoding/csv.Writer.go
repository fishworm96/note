package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	w := csv.NewWriter(os.Stdout)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("eror writing record to csv:", err)
		}
	}

	// 将缓冲数据写入底层写入器（标准输出）。
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}