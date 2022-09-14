package database

import (
	"server/graph/model"
)

type NewKeyPair struct {
	PublicKey string `validate:"nonzero" bson:"publicKey"`
	SecretKey string `validate:"nonzero" bson:"secretKey"`
}

type DBLocation struct {
	Longitude float32 `validate:"nonezero" bson:"longitude"`
	Latitude  float32 `validate:"nonezero" bson:"latitude"`
}

type DBAddress struct {
	Title     string     `validate:"nonzero" bson:"title"`
	Location  DBLocation `validate:"nonzero" bson:"location"`
	City      string     `validate:"nonzero" bson:"city"`
	Zip       string     `validate:"nonzero" bson:"zip"`
	Country   string     `validate:"nonzero" bson:"country"`
	Is_chosed bool       `bson:"is_chosed"`
}

type BirthDate struct {
	Day   int    `validate:"nonzero" bson:"day"`
	Month int    `validate:"nonzero" bson:"month"`
	Year  int    `validate:"nonzero" bson:"year"`
	Iso   string `bson:"iso"`
}

type NewDBUserMongo struct {
	FirstName         string        `validate:"nonzero" bson:"first_name"`
	Last_name         string        `validate:"nonzero" bson:"last_name"`
	Email             string        `validate:"nonzero" bson:"email"`
	Phone_number      string        `validate:"nonzero" bson:"phonenumber"`
	Adresses          []DBAddress   `bson:"adresses"`
	FirebaseUID       string        `validate:"nonzero" bson:"firebaseUID"`
	CreatedAt         string        `validate:"nonzero" bson:"createdAt"`
	UpdatedAt         string        `validate:"nonzero" bson:"updatedAt"`
	FcmToken          string        `bson:"fcmToken"`
	PhotoUrl          string        `validate:"nonzero" bson:"photoUrl"`
	Keypair           NewKeyPair    `bson:"keypair"`
	Is_terms_accepted bool          `bson:"is_terms_accepted"`
	Country           string        `bson:"country"`
	DefaultCurrency   string        `validate:"nonzero" bson:"defaultCurrency"`
	User_baned        bool          `bson:"user_baned"`
	Birth_date        BirthDate     `bson:"birth_date"`
	Contacts          []string      `bson:"contacts"`
	ShortId           string        `bson:"shortId"`
	Permissions       []*model.Role `bson:"permissions"`
	Fee               int           `bson:"fee"`
	Is_online         bool          `bson:"is_online"`
	PinCode           string        `validate:"nonzero" bson="pinCode"`
	InvitedBy         string        `bson:"invitedBy "`
	AccountFrozen     bool          `bson:"accountFrozen"`
	Deleted           bool          `bson:"deleted"`
	Ip                *ConnectionDB `bson:"iplocation"`
}
