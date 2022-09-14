package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type DBReportTransaction struct {
	Transaction primitive.ObjectID  `validate:"nonzero" bson:"transaction"`
	Message     string              `validate:"nonzero" bson:"message"`
	Status      string              `validate:"nonzero" bson:"status"`
	ReportedBy  primitive.ObjectID  `validate:"nonzero" bson:"reportedBy"`
	RevolvedBy  *primitive.ObjectID `validate:"nonzero" bson:"revolvedBy"`
	CreatedAt   string              `validate:"nonzero" bson:"createdAt"`
	UpdatedAt   string              `validate:"nonzero" bson:"updatedAt"`
}

// transaction: {type: Schema.Types.ObjectId, ref: 'Transactions', required: true},
// message: {type: String, required: true},
// status: {type: String, enum: ["RESOLVED", "OPEN", "NOT_RESOLVED", "CLOSED"], default: "OPEN"},
// reportedBy: {type: Schema.Types.ObjectId, ref: 'Users', required: true},
// revolvedBy: {type: Schema.Types.ObjectId, ref: 'Users', required: false},
// createdAt: { type: String, default: new Date().toISOString(), required: true},
// updatedAt: { type: String, default: new Date().toISOString(), required: true},
