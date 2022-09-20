package transaction

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"server/database"
	"server/graph/model"
	"server/user"
	"server/utils"
	"time"

	"server/finance"
	"server/notification"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	utils.LoadEnv()
}

func reverseFloats(input []*float64) []*float64 {
	if len(input) == 0 {
		return input
	}
	return append(reverseFloats(input[1:]), input[0])
}

func LoadTransactiosnForUserID(ctx context.Context, userID string) ([]*model.Paiement, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("transactions")
	objectId, _ := primitive.ObjectIDFromHex(userID)

	_match := bson.D{{Key: "$match", Value: bson.D{{Key: "For", Value: objectId}}}}

	cursor, err := _collections.Aggregate(ctx, mongo.Pipeline{_match})
	if err != nil {
		return nil, err
	}

	var payments []*model.Paiement
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var singlePayment *model.Paiement
		if err = cursor.Decode(&singlePayment); err != nil {
			log.Fatal(err)
		}
		payments = append(payments, singlePayment)
	}

	return payments, nil
}

func GetTransactionByIdAgent(ctx context.Context, transactionId string) (*model.Paiement, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("transactions")
	objectId, _ := primitive.ObjectIDFromHex(transactionId)

	var payment *model.Paiement
	err := _collections.FindOne(ctx, bson.D{{Key: "_id", Value: objectId}}).Decode(&payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func GetTransactionByIdUnauthed(ctx context.Context, transactionId string) (*model.Paiement, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("transactions")
	objectId, _ := primitive.ObjectIDFromHex(transactionId)

	var payment *model.Paiement
	err := _collections.FindOne(ctx, bson.D{{Key: "_id", Value: objectId}}).Decode(&payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func GetTransactionById(ctx context.Context, transactionId string, userId string) (*model.Paiement, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("transactions")
	objectId, _ := primitive.ObjectIDFromHex(transactionId)
	objectIdUser, _ := primitive.ObjectIDFromHex(userId)

	var payment *model.Paiement
	err := _collections.FindOne(ctx, bson.D{{Key: "_id", Value: objectId}, {Key: "$or", Value: bson.A{
		bson.M{"destinationUserId": objectIdUser},
		bson.M{"creatorId": objectIdUser},
	}}}).Decode(&payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func CreateTransfer(ctx context.Context, address *string, token string, amount float64, destinationUser string, _user model.User, ip string) (*bool, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("transactions")
	payee, err := user.GetUserById(destinationUser)
	if err != nil {
		return nil, err
	}

	_amount := finance.ToWei(amount)
	fee := *finance.ToWei(finance.CalculateFee(amount))
	receipt, err := finance.Transfer(_user.Keypair.PublicKey, *address, _amount, _user.Keypair.SecretKey)
	if err != nil {
		return nil, err
	}

	_destinationUser, _ := primitive.ObjectIDFromHex(payee.ID)
	_creator, _ := primitive.ObjectIDFromHex(_user.ID)
	_time := time.Now().UTC().Format(time.RFC3339)
	_id := primitive.NewObjectID()

	input := database.DBTransaction{
		ID:                &_id,
		TransactionId:     receipt.TxHash.String(),
		Source:            _user.Keypair.PublicKey,
		Destination:       *address,
		Fee:               fee.String(),
		Amount:            _amount.String(),
		Token:             token,
		Description:       "",
		DestinationUserID: &_destinationUser,
		Type:              model.PaymentTypeTransfert.String(),
		Status:            model.PaymentStatusDone.String(),
		CreatorID:         &_creator,
		Country:           "",
		CreatedAt:         _time,
		UpdatedAt:         _time,
	}

	_result, err := _collections.InsertOne(ctx, input)
	if err != nil {
		return nil, err
	}

	if _result.InsertedID == nil {
		return nil, fmt.Errorf("insertedId error")
	}

	result := true

	notification.CreateDBNotification(model.PaymentTypeTransfert, _user, *payee, input)
	return &result, nil
}

func AddTopup(ctx context.Context, topup model.TopUpInput, pinCode string, payerDB model.User, ip string) (string, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("transactions")
	destinationUser, _ := primitive.ObjectIDFromHex(topup.DestinationUser)
	id := primitive.NewObjectID()

	_amount := finance.ToWei(topup.Amount)
	receipt, err := finance.Deposit(id.Hex(), _amount, payerDB.Keypair.SecretKey, payerDB.Keypair.PublicKey)
	if err != nil {
		return "", err
	}
	agency, _ := primitive.ObjectIDFromHex(topup.Agency)
	creator, _ := primitive.ObjectIDFromHex(payerDB.ID)
	_time := time.Now().UTC().Format(time.RFC3339)
	transaction := database.DBTransaction{
		ID:                &id,
		TransactionId:     receipt.TxHash.String(),
		Source:            "",
		Destination:       topup.Destination,
		AgencyID:          &agency,
		Fee:               big.NewInt(0).String(),
		Amount:            _amount.String(),
		Token:             *payerDB.DefaultCurrency,
		Description:       "",
		DestinationUserID: &destinationUser,
		ShortId:           "",
		Type:              model.PaymentTypeTopup.String(),
		Status:            model.PaymentStatusInProgress.String(),
		CreatorID:         &creator,
		EnterpriseID:      nil,
		Country:           "CG",
		CreatedAt:         _time,
		UpdatedAt:         _time,
	}
	_collections.InsertOne(context.Background(), transaction)
	return id.Hex(), nil
}

func AddWithdraw(ctx context.Context, withdraw model.WithdrawInput, pinCode string, payerDB model.User, ip string) (string, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("transactions")
	destinationUser, _ := primitive.ObjectIDFromHex(withdraw.DestinationUser)
	id := primitive.NewObjectID()

	_amount := finance.ToWei(withdraw.Amount)
	receipt, err := finance.AskWithdraw(_amount, payerDB.Keypair.SecretKey, id.Hex(), payerDB.Keypair.PublicKey)
	if err != nil {
		return "", err
	}

	agency, _ := primitive.ObjectIDFromHex(withdraw.Agency)
	creator, _ := primitive.ObjectIDFromHex(payerDB.ID)
	_time := time.Now().UTC().Format(time.RFC3339)
	transaction := database.DBTransaction{
		ID:                &id,
		TransactionId:     receipt.TxHash.String(),
		Source:            "",
		Destination:       withdraw.Destination,
		AgencyID:          &agency,
		ValidatorID:       nil,
		CancellorID:       nil,
		Fee:               big.NewInt(0).String(),
		Amount:            _amount.String(),
		Token:             *payerDB.DefaultCurrency,
		Description:       "",
		DestinationUserID: &destinationUser,
		ShortId:           "",
		Type:              model.PaymentTypeWithdraw.String(),
		Status:            model.PaymentStatusInProgress.String(),
		CreatorID:         &creator,
		EnterpriseID:      nil,
		Country:           "CG",
		CreatedAt:         _time,
		UpdatedAt:         _time,
	}
	_collections.InsertOne(context.Background(), transaction)
	return id.Hex(), nil
}

func GetAllTransactionByEnterpriseId(ctx context.Context, enterpriseID string, from string, to string, limit float64, skip float64, userId string) (*model.TransactionWithPageInfo, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("transactions")
	objectId, _ := primitive.ObjectIDFromHex(enterpriseID)
	objectIdUser, _ := primitive.ObjectIDFromHex(userId)

	_match := bson.D{{Key: "$match", Value: bson.D{{Key: "enterprise", Value: objectId}, {Key: "creator", Value: objectIdUser}, {Key: "createdAt", Value: bson.D{{Key: "$gte", Value: ""}, {Key: "$lte", Value: ""}}}}}}
	_limit := bson.D{{Key: "$limit", Value: limit}}
	_skip := bson.D{{Key: "$skip", Value: limit * skip}}

	cursor, err := _collections.Aggregate(ctx, mongo.Pipeline{_match, _limit, _skip})
	if err != nil {
		return nil, err
	}
	var paiements []*model.Paiement
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var singlePaiement *model.Paiement
		if err = cursor.Decode(&singlePaiement); err != nil {
			log.Fatal(err)
		}
		paiements = append(paiements, singlePaiement)
	}

	count, err := _collections.CountDocuments(ctx, bson.D{{Key: "enterprise", Value: objectId}, {Key: "creator", Value: objectIdUser}, {Key: "createdAt", Value: bson.D{{Key: "$gte", Value: from}, {Key: "$lte", Value: to}}}})
	if err != nil {
		return nil, err
	}
	_count := float64(count)
	output := model.TransactionWithPageInfo{
		Transactions: paiements,
		PageTotal:    &_count,
	}

	return &output, nil
}

func CancelTransactionUser(ctx context.Context, transactionID string, typeArg model.PaymentType, pinCode string, canceller model.User, ip string) (bool, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("transactions")
	objectId, _ := primitive.ObjectIDFromHex(transactionID)
	_time := time.Now().UTC().Format(time.RFC3339)
	if typeArg == model.PaymentTypeTopup {
		_, err := finance.CancelWithDraw(transactionID, canceller.Keypair.SecretKey, canceller.Keypair.PublicKey)
		if err != nil {
			return false, err
		}
	} else if typeArg == model.PaymentTypeWithdraw {
		_, err := finance.CancelDeposit(transactionID, canceller.Keypair.SecretKey, canceller.Keypair.PublicKey)
		if err != nil {
			return false, err
		}
	}

	_, err := _collections.UpdateOne(ctx, bson.D{{Key: "_id", Value: objectId}}, bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: model.PaymentStatusCancelledUser}, {Key: "updatedAt", Value: _time}}}})
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetSuccessFullTransactionByEnterpriseId(ctx context.Context, enterpriseID string, from string, to string, limit float64, skip float64, userId string) (*model.TransactionWithPageInfo, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("transactions")
	objectId, _ := primitive.ObjectIDFromHex(enterpriseID)
	objectIdUser, _ := primitive.ObjectIDFromHex(userId)

	_match := bson.D{{Key: "$match", Value: bson.D{{Key: "enterprise", Value: objectId}, {Key: "creator", Value: objectIdUser}, {Key: "createdAt", Value: bson.D{{Key: "$gte", Value: from}, {Key: "$lte", Value: to}}}, {Key: "status", Value: model.PaymentStatusDone}}}}
	_limit := bson.D{{Key: "$limit", Value: limit}}
	_skip := bson.D{{Key: "$skip", Value: limit * skip}}

	cursor, err := _collections.Aggregate(ctx, mongo.Pipeline{_match, _limit, _skip})
	if err != nil {
		return nil, err
	}

	var paiements []*model.Paiement
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var singlePaiement *model.Paiement
		if err = cursor.Decode(&singlePaiement); err != nil {
			log.Fatal(err)
		}
		paiements = append(paiements, singlePaiement)
	}

	count, err := _collections.CountDocuments(ctx, bson.D{{Key: "enterprise", Value: objectId}, {Key: "creator", Value: objectIdUser}, {Key: "createdAt", Value: bson.D{{Key: "$gte", Value: from}, {Key: "$lte", Value: to}}}, {Key: "status", Value: model.PaymentStatusDone}})
	if err != nil {
		return nil, err
	}

	_count := float64(count)
	output := model.TransactionWithPageInfo{
		Transactions: paiements,
		PageTotal:    &_count,
	}

	return &output, nil
}

func GetRefundedTransactionByEnterpriseId(ctx context.Context, enterpriseID string, from string, to string, limit float64, skip float64, userId string) (*model.TransactionWithPageInfo, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("transactions")
	objectId, _ := primitive.ObjectIDFromHex(enterpriseID)
	objectIdUser, _ := primitive.ObjectIDFromHex(userId)

	_match := bson.D{{Key: "$match", Value: bson.D{{Key: "enterprise", Value: objectId}, {Key: "creator", Value: objectIdUser}, {Key: "createdAt", Value: bson.D{{Key: "$gte", Value: from}, {Key: "$lte", Value: to}}}, {Key: "status", Value: model.PaymentStatusRefunded}}}}
	_limit := bson.D{{Key: "$limit", Value: limit}}
	_skip := bson.D{{Key: "$skip", Value: limit * skip}}

	cursor, err := _collections.Aggregate(ctx, mongo.Pipeline{_match, _limit, _skip})
	if err != nil {
		return nil, err
	}

	var paiements []*model.Paiement
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var singlePaiement *model.Paiement
		if err = cursor.Decode(&singlePaiement); err != nil {
			log.Fatal(err)
		}
		paiements = append(paiements, singlePaiement)
	}

	count, err := _collections.CountDocuments(ctx, bson.D{{Key: "enterprise", Value: objectId}, {Key: "creator", Value: objectIdUser}, {Key: "createdAt", Value: bson.D{{Key: "$gte", Value: from}, {Key: "$lte", Value: to}}}, {Key: "status", Value: model.PaymentStatusRefunded}})
	if err != nil {
		return nil, err
	}
	_count := float64(count)
	output := model.TransactionWithPageInfo{
		Transactions: paiements,
		PageTotal:    &_count,
	}

	return &output, nil
}

func GetNonCapturedTransactionByEnterpriseId(ctx context.Context, enterpriseID string, from string, to string, limit float64, skip float64, userId string) (*model.TransactionWithPageInfo, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("transactions")
	objectId, _ := primitive.ObjectIDFromHex(enterpriseID)
	objectIdUser, _ := primitive.ObjectIDFromHex(userId)

	_match := bson.D{{Key: "$match", Value: bson.D{{Key: "enterprise", Value: objectId}, {Key: "creator", Value: objectIdUser}, {Key: "createdAt", Value: bson.D{{Key: "$gte", Value: from}, {Key: "$lte", Value: to}}}, {Key: "$or", Value: bson.A{bson.D{{Key: "status", Value: model.PaymentStatusRequiresAction}, {Key: "status", Value: model.PaymentStatusRequiresConfirmation}, {Key: "status", Value: model.PaymentStatusRequiresPaiement}, {Key: "status", Value: model.PaymentStatusOngoing}}}}}}}
	_limit := bson.D{{Key: "$limit", Value: limit}}
	_skip := bson.D{{Key: "$skip", Value: limit * skip}}

	cursor, err := _collections.Aggregate(ctx, mongo.Pipeline{_match, _limit, _skip})
	if err != nil {
		return nil, err
	}
	var paiements []*model.Paiement
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var singlePaiement *model.Paiement
		if err = cursor.Decode(&singlePaiement); err != nil {
			log.Fatal(err)
		}
		paiements = append(paiements, singlePaiement)
	}

	count, err := _collections.CountDocuments(ctx, bson.D{{Key: "enterprise", Value: objectId}, {Key: "creator", Value: objectIdUser}, {Key: "createdAt", Value: bson.D{{Key: "$gte", Value: from}, {Key: "$lte", Value: to}}}, {Key: "$or", Value: bson.A{bson.D{{Key: "status", Value: model.PaymentStatusRequiresAction}, {Key: "status", Value: model.PaymentStatusRequiresConfirmation}, {Key: "status", Value: model.PaymentStatusRequiresPaiement}, {Key: "status", Value: model.PaymentStatusOngoing}}}}})
	if err != nil {
		return nil, err
	}
	_count := float64(count)
	output := model.TransactionWithPageInfo{
		Transactions: paiements,
		PageTotal:    &_count,
	}

	return &output, nil
}

func GetFailedTransactionByEnterpriseId(ctx context.Context, enterpriseID string, from string, to string, limit float64, skip float64, userId string) (*model.TransactionWithPageInfo, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("transactions")
	objectId, _ := primitive.ObjectIDFromHex(enterpriseID)
	objectIdUser, _ := primitive.ObjectIDFromHex(userId)

	_match := bson.D{{Key: "$match", Value: bson.D{{Key: "enterprise", Value: objectId}, {Key: "creator", Value: objectIdUser}, {Key: "createdAt", Value: bson.D{{Key: "$gte", Value: from}, {Key: "$lte", Value: to}}}, {Key: "status", Value: model.PaymentStatusFailed}}}}
	_limit := bson.D{{Key: "$limit", Value: limit}}
	_skip := bson.D{{Key: "$skip", Value: limit * skip}}

	cursor, err := _collections.Aggregate(ctx, mongo.Pipeline{_match, _limit, _skip})
	if err != nil {
		return nil, err
	}
	var paiements []*model.Paiement
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var singlePaiement *model.Paiement
		if err = cursor.Decode(&singlePaiement); err != nil {
			log.Fatal(err)
		}
		paiements = append(paiements, singlePaiement)
	}

	count, err := _collections.CountDocuments(ctx, bson.D{{Key: "enterprise", Value: objectId}, {Key: "creator", Value: objectIdUser}, {Key: "createdAt", Value: bson.D{{Key: "$gte", Value: ""}, {Key: "$lte", Value: ""}}}, {Key: "status", Value: model.PaymentStatusFailed}})
	if err != nil {
		return nil, err
	}
	_count := float64(count)
	output := model.TransactionWithPageInfo{
		Transactions: paiements,
		PageTotal:    &_count,
	}

	return &output, nil
}

func buildChartDataWithArrayStatus(ctx context.Context, enterpriseID string, from string, to string, paiements []bson.M) []*float64 {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("transactions")
	var data []*float64
	if len(paiements) > 0 {
		total := paiements[0]["total"].(*float64)
		data = append(data, total)
	} else {
		f := float64(0)
		data = append(data, &f)
	}

	date, _ := time.Parse(time.RFC3339, from)
	_from := date
	_to, _ := time.Parse(time.RFC3339, to)
	objectId, _ := primitive.ObjectIDFromHex(enterpriseID)
	for i := 0; i <= 5; i++ {
		sub := _from.Add(-_from.Sub(_to))
		_match := bson.D{{Key: "$match", Value: bson.D{{Key: "enterprise", Value: objectId}, {Key: "createdAt", Value: bson.D{{Key: "$gte", Value: sub.UTC().Format(time.RFC3339)}, {Key: "$lte", Value: _from.UTC().Format(time.RFC3339)}}}, {Key: "status", Value: bson.A{model.PaymentStatusRequiresConfirmation.String(), model.PaymentStatusRequiresAction.String(), model.PaymentStatusRequiresPaiement.String()}}}}}
		_group := bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: nil}, {Key: "total", Value: bson.D{{Key: "$sum", Value: bson.D{{Key: "$toInt", Value: "$amount"}}}}}}}}
		cursor, err := _collections.Aggregate(ctx, mongo.Pipeline{_match, _group})
		if err != nil {
			log.Fatal(err)
		}
		var results []bson.M
		for cursor.Next(ctx) {
			var singlePaiement *bson.M
			if err = cursor.Decode(&singlePaiement); err != nil {
				log.Fatal(err)
			}
			results = append(results, *singlePaiement)
		}

		if len(results) > 0 {
			total := results[0]["total"].(*float64)
			data = append(data, total)
		} else {
			f := float64(0)
			data = append(data, &f)
		}
		_from = sub
	}

	return data

}

func buildChartData(ctx context.Context, enterpriseID string, from string, to string, paiements []bson.M) []*float64 {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("transactions")
	var data []*float64
	if len(paiements) > 0 {
		total := paiements[0]["total"].(float64)
		data = append(data, &total)
	} else {
		f := float64(0)
		data = append(data, &f)
	}

	date, _ := time.Parse(time.RFC3339, from)
	_from := date
	_to, _ := time.Parse(time.RFC3339, to)
	objectId, _ := primitive.ObjectIDFromHex(enterpriseID)
	for i := 0; i <= 5; i++ {
		sub := _from.Add(-_from.Sub(_to))
		_match := bson.D{{Key: "$match", Value: bson.D{{Key: "enterprise", Value: objectId}, {Key: "createdAt", Value: bson.D{{Key: "$gte", Value: sub.UTC().Format(time.RFC3339)}, {Key: "$lte", Value: _from.UTC().Format(time.RFC3339)}}}, {Key: "status", Value: model.PaymentStatusDone.String()}}}}
		_group := bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: nil}, {Key: "total", Value: bson.D{{Key: "$sum", Value: bson.D{{Key: "$toInt", Value: "$amount"}}}}}}}}
		cursor, err := _collections.Aggregate(ctx, mongo.Pipeline{_match, _group})
		if err != nil {
			log.Fatal(err)
		}
		var results []bson.M
		for cursor.Next(ctx) {
			var singlePaiement *bson.M
			if err = cursor.Decode(&singlePaiement); err != nil {
				log.Fatal(err)
			}
			results = append(results, *singlePaiement)
		}

		if len(results) > 0 {
			total := results[0]["total"].(float64)
			data = append(data, &total)
		} else {
			f := float64(0)
			data = append(data, &f)
		}
		_from = sub
	}
	return data
}

func diffPercent(current float64, former float64) *float64 {
	if current < former {
		_former := big.NewFloat(former)
		_current := big.NewFloat(current)
		sub := new(big.Float).Sub(_former, _current)
		mul := new(big.Float).Mul(sub, big.NewFloat(100))
		div := new(big.Float).Quo(mul, _current)
		res := new(big.Float).Mul(div, big.NewFloat(-1))
		output, _ := res.Float64()

		return &output
	}
	_former := big.NewFloat(former)
	_current := big.NewFloat(current)
	sub := new(big.Float).Sub(_current, _former)
	mul := new(big.Float).Mul(sub, big.NewFloat(100))
	div := new(big.Float).Quo(mul, _former)
	output, _ := div.Float64()
	return &output

}

func GetProfilNetChartData(ctx context.Context, enterpriseID string, from string, to string, userId string) (*model.ChartData, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("transactions")
	objectId, _ := primitive.ObjectIDFromHex(enterpriseID)

	_match := bson.D{{Key: "$match", Value: bson.D{{Key: "enterprise", Value: objectId}, {Key: "createdAt", Value: bson.D{{Key: "$gte", Value: from}, {Key: "$lte", Value: to}}}, {Key: "status", Value: model.PaymentStatusDone.String()}}}}
	_group := bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: nil}, {Key: "total", Value: bson.D{{Key: "$sum", Value: bson.D{{Key: "$toDouble", Value: "$amount"}}}}}}}}
	cursor, err := _collections.Aggregate(ctx, mongo.Pipeline{_match, _group})
	if err != nil {
		return nil, err
	}
	var _currentTotal []bson.M
	for cursor.Next(ctx) {
		var singlePaiement *bson.M
		if err = cursor.Decode(&singlePaiement); err != nil {
			log.Fatal(err)
		}
		_currentTotal = append(_currentTotal, *singlePaiement)
	}

	_from, _ := time.Parse(time.RFC3339, from)
	_to, _ := time.Parse(time.RFC3339, to)
	sub := _from.Add(-_from.Sub(_to))
	_match_2 := bson.D{{Key: "$match", Value: bson.D{{Key: "enterprise", Value: objectId}, {Key: "createdAt", Value: bson.D{{Key: "$gte", Value: sub.UTC().Format(time.RFC3339)}, {Key: "$lte", Value: from}}}, {Key: "status", Value: model.PaymentStatusDone.String()}}}}
	_group_2 := bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: nil}, {Key: "total", Value: bson.D{{Key: "$sum", Value: bson.D{{Key: "$toDouble", Value: "$amount"}}}}}}}}
	cursor_2, err := _collections.Aggregate(ctx, mongo.Pipeline{_match_2, _group_2})
	if err != nil {
		return nil, err
	}
	var _lastperiodTotal []bson.M
	for cursor_2.Next(ctx) {
		var singlePaiement *bson.M
		if err = cursor.Decode(&singlePaiement); err != nil {
			log.Fatal(err)
		}
		_lastperiodTotal = append(_lastperiodTotal, *singlePaiement)
	}

	chart := buildChartData(ctx, enterpriseID, from, to, _currentTotal)
	var current float64
	if len(_currentTotal) > 0 {
		current = _currentTotal[0]["total"].(float64)
	} else {
		current = 0
	}

	var former float64
	if len(_lastperiodTotal) > 0 {
		former = _lastperiodTotal[0]["total"].(float64)
	} else {
		former = 0
	}

	var diff *float64
	if current > 0 && former > 0 {
		diff = diffPercent(current, former)
	} else {
		diff = new(float64)
	}

	var sign bool
	if math.Signbit(*diff) {
		sign = false
	} else {
		sign = true
	}

	_diff := math.Abs(*diff)
	output := model.ChartData{
		CurrentTotal:          &current,
		FormerTotal:           &former,
		PourcentageDifference: &_diff,
		IsPositive:            &sign,
		Chart:                 reverseFloats(chart),
	}

	return &output, nil
}

func GetProfilBrutChartData(ctx context.Context, enterpriseID string, from string, to string, userId string) (*model.ChartData, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("transactions")
	objectId, _ := primitive.ObjectIDFromHex(enterpriseID)

	_match := bson.D{{Key: "$match", Value: bson.D{{Key: "enterprise", Value: objectId}, {Key: "createdAt", Value: bson.D{{Key: "$gte", Value: from}, {Key: "$lte", Value: to}}}, {Key: "status", Value: model.PaymentStatusDone.String()}}}}
	_group := bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: nil}, {Key: "total", Value: bson.D{{Key: "$sum", Value: bson.D{{Key: "$toInt", Value: "$amount"}}}}}}}}
	cursor, err := _collections.Aggregate(ctx, mongo.Pipeline{_match, _group})
	if err != nil {
		return nil, err
	}
	var _currentTotal []bson.M
	for cursor.Next(ctx) {
		var singlePaiement *bson.M
		if err = cursor.Decode(&singlePaiement); err != nil {
			log.Fatal(err)
		}
		_currentTotal = append(_currentTotal, *singlePaiement)
	}

	_from, err := time.Parse(time.RFC3339, from)
	if err != nil {
		return nil, err
	}
	_to, err := time.Parse(time.RFC3339, to)
	if err != nil {
		return nil, err
	}
	sub := _from.Add(-_from.Sub(_to))
	_match_2 := bson.D{{Key: "$match", Value: bson.D{{Key: "enterprise", Value: objectId}, {Key: "createdAt", Value: bson.D{{Key: "$gte", Value: sub.UTC().Format(time.RFC3339)}, {Key: "$lte", Value: from}}}, {Key: "status", Value: model.PaymentStatusDone.String()}}}}
	_group_2 := bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: nil}, {Key: "total", Value: bson.D{{Key: "$sum", Value: bson.D{{Key: "$toInt", Value: "$amount"}}}}}}}}
	cursor_2, err := _collections.Aggregate(ctx, mongo.Pipeline{_match_2, _group_2})
	if err != nil {
		return nil, err
	}
	var _lastperiodTotal []bson.M
	for cursor_2.Next(ctx) {
		var singlePaiement *bson.M
		if err = cursor.Decode(&singlePaiement); err != nil {
			log.Fatal(err)
		}
		_lastperiodTotal = append(_lastperiodTotal, *singlePaiement)
	}

	chart := buildChartData(ctx, enterpriseID, from, to, _currentTotal)
	var current float64
	if len(_currentTotal) > 0 {
		current = _currentTotal[0]["total"].(float64)
	} else {
		current = 0
	}

	var former float64
	if len(_lastperiodTotal) > 0 {
		former = _lastperiodTotal[0]["total"].(float64)
	} else {
		former = 0
	}

	var diff *float64
	if current > 0 && former > 0 {
		diff = diffPercent(current, former)
	} else {
		diff = new(float64)
	}

	var sign bool
	if math.Signbit(*diff) {
		sign = false
	} else {
		sign = true
	}

	_diff := math.Abs(*diff)
	output := model.ChartData{
		CurrentTotal:          &current,
		FormerTotal:           &former,
		PourcentageDifference: &_diff,
		IsPositive:            &sign,
		Chart:                 reverseFloats(chart),
	}

	return &output, nil

}

func GetProfilNonCarpturedChartData(ctx context.Context, enterpriseID string, from string, to string, userId string) (*model.ChartData, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("transactions")
	objectId, _ := primitive.ObjectIDFromHex(enterpriseID)

	_match := bson.D{{Key: "$match", Value: bson.D{{Key: "enterprise", Value: objectId}, {Key: "createdAt", Value: bson.D{{Key: "$gte", Value: from}, {Key: "$lte", Value: to}}}, {Key: "status", Value: bson.A{model.PaymentStatusRequiresConfirmation.String(), model.PaymentStatusRequiresAction.String(), model.PaymentStatusRequiresPaiement.String()}}}}}
	_group := bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: nil}, {Key: "total", Value: bson.D{{Key: "$sum", Value: bson.D{{Key: "$toInt", Value: "$amount"}}}}}}}}
	cursor, err := _collections.Aggregate(ctx, mongo.Pipeline{_match, _group})
	if err != nil {
		return nil, err
	}
	var _currentTotal []bson.M
	for cursor.Next(ctx) {
		var singlePaiement *bson.M
		if err = cursor.Decode(&singlePaiement); err != nil {
			log.Fatal(err)
		}
		_currentTotal = append(_currentTotal, *singlePaiement)
	}

	_from, err := time.Parse(time.RFC3339, from)
	if err != nil {
		return nil, err
	}
	_to, err := time.Parse(time.RFC3339, to)
	if err != nil {
		return nil, err
	}
	sub := _from.Add(-_from.Sub(_to))
	_match_2 := bson.D{{Key: "$match", Value: bson.D{{Key: "enterprise", Value: objectId}, {Key: "createdAt", Value: bson.D{{Key: "$gte", Value: sub.UTC().Format(time.RFC3339)}, {Key: "$lte", Value: from}}}, {Key: "status", Value: bson.A{model.PaymentStatusRequiresConfirmation.String(), model.PaymentStatusRequiresAction.String(), model.PaymentStatusRequiresPaiement.String()}}}}}
	_group_2 := bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: nil}, {Key: "total", Value: bson.D{{Key: "$sum", Value: bson.D{{Key: "$toInt", Value: "$amount"}}}}}}}}
	cursor_2, err := _collections.Aggregate(ctx, mongo.Pipeline{_match_2, _group_2})
	if err != nil {
		return nil, err
	}
	var _lastperiodTotal []bson.M
	for cursor_2.Next(ctx) {
		var singlePaiement *bson.M
		if err = cursor.Decode(&singlePaiement); err != nil {
			log.Fatal(err)
		}
		_lastperiodTotal = append(_lastperiodTotal, *singlePaiement)
	}

	chart := buildChartDataWithArrayStatus(ctx, enterpriseID, from, to, _currentTotal)
	var current float64
	if len(_currentTotal) > 0 {
		current = _currentTotal[0]["total"].(float64)
	} else {
		current = 0
	}

	var former float64
	if len(_lastperiodTotal) > 0 {
		former = _lastperiodTotal[0]["total"].(float64)
	} else {
		former = 0
	}

	var diff *float64
	if current > 0 && former > 0 {
		diff = diffPercent(current, former)
	} else {
		diff = new(float64)
	}

	var sign bool
	if math.Signbit(*diff) {
		sign = false
	} else {
		sign = true
	}

	_diff := math.Abs(*diff)
	output := model.ChartData{
		CurrentTotal:          &current,
		FormerTotal:           &former,
		PourcentageDifference: &_diff,
		IsPositive:            &sign,
		Chart:                 reverseFloats(chart),
	}

	return &output, nil

}

func AuthenticateForTransaction(ctx context.Context, amount float64, ref *string, enterprise *model.Enterprise, ip string) (model.QRCodeOwner, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("transactions")
	objectId, _ := primitive.ObjectIDFromHex(enterprise.ID)

	fee := finance.CalculateFee(amount)
	toWei := finance.ToWei(fee)
	_amountToWei := finance.ToWei(amount)
	_time := time.Now().UTC().Format(time.RFC3339)
	newId := primitive.NewObjectID()
	input := database.DBTransaction{
		ID:            &newId,
		Destination:   enterprise.WalletPublicKey,
		Fee:           big.NewInt(0).String(),
		FeeEnterprise: toWei.String(),
		Amount:        _amountToWei.String(),
		Token:         os.Getenv("DEFAULT_CURRENCY"),
		Type:          model.PaymentTypeCommerce.String(),
		Status:        model.PaymentStatusRequiresPaiement.String(),
		EnterpriseID:  &objectId,
		Country:       "CG",
		CreatedAt:     _time,
		UpdatedAt:     _time,
	}

	result, err := _collections.InsertOne(ctx, input)
	if err != nil {
		return nil, err
	}

	insertedID := result.InsertedID.(primitive.ObjectID).Hex()
	receipt, err := finance.AddUnConfirmedTransaction(insertedID, enterprise.WalletPublicKey, _amountToWei)
	if err != nil && receipt.TxHash.String() != "" {
		return nil, err
	}

	payment, err := GetTransactionByIdUnauthed(ctx, insertedID)
	if err != nil {
		return nil, err
	}
	return model.QRCodeOwner(payment), err

}
