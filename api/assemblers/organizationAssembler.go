package assemblers

import (
	api "organizationRestService/api/gen"
	"organizationRestService/models"
)

// AssembleOrganizationApiSnapshots will convert an []entity into a []snapshot
func AssembleOrganizationApiSnapshots(organizations []models.Organization) api.Organizations {
	var snapshots api.Organizations
	for _, organization := range organizations {
		snapshot := AssembleOrganizationApiSnapshot(&organization)
		snapshots = append(snapshots, *snapshot)
	}

	return snapshots
}

// AssembleOrganizationApiSnapshot will convert an entity into an api snapshot
func AssembleOrganizationApiSnapshot(organization *models.Organization) *api.Organization {
	snapshot := new(api.Organization)

	snapshot.Address = *AssembleAddressApiSnapshot(organization.Address)

	intID := int(organization.ID)
	snapshot.ID = &intID

	if organization.ShortName != nil {
		snapshot.Name = organization.ShortName
	}
	if organization.LegalName != nil {
		snapshot.LegalName = organization.LegalName
	}
	if organization.XRefCode1 != nil {
		snapshot.XRefCode = organization.XRefCode1
	}
	if organization.XRefCode2 != nil {
		snapshot.XRefCode2 = organization.XRefCode2
	}
	if organization.LegalEntityIdentifier != nil {
		snapshot.LegalEntityIdentifier = organization.LegalEntityIdentifier
	}
	if organization.CftcInterimCompliantIdentifier != nil {
		snapshot.CftcInterimCompliantIdentifier = organization.CftcInterimCompliantIdentifier
	}

	return snapshot
}

// AssembleOrganizationEntity will convert an api snapshot into an entity
func AssembleOrganizationEntity(snapshot api.Organization) *models.Organization {
	organization := new(models.Organization)

	organization.Address = *AssembleAddressEntity(snapshot.Address)

	if snapshot.ID != nil {
		organization.ID = uint(*snapshot.ID)
	}

	if snapshot.Name != nil {
		organization.ShortName = snapshot.Name
	}
	if snapshot.LegalName != nil {
		organization.LegalName = snapshot.LegalName
	}
	if snapshot.XRefCode != nil {
		organization.XRefCode1 = snapshot.XRefCode
	}
	if snapshot.XRefCode2 != nil {
		organization.XRefCode2 = snapshot.XRefCode2
	}
	if snapshot.LegalEntityIdentifier != nil {
		organization.LegalEntityIdentifier = snapshot.LegalEntityIdentifier
	}
	if snapshot.CftcInterimCompliantIdentifier != nil {
		organization.CftcInterimCompliantIdentifier = snapshot.CftcInterimCompliantIdentifier
	}

	return organization
}
