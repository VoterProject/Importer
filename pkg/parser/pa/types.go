package pa_parser

import "time"

type Record struct {
	ID                *string
	Title             *string
	LastName          *string
	FirstName         *string
	MiddleName        *string
	Suffix            *string
	Gender            *string
	DOB               *time.Time
	RegistrationDate  *time.Time
	VoterStatus       *string
	StatusChangeDate  *time.Time
	PartyCode         *string
	HouseNumber       *string
	HouseNumberSuffix *string
	StreetName        *string
	ApartmentNumber   *string
	AddressLine2      *string
	City              *string
	State             *string
	Zip               *string
	MailAddress1      *string
	MailAddress2      *string
	MailCity          *string
	MailState         *string
	MailZip           *string
	LastVoteDate      *time.Time
	PrecinctCode      *string
	PrecinctSplitID   *string
	DateLastChanged   *time.Time
	CustomData1       *string
	Districts         map[int]*string
	Elections         map[int]Election
	HomePhone         *string
	County            *string
	MailCountry       *string
}

type Election struct {
	VoteMethod *string
	Party      *string
}
