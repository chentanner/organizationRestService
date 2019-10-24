package assemblers

import (
	api "organizationRestService/api/gen"
	"organizationRestService/models"
)

// AssembleAddressApiSnapshot will convert an entity address into an api snapshot address
func AssembleAddressApiSnapshot(entity models.Address) *api.Address {
	snapshot := new(api.Address)

	if entity.SuiteNumberID != nil {
		snapshot.Suite = entity.SuiteNumberID
	}
	if entity.StreetAddress != nil {
		snapshot.StreetAddress = entity.StreetAddress
	}
	if entity.City != nil {
		snapshot.City = entity.City
	}
	if entity.ProvinceStateCode != nil {
		snapshot.ProvState = entity.ProvinceStateCode
	}
	if entity.CountryCode != nil {
		snapshot.CountryCode = entity.CountryCode
	}
	if entity.PostalCode != nil {
		snapshot.PostalCode = entity.PostalCode
	}

	return snapshot
}

// AssembleAddressEntity will convert an api snapshot address into an entity address
func AssembleAddressEntity(snapshot api.Address) *models.Address {
	entity := new(models.Address)

	if snapshot.Suite != nil {
		entity.SuiteNumberID = snapshot.Suite
	}
	if snapshot.StreetAddress != nil {
		entity.StreetAddress = snapshot.StreetAddress
	}
	if snapshot.City != nil {
		entity.City = snapshot.City
	}
	if snapshot.ProvState != nil {
		entity.ProvinceStateCode = snapshot.ProvState
	}
	if snapshot.CountryCode != nil {
		entity.CountryCode = snapshot.CountryCode
	}
	if snapshot.PostalCode != nil {
		entity.PostalCode = snapshot.PostalCode
	}

	return entity
}
