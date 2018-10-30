package pa_parser

import (
	"bufio"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func ParseDirectory(path string) {
	fmt.Println(path)

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !strings.Contains(path, "FVE") {
			return nil
		}
		if !strings.Contains(path, "MONTGOMERY") {
			return nil
		}
		parseFile(path)
		return nil
	})
}

func parseFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parseLine(scanner.Text())
		//return
	}

}

func parseLine(line string) *Record {
	fields := strings.Split(line, "\t")

	// Districts are fields 30-70
	districts := map[int]*string{}
	for i := 30; i < 70; i++ {
		districts[i-30] = parseString(fields[i])
	}

	// Elections are fields 71-150, grouped in pairs
	elections := map[int]Election{}
	for i := 70; i < 150; i = i + 2 {
		election := Election{
			parseString(fields[i]),
			parseString(fields[i+1]),
		}
		elections[i-70] = election
	}
	record := Record{
		parseString(fields[0]),
		parseString(fields[1]),
		parseString(fields[2]),
		parseString(fields[3]),
		parseString(fields[4]),
		parseString(fields[5]),
		parseString(fields[6]),
		parseTime(fields[7]),
		parseTime(fields[8]),
		parseString(fields[9]),
		parseTime(fields[10]),
		parseString(fields[11]),
		parseString(fields[12]),
		parseString(fields[13]),
		parseString(fields[14]),
		parseString(fields[15]),
		parseString(fields[16]),
		parseString(fields[17]),
		parseString(fields[18]),
		parseString(fields[19]),
		parseString(fields[20]),
		parseString(fields[21]),
		parseString(fields[22]),
		parseString(fields[23]),
		parseString(fields[24]),
		parseTime(fields[25]),
		parseString(fields[26]),
		parseString(fields[27]),
		parseTime(fields[28]),
		parseString(fields[29]),
		districts,
		elections,
		parseString(fields[150]),
		parseString(fields[151]),
		parseString(fields[152]),
	}

	if record.FirstName != nil && *(record.FirstName) == "AMIR" {
		spew.Println(record)
	}
	return &record
}

func parseTime(field string) *time.Time {
	t, err := time.Parse("01/02/2006", field)
	if err != nil {
		return nil
	}
	return &t
}
func parseString(field string) *string {
	field = strings.Trim(field, "\"")
	if field == "" {
		return nil
	}
	return &field
}
