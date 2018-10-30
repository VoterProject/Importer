package pa_parser

import (
	"bufio"
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/voterproject/importer/pkg/sql"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type pa_parser struct {
	db     *sql.VoterDB
	Tables bool
}

func ParseDirectory(path string, db *sql.VoterDB) {

	pa := pa_parser{db: db}

	fmt.Println(path)

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !strings.Contains(path, "FVE") {
			return nil
		}
		if !strings.Contains(path, "MONTGOMERY") {
			return nil
		}
		pa.parseFile(path)
		return nil
	})
}

func (pa *pa_parser) parseFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)

	var records []Record
	var elections []Election
	var districts []District
	for scanner.Scan() {
		x := pa.parseLine(scanner.Text())
		records = append(records, *x.Record)
		elections = append(elections, x.Election...)
		districts = append(districts, x.District...)
		//return
	}

	pa.db.DB.CreateTable(&(records[0]))
	pa.db.DB.CreateTable(&(elections[0]))
	pa.db.DB.CreateTable(&(districts[0]))

	f1, err := os.Create("./pa_records.csv")
	f2, err := os.Create("./pa_elections.csv")
	f3, err := os.Create("./pa_districts.csv")
	err = gocsv.MarshalFile(&records, f1)
	err = gocsv.MarshalFile(&elections, f2)
	err = gocsv.MarshalFile(&districts, f3)

}

func (pa *pa_parser) parseLine(line string) *Results {
	fields := strings.Split(line, "\t")
	ID := parseString(fields[0])

	if ID == nil {
		panic("Nil ID")
	}

	//Districts are fields 30-70
	districts := make([]District, 40)
	for i := 30; i < 70; i++ {
		j := i - 30
		districts[j] = District{
			RecordID: *ID,
			Number:   j + 1,
			District: parseString(fields[i])}
	}

	// Elections are fields 71-150, grouped in pairs
	elections := make([]Election, 40)
	for i := 70; i < 150; i = i + 2 {
		j := (i - 70) / 2
		election := Election{
			RecordID:   *ID,
			Number:     j + 1,
			VoteMethod: parseString(fields[i]),
			Party:      parseString(fields[i+1]),
		}
		elections[(i-70)/2] = election
	}
	record := Record{
		*ID,
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
		parseString(fields[150]),
		parseString(fields[151]),
		parseString(fields[152]),
	}
	return &Results{
		&record,
		elections,
		districts,
	}
}

func (pa *pa_parser) ToCSV(records []Record) {

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
