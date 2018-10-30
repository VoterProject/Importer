package pa_parser

import (
	"bufio"
	"fmt"
	"github.com/Jeffail/tunny"
	"github.com/voterproject/importer/pkg/sql"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type pa_parser struct {
	db     *sql.VoterDB
	Tables bool
}

func ParseDirectory(path string, db *sql.VoterDB) {
	db.DB.AutoMigrate(&Record{}, &District{}, &Election{})

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

	var wg sync.WaitGroup
	pool := tunny.NewFunc(8, func(i interface{}) interface{} {
		return pa.parseLine(i.(string), &wg)
	})
	for scanner.Scan() {
		wg.Add(1)
		pool.Process(scanner.Text())
		//recordQueue := make([]*Record, 0)
		//record := pa.parseLine(scanner.Text())
		//if record != nil {
		//	recordQueue = append(recordQueue, record)
		//	if len(recordQueue) == 100 {
		//		fmt.Println("Start")
		//		tx := pa.db.DB.Begin()
		//		for _, v := range recordQueue {
		//			tx.Save(v)
		//		}
		//		tx.Commit()
		//		recordQueue = make([]*Record, 0)
		//		fmt.Println("Commit 100")
		//		return
		//	}
		//}
	}
	fmt.Println("Done breaking apart")
	wg.Wait()
	fmt.Println("Done import")
}

func (pa *pa_parser) parseLine(line string, wg *sync.WaitGroup) *Record {
	fields := strings.Split(line, "\t")

	// Districts are fields 30-70
	districts := make([]District, 40)
	for i := 30; i < 70; i++ {
		j := i - 30
		districts[j] = District{
			Number:   j + 1,
			District: parseString(fields[i])}
	}

	// Elections are fields 71-150, grouped in pairs
	elections := make([]Election, 40)
	for i := 70; i < 150; i = i + 2 {
		j := (i - 70) / 2
		election := Election{
			Number:     j + 1,
			VoteMethod: parseString(fields[i]),
			Party:      parseString(fields[i+1]),
		}
		elections[(i-70)/2] = election
	}

	id := parseString(fields[0])
	if id == nil {
		return nil
	}

	record := Record{
		ID:                *id,
		Title:             parseString(fields[1]),
		LastName:          parseString(fields[2]),
		FirstName:         parseString(fields[3]),
		MiddleName:        parseString(fields[4]),
		Suffix:            parseString(fields[5]),
		Gender:            parseString(fields[6]),
		DOB:               parseTime(fields[7]),
		RegistrationDate:  parseTime(fields[8]),
		VoterStatus:       parseString(fields[9]),
		StatusChangeDate:  parseTime(fields[10]),
		PartyCode:         parseString(fields[11]),
		HouseNumber:       parseString(fields[12]),
		HouseNumberSuffix: parseString(fields[13]),
		StreetName:        parseString(fields[14]),
		ApartmentNumber:   parseString(fields[15]),
		AddressLine2:      parseString(fields[16]),
		City:              parseString(fields[17]),
		State:             parseString(fields[18]),
		Zip:               parseString(fields[19]),
		MailAddress1:      parseString(fields[20]),
		MailAddress2:      parseString(fields[21]),
		MailCity:          parseString(fields[22]),
		MailState:         parseString(fields[23]),
		MailZip:           parseString(fields[24]),
		LastVoteDate:      parseTime(fields[25]),
		PrecinctCode:      parseString(fields[26]),
		PrecinctSplitID:   parseString(fields[27]),
		DateLastChanged:   parseTime(fields[28]),
		CustomData1:       parseString(fields[29]),
		HomePhone:         parseString(fields[150]),
		County:            parseString(fields[151]),
		MailCountry:       parseString(fields[152]),
	}

	pa.db.DB.Save(&record)

	tx := pa.db.DB.Begin()
	for _, v := range elections {
		if v.Party == nil {
			continue
		}
		v.RecordID = record.ID
		tx.Save(&v)
	}
	for _, v := range districts {
		if v.District == nil {
			continue
		}
		v.RecordID = record.ID
		tx.Save(&v)
	}
	tx.Commit()
	wg.Done()
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
