// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"github.com/google/uuid"
)

type Competitor struct {
	ID           uuid.UUID
	CompetitorNo string
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
}
