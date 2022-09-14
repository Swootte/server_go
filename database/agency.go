package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type DBAgency struct {
	ShortId   string
	Title     string
	Address   string
	Status    string
	Creator   primitive.ObjectID
	CreatedAt string
	UpdatedAt string
	City      string
	Country   string
}
