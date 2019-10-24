package assemblers

import (
	api "organizationRestService/api/gen"
	"organizationRestService/models"
)

// AssembleSnapshotCollection will convert an entity collection into an api snapshot collection
func AssembleSnapshotCollection(entityCollection models.ResourceCollection) *api.BaseCollection {
	snapshotCollection := new(api.BaseCollection)

	snapshotCollection.Count = new(int)
	*snapshotCollection.Count = entityCollection.Count
	snapshotCollection.Limit = new(int)
	*snapshotCollection.Limit = entityCollection.Limit
	snapshotCollection.Start = new(int)
	*snapshotCollection.Start = entityCollection.Start
	snapshotCollection.ErrorCode = new(string)
	*snapshotCollection.ErrorCode = entityCollection.ErrorCode
	snapshotCollection.ErrorMessage = new(string)
	*snapshotCollection.ErrorMessage = entityCollection.ErrorMessage

	// snapshotCollection.Links = entityCollection.Links

	return snapshotCollection
}

// AssembleEntityCollection will convert an api snapshot collection into an entity collection
func AssembleEntityCollection(snapshotCollection api.BaseCollection) *models.ResourceCollection {
	entityCollection := new(models.ResourceCollection)

	entityCollection.Count = *snapshotCollection.Count
	entityCollection.ErrorCode = *snapshotCollection.ErrorCode
	entityCollection.ErrorMessage = *snapshotCollection.ErrorMessage
	entityCollection.Limit = *snapshotCollection.Limit
	entityCollection.Start = *snapshotCollection.Start

	// entityCollection.Links = snapshotCollection.Links

	return entityCollection
}
