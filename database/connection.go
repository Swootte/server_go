package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type ConnectionDB struct {
	ID        *primitive.ObjectID `json:"_id" bson:"_id"`
	IpAddress string              `json:"ipAddress" bson:"ipAddress"`
	CreatedAt string              `json:"createdAt" bson:"createdAt"`
	DeviceId  string              `json:"deviceId" bson:"deviceId"`
	Location  DBLocation          `json:"location" bson:"location"`
	Country   string              `json:"country" bson:"country"`
	City      string              `json:"city" bson:"city"`
	Zip       string              `json:"zip" bson:"zip"`
	Region    string              `json:"region" bson:"region"`
}
