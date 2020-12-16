package model

type FavouriteMotel struct {
	UserCode string  `json:"UserCode" bson:"UserCode"`
	Motels   []Motel `json:"Motels" bson:"Motels"`
}
