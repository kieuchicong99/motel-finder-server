package repository

import (
	"server/infrastructure"
	"server/model"
)

type reportRepository struct {
	baseRepository
}

func NewReportRepository() model.ReportRepositoryInterface {
	var mongoGoDriver = &infrastructure.MongoGoDriverConnector{}
	client, err := mongoGoDriver.Connect()
	if err != nil {
		return nil
	}

	return &reviewRepository{
		baseRepository{
			primaryCollection: client.Database("motelfinder").Collection("reports"),
		},
	}
}
