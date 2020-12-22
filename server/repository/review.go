package repository

import (
	"server/infrastructure"
	"server/model"
)

type reviewRepository struct {
	baseRepository
}

func NewReviewRepository() model.ReviewRepositoryInterface {
	var mongoGoDriver = &infrastructure.MongoGoDriverConnector{}
	client, err := mongoGoDriver.Connect()
	if err != nil {
		return nil
	}

	return &reviewRepository{
		baseRepository{
			primaryCollection: client.Database("motelfinder").Collection("reviews"),
		},
	}
}
