// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
)

type Catch struct {
	ID           uuid.UUID
	EventID      sql.NullInt32
	CompetitorID sql.NullInt32
	SpeciesID    sql.NullInt32
	Size         sql.NullInt32
	CaughtAt     interface{}
	Bait         string
	Location     string
	Latitude     interface{}
	Longitude    interface{}
	Marshall     string
	MarshallID   sql.NullInt32
	Status       sql.NullInt32
}

type Club struct {
	ID              uuid.UUID
	Name            sql.NullString
	BillingAddress1 sql.NullString
	BillingAddress2 sql.NullString
	BillingSuburb   sql.NullString
	BillingState    sql.NullString
	BillingPostcode sql.NullString
	StripeBillingID sql.NullString
	Owner           sql.NullInt32
	Settings        json.RawMessage
}

type Competition struct {
	ID             uuid.UUID
	OrganisationID sql.NullInt32
	ShortName      sql.NullString
	Name           sql.NullString
	LogoUrl        sql.NullString
	CustomDomain   sql.NullString
	CurrentEvent   sql.NullInt32
	Settings       json.RawMessage
}

type Competitor struct {
	ID           uuid.UUID
	EventID      uuid.UUID
	CompetitorNo sql.NullString
	Firstname    string
	Lastname     string
	Email        string
	Address1     string
	Address2     string
	Suburb       string
	State        string
	Postcode     string
	Phone        string
	Mobile       string
	Paid         interface{}
	Registered   interface{}
	Checkin      interface{}
	Ticket       sql.NullInt32
	TeamID       sql.NullInt32
	UserID       sql.NullInt32
}

type Event struct {
	ID            uuid.UUID
	CompetitionID sql.NullInt32
	Slug          sql.NullString
	Name          sql.NullString
	StartDate     sql.NullTime
	EndDate       sql.NullTime
	Location      sql.NullString
	Status        sql.NullInt32
	Settings      json.RawMessage
}

type OrganisationUser struct {
	OrganisationID sql.NullInt32
	UserID         sql.NullInt32
	Admin          interface{}
	Marshall       interface{}
}

type Species struct {
	ID             uuid.UUID
	Slug           string
	CommonName     string
	ScientificName string
	PhotoUrl       string
}

type Team struct {
	ID       uuid.UUID
	EventID  sql.NullInt32
	TeamNo   sql.NullInt32
	Name     sql.NullString
	BoatRego sql.NullString
}

type Ticket struct {
	ID                uuid.UUID
	EventID           sql.NullInt32
	Name              sql.NullString
	StartCompetitorNo sql.NullInt32
	NextCompetitorNo  sql.NullInt32
	Price             sql.NullInt32
	StripeProductID   sql.NullInt32
	MaxNoCompetitors  sql.NullInt32
}

type User struct {
	ID              uuid.UUID
	Username        string
	Password        string
	Firstname       sql.NullString
	Lastname        sql.NullString
	Email           sql.NullString
	Mobile          sql.NullString
	ApiToken        sql.NullString
	Address1        sql.NullString
	Address2        sql.NullString
	Suburb          sql.NullString
	State           sql.NullString
	Postcode        sql.NullString
	StripeBillingID sql.NullString
	Settings        json.RawMessage
}
