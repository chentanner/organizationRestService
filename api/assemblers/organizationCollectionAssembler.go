package assemblers

import (
	api "organizationRestService/api/gen"
	"organizationRestService/models"
)

// AssembleSnapshotCollection will convert an entity collection into an api snapshot collection
func AssembleOrganizationSnapshotCollection(entityCollection models.ResourceCollection, items []models.Organization) *api.OrganizationCollection {
	snapshotCollection := new(api.OrganizationCollection)
	snapshotCollection.BaseCollection = *AssembleSnapshotCollection(entityCollection)
	if items != nil {
		snapshotCollection.Items = new(api.Organizations)
		*snapshotCollection.Items = AssembleOrganizationApiSnapshots(items)
	}
	return snapshotCollection
}
