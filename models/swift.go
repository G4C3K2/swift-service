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
}
