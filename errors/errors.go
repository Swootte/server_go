package errors

import (
	"context"
	"os"
	"reflect"
	"runtime"
	"server/connection"
	"server/database"
	"strings"
	"time"
)

func SaveError(ctx context.Context, err error, line string, _package string, ip string, deviceId string) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("errors")
	_time := time.Now().UTC().Format(time.RFC3339)
	iptoLocation := connection.IpToLocation(ip)
	status := PENDING
	errIndb := ErrorDB{
		Status:        &status,
		Error:         &err,
		Line:          line,
		Package:       _package,
		UpdatedAt:     &_time,
		CreatedAt:     &_time,
		IpGeolocation: &database.ConnectionDB{IpAddress: ip, CreatedAt: _time, DeviceId: deviceId, Location: database.DBLocation{}, Country: iptoLocation.Country_long, City: iptoLocation.Country_long, Zip: iptoLocation.Zipcode, Region: iptoLocation.Region},
	}

	_collections.InsertOne(ctx, errIndb)
}

func GetPackageName(temp interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	strs = strings.Split(strs[len(strs)-2], "/")
	return strs[len(strs)-1]
}

func GetEthereumError(err string) string {
	switch err {
	case "0":
		return "The transaction passed the precheck validations"
	case "1":
		return "For any error not handled by specific error codes listed below"
	case "2":
		return "Payer account does not exist"
	case "3":
		return "Node Account provided does not match the node account of the node the transaction was submitted to"
	case "4":
		return "Pre-Check error when TransactionValidStart + transactionValidDuration is less than current consensus time"
	case "5":
		return "Transaction start time is greater than current consensus time"
	case "6":
		return "Valid transaction duration is a positive non zero number that does not exceed 120 seconds"
	case "7":
		return "The transaction signature is not valid"
	case "8":
		return "Transaction memo size exceeded 100 bytes"
	case "9":
		return "The fee provided in the transaction is insufficient for this type of transaction"
	case "10":
		return "The payer account has insufficient cryptocurrency to pay the transaction fee"
	case "11":
		return "This transaction ID is a duplicate of one that was submitted to this node or reached consensus in the last 180 seconds (receipt period)"
	case "12":
		return "The API is throttled out"
	case "13":
		return "The API is not currently supported"
	case "14":
		return "The user don't have the role to perform this action"
	case "15":
		return "The user don't have the role to perform this action"
	case "16":
		return "The user don't have the role to perform this action"
	case "17":
		return "The user is not allowed to perform this action"
	case "18":
		return "This account is frozen"
	case "19":
		return "This account is frozen"
	}
	return "An error has occurred please try again"
}
