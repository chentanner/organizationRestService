package services

import (
	"fmt"
	"organizationRestService/api/assemblers"
	api "organizationRestService/api/gen"
	"organizationRestService/appErrors"
	"organizationRestService/models"
	"organizationRestService/models/stores"
)

//OrganizationServiceInterface interface for service
type OrganizationServiceInterface interface {
	GetOrganizations(start int, limit int) (*api.OrganizationCollection, error)
	GetOrganization(id int) (*api.Organization, error)
	Create(org *api.Organization) (*models.Organization, error)
}

// OrganizationService type containing organizationService
type OrganizationService struct {
	// Needs a reference to the stores/data access layer
	organizationStore stores.OrganizationStoreInterface
	validator         ServiceValidatorInterface
}

// NewOrganizationService creates a new instance
func NewOrganizationService(organizationStore stores.OrganizationStoreInterface, validator ServiceValidatorInterface) *OrganizationService {
	service := new(OrganizationService)
	service.organizationStore = organizationStore
	service.validator = validator
	return service
}

// GetOrganizations get organizations
func (service *OrganizationService) GetOrganizations(start int, limit int) (*api.OrganizationCollection, error) {
	organizations, count, err := service.organizationStore.GetOrganizations(start, limit)

	if err != nil {
		fmt.Printf("\nOrganizationService: GetOrganizations error: %v\n", err)
		collection := models.ResourceCollection{ErrorCode: "ABCD", ErrorMessage: err.Error()}
		assemblers.AssembleOrganizationSnapshotCollection(collection, organizations)
		return assemblers.AssembleOrganizationSnapshotCollection(collection, organizations), err
	}

	// convert to anonymous interface type
	organizationsSlice := make([]interface{}, len(organizations))
	for i, v := range organizations {
		organizationsSlice[i] = v
	}

	resourceCollection := models.ResourceCollection{
		Items: organizationsSlice,
		Count: count,
		Start: start,
		Limit: limit,
	}
	content := assemblers.AssembleOrganizationSnapshotCollection(resourceCollection, organizations)
	fmt.Printf("%v", content)
	return content, nil
}

// GetOrganization from the DB
func (service *OrganizationService) GetOrganization(id int) (*api.Organization, error) {
	org, err := service.organizationStore.GetOrganization(id)
	if err != nil { //Org not found!
		return nil, err
	}

	orgSnapshot := assemblers.AssembleOrganizationApiSnapshot(org)

	return orgSnapshot, nil
}

// Create will create a new organization
func (service *OrganizationService) Create(orgSnapshot *api.Organization) (*models.Organization, error) {
	orgEntity := assemblers.AssembleOrganizationEntity(*orgSnapshot)

	isUniqueErr := service.validateUniqueOrganizationNames(orgEntity)
	if isUniqueErr != nil {
		return nil, isUniqueErr
	}
	err := service.validator.ValidateStruct(orgEntity)
	if err != nil {
		return nil, err
	}

	createErr := service.organizationStore.CreateOrganization(orgEntity)

	return orgEntity, createErr
}

func (service *OrganizationService) validateUniqueOrganizationNames(newOrganization *models.Organization) error {

	org, err := service.organizationStore.FindOrganizationsByShortName(*newOrganization.ShortName)
	if err != nil {
		return err
	}
	if org != nil {
		return appErrors.NewRuntimeError(appErrors.DUPLICATE_ORG_SHORT_NAME)
	}

	org, err = service.organizationStore.FindOrganizationsByLegalName(*newOrganization.LegalName)
	if err != nil {
		return err
	}
	if org != nil {
		return appErrors.NewRuntimeError(appErrors.DUPLICATE_ORG_LEGAL_NAME)
	}

	return nil
}
