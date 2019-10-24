package models

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB //database

//AbstractModel base model
type AbstractModel struct {
	Expired string `json:"isExpired"`
	Version string `json:"version"`
}

// ResourceCollection base response resource
type ResourceCollection struct {
	ErrorCode    string        `json:"errorCode"`
	ErrorMessage string        `json:"errorMessage"`
	Items        []interface{} `json:"items"`
	Count        int           `json:"count"`
	Start        int           `json:"start"`
	Limit        int           `json:"limit"`
	Links        []Link        `json:"links"`
}

// Link restful link
type Link struct {
	Method           string
	Rel              string
	Href             string
	URI              string
	LinkType         string
	ResponseType     string
	ItemType         string
	ResponseItemType string
	Title            string
}
