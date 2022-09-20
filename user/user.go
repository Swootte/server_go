package user

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"os"
	database "server/database"
	"server/ethereum"
	"server/finance"
	fb "server/firebase"
	"server/graph/model"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/validator.v2"

	"server/utils"
)

var firebase = fb.Connect()
var auth = firebase.Connect()

func init() {
	utils.LoadEnv()
}

func CreateUser(ctx context.Context, user *model.UserInput, uid string, ip string) (*model.UserCreated, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PinCode), 10)
	if err != nil {
		fmt.Println(err)
	}

	photourl := "https://firebasestorage.googleapis.com/v0/b/swoosh-97759.appspot.com/o/tinda_swootte_logo.png?alt=media&token=ba9ff652-78e6-4984-b0b7-2dd5d4503f3b"
	user.PhotoURL = &photourl

	// fbToken := fb.Connect().VerifyIdToken(ctx, user.Token)
	signedUser, err := auth.CreateUser(ctx, user.Email, user.Phonenumber, user.Password, *user.DisplayName, *user.PhotoURL, uid)
	if err != nil {
		return nil, err
	}

	customToken, err := firebase.CreateCustomToken(ctx, signedUser.UID)
	if err != nil {
		return nil, err
	}

	privateKey := ethereum.CreateWallet()
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyDB := hexutil.Encode(privateKeyBytes)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	secret, err := utils.Encrypt(privateKeyDB, os.Getenv("SERVER_SECRET"))
	if err != nil {
		return nil, err
	}

	_time := time.Now().UTC().Format(time.RFC3339)

	var _permissions []*model.Role
	_permissions = append(_permissions, &model.AllRole[2])

	var addresses []database.DBAddress
	addresses = append(addresses, database.DBAddress{
		Title: *user.Adress.Title,
		City:  *user.Adress.City,
		Location: database.DBLocation{
			Longitude: 0,
			Latitude:  0,
		},
		Is_chosed: true,
	})

	_userDb := database.NewDBUserMongo{
		FirstName:            user.FirstName,
		Last_name:            user.LastName,
		Email:                user.Email,
		Phone_number:         user.Phonenumber,
		Adresses:             addresses,
		FirebaseUID:          signedUser.UID,
		CreatedAt:            _time,
		UpdatedAt:            _time,
		FcmToken:             *user.FcmToken,
		PhotoUrl:             *user.PhotoURL,
		Keypair:              database.NewKeyPair{PublicKey: address, SecretKey: secret},
		Is_terms_accepted:    true,
		Country:              user.Country,
		DefaultCurrency:      os.Getenv("DEFAULT_CURRENCY"),
		User_baned:           false,
		Birth_date:           database.BirthDate(*user.BirthDate),
		Contacts:             make([]string, 0),
		Permissions:          _permissions,
		Fee:                  5,
		Is_online:            false,
		PinCode:              string(hashedPassword),
		InvitedBy:            *user.InvitedBy,
		AccountFrozen:        false,
		Deleted:              false,
		IndentityStatus:      model.IdentityStatusNotUploaded,
		ResidenceProofStatus: model.ResidenceProofStatusNotUploaded,
	}

	if errs := validator.Validate(&_userDb); errs != nil {
		//fmt.Println("didnt validate this shit", errs)
	}

	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := _collections.InsertOne(ctx, _userDb)
	if err != nil {
		return nil, err
	}

	insertedID := res.InsertedID.(primitive.ObjectID).Hex()
	returnedUser, _ := GetUserById(insertedID)

	result := model.UserCreated{
		User:        returnedUser,
		CustomToken: customToken,
	}

	return &result, nil

}

func GetUserByFirebaseId(uid string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("users")
	var singleUser *model.User
	err := _collections.FindOne(ctx, bson.D{{Key: "firebaseUID", Value: uid}}).Decode(&singleUser)
	return singleUser, err
}

func MigrateAllUsersWallet() {
	log.Printf("migration started")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("users")
	cursor, err := _collections.Find(context.Background(), bson.M{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("migration started")
	for cursor.Next(ctx) {
		var singleUser *model.User
		if err = cursor.Decode(&singleUser); err != nil {
			log.Printf("err:%s", err.Error())
			fmt.Println(err)
		}
		if *singleUser.Phonenumber != "+33782798614" {
			privateKey, address, _ := finance.CreateAccount()
			secret, err := utils.Encrypt(privateKey, os.Getenv("SERVER_SECRET"))
			if err != nil {
				log.Printf("err:%s", err.Error())
				fmt.Println(err)
			}
			objectId, _ := primitive.ObjectIDFromHex(singleUser.ID)
			_collections.UpdateOne(context.Background(), bson.D{{Key: "_id", Value: objectId}}, bson.D{{Key: "$set", Value: bson.D{{Key: "keypair.publicKey", Value: address}, {Key: "keypair.secretKey", Value: secret}, {Key: "defaultCurrency", Value: os.Getenv("DEFAULT_CURRENCY")}, {Key: "deleted", Value: false}}}})
		} else {
			secret, _ := utils.Encrypt(os.Getenv("CHAIN_PRIVATE_KEY"), os.Getenv("SERVER_SECRET"))
			objectId, _ := primitive.ObjectIDFromHex(singleUser.ID)
			_collections.UpdateOne(context.Background(), bson.D{{Key: "_id", Value: objectId}}, bson.D{{Key: "$set", Value: bson.D{{Key: "keypair.publicKey", Value: os.Getenv("CHAIN_ADDRESS")}, {Key: "keypair.secretKey", Value: secret}, {Key: "defaultCurrency", Value: os.Getenv("DEFAULT_CURRENCY")}, {Key: "deleted", Value: false}}}})
		}
	}

	fmt.Println("migration done")
}

func UserExist(uid string) (bool, error) {
	res, err := GetUserByFirebaseId(uid)
	if res.Deleted != nil && !*res.Deleted {
		return true, err
	}
	return false, err
}

func UsersExist(userdb model.User, args interface{}) *model.User {
	if userdb.Deleted != nil && !*userdb.Deleted {
		return &userdb
	}
	return nil
}

func changeFirstName(userdb model.User, _id string, firstname string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_time := time.Now().UTC().Format(time.RFC3339)
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("users")
	objectId, _ := primitive.ObjectIDFromHex(_id)
	_, err := _collections.UpdateOne(ctx, bson.M{"_id": objectId}, bson.D{{Key: "$set", Value: bson.M{"updatedAt": _time, "first_name": firstname}}})
	if err != nil {
		panic(err)
	}
	userdb.FirstName = &firstname
	return &userdb, err
}

func changeLastName(userdb model.User, _id string, lastname string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_time := time.Now().UTC().Format(time.RFC3339)
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("users")
	objectId, _ := primitive.ObjectIDFromHex(_id)
	_, err := _collections.UpdateOne(ctx, bson.M{"_id": objectId}, bson.D{{Key: "$set", Value: bson.M{"updatedAt": _time, "last_name": lastname}}})
	if err != nil {
		panic(err)
	}

	userdb.LastName = &lastname
	return &userdb, err
}

func changeEmail(userdb model.User, _id string, email string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_time := time.Now().UTC().Format(time.RFC3339)
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("users")
	objectId, _ := primitive.ObjectIDFromHex(_id)
	_, err := _collections.UpdateOne(ctx, bson.M{"_id": objectId}, bson.D{{Key: "$set", Value: bson.M{"updatedAt": _time, "email": email}}})
	if err != nil {
		panic(err)
	}

	return true
}

func changeAddress(userdb model.User, address *model.Adress) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_time := time.Now().UTC().Format(time.RFC3339)
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("users")
	objectId, _ := primitive.ObjectIDFromHex(userdb.ID)
	_, err := _collections.UpdateOne(ctx, bson.M{"_id": objectId}, bson.D{{Key: "$set", Value: bson.M{"updatedAt": _time, "addresses": bson.M{"$push": address}}}})
	if err != nil {
		panic(err)
	}

	return true
}

func beforeAllTest() {
	panic(fmt.Errorf("not implemented"))
}

func LoadAllActivities(userdb *model.User) ([]*model.Paiement, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	opts := options.Find().SetSort(bson.D{{Key: "createdAt", Value: -1}})
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("transactions")
	_pipeline := bson.D{
		{
			Key: "$or", Value: bson.A{
				bson.M{"destination": userdb.Keypair.PublicKey},
				bson.M{"source": userdb.Keypair.PublicKey},
			},
		},
	}

	var payments []*model.Paiement
	res, err := _collections.Find(ctx, _pipeline, opts)

	if err != nil {
		return nil, err
	}

	defer res.Close(ctx)
	for res.Next(ctx) {
		var singlePayment *model.Paiement

		if err = res.Decode(&singlePayment); err != nil {
			log.Fatal(err)
		}

		payments = append(payments, singlePayment)

	}

	return payments, err

}

func GetAllParticipatingTransactions(userdb model.User) ([]*model.Paiement, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("transactions")

	objectId, _ := primitive.ObjectIDFromHex(userdb.ID)
	_pipeline := bson.M{
		"$or": []interface{}{
			bson.M{"destination": userdb.Keypair.PublicKey, "token": *userdb.DefaultCurrency},
			bson.M{"creator": objectId},
		},
		"$sort": bson.M{
			"creadtedAt": -1,
		},
	}

	var payments []*model.Paiement
	res, err := _collections.Aggregate(ctx, _pipeline)
	if err != nil {
		return nil, err
	}
	defer res.Close(ctx)
	for res.Next(ctx) {
		var singlePayment *model.Paiement
		if err = res.Decode(&singlePayment); err != nil {
			log.Fatal(err)
		}

		payments = append(payments, singlePayment)

	}

	return payments, err

}

func saveImage(userdb model.User, _type string, stream string) {
	panic(fmt.Errorf("not implemented"))
}

func gethAllAuthedUsers() {
	panic(fmt.Errorf("not implemented"))
}

func DeleteUser(userdb model.User, ip string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("users")
	objectId, _ := primitive.ObjectIDFromHex(userdb.ID)

	_time := time.Now().UTC().Format(time.RFC3339)
	_match := bson.D{{Key: "_id", Value: objectId}}
	_set := bson.D{{Key: "$set", Value: bson.D{{Key: "deleted", Value: true}, {Key: "updatedAt", Value: _time}}}}
	_, err := _collections.UpdateOne(ctx, _match, _set)
	if err != nil {
		return false, err
	}

	return true, nil
}

func saveDocument() {
	panic(fmt.Errorf("not implemented"))
}

func UpdateFcmToken(userdb *model.User, fcmToken *string, ip string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("users")
	//objectId, _ := primitive.ObjectIDFromHex(userdb.ID)
	_time := time.Now().UTC().Format(time.RFC3339)

	_set := bson.D{{Key: "$set", Value: bson.D{{Key: "fcmToken", Value: fcmToken}, {Key: "updatedAt", Value: _time}}}}
	_, err := _collections.UpdateOne(ctx, bson.D{{Key: "firebaseUID", Value: userdb.FirebaseUID}}, _set)
	if err != nil {
		return false, err
	}

	return true, nil
}

func UpdateProfilPicture(userdb model.User, photoUrl string, ip string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("users")
	objectId, _ := primitive.ObjectIDFromHex(userdb.ID)

	_time := time.Now().UTC().Format(time.RFC3339)
	_, err := _collections.UpdateOne(ctx, bson.M{"_id": objectId}, bson.M{"$set": bson.M{"photoUrl": photoUrl, "updatedAt": _time}})
	if err != nil {
		return false, err
	}

	return true, nil
}

func ChangePinCode(userdb model.User, pinCode string, ip string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("users")
	objectId, _ := primitive.ObjectIDFromHex(userdb.ID)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pinCode), 10)
	if err != nil {
		return false, err
	}

	_time := time.Now().UTC().Format(time.RFC3339)
	_, er := _collections.UpdateOne(ctx, bson.M{"_id": objectId}, bson.M{"$set": bson.M{"pinCode": hashedPassword, "updatedAt": _time}})
	if er != nil {
		return false, er
	}

	return true, nil
}

func ToggleOnlineStatus(userdb model.User, toogle bool, ip string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("users")
	objectId, _ := primitive.ObjectIDFromHex(userdb.ID)

	_time := time.Now().UTC().Format(time.RFC3339)
	_, err := _collections.UpdateOne(ctx, bson.M{"_id": objectId}, bson.D{{Key: "$set", Value: bson.M{"is_online": toogle, "updatedAt": _time}}})
	if err != nil {
		panic(err)
	}

	_result := true
	return _result, nil
}

func GetUserById(user string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("users")
	objectId, _ := primitive.ObjectIDFromHex(user)

	var signleUser *model.User
	err := _collections.FindOne(ctx, bson.M{"_id": objectId}).Decode(&signleUser)
	if err != nil {
		fmt.Println(err, "error hype")
	}

	return signleUser, err
}

func UpdateDefaultCurrencyOnUser(user *model.User, defaultCurrency string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("users")
	objectId, _ := primitive.ObjectIDFromHex(user.ID)

	_time := time.Now().UTC().Format(time.RFC3339)
	_, err := _collections.UpdateOne(ctx, bson.M{"_id": objectId}, bson.D{{Key: "$set", Value: bson.M{"defaultCurrency": defaultCurrency, "updatedAt": _time}}})
	user.DefaultCurrency = &defaultCurrency
	return user, err
}

func usersDataloader() {
	panic(fmt.Errorf("not implemented"))
}

func findUsersbyIds() {
	panic(fmt.Errorf("not implemented"))
}

func getIsTokenOwner() {

}

func SearchUser(user *model.User, searchText string) ([]*model.UserSmall, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("users")
	objectId, _ := primitive.ObjectIDFromHex(user.ID)

	user_contacts := bson.A{}
	if len(user.Contacts) > 0 {
		for i := 0; i < len(user.Contacts); i++ {
			_conctact_objectId, _ := primitive.ObjectIDFromHex(*user.Contacts[i])
			user_contacts = append(user_contacts, _conctact_objectId)
		}
	}

	if searchText == "" {
		_match := bson.D{
			{
				Key: "$match", Value: bson.M{
					"_id": bson.M{
						"$in": user_contacts,
						"$ne": objectId,
					},
				},
			},
		}

		_limit := bson.D{
			{
				Key: "$limit", Value: 15,
			},
		}
		res, err := _collections.Aggregate(ctx, mongo.Pipeline{_match, _limit})
		if err != nil {
			return nil, err
		}
		var users []*model.UserSmall
		defer res.Close(ctx)
		for res.Next(ctx) {
			var singleUser *model.UserSmall
			if err = res.Decode(&singleUser); err != nil {
				log.Fatal(err)
			}

			users = append(users, singleUser)

		}

		return users, err
	}

	_search := bson.D{
		{
			Key: "$search", Value: bson.D{
				{
					Key: "text", Value: bson.D{
						{
							Key: "query", Value: bson.A{searchText},
						},
						{
							Key: "path", Value: bson.D{
								{
									Key: "wildcard", Value: "*",
								},
							},
						},
					},
				},
			},
		},
	}

	_match := bson.D{
		{
			Key: "$match", Value: bson.D{
				{
					Key: "_id", Value: bson.D{
						{
							Key: "$in", Value: user_contacts,
						},
						{
							Key: "$ne", Value: objectId,
						},
					},
				},
			},
		},
	}
	res, err := _collections.Aggregate(ctx, mongo.Pipeline{_search, _match})
	if err != nil {
		return nil, err
	}
	var users []*model.UserSmall
	defer res.Close(ctx)
	for res.Next(ctx) {
		var singleUser *model.UserSmall
		if err = res.Decode(&singleUser); err != nil {
			log.Fatal(err)
		}

		users = append(users, singleUser)

	}

	return users, err

}

func SearchUserUnAdded(user *model.User, searchText string) ([]*model.UserSmall, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("users")
	objectId, _ := primitive.ObjectIDFromHex(user.ID)

	user_contacts := bson.A{}
	if len(user.Contacts) > 0 {
		for i := 0; i < len(user.Contacts); i++ {
			_conctact_objectId, _ := primitive.ObjectIDFromHex(*user.Contacts[i])
			user_contacts = append(user_contacts, _conctact_objectId)
		}
	}

	if searchText == "" {
		_match := bson.D{
			{
				Key: "$match", Value: bson.M{
					"_id": bson.M{
						"$nin": user_contacts,
						"$ne":  objectId,
					},
				},
			},
		}

		_limit := bson.D{
			{
				Key: "$limit", Value: 15,
			},
		}
		res, err := _collections.Aggregate(ctx, mongo.Pipeline{_match, _limit})

		var users []*model.UserSmall
		defer res.Close(ctx)
		for res.Next(ctx) {
			var singleUser *model.UserSmall
			if err = res.Decode(&singleUser); err != nil {
				log.Fatal(err)
			}

			users = append(users, singleUser)

		}

		return users, err
	}

	_search := bson.D{
		{
			Key: "$search", Value: bson.D{
				{
					Key: "text", Value: bson.D{
						{
							Key: "query", Value: bson.A{searchText},
						},
						{
							Key: "path", Value: bson.D{
								{
									Key: "wildcard", Value: "*",
								},
							},
						},
					},
				},
			},
		},
	}

	_match := bson.D{
		{
			Key: "$match", Value: bson.M{
				"_id": bson.M{
					"$nin": user_contacts,
					"$ne":  objectId,
				},
			},
		},
	}
	res, err := _collections.Aggregate(ctx, mongo.Pipeline{_search, _match})
	var users []*model.UserSmall
	defer res.Close(ctx)
	for res.Next(ctx) {
		var singleUser *model.UserSmall
		if err = res.Decode(&singleUser); err != nil {
			log.Fatal(err)
		}

		users = append(users, singleUser)

	}

	return users, err

}

func AddContact(user model.User, contact string, ip string) (*bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("users")
	objectId, _ := primitive.ObjectIDFromHex(user.ID)

	_constraint := bson.M{
		"_id": objectId,
	}

	_time := time.Now().UTC().Format(time.RFC3339)
	_updatePipeline := bson.D{{Key: "$set", Value: bson.M{
		"updatedAt": _time,
	}}, {
		Key: "$push", Value: bson.M{
			"contacts": contact,
		},
	}}
	_, err := _collections.UpdateOne(ctx, _constraint, _updatePipeline)
	if err != nil {
		return nil, err
	}

	_result := true

	return &_result, nil

}

func RemoveContact(user model.User, contact string, ip string) (*bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("users")
	objectId, _ := primitive.ObjectIDFromHex(user.ID)

	_constraint := bson.M{
		"_id": objectId,
	}

	_time := time.Now().UTC().Format(time.RFC3339)
	_updatePipeline := bson.D{{Key: "$set", Value: bson.M{
		"updatedAt": _time,
	}}, {
		Key: "$pull", Value: bson.M{
			"contacts": contact,
		},
	}}
	_, err := _collections.UpdateOne(ctx, _constraint, _updatePipeline)
	if err != nil {
		return nil, err
	}

	_result := true
	return &_result, nil

}

func LoadBalance(publicKey string) (*model.Wallet, error) {
	balance, err := finance.GetBalanceOnContractInstance(publicKey)
	if err != nil {
		return nil, err
	}
	_balance, _ := balance.Float64()
	output := model.Wallet{
		Address:  &publicKey,
		Amount:   &_balance,
		IsFrozen: new(bool),
	}
	return &output, nil

}

func LoadQrCode(user model.User) string {
	return user.Keypair.PublicKey
}
