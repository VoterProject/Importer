package pa_parser

import "time"

type Record struct {
	ID                string `gorm:"primary_key"`
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
	Districts         []District `gorm:"foreignkey:RecordID;association_foreignkey:ID"`
	Elections         []Election `gorm:"foreignkey:RecordID;association_foreignkey:ID"`
	HomePhone         *string
	County            *string
	MailCountry       *string
}

type Election struct {
	ElectionID uint `gorm:"primary_key"`
	RecordID   string
	Number     int
	VoteMethod *string
	Party      *string
}

type District struct {
	DistrictID uint `gorm:"primary_key"`
	RecordID   string
	Number     int
	District   *string
}
