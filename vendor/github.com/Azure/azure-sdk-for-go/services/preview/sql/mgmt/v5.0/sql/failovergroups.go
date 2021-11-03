package sql

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// FailoverGroupsClient is the the Azure SQL Database management API provides a RESTful set of web services that
// interact with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update,
// and delete databases.
type FailoverGroupsClient struct {
	BaseClient
}

// NewFailoverGroupsClient creates an instance of the FailoverGroupsClient client.
func NewFailoverGroupsClient(subscriptionID string) FailoverGroupsClient {
	return NewFailoverGroupsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewFailoverGroupsClientWithBaseURI creates an instance of the FailoverGroupsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewFailoverGroupsClientWithBaseURI(baseURI string, subscriptionID string) FailoverGroupsClient {
	return FailoverGroupsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a failover group.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server containing the failover group.
// failoverGroupName - the name of the failover group.
// parameters - the failover group parameters.
func (client FailoverGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string, parameters FailoverGroup) (result FailoverGroupsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FailoverGroupsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.FailoverGroupProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.FailoverGroupProperties.ReadWriteEndpoint", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.FailoverGroupProperties.PartnerServers", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("sql.FailoverGroupsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serverName, failoverGroupName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client FailoverGroupsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string, parameters FailoverGroup) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"failoverGroupName": autorest.Encode("path", failoverGroupName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.Location = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/failoverGroups/{failoverGroupName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client FailoverGroupsClient) CreateOrUpdateSender(req *http.Request) (future FailoverGroupsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client FailoverGroupsClient) CreateOrUpdateResponder(resp *http.Response) (result FailoverGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a failover group.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server containing the failover group.
// failoverGroupName - the name of the failover group.
func (client FailoverGroupsClient) Delete(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string) (result FailoverGroupsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FailoverGroupsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, serverName, failoverGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client FailoverGroupsClient) DeletePreparer(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"failoverGroupName": autorest.Encode("path", failoverGroupName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/failoverGroups/{failoverGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client FailoverGroupsClient) DeleteSender(req *http.Request) (future FailoverGroupsDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client FailoverGroupsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Failover fails over from the current primary server to this server.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server containing the failover group.
// failoverGroupName - the name of the failover group.
func (client FailoverGroupsClient) Failover(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string) (result FailoverGroupsFailoverFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FailoverGroupsClient.Failover")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.FailoverPreparer(ctx, resourceGroupName, serverName, failoverGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "Failover", nil, "Failure preparing request")
		return
	}

	result, err = client.FailoverSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "Failover", result.Response(), "Failure sending request")
		return
	}

	return
}

// FailoverPreparer prepares the Failover request.
func (client FailoverGroupsClient) FailoverPreparer(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"failoverGroupName": autorest.Encode("path", failoverGroupName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/failoverGroups/{failoverGroupName}/failover", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// FailoverSender sends the Failover request. The method will close the
// http.Response Body if it receives an error.
func (client FailoverGroupsClient) FailoverSender(req *http.Request) (future FailoverGroupsFailoverFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// FailoverResponder handles the response to the Failover request. The method always
// closes the http.Response Body.
func (client FailoverGroupsClient) FailoverResponder(resp *http.Response) (result FailoverGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ForceFailoverAllowDataLoss fails over from the current primary server to this server. This operation might result in
// data loss.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server containing the failover group.
// failoverGroupName - the name of the failover group.
func (client FailoverGroupsClient) ForceFailoverAllowDataLoss(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string) (result FailoverGroupsForceFailoverAllowDataLossFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FailoverGroupsClient.ForceFailoverAllowDataLoss")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ForceFailoverAllowDataLossPreparer(ctx, resourceGroupName, serverName, failoverGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "ForceFailoverAllowDataLoss", nil, "Failure preparing request")
		return
	}

	result, err = client.ForceFailoverAllowDataLossSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "ForceFailoverAllowDataLoss", result.Response(), "Failure sending request")
		return
	}

	return
}

// ForceFailoverAllowDataLossPreparer prepares the ForceFailoverAllowDataLoss request.
func (client FailoverGroupsClient) ForceFailoverAllowDataLossPreparer(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"failoverGroupName": autorest.Encode("path", failoverGroupName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/failoverGroups/{failoverGroupName}/forceFailoverAllowDataLoss", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ForceFailoverAllowDataLossSender sends the ForceFailoverAllowDataLoss request. The method will close the
// http.Response Body if it receives an error.
func (client FailoverGroupsClient) ForceFailoverAllowDataLossSender(req *http.Request) (future FailoverGroupsForceFailoverAllowDataLossFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// ForceFailoverAllowDataLossResponder handles the response to the ForceFailoverAllowDataLoss request. The method always
// closes the http.Response Body.
func (client FailoverGroupsClient) ForceFailoverAllowDataLossResponder(resp *http.Response) (result FailoverGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets a failover group.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server containing the failover group.
// failoverGroupName - the name of the failover group.
func (client FailoverGroupsClient) Get(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string) (result FailoverGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FailoverGroupsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, serverName, failoverGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client FailoverGroupsClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"failoverGroupName": autorest.Encode("path", failoverGroupName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/failoverGroups/{failoverGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client FailoverGroupsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client FailoverGroupsClient) GetResponder(resp *http.Response) (result FailoverGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByServer lists the failover groups in a server.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server containing the failover group.
func (client FailoverGroupsClient) ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result FailoverGroupListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FailoverGroupsClient.ListByServer")
		defer func() {
			sc := -1
			if result.fglr.Response.Response != nil {
				sc = result.fglr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByServerNextResults
	req, err := client.ListByServerPreparer(ctx, resourceGroupName, serverName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "ListByServer", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.fglr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "ListByServer", resp, "Failure sending request")
		return
	}

	result.fglr, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "ListByServer", resp, "Failure responding to request")
		return
	}
	if result.fglr.hasNextLink() && result.fglr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByServerPreparer prepares the ListByServer request.
func (client FailoverGroupsClient) ListByServerPreparer(ctx context.Context, resourceGroupName string, serverName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/failoverGroups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServerSender sends the ListByServer request. The method will close the
// http.Response Body if it receives an error.
func (client FailoverGroupsClient) ListByServerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByServerResponder handles the response to the ListByServer request. The method always
// closes the http.Response Body.
func (client FailoverGroupsClient) ListByServerResponder(resp *http.Response) (result FailoverGroupListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByServerNextResults retrieves the next set of results, if any.
func (client FailoverGroupsClient) listByServerNextResults(ctx context.Context, lastResults FailoverGroupListResult) (result FailoverGroupListResult, err error) {
	req, err := lastResults.failoverGroupListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "listByServerNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "listByServerNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "listByServerNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByServerComplete enumerates all values, automatically crossing page boundaries as required.
func (client FailoverGroupsClient) ListByServerComplete(ctx context.Context, resourceGroupName string, serverName string) (result FailoverGroupListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FailoverGroupsClient.ListByServer")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByServer(ctx, resourceGroupName, serverName)
	return
}

// Update updates a failover group.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server containing the failover group.
// failoverGroupName - the name of the failover group.
// parameters - the failover group parameters.
func (client FailoverGroupsClient) Update(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string, parameters FailoverGroupUpdate) (result FailoverGroupsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FailoverGroupsClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, serverName, failoverGroupName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.FailoverGroupsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client FailoverGroupsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string, parameters FailoverGroupUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"failoverGroupName": autorest.Encode("path", failoverGroupName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/failoverGroups/{failoverGroupName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client FailoverGroupsClient) UpdateSender(req *http.Request) (future FailoverGroupsUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client FailoverGroupsClient) UpdateResponder(resp *http.Response) (result FailoverGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
