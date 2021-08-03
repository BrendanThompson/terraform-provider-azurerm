// Package searchservice implements the Azure ARM Searchservice service API version 2019-05-06.
//
// Client that can be used to manage and query indexes and documents, as well as manage other resources, on a search
// service.
package searchservice

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"github.com/gofrs/uuid"
)

const (
	// DefaultSearchDNSSuffix is the default value for search dns suffix
	DefaultSearchDNSSuffix = "search.windows.net"
)

// BaseClient is the base client for Searchservice.
type BaseClient struct {
	autorest.Client
	SearchServiceName string
	SearchDNSSuffix   string
}

// New creates an instance of the BaseClient client.
func New(searchServiceName string) BaseClient {
	return NewWithoutDefaults(searchServiceName, DefaultSearchDNSSuffix)
}

// NewWithoutDefaults creates an instance of the BaseClient client.
func NewWithoutDefaults(searchServiceName string, searchDNSSuffix string) BaseClient {
	return BaseClient{
		Client:            autorest.NewClientWithUserAgent(UserAgent()),
		SearchServiceName: searchServiceName,
		SearchDNSSuffix:   searchDNSSuffix,
	}
}

// GetServiceStatistics gets service level statistics for a search service.
// Parameters:
// clientRequestID - the tracking ID sent with the request to help with debugging.
func (client BaseClient) GetServiceStatistics(ctx context.Context, clientRequestID *uuid.UUID) (result ServiceStatistics, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GetServiceStatistics")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetServiceStatisticsPreparer(ctx, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.BaseClient", "GetServiceStatistics", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetServiceStatisticsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "searchservice.BaseClient", "GetServiceStatistics", resp, "Failure sending request")
		return
	}

	result, err = client.GetServiceStatisticsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.BaseClient", "GetServiceStatistics", resp, "Failure responding to request")
		return
	}

	return
}

// GetServiceStatisticsPreparer prepares the GetServiceStatistics request.
func (client BaseClient) GetServiceStatisticsPreparer(ctx context.Context, clientRequestID *uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"searchDnsSuffix":   client.SearchDNSSuffix,
		"searchServiceName": client.SearchServiceName,
	}

	const APIVersion = "2019-05-06"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{searchServiceName}.{searchDnsSuffix}", urlParameters),
		autorest.WithPath("/servicestats"),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetServiceStatisticsSender sends the GetServiceStatistics request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GetServiceStatisticsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetServiceStatisticsResponder handles the response to the GetServiceStatistics request. The method always
// closes the http.Response Body.
func (client BaseClient) GetServiceStatisticsResponder(resp *http.Response) (result ServiceStatistics, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
