package model

type Notification struct {
	ID        *string `json:"_id" bson:"_id"`
	Text      *string `json:"text" bson:"text"`
	Type      *string `json:"type" bson:"type"`
	ImgURL    *string `json:"imgUrl" bson:"imgUrl"`
	IsRead    *bool   `json:"isRead" bson:"isRead"`
	FromID    *string `json:"fromId" bson:"fromId"`
	CreatedAt *string `json:"createdAt" bson:"createdAt"`
}
