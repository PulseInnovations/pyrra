/*
 * Pyrra
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
)

// ObjectivesApiRouter defines the required methods for binding the api requests to a responses for the ObjectivesApi
// The ObjectivesApiRouter implementation should parse necessary information from the http request,
// pass the data to a ObjectivesApiServicer to perform the required actions, then write the service results to the http response.
type ObjectivesApiRouter interface {
	GetMultiBurnrateAlerts(http.ResponseWriter, *http.Request)
	GetObjectiveErrorBudget(http.ResponseWriter, *http.Request)
	GetObjectiveStatus(http.ResponseWriter, *http.Request)
	GetREDErrors(http.ResponseWriter, *http.Request)
	GetREDRequests(http.ResponseWriter, *http.Request)
	ListObjectives(http.ResponseWriter, *http.Request)
}

// ObjectivesApiServicer defines the api actions for the ObjectivesApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type ObjectivesApiServicer interface {
	GetMultiBurnrateAlerts(context.Context, string, string, bool, bool) (ImplResponse, error)
	GetObjectiveErrorBudget(context.Context, string, string, int32, int32) (ImplResponse, error)
	GetObjectiveStatus(context.Context, string, string) (ImplResponse, error)
	GetREDErrors(context.Context, string, string, int32, int32) (ImplResponse, error)
	GetREDRequests(context.Context, string, string, int32, int32) (ImplResponse, error)
	ListObjectives(context.Context, string) (ImplResponse, error)
}
