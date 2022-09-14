package errors

import (
	"server/database"
)

type Status string

const (
	PENDING Status = "PENDING"
	Fixed   Status = "FIXED"
)

type ErrorDB struct {
	Status        *Status                `json="status" bson:"status"`
	Error         *error                 `json="error" bson:"error"`
	Line          string                 `json="line" bson:"line"`
	Package       string                 `json="package" bson:"package"`
	UpdatedAt     *string                `json="updatedAt" bson:"updatedAt"`
	CreatedAt     *string                `json="createdAt" bson:"createdAt"`
	IpGeolocation *database.ConnectionDB `json="ipGeolocation" bson="ipGeolocation"`
}
