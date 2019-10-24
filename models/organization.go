package models

import (
	"github.com/jinzhu/gorm"
)

//Organization a struct to rep an Organization
type Organization struct {
	gorm.Model
	Address
	ShortName                      *string `json:"shortName,omitempty" validate:"required"`
	LegalName                      *string `json:"legalName,omitempty" validate:"required"`
	XRefCode1                      *string `json:"xrefCode1,omitempty"`
	XRefCode2                      *string `json:"xrefCode2,omitempty"`
	LegalEntityIdentifier          *string `json:"legalEntityIdentifier,omitempty" validate:"len=20"`
	CftcInterimCompliantIdentifier *string `json:"cftcIdentifier,omitempty validate:"len=20""`
}
