package wa_parser

import (
	"bufio"
	"fmt"
	"github.com/voterproject/importer/pkg/sql"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type wa_parser struct {
	db       *sql.VoterDB
	lock     *sync.Mutex
	FirstRun bool
}

func ParseDirectory(path string, db *sql.VoterDB) {
	wa := wa_parser{db: db, lock: &sync.Mutex{}, FirstRun: true}

	fmt.Println(path)
	var wg sync.WaitGroup

	//pool := tunny.NewFunc(6, func(i interface{}) interface{} {
	//	wa.parseRecords(i.(string), &wg)
	//	debug.FreeOSMemory()
	//	return nil
	//})

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, "VRDB") {
			wa.parseRecords(path, &wg)
		} else {
			return nil
		}

		return nil
	})

	wg.Wait()
	fmt.Println("Done")
}

func (wa *wa_parser) parseRecords(path string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)

	var records []Record
	i := 0
	for scanner.Scan() {
		r := wa.parseLine(scanner.Text(), i)
		records = append(records, *r)
		i++
	}
	fmt.Println(records)

}

func (wa *wa_parser) parseLine(line string, i int) *Record {
	fields := strings.Split(line, "\t")

	StateID := parseString(fields[0])
	CountyID := parseString(fields[1])
	if StateID == nil || CountyID == nil {
		panic("Nil ID " + line)
	}

	record := Record{
		*StateID,
		*CountyID,
		parseString(fields[2]),
		parseString(fields[3]),
		parseString(fields[4]),
		parseString(fields[5]),
		parseString(fields[6]),
		parseTime(fields[7]),
		parseString(fields[8]),
		parseString(fields[9]),
		parseString(fields[10]),
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
		parseString(fields[25]),
		parseString(fields[26]),
		parseString(fields[27]),
		parseString(fields[28]),
		parseString(fields[29]),
		parseString(fields[30]),
		parseString(fields[31]),
		parseString(fields[32]),
		parseTime(fields[33]),
		parseString(fields[34]),
		parseTime(fields[35]),
		parseString(fields[36]),
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
