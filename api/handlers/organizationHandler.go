package handlers

import (
	api "organizationRestService/api/gen"
	"organizationRestService/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

// OrganizationHandlers type containing organizationService
type OrganizationHandlers struct {
	organizationService services.OrganizationServiceInterface
}

// NewOrganizationHandler creates new instance of the handler
func NewOrganizationHandler(organizationService services.OrganizationServiceInterface) *OrganizationHandlers {
	handler := new(OrganizationHandlers)
	handler.organizationService = organizationService
	return handler
}

//GetOrganizations gets all the orgs
func (orgHandler *OrganizationHandlers) GetOrganizations(ctx echo.Context, params api.GetOrganizationsParams) error {
	start := 0
	if params.Start != nil {
		start = *params.Start
	}
	limit := 100
	if params.Limit != nil {
		limit = *params.Limit
	}

	content, err := orgHandler.organizationService.GetOrganizations(start, limit)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &content)
	}

	return ctx.JSON(http.StatusOK, &content)
}

// GetOrganization Gets organization by id (GET /organizations/{id})
func (orgHandler *OrganizationHandlers) GetOrganization(ctx echo.Context, id int) error {
	content, err := orgHandler.organizationService.GetOrganization(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &content)
	}

	return ctx.JSON(http.StatusOK, &content)
}

// CreateOrganization create an organization
func (orgHandler *OrganizationHandlers) CreateOrganization(ctx echo.Context) error {
	organization := new(api.Organization)
	if err := ctx.Bind(organization); err != nil {
		return err
	}

	resp, err := orgHandler.organizationService.Create(organization)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err})
	}

	return ctx.JSON(http.StatusCreated, &resp)
}
