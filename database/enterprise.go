package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DBPerson struct {
	First_name string `validate:"nonzero" bson:"first_name"`
	Last_name  string `validate:"nonzero" bson:"last_name"`
	Email      string `validate:"nonzero" bson:"email"`
	Address    string `validate:"nonzero" bson:"address"`
	City       string `validate:"nonzero" bson:"city"`
	Zip        string `validate:"nonzero" bson:"zip"`
	State      string `validate:"nonzero" bson:"state"`
}

type DBPhone struct {
	DialCode string `validate:"nonzero" bson:"dial_code"`
	Phone    string `validate:"nonzero" bson:"phone"`
}

type DBEnterprise struct {
	Type                 string             `validate:"nonzero" bson:"type"`
	Name                 string             `validate:"nonzero" bson:"name"`
	Rccm                 string             `validate:"nonzero" bson:"rccm"`
	Website              string             `validate:"nonzero" bson:"website"`
	LogoUrl              string             `validate:"nonzero" bson:"logoUrl"`
	Creator              primitive.ObjectID `validate:"nonzero" bson:"creator"`
	CreatedAt            string             `validate:"nonzero" bson:"createdAt"`
	UpdatedAt            string             `validate:"nonzero" bson:"updatedAt"`
	PublishableKey       string             `validate:"nonzero" bson:"publishableKey"`
	Private_key          string             `validate:"nonzero" bson:"private_key"`
	WalletPublicKey      string             `validate:"nonzero" bson:"walletPublicKey"`
	WalletSecretKey      string             `validate:"nonzero" bson:"walletSecretKey"`
	Country              string             `validate:"nonzero" bson:"country"`
	Address              DBAddress          `bson:"adress"`
	Removed              bool               `validate:"nonzero" bson:"removed"`
	DeletedAt            string             `validate:"nonzero" bson:"DeletedAt"`
	Default_enterprise   bool               `validate:"nonzero" bson:"default_enterprise"`
	Sector               string             `validate:"nonzero" bson:"sector"`
	Description          string             `validate:"nonzero" bson:"description"`
	SellingPhysicalGoods bool               `validate:"nonzero" bson:"sellingPhysicalGoods"`
	TransactionLibele    string             `validate:"nonzero" bson:"transactionLibele"`
	AbregedLibele        string             `validate:"nonzero" bson:"abregedLibele"`
	Phone                DBPhone            `validate:"nonzero" bson:"phone"`
	Email                string             `validate:"nonzero" bson:"email"`
	Bordereau            string             `validate:"nonzero" bson:"bordereau"`
	Person               DBPerson           `validate:"nonzero" bson:"person"`
}
