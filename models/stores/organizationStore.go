package stores

import (
	"errors"
	"fmt"
	"organizationRestService/models"

	"github.com/jinzhu/gorm"
)

// OrganizationStoreInterface is an interface to be implemented by a Store to perform operations and manage persistance on the data.
type OrganizationStoreInterface interface {
	GetOrganizations(start int, limit int) ([]models.Organization, int, error)
	GetOrganization(id int) (*models.Organization, error)
	CreateOrganization(*models.Organization) error
	FindOrganizationsByShortName(shortName string) (*models.Organization, error)
	FindOrganizationsByLegalName(legalName string) (*models.Organization, error)
}

// OrganizationStore struct
type OrganizationStore struct {
	db *gorm.DB
}

// NewOrganizationStore constructor
func NewOrganizationStore(db *gorm.DB) *OrganizationStore {
	store := new(OrganizationStore)
	store.db = db
	db.LogMode(true)
	return store
}

// GetOrganizations from the DB. Supports specifying a start and limit paramater to access organizations in a paged fashion.
func (store *OrganizationStore) GetOrganizations(start int, limit int) ([]models.Organization, int, error) {
	var organizations []models.Organization
	var count int
	db := store.db.Model(&models.Organization{})
	countErr := db.Count(&count).Error
	err := db.Offset(start).Limit(limit).Find(&organizations).Error

	if err != nil {
		fmt.Print("OrganizationStore: GetOrganizations error: %v", err)
		return nil, -1, err
		return organizations, count, err
	}

	if countErr != nil {
		fmt.Print("OrganizationStore: GetOrganizations count error: %v", countErr)
		return organizations, count, countErr
	}

	return organizations, count, nil
}

// GetOrganization from the DB
func (store *OrganizationStore) GetOrganization(id int) (*models.Organization, error) {

	org := &models.Organization{}
	store.db.Table("organizations").Where("ID = ?", id).First(org)
	if org.ShortName == nil || *org.ShortName == "" { //Org not found!
		return nil, errors.New("organization not found")
	}

	return org, nil
}

// CreateOrganization will access the database and create an organization
func (store *OrganizationStore) CreateOrganization(organization *models.Organization) error {
	return store.db.Create(organization).Error
}

// FindOrganizationsByShortName will find an organization by short name.
func (store *OrganizationStore) FindOrganizationsByShortName(shortName string) (*models.Organization, error) {
	org := &models.Organization{}
	if store.db.Table("organizations").Where("short_name = ?", shortName).First(org).RecordNotFound() {
		return nil, nil
	}

	return org, nil
}

// FindOrganizationsByLegalName will find an organization by short name.
func (store *OrganizationStore) FindOrganizationsByLegalName(legalName string) (*models.Organization, error) {
	org := &models.Organization{}
	if store.db.Table("organizations").Where("legal_name = ?", legalName).First(org).RecordNotFound() {
		return nil, nil
	}

	return org, nil
}
