package models

type SwiftEntry struct {
	SwiftCode       string `json:"swiftCode" bson:"swiftCode"`
	BankName        string `json:"bankName" bson:"bankName"`
	Address         string `json:"address" bson:"address"`
	CountryISO2     string `json:"countryISO2" bson:"countryISO2"`
	CountryName     string `json:"countryName" bson:"countryName"`
	IsHeadquarter   bool   `json:"isHeadquarter" bson:"isHeadquarter"`
	HeadquarterCode string `json:"headquarterCode,omitempty" bson:"headquarterCode,omitempty"` // dla branchy: 8 znaków HQ
}
