package models

type SwiftEntry struct {
	SwiftCode     string  `bson:"swift_code"`
	CodeType      string  `bson:"code_type"`
	Name          string  `bson:"name"`
	Address       *string `bson:"address"`
	TownName      string  `bson:"town_name"`
	CountryCode   string  `bson:"country_code"`
	CountryName   string  `bson:"country_name"`
	TimeZone      string  `bson:"time_zone"`
	IsHeadquarter bool    `bson:"is_headquarter"`
	HqCode        *string `bson:"hqCode,omiteempty"`
}

type SwiftCodeResponse struct {
	Address       *string       `json:"address"`
	BankName      string        `json:"bankName"`
	CountryISO2   string        `json:"countryISO2"`
	CountryName   string        `json:"countryName"`
	IsHeadquarter bool          `json:"isHeadquarter"`
	SwiftCode     string        `json:"swiftCode"`
	Branches      []SwiftBranch `json:"branches,omitempty"`
}

type SwiftBranch struct {
	Address       *string `json:"address"`
	BankName      string  `json:"bankName"`
	CountryISO2   string  `json:"countryISO2"`
	IsHeadquarter bool    `json:"isHeadquarter"`
	SwiftCode     string  `json:"swiftCode"`
}

type CountryISO2CodeResponse struct {
	CountryISO2 string        `json:"countryISO2"`
	CountryName string        `json:"countryName"`
	SwiftCodes  []SwiftBranch `json:"swifts,omitempty"`
}

type CountryShort struct {
	CountryCode string `bson:"country_code"`
	CountryName string `bson:"country_name"`
}

type AddSwiftCodeRequest struct {
	Address       string `json:"address"`
	BankName      string `json:"bankName"`
	CountryISO2   string `json:"countryISO2"`
	CountryName   string `json:"countryName"`
	IsHeadquarter bool   `json:"isHeadquarter"`
	SwiftCode     string `json:"swiftCode"`
}
