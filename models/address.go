package models

//Address a struct to rep an address
type Address struct {
	SuiteNumberID     *string `json:"suiteNumber"`
	StreetAddress     *string `json:"streetAddress"`
	City              *string `json:"city"`
	ProvinceStateCode *string `json:"provinceStateCode"`
	CountryCode       *string `json:"countryCode"`
	PostalCode        *string `json:"postalCode"`
}
