package model

type Phone struct {
	DialCode string `json:"dialcode" bson:"dial_code"`
	Phone    string `json:"phone" bson:"phone"`
}

type Enterprise struct {
	ID                   string  `json:"_id" bson:"_id"`
	Type                 *string `json:"type" bson:"type"`
	Name                 *string `json:"name" bson:"name"`
	Website              *string `json:"website" bson:"website"`
	LogoURL              *string `json:"logoUrl" bson:"logoUrl"`
	Creator              string  `json:"creator" bson:"creator"`
	CreatedAt            string  `json:"createdAt" bson:"createdAt"`
	UpdatedAt            *string `json:"updatedAt" bson:"updatedAt"`
	Person               *Person `json:"person" bson:"person"`
	PublishableKeyString string  `json:"publishableKey" bson:"publishableKey"`
	PrivateKeystring     string  `json:"private_key" bson:"private_key"`
	WalletPublicKey      string  `json:"walletPublicKey" bson:"walletPublicKey"`
	WalletSecretKey      *string `json:"walletSecretKey" bson:"walletSecretKey"`
	Country              *string `json:"country" bson:"country"`
	Address              *Adress `json:"address" bson:"address"`
	DefaultEnterprise    bool    `json:"default_enterprise" bson:"default_enterprise"`
	Description          *string `json:"description" bson:"description"`
	SellingPhysicalGoods *bool   `json:"sellingPhysicalGoods" bson:"sellingPhysicalGoods"`
	SelfShippingProduct  *bool   `json:"selfShippingProduct" bson:"selfShippingProduct"`
	ShippingDelay        *string `json:"shippingDelay" bson:"shippingDelay"`
	TransactionLibele    *string `json:"transactionLibele" bson:"transactionLibele"`
	AbregedLibele        *string `json:"abregedLibele" bson:"abregedLibele"`
	Phone                *Phone  `json:"phone" bson:"phone"`
	Email                *string `json:"email" bson:"email"`
	Rccm                 *string `json:"rccm" bson:"rccm"`
	Sector               *string `json:"sector" bson:"sector"`
}
