package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateReportPayload struct {
	BaseReport
}
type Report struct {
	ID        primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	BaseReport `bson:",inline"`
}
type BaseReport struct {
	MotelCode   primitive.ObjectID `json:"MotelCode" bson:"MotelCode"`
	Description string `json:"Description" bson:"Description"`
	UserCode    primitive.ObjectID `json:"UserCode" bson:"UserCode"`
}
type ReportRepositoryInterface interface {
	BaseRepository
}