package model

type Paiement struct {
	ID                 string        `json:"_id" bson:"_id"`
	Status             PaymentStatus `json:"status"`
	Type               PaymentType   `json:"type"`
	CreatorID          *string       `json:"creatorId"`
	Token              string        `json:"token"`
	AmountInt64        string        `json:"amount" bson:"amount"`
	FeeInt64           string        `json:"fee" bson:"fee"`
	FeeEnterpriseInt64 string        `json:"feeEnterprise" bson:"feeEnterprise"`
	CancellorID        *string       `json:"cancellorId"`
	AgencyID           *string       `json:"agencyId"`
	Destination        string        `json:"destination"`
	ValidatorID        *string       `json:"validatorId"`
	TransactionID      *string       `json:"transactionId"`
	CreatedAt          *string       `json:"createdAt"`
	Description        *string       `json:"description"`
	UpdatedAt          *string       `json:"updatedAt"`
	ShortID            string        `json:"shortId"`
	DestinationUserID  *string       `json:"destinationUserId"`
	EnterpriseID       *string       `json:"enterpriseId"`
}

func (Paiement) IsQRCodeOwner() {}
