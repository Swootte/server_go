package connection

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"server/database"
	"time"

	"github.com/ip2location/ip2location-go/v9"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func WarnOfNewConnectionOccured(ctx context.Context, ip string, deviceId string) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("connections")
	record := IpToLocation(ip)
	id := primitive.NewObjectID()
	_time := time.Now().UTC().Format(time.RFC3339)
	newLogin := database.ConnectionDB{
		ID:        &id,
		IpAddress: ip,
		CreatedAt: _time,
		DeviceId:  deviceId,
		Location:  database.DBLocation{Longitude: record.Longitude, Latitude: record.Latitude},
		Country:   record.Country_long,
		City:      record.City,
		Zip:       record.Zipcode,
		Region:    record.Region,
	}

	_, err := _collections.InsertOne(ctx, newLogin)
	if err != nil {
		fmt.Println(err)
	}
}

func IpToLocation(ip string) *ip2location.IP2Locationrecord {
	const projectDirName = "server"
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))
	db, err := ip2location.OpenDB(string(rootPath) + "/IP2LOCATION-LITE-DB11.IPV6.BIN")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	results, err := db.Get_all(ip)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &results

}
