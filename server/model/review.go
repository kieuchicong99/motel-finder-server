package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Review struct {
	ID        primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	BaseReview `bson:",inline"`
}
type BaseReview struct {

	MotelCode primitive.ObjectID `json:"MotelCode" bson:"MotelCode"`
	Comment   string             `json:"Comment" bson:"Comment"`
	Rating    int                `json:"Rating" bson:"Rating"`
	Status    bool               `json:"Status" bson:"Status"`
	UserCode    primitive.ObjectID `json:"UserCode" bson:"UserCode"`
}
type CreateReviewPayload struct {
	BaseReview `bson:",inline"`
}
type UpdateStatusReview struct {
	Status bool `json:"Status"`
}
type ReviewRepositoryInterface interface {
	BaseRepository
}
