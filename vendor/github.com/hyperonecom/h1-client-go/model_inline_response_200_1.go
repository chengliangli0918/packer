/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// InlineResponse2001 struct for InlineResponse2001
type InlineResponse2001 struct {
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Value     string    `json:"value"`
	CreatedOn time.Time `json:"createdOn"`
}