package wa_parser

import "time"

type Record struct {
	StateID               string
	CountyID              string
	Title                 *string
	FirstName             *string
	MiddleName            *string
	LastName              *string
	Suffix                *string
	DOB                   *time.Time
	Gender                *string
	RegStNum              *string
	RegStFrac             *string
	RegStName             *string
	RegStType             *string
	RegUnitType           *string
	RegStPreDirection     *string
	RegStPostDirection    *string
	RegUnitNum            *string
	RegCity               *string
	RegState              *string
	RegZipCode            *string
	CountyCode            *string
	PrecinctCode          *string
	PrecinctPart          *string
	LegislativeDistrict   *string
	CongressionalDistrict *string
	Mail1                 *string
	Mail2                 *string
	Mail3                 *string
	Mail4                 *string
	MailCity              *string
	MailZip               *string
	MailState             *string
	MailCountry           *string
	RegistrationDate      *time.Time
	AbsenteeType          *string
	LastVoted             *time.Time
	StatusCode            *string
}

type History struct {
	CountyCode      *string
	StateID         string
	ElectionDate    *time.Time
	VotingHistoryID *string
}
