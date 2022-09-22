package enterprise

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"server/database"
	"server/graph/model"
	"server/transaction"
	"server/utils"
	"time"

	"server/finance"

	snippets "server/firebase"
	"server/pdf"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	utils.LoadEnv()
}

func CreateEnterprise(ctx context.Context, enterprise model.EnterpriseInput, owner model.User, ip string) (*model.Enterprise, error) {
	passed, err := bulkDefaultEnterpriseToFalse(ctx, owner.ID)
	if err != nil && !passed {
		return nil, err
	}

	privateKey, publicKey, isError := finance.CreateAccount()
	if isError {

	}
	publishableKey := utils.Ase256Encode(os.Getenv("SERVER_SECRET_KEY"), uuid.NewString())
	private_key := utils.Ase256Encode(os.Getenv("SERVER_SECRET_KEY"), uuid.NewString())

	secret, err := utils.Encrypt(privateKey, os.Getenv("SERVER_SECRET"))
	if err != nil {
		return nil, err
	}

	bytes, err := pdf.CreatePdfFile(publicKey)
	if err != nil {
		return nil, err
	}
	url, err := snippets.Connect().UploadFile(bytes.Bytes(), "enterprises/"+*owner.FirebaseUID+"/"+publicKey+".pdf")
	if err != nil {
		return nil, err
	}

	_time := time.Now().UTC().Format(time.RFC3339)
	objectId, _ := primitive.ObjectIDFromHex(owner.ID)

	enterpriseDB := database.DBEnterprise{
		Type:            string(enterprise.Type),
		Name:            enterprise.Name,
		Rccm:            enterprise.Rccm,
		Website:         *enterprise.Website,
		LogoUrl:         *enterprise.LogoURL,
		Creator:         objectId,
		CreatedAt:       _time,
		UpdatedAt:       _time,
		PublishableKey:  publishableKey,
		Private_key:     private_key,
		WalletPublicKey: publicKey,
		WalletSecretKey: secret,
		Country:         enterprise.Country,
		Address: database.DBAddress{
			Title: *enterprise.Address.Title,
			Location: database.DBLocation{
				Longitude: 0,
				Latitude:  0,
			},
			City:      *enterprise.Address.City,
			Zip:       *enterprise.Person.Zip,
			Country:   enterprise.Country,
			Is_chosed: true,
		},
		Removed:              false,
		DeletedAt:            "",
		Default_enterprise:   true,
		Sector:               enterprise.ActivitySector,
		Description:          *enterprise.Description,
		SellingPhysicalGoods: *enterprise.SellingPhysicalGoods,
		TransactionLibele:    enterprise.TransactionLibele,
		AbregedLibele:        enterprise.AbregedLibele,
		Phone:                enterprise.Phone,
		Email:                enterprise.Email,
		Bordereau:            url,
		Person: database.DBPerson{
			First_name: enterprise.Person.FirstName,
			Last_name:  enterprise.Person.LastName,
			Email:      enterprise.Person.Email,
			Address:    enterprise.Person.Address,
			City:       enterprise.Person.City,
			Zip:        *enterprise.Person.Zip,
			State:      *enterprise.Person.State,
		},
	}
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("entreprises")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := _collections.InsertOne(ctx, enterpriseDB)
	if err != nil {
		return nil, err
	}

	insertedID := res.InsertedID.(primitive.ObjectID).Hex()
	return GetEnterpriseByIdWithUseriD(ctx, insertedID, owner.ID)

}

func bulkDefaultEnterpriseToFalse(ctx context.Context, userId string) (bool, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("entreprises")
	objectId, _ := primitive.ObjectIDFromHex(userId)

	_match := bson.D{{Key: "creator", Value: objectId}, {Key: "default_enterprise", Value: true}}

	_time := time.Now().UTC().Format(time.RFC3339)
	_update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "default_enterprise", Value: false},
		{Key: "updatedAt", Value: _time}}}}

	_, err := _collections.UpdateMany(ctx, _match, _update)
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetAllEnterpriseForAUser(ctx context.Context, userId string) ([]*model.Enterprise, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("entreprises")
	objectId, _ := primitive.ObjectIDFromHex(userId)

	cursor, err := _collections.Find(ctx, bson.D{
		{Key: "creator", Value: objectId},
		{Key: "removed", Value: false}})
	if err != nil {
		return nil, err
	}

	var enterprises []*model.Enterprise
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var singleEnterprise *model.Enterprise
		if err = cursor.Decode(&singleEnterprise); err != nil {
			log.Fatal(err)
		}
		enterprises = append(enterprises, singleEnterprise)
	}

	return enterprises, nil

}

func ChangeDefaultEnterprise(ctx context.Context, enterpriseId, userId string, ip string) ([]*model.Enterprise, error) {
	_, err := bulkDefaultEnterpriseToFalse(ctx, userId)
	if err != nil {
		return nil, err
	}

	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("entreprises")
	objectId, _ := primitive.ObjectIDFromHex(enterpriseId)
	_match := bson.D{
		{Key: "_id", Value: objectId},
		{Key: "removed", Value: false}}

	_time := time.Now().UTC().Format(time.RFC3339)
	_update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "default_enterprise", Value: true},
		{Key: "updatedAt", Value: _time}}}}

	result, err := _collections.UpdateOne(ctx, _match, _update)
	if err != nil {
		return nil, err
	}

	if result.ModifiedCount == 0 {
		return nil, fmt.Errorf("could not modify this file")
	}

	return GetAllEnterpriseForAUser(ctx, userId)
}

func GetEnterpriseByIdWithUseriD(ctx context.Context, enterpriseId string, userID string) (*model.Enterprise, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("entreprises")
	objectIdEnterprise, _ := primitive.ObjectIDFromHex(enterpriseId)
	objectId, _ := primitive.ObjectIDFromHex(userID)

	var enterprise *model.Enterprise
	err := _collections.FindOne(ctx, bson.D{{Key: "_id", Value: objectIdEnterprise}, {Key: "creator", Value: objectId}, {Key: "removed", Value: false}}).Decode(&enterprise)
	if err != nil {
		return nil, err
	}
	return enterprise, nil
}

func GetEnterpriseById(ctx context.Context, enterpriseId string) (*model.Enterprise, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("entreprises")
	objectIdEnterprise, _ := primitive.ObjectIDFromHex(enterpriseId)

	var enterprise *model.Enterprise
	err := _collections.FindOne(ctx, bson.D{{Key: "_id", Value: objectIdEnterprise}, {Key: "removed", Value: false}}).Decode(&enterprise)
	if err != nil {
		return nil, err
	}

	return enterprise, nil
}

func RemoveEnterprise(ctx context.Context, enterpriseId string, userID string, ip string) ([]*model.Enterprise, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("entreprises")
	objectIdEnterprise, _ := primitive.ObjectIDFromHex(enterpriseId)
	objectId, _ := primitive.ObjectIDFromHex(userID)

	_time := time.Now().UTC().Format(time.RFC3339)
	_, err := _collections.UpdateOne(ctx, bson.D{{Key: "_id", Value: objectIdEnterprise}, {Key: "creator", Value: objectId}, {Key: "removed", Value: false}}, bson.D{{Key: "$set", Value: bson.D{{Key: "removed", Value: true}, {Key: "removedAt", Value: _time}, {Key: "updatedAt", Value: _time}}}})
	if err != nil {
		return nil, err
	}

	cursor, err := GetAllEnterpriseForAUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	if len(cursor) > 0 {
		id, _ := primitive.ObjectIDFromHex(cursor[0].ID)
		_collections.UpdateOne(ctx, bson.D{{Key: "_id", Value: id}, {Key: "creator", Value: objectId}, {Key: "removed", Value: false}}, bson.D{{Key: "$set", Value: bson.D{{Key: "default_enterprise", Value: true}, {Key: "updatedAt", Value: _time}}}})
	} else {
		return cursor, nil
	}
	return GetAllEnterpriseForAUser(ctx, userID)
}

func GetEnterpriseBalance(ctx context.Context, enterpriseId string, userID string) *float64 {
	enterprise, err := GetEnterpriseByIdWithUseriD(ctx, enterpriseId, userID)
	if err != nil {
		return nil
	}

	balance, err := finance.GetBalanceOnContractInstance(enterprise.WalletPublicKey)
	if err != nil {
		return nil
	}

	result, _ := balance.Float64()
	return &result
}

func CancelEnterpriseTransaction(ctx context.Context, enterpriseId string, transactionnID string, ip string) (*bool, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("transactions")
	objectIdEnterprise, _ := primitive.ObjectIDFromHex(enterpriseId)
	objectIdTransaction, _ := primitive.ObjectIDFromHex(transactionnID)

	_time := time.Now().UTC().Format(time.RFC3339)
	var constraint []model.PaymentStatus
	constraint = append(constraint, model.PaymentStatusRequiresPaiement)
	constraint = append(constraint, model.PaymentStatusRequiresConfirmation)
	constraint = append(constraint, model.PaymentStatusRequiresAction)
	constraint = append(constraint, model.PaymentStatusOngoing)
	constraint = append(constraint, model.PaymentStatusInProgress)
	_, err := _collections.UpdateOne(ctx, bson.D{{Key: "_id", Value: objectIdTransaction}, {Key: "enterprise", Value: objectIdEnterprise}, {Key: "status", Value: bson.D{{Key: "$in", Value: constraint}}}}, bson.D{{Key: "status", Value: "CANCELLED"}, {Key: "updated", Value: _time}})
	if err != nil {
		return nil, err
	}

	result := true
	return &result, nil
}

func SendMoney(ctx context.Context, enterpriseId string, pinCode string, publicKey string, amount float64, userID string, enterprise *model.Enterprise, ip string) (*bool, error) {
	_amount := finance.ToWei(amount)
	_receipt, err := finance.Transfer(enterprise.WalletPublicKey, publicKey, _amount, *enterprise.WalletSecretKey)
	if err != nil {
		return nil, err
	}
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("transactions")
	txHash := _receipt.BlockHash.Hex()
	objectIdEnterprise, _ := primitive.ObjectIDFromHex(enterpriseId)
	objectIduser, _ := primitive.ObjectIDFromHex(userID)
	_time := time.Now().UTC().Format(time.RFC3339)

	input := database.DBTransaction{
		TransactionId: txHash,
		Source:        enterprise.WalletPublicKey,
		Destination:   publicKey,
		Fee:           big.NewInt(0).String(),
		FeeEnterprise: finance.ToWei(finance.CalculateFee(amount)).String(),
		Amount:        _amount.String(),
		Token:         os.Getenv("DEFAULT_CURRENCY"),
		Description:   "",
		//DestinationUser: &[12]byte{},
		ShortId:      "",
		Type:         model.PaymentTypeTransfert.String(),
		Status:       model.PaymentStatusDone.String(),
		CreatorID:    &objectIduser,
		EnterpriseID: &objectIdEnterprise,
		Country:      "",
		CreatedAt:    _time,
		UpdatedAt:    _time,
		Ip: &database.ConnectionDB{
			IpAddress: ip,
			CreatedAt: _time,
			DeviceId:  "",
		},
	}

	_result, err := _collections.InsertOne(ctx, input)
	if err != nil {
		return nil, err
	}

	result := false
	if _result.InsertedID != nil {
		result = true
	}

	return &result, nil

}

func PayEnterprise(ctx context.Context, enterpriseId string, amount float64, pinCode string, user model.User, ip string) (*model.Paiement, error) {
	enterprise, err := GetEnterpriseById(ctx, enterpriseId)
	if err != nil {
		return nil, err
	}

	id := primitive.NewObjectID()
	_amount := finance.ToWei(amount)
	receipt, err := finance.CommercePay(id.Hex(), enterprise.WalletPublicKey, _amount, user.Keypair.SecretKey, user.Keypair.PublicKey)
	if err != nil {
		return nil, err
	}

	txhash := receipt.BlockHash.Hex()
	_time := time.Now().UTC().Format(time.RFC3339)

	objectIdEnterprise, _ := primitive.ObjectIDFromHex(enterpriseId)
	objectIduser, _ := primitive.ObjectIDFromHex(user.ID)

	input := database.DBTransaction{
		ID:            &id,
		TransactionId: txhash,
		Source:        user.Keypair.PublicKey,
		Destination:   enterprise.WalletPublicKey,
		Fee:           big.NewInt(0).String(),
		FeeEnterprise: finance.ToWei(finance.CalculateFee(amount)).String(),
		Amount:        _amount.String(),
		Token:         os.Getenv("DEFAULT_CURRENCY"),
		Description:   "",
		Type:          model.PaymentTypeCommerce.String(),
		Status:        model.PaymentStatusDone.String(),
		CreatorID:     &objectIduser,
		EnterpriseID:  &objectIdEnterprise,
		Country:       "",
		CreatedAt:     _time,
		UpdatedAt:     _time,
	}
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("transactions")
	_result, err := _collections.InsertOne(ctx, input)
	if err != nil {
		return nil, err
	}

	insertedID := _result.InsertedID.(primitive.ObjectID).Hex()

	return transaction.GetTransactionByIdUnauthed(ctx, insertedID)
}

func RecreatePublishableKey(ctx context.Context, enterpriseId string, userID string, ip string) ([]*model.Enterprise, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("entreprises")
	objectIdEnterprise, _ := primitive.ObjectIDFromHex(enterpriseId)
	objectId, _ := primitive.ObjectIDFromHex(userID)

	_time := time.Now().UTC().Format(time.RFC3339)
	encrypted := utils.Ase256Encode(os.Getenv("SERVER_SECRET_KEY"), uuid.NewString())
	_collections.UpdateOne(ctx, bson.D{{Key: "_id", Value: objectIdEnterprise}, {Key: "creator", Value: objectId}}, bson.D{{Key: "$set", Value: bson.D{{Key: "publishableKey", Value: encrypted}, {Key: "updatedAt", Value: _time}}}})
	return GetAllEnterpriseForAUser(ctx, userID)
}

func RcreatePrivateKey(ctx context.Context, enterpriseId string, userID string, ip string) ([]*model.Enterprise, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("entreprises")
	objectIdEnterprise, _ := primitive.ObjectIDFromHex(enterpriseId)
	objectId, _ := primitive.ObjectIDFromHex(userID)

	_time := time.Now().UTC().Format(time.RFC3339)
	encrypted := utils.Ase256Encode(os.Getenv("SERVER_SECRET_KEY"), uuid.NewString())
	_collections.UpdateOne(ctx, bson.D{{Key: "_id", Value: objectIdEnterprise}, {Key: "creator", Value: objectId}}, bson.D{{Key: "$set", Value: bson.D{{Key: "private_key", Value: encrypted}, {Key: "updatedAt", Value: _time}}}})
	return GetAllEnterpriseForAUser(ctx, userID)
}

func KeysExist(ctx context.Context, publishableKey string, private_key string) (*model.Enterprise, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("entreprises")
	encrypted := utils.Ase256Encode(os.Getenv("SERVER_SECRET_KEY"), publishableKey)
	encrypted_2 := utils.Ase256Encode(os.Getenv("SERVER_SECRET_KEY"), private_key)

	var enterprise *model.Enterprise

	failure := _collections.FindOne(ctx, bson.D{{Key: "publishableKey", Value: encrypted}, {Key: "private_key", Value: encrypted_2}, {Key: "removed", Value: false}}).Decode(&enterprise)
	if failure != nil {
		return nil, failure
	}

	return enterprise, nil
}

func UpdateEnterpriseType(ctx context.Context, enterpriseId string, userID string, typeArg string, country string, ip string) ([]*model.Enterprise, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("entreprises")
	objectIdEnterprise, _ := primitive.ObjectIDFromHex(enterpriseId)
	objectId, _ := primitive.ObjectIDFromHex(userID)

	_time := time.Now().UTC().Format(time.RFC3339)
	_update := bson.D{{Key: "$set", Value: bson.D{{Key: "type", Value: typeArg}, {Key: "country", Value: country}, {Key: "updatedAt", Value: _time}}}}

	_, err := _collections.UpdateOne(ctx, bson.D{
		{Key: "_id", Value: objectIdEnterprise}, {Key: "creator", Value: objectId},
	}, _update)
	if err != nil {
		return nil, err
	}

	return GetAllEnterpriseForAUser(ctx, userID)

}

func UpdatePersonnalInformation(ctx context.Context, enterpriseId string, userID string, firstName string, lastName string, email string, address string, city string, state string, zip string, ip string) ([]*model.Enterprise, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("entreprises")
	objectIdEnterprise, _ := primitive.ObjectIDFromHex(enterpriseId)
	objectId, _ := primitive.ObjectIDFromHex(userID)

	_time := time.Now().UTC().Format(time.RFC3339)
	_update := bson.D{{Key: "$set", Value: bson.D{{Key: "person", Value: bson.D{
		{Key: "first_name", Value: firstName},
		{Key: "last_name", Value: lastName},
		{Key: "email", Value: email},
		{Key: "address", Value: address},
		{Key: "city", Value: city},
		{Key: "state", Value: state},
		{Key: "zip", Value: zip},
	}}, {Key: "updatedAt", Value: _time}}}}

	_, err := _collections.UpdateOne(ctx, bson.D{{Key: "_id", Value: objectIdEnterprise}, {Key: "creator", Value: objectId}}, _update)
	if err != nil {
		return nil, err
	}

	return GetAllEnterpriseForAUser(ctx, userID)

}

func UpdateEnterpriseInformation(ctx context.Context, enterpriseId string, userID string, rccm string, sector string, website *string, description *string, ip string) ([]*model.Enterprise, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("entreprises")
	objectIdEnterprise, _ := primitive.ObjectIDFromHex(enterpriseId)
	objectId, _ := primitive.ObjectIDFromHex(userID)

	_time := time.Now().UTC().Format(time.RFC3339)
	_update := bson.D{{Key: "$set", Value: bson.D{{Key: "rccm", Value: rccm}, {Key: "sector", Value: sector}, {Key: "website", Value: website}, {Key: "description", Value: description}, {Key: "updatedAt", Value: _time}}}}

	_, err := _collections.UpdateOne(ctx, bson.D{
		{Key: "_id", Value: objectIdEnterprise}, {Key: "creator", Value: objectId},
	}, _update)
	if err != nil {
		return nil, err
	}

	return GetAllEnterpriseForAUser(ctx, userID)

}

func UpdateExecutionInformation(ctx context.Context, enterpriseId string, userID string, sellingPyshicalGoods *bool, selfShipping *bool, shippingDelay *string, ip string) ([]*model.Enterprise, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("entreprises")
	objectIdEnterprise, _ := primitive.ObjectIDFromHex(enterpriseId)
	objectId, _ := primitive.ObjectIDFromHex(userID)

	_time := time.Now().UTC().Format(time.RFC3339)
	_update := bson.D{{Key: "$set", Value: bson.D{{Key: "sellingPhysicalGoods", Value: sellingPyshicalGoods}, {Key: "selfShippingProduct", Value: selfShipping}, {Key: "shippingDelay", Value: shippingDelay}, {Key: "updatedAt", Value: _time}}}}

	_, err := _collections.UpdateOne(ctx, bson.D{
		{Key: "_id", Value: objectIdEnterprise}, {Key: "creator", Value: objectId},
	}, _update)
	if err != nil {
		return nil, err
	}

	return GetAllEnterpriseForAUser(ctx, userID)

}

func UpdatePublicInformation(ctx context.Context, enterpriseId string, userID string, name string, libelle string, libelleAbreged string, email *string, phone string, ip string) ([]*model.Enterprise, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("entreprises")
	objectIdEnterprise, _ := primitive.ObjectIDFromHex(enterpriseId)
	objectId, _ := primitive.ObjectIDFromHex(userID)

	_time := time.Now().UTC().Format(time.RFC3339)
	_update := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: name}, {Key: "transactionLibele", Value: libelle}, {Key: "abregedLibele", Value: libelleAbreged}, {Key: "email", Value: email}, {Key: "phone", Value: phone}, {Key: "updatedAt", Value: _time}}}}

	_, err := _collections.UpdateOne(ctx, bson.D{
		{Key: "_id", Value: objectIdEnterprise}, {Key: "creator", Value: objectId},
	}, _update)
	if err != nil {
		return nil, err
	}

	return GetAllEnterpriseForAUser(ctx, userID)
}

func PayUnConfirmedTransaction(ctx context.Context, enterpriseId string, transactionID string, user model.User, ip string) (*model.Paiement, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("transactions")
	enterprise, err := GetEnterpriseById(ctx, enterpriseId)
	if err != nil {
		return nil, err
	}

	receipt, err := finance.ConfirmUnformedTransaction(transactionID, user.Keypair.SecretKey, user.Keypair.PublicKey)
	if err != nil {
		return nil, err
	}

	objectId, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return nil, err
	}
	objectIdEnterprise, err := primitive.ObjectIDFromHex(enterprise.ID)
	if err != nil {
		return nil, err
	}

	tx_time := time.Now().UTC().Format(time.RFC3339)

	objectID, _ := primitive.ObjectIDFromHex(transactionID)
	f, err := _collections.UpdateOne(ctx, bson.M{"_id": objectID}, bson.D{{Key: "$set", Value: bson.M{"transactionID": receipt.BlockHash.Hex(), "status": model.PaymentStatusDone.String(),
		"updatedAt": tx_time, "enterpriseId": objectIdEnterprise, "destination": enterprise.WalletPublicKey, "source": user.Keypair.PublicKey, "creatorId": objectId, "FeeEnterprise": 0}}})
	if err != nil {
		return nil, err
	}

	if f.ModifiedCount == 0 {
		return nil, fmt.Errorf("no dcoument was updated")
	}

	return transaction.GetTransactionByIdUnauthed(ctx, transactionID)

}

func RefundTransaction(ctx context.Context, enterpriseId string, transactionID string, enterprise model.Enterprise, ip string) (*bool, error) {
	receipt, err := finance.RefundCommercePay(transactionID, *enterprise.WalletSecretKey, enterprise.WalletPublicKey)
	if err != nil {
		return nil, err
	}

	objectId, _ := primitive.ObjectIDFromHex(transactionID)

	_time := time.Now().UTC().Format(time.RFC3339)
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("transactions")
	result, err := _collections.UpdateOne(ctx, bson.M{"_id": objectId}, bson.D{{Key: "$set", Value: bson.M{"status": model.PaymentStatusRefunded, "updatedAt": _time, "errHash": receipt.BlockHash.Hex()}}})
	if err != nil {
		return nil, err
	}

	if result.ModifiedCount == 0 {
		return nil, fmt.Errorf("this is the error")
	}

	_result := true

	return &_result, nil
}

func GetEnterprisePDF(ctx context.Context, enterpriseId string) (string, error) {
	enterprise, err := GetEnterpriseById(ctx, enterpriseId)
	if err != nil {
		return "", err
	}
	pdf, err := pdf.CreatePdfFile(enterprise.WalletPublicKey)
	if err != nil {
		return "", err
	}

	output := pdf.String()
	return output, nil
}
