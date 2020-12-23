package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Motel struct {
	MotelCode        primitive.ObjectID `json:"MotelCode" bson:"MotelCode,omitempty"`
	Address          string             `json:"Address" bson:"Address,omitempty"`
	Images           []string           `json:"Images" bson:"Images,omitempty"`
	Image            string             `json:"Image" bson:"Image,omitempty"`
	Tags             []string           `json:"Tags" bson:"Tags,omitempty"`
	Description      string             `json:"Description" bson:"Description"`
	Title            string             `json:"Title" bson:"Title,omitempty"`
	Cost             float32            `json:"Cost" bson:"Cost,omitempty"`
	Latitude         string             `json:"Latitude" bson:"Latitude,omitempty"`
	Longitude        string             `json:"Longitude" bson:"Longitude,omitempty"`
	CreatedAt        time.Time          `json:"CreatedAt" bson:"CreatedAt,omitempty"`
	FinishedAt       time.Time          `json:"FinishedAt" bson:"FinishedAt,omitempty"`
	Status           bool               `json:"Status" bson:"Status"`       // trạng thái để admin duyệt bài
	Available        bool               `json:"Available" bson:"Available"` // trạng thái phòng đã có người thuê hay chưa
	Bathroom         Bathroom           `json:"Bathroom" bson:"Bathroom"`
	Kitchen          string             `json:"kitchen" bson:"Kitchen"`
	HasAirCondition  bool               `json:"HasAirCondition" bson:"HasAirCondition"`
	HasBalcony       bool               `json:"HasBalcony" bson:"HasBalcony"`
	ElectricityPrice uint               `json:"ElectricityPrice" bson:""`
	WaterPrice       uint               `json:"WaterPrice" bson:"WaterPrice"`
	UserCode         primitive.ObjectID `json:"UserCode" bson:"UserCode,omitempty"`
	Acreage          float64   `json:"Acreage" bson:"Acreage"`
}
type Bathroom struct {
	Type           string `json:"Type"`
	HasWaterHeater bool   `json:"HasWaterHeater"`
}
type MotelRepositoryInterface interface {
	MakeIndexes()
	Insert(motel *Motel) (insertResponse *InsertResponse, httpCode int)
	UpdateInfo(code string, motel *Motel) (updateResponse *UpdateResponse, httpCode int)
	UpdateStatus(code string, motel *Motel) (updateResponse *UpdateResponse, httpCode int)
	UpdateAvailable(code string, motel *Motel) (updateResponse *UpdateResponse, httpCode int)
	GetAll(page, pageSize int, address string, fromCost, toCost uint, fromAcreage, toAcreage float64, hasKitchen string, hasAirCondition, hasWaterHeater, hasBalcony *bool) (GetManyResponse *GetManyResponse, totalResult, httpCode int)
	GetAllByOwnerCode(code string ) (GetManyResponse *GetManyResponse, totalResult int, httpCode int)
	GetByCode(code string) (getOne *GetOneResponse, httpCode int)
}

type CreateMotelPayload struct {
	Address         string   `json:"Address" bson:"Address,omitempty" example:"Ngõ 2, Phạm Văn Đồng, Cầu Giấy, Hà Nội"`
	Images          []string `json:"Images" bson:"Images,omitempty" example:"https://file4.batdongsan.com.vn/2020/08/28/20200828202318-987e_wm.jpg"`
	Image           string   `json:"Image" bson:"Image,omitempty" example:"https://file4.batdongsan.com.vn/2020/08/28/20200828202318-987e_wm.jpg"`
	Tags            []string `json:"Tags" bson:"Tags,omitempty" example:"hanoi"`
	Description     string   `json:"Description" bson:"Description,omitempty" example:"Nhà rộng rãi, thoáng mát"`
	Title           string   `json:"Title" bson:"Title,omitempty" example:"CCMN SỐ 1A NGÕ 49 TRIỀU KHÚC, FULL HẾT ĐỒ THANG MÁY BAN CÔNG, BẾP TÁCH RIÊNG, 26M2 GIÁ TỪ 3,3 TR/TH"`
	Cost            float32  `json:"Cost" bson:"Cost,omitempty" example:"3.3"`
	Latitude        string   `json:"Latitude" bson:"Latitude,omitempty" example:"21.0286755"`
	Longitude       string   `json:"Longitude" bson:"Longitude,omitempty" example:"105.7558943,13z"`
	Bathroom        Bathroom `json:"Bathroom" bson:"Bathroom"`
	Kitchen         string   `json:"kitchen" bson:"Kitchen"`
	HasAirCondition bool     `json:"HasAirCondition" bson:"HasAirCondition"`
	HasBalcony      bool     `json:"HasBalcony" bson:"HasBalcony"`
	Acreage          float64   `json:"Acreage" bson:"Acreage"`
	ElectricityPrice uint `json:"ElectricityPrice" bson:""`
	WaterPrice       uint `json:"WaterPrice" bson:"WaterPrice"`
}

type UpdateMotelPayload struct {
	Address          string    `json:"Address" bson:"Address,omitempty" example:"Ngõ 2, Phạm Văn Đồng, Cầu Giấy, Hà Nội"`
	Images           []string  `json:"Images" bson:"Images,omitempty" example:"https://file4.batdongsan.com.vn/2020/08/28/20200828202318-987e_wm.jpg"`
	Image            string    `json:"Image" bson:"Image,omitempty" example:"https://file4.batdongsan.com.vn/2020/08/28/20200828202318-987e_wm.jpg"`
	Tags             []string  `json:"Tags" bson:"Tags,omitempty" example:"hanoi"`
	Description      string    `json:"Description" bson:"Description,omitempty" example:"Nhà rộng rãi, thoáng mát"`
	Title            string    `json:"Title" bson:"Title,omitempty" example:"CCMN SỐ 1A NGÕ 49 TRIỀU KHÚC, FULL HẾT ĐỒ THANG MÁY BAN CÔNG, BẾP TÁCH RIÊNG, 26M2 GIÁ TỪ 3,3 TR/TH"`
	Cost             float32   `json:"Cost" bson:"Cost,omitempty" example:"3.3"`
	Latitude         string    `json:"Latitude" bson:"Latitude,omitempty" example:"21.0286755"`
	Longitude        string    `json:"Longitude" bson:"Longitude,omitempty" example:"105.7558943,13z"`
	FinishedAt       time.Time `json:"FinishedAt" bson:"FinishedAt,omitempty" example:"2021-01-15T22:20:28.227+00:00"`
	Status           bool      `json:"Status" bson:"Status,omitempty" example:"true"`
	Bathroom         Bathroom  `json:"Bathroom" bson:"Bathroom"`
	Kitchen          string    `json:"kitchen" bson:"Kitchen"`
	HasAirCondition  bool      `json:"HasAirCondition" bson:"HasAirCondition"`
	HasBalcony       bool      `json:"HasBalcony" bson:"HasBalcony"`
	Acreage          float64   `json:"Acreage" bson:"Acreage"`
	ElectricityPrice uint      `json:"ElectricityPrice" bson:""`
	WaterPrice       uint      `json:"WaterPrice" bson:"WaterPrice"`
}

type UpdateMotelAvailablePayload struct {
	Available bool `json:"Available" bson:"Available" example:"false"`
}

type UpdateMotelStatusPayload struct {
	Status bool `json:"Status" bson:"Status" example:"false"`
}

type MotelQuery struct {
	Address         string  `json:"address" schema:"address"`
	FromCost        uint    `json:"from_cost" schema:"from_cost"`
	ToCost          uint    `json:"to_cost"  schema:"to_cost"`
	FromAcreage     float64 `json:"from_acreage" schema:"from_acreage"`
	ToAcreage       float64 `json:"to_acreage" schema:"to_acreage"`
	HasKitchen      string  `json:"has_kitchen" schema:"has_kitchen"`
	HasAirCondition *bool   `json:"has_air_condition" schema:"has_air_condition"`
	HasWaterHeater  *bool   `json:"has_water_heater" schema:"has_water_heater"`
	HasBalcony      *bool   `json:"has_balcony" schema:"has_balcony"`
}