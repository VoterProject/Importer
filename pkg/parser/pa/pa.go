package pa_parser

import (
	"bufio"
	"fmt"
	"github.com/Jeffail/tunny"
	"github.com/gocarina/gocsv"
	"github.com/voterproject/importer/pkg/sql"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

type pa_parser struct {
	db       *sql.VoterDB
	lock     *sync.Mutex
	FirstRun bool
}

func ParseDirectory(path string, db *sql.VoterDB) {
	pa := pa_parser{db: db, lock: &sync.Mutex{}, FirstRun: true}

	fmt.Println(path)
	var wg sync.WaitGroup

	pool := tunny.NewFunc(6, func(i interface{}) interface{} {
		pa.parseFile(i.(string), &wg)
		debug.FreeOSMemory()
		return nil
	})

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !strings.Contains(path, "FVE") {
			return nil
		}
		wg.Add(1)

		if pa.FirstRun {
			pa.parseFile(path, &wg)
		} else {
			go pool.Process(path)
		}

		return nil
	})

	wg.Wait()
	fmt.Println("Done")

}

func (pa *pa_parser) parseFile(path string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Parsing: %s\n", path)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)

	var store sync.Map

	i := 0
	for scanner.Scan() {
		pa.parseLine(scanner.Text(), i, &store)
		i++
	}

	// Close file, don't wait too long.
	file.Close()

	fmt.Printf("\tDone parsing: %s\n", path)

	var records []Record
	var elections []Election
	var districts []District

	store.Range(func(key, value interface{}) bool {
		result := value.(*Results)
		records = append(records, *result.Record)
		elections = append(elections, result.Election...)
		districts = append(districts, result.District...)
		store.Delete(key)
		return true
	})
	fmt.Printf("\tDone listing: %s\n", path)

	if pa.FirstRun {
		pa.db.DB.CreateTable(&(records[0]))
		pa.db.DB.CreateTable(&(elections[0]))
		pa.db.DB.CreateTable(&(districts[0]))
	}

	pa.lock.Lock()
	defer pa.lock.Unlock()

	var f1, f2, f3 *os.File
	if pa.FirstRun {
		f1, err = os.OpenFile("./pa_records.csv", os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0777)
		f2, err = os.OpenFile("./pa_elections.csv", os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0777)
		f3, err = os.OpenFile("./pa_districts.csv", os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0777)
	} else {
		f1, err = os.OpenFile("./pa_records.csv", os.O_APPEND|os.O_RDWR, 0777)
		f2, err = os.OpenFile("./pa_elections.csv", os.O_APPEND|os.O_RDWR, 0777)
		f3, err = os.OpenFile("./pa_districts.csv", os.O_APPEND|os.O_RDWR, 0777)
	}

	defer f1.Close()
	defer f2.Close()
	defer f3.Close()

	if pa.FirstRun {
		err = gocsv.MarshalFile(&(records), f1)
		err = gocsv.MarshalFile(&(elections), f2)
		err = gocsv.MarshalFile(&(districts), f3)
		pa.FirstRun = false
	} else {
		err = gocsv.MarshalWithoutHeaders(&(records), f1)
		err = gocsv.MarshalWithoutHeaders(&(elections), f2)
		err = gocsv.MarshalWithoutHeaders(&(districts), f3)
	}

	records = nil
	elections = nil
	districts = nil

	fmt.Printf("\tDone CSV: %s\n", path)

	if pa.FirstRun {
		pa.FirstRun = false
	}
}

func (pa *pa_parser) parseLine(line string, i int, store *sync.Map) {
	fields := strings.Split(line, "\t")
	ID := parseString(fields[0])

	if ID == nil {
		panic("Nil ID")
	}

	//Districts are fields 30-70
	var districts []District
	for i := 30; i < 70; i++ {
		j := i - 30

		district := parseString(fields[i])
		if district == nil {
			continue
		}
		d := District{
			RecordID: *ID,
			Number:   j + 1,
			District: district,
		}
		districts = append(districts, d)
	}

	// Elections are fields 71-150, grouped in pairs
	var elections []Election
	for i := 70; i < 150; i = i + 2 {
		j := (i - 70) / 2
		voteMethod := parseString(fields[i])
		party := parseString(fields[i+1])

		if voteMethod == nil && party == nil {
			continue
		}
		election := Election{
			RecordID:   *ID,
			Number:     j + 1,
			VoteMethod: voteMethod,
			Party:      party,
		}
		elections = append(elections, election)
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

	store.Store(i, &Results{&record, elections, districts})
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
