package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateReportPayload struct {
	BaseReport
}
type Report struct {
	BaseReport `bson:",inline"`
}
type BaseReport struct {
	MotelCode   primitive.ObjectID `json:"motel_code"`
	Description string
	UserCode    primitive.ObjectID `json:"UserCode" bson:"UserCode"`
}
