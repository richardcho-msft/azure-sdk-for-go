package siterecovery

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ReplicationStorageClassificationMappingsClient is the client for the ReplicationStorageClassificationMappings
// methods of the Siterecovery service.
type ReplicationStorageClassificationMappingsClient struct {
	BaseClient
}

// NewReplicationStorageClassificationMappingsClient creates an instance of the
// ReplicationStorageClassificationMappingsClient client.
func NewReplicationStorageClassificationMappingsClient(subscriptionID string, resourceGroupName string, resourceName string) ReplicationStorageClassificationMappingsClient {
	return NewReplicationStorageClassificationMappingsClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName, resourceName)
}

// NewReplicationStorageClassificationMappingsClientWithBaseURI creates an instance of the
// ReplicationStorageClassificationMappingsClient client using a custom endpoint.  Use this when interacting with an
// Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewReplicationStorageClassificationMappingsClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string, resourceName string) ReplicationStorageClassificationMappingsClient {
	return ReplicationStorageClassificationMappingsClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName, resourceName)}
}

// Create the operation to create a storage classification mapping.
// Parameters:
// fabricName - fabric name.
// storageClassificationName - storage classification name.
// storageClassificationMappingName - storage classification mapping name.
// pairingInput - pairing input.
func (client ReplicationStorageClassificationMappingsClient) Create(ctx context.Context, fabricName string, storageClassificationName string, storageClassificationMappingName string, pairingInput StorageClassificationMappingInput) (result ReplicationStorageClassificationMappingsCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationStorageClassificationMappingsClient.Create")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, fabricName, storageClassificationName, storageClassificationMappingName, pairingInput)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationMappingsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationMappingsClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ReplicationStorageClassificationMappingsClient) CreatePreparer(ctx context.Context, fabricName string, storageClassificationName string, storageClassificationMappingName string, pairingInput StorageClassificationMappingInput) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":                       autorest.Encode("path", fabricName),
		"resourceGroupName":                autorest.Encode("path", client.ResourceGroupName),
		"resourceName":                     autorest.Encode("path", client.ResourceName),
		"storageClassificationMappingName": autorest.Encode("path", storageClassificationMappingName),
		"storageClassificationName":        autorest.Encode("path", storageClassificationName),
		"subscriptionId":                   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications/{storageClassificationName}/replicationStorageClassificationMappings/{storageClassificationMappingName}", pathParameters),
		autorest.WithJSON(pairingInput),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationStorageClassificationMappingsClient) CreateSender(req *http.Request) (future ReplicationStorageClassificationMappingsCreateFuture, err error) {
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

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ReplicationStorageClassificationMappingsClient) CreateResponder(resp *http.Response) (result StorageClassificationMapping, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the operation to delete a storage classification mapping.
// Parameters:
// fabricName - fabric name.
// storageClassificationName - storage classification name.
// storageClassificationMappingName - storage classification mapping name.
func (client ReplicationStorageClassificationMappingsClient) Delete(ctx context.Context, fabricName string, storageClassificationName string, storageClassificationMappingName string) (result ReplicationStorageClassificationMappingsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationStorageClassificationMappingsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, fabricName, storageClassificationName, storageClassificationMappingName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationMappingsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationMappingsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ReplicationStorageClassificationMappingsClient) DeletePreparer(ctx context.Context, fabricName string, storageClassificationName string, storageClassificationMappingName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":                       autorest.Encode("path", fabricName),
		"resourceGroupName":                autorest.Encode("path", client.ResourceGroupName),
		"resourceName":                     autorest.Encode("path", client.ResourceName),
		"storageClassificationMappingName": autorest.Encode("path", storageClassificationMappingName),
		"storageClassificationName":        autorest.Encode("path", storageClassificationName),
		"subscriptionId":                   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications/{storageClassificationName}/replicationStorageClassificationMappings/{storageClassificationMappingName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationStorageClassificationMappingsClient) DeleteSender(req *http.Request) (future ReplicationStorageClassificationMappingsDeleteFuture, err error) {
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
func (client ReplicationStorageClassificationMappingsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the details of the specified storage classification mapping.
// Parameters:
// fabricName - fabric name.
// storageClassificationName - storage classification name.
// storageClassificationMappingName - storage classification mapping name.
func (client ReplicationStorageClassificationMappingsClient) Get(ctx context.Context, fabricName string, storageClassificationName string, storageClassificationMappingName string) (result StorageClassificationMapping, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationStorageClassificationMappingsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, fabricName, storageClassificationName, storageClassificationMappingName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationMappingsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationMappingsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationMappingsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ReplicationStorageClassificationMappingsClient) GetPreparer(ctx context.Context, fabricName string, storageClassificationName string, storageClassificationMappingName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":                       autorest.Encode("path", fabricName),
		"resourceGroupName":                autorest.Encode("path", client.ResourceGroupName),
		"resourceName":                     autorest.Encode("path", client.ResourceName),
		"storageClassificationMappingName": autorest.Encode("path", storageClassificationMappingName),
		"storageClassificationName":        autorest.Encode("path", storageClassificationName),
		"subscriptionId":                   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications/{storageClassificationName}/replicationStorageClassificationMappings/{storageClassificationMappingName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationStorageClassificationMappingsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ReplicationStorageClassificationMappingsClient) GetResponder(resp *http.Response) (result StorageClassificationMapping, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the storage classification mappings in the vault.
func (client ReplicationStorageClassificationMappingsClient) List(ctx context.Context) (result StorageClassificationMappingCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationStorageClassificationMappingsClient.List")
		defer func() {
			sc := -1
			if result.scmc.Response.Response != nil {
				sc = result.scmc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationMappingsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.scmc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationMappingsClient", "List", resp, "Failure sending request")
		return
	}

	result.scmc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationMappingsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.scmc.hasNextLink() && result.scmc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ReplicationStorageClassificationMappingsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationStorageClassificationMappings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationStorageClassificationMappingsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ReplicationStorageClassificationMappingsClient) ListResponder(resp *http.Response) (result StorageClassificationMappingCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ReplicationStorageClassificationMappingsClient) listNextResults(ctx context.Context, lastResults StorageClassificationMappingCollection) (result StorageClassificationMappingCollection, err error) {
	req, err := lastResults.storageClassificationMappingCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationMappingsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationMappingsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationMappingsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ReplicationStorageClassificationMappingsClient) ListComplete(ctx context.Context) (result StorageClassificationMappingCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationStorageClassificationMappingsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListByReplicationStorageClassifications lists the storage classification mappings for the fabric.
// Parameters:
// fabricName - fabric name.
// storageClassificationName - storage classification name.
func (client ReplicationStorageClassificationMappingsClient) ListByReplicationStorageClassifications(ctx context.Context, fabricName string, storageClassificationName string) (result StorageClassificationMappingCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationStorageClassificationMappingsClient.ListByReplicationStorageClassifications")
		defer func() {
			sc := -1
			if result.scmc.Response.Response != nil {
				sc = result.scmc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByReplicationStorageClassificationsNextResults
	req, err := client.ListByReplicationStorageClassificationsPreparer(ctx, fabricName, storageClassificationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationMappingsClient", "ListByReplicationStorageClassifications", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByReplicationStorageClassificationsSender(req)
	if err != nil {
		result.scmc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationMappingsClient", "ListByReplicationStorageClassifications", resp, "Failure sending request")
		return
	}

	result.scmc, err = client.ListByReplicationStorageClassificationsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationMappingsClient", "ListByReplicationStorageClassifications", resp, "Failure responding to request")
		return
	}
	if result.scmc.hasNextLink() && result.scmc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByReplicationStorageClassificationsPreparer prepares the ListByReplicationStorageClassifications request.
func (client ReplicationStorageClassificationMappingsClient) ListByReplicationStorageClassificationsPreparer(ctx context.Context, fabricName string, storageClassificationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":                autorest.Encode("path", fabricName),
		"resourceGroupName":         autorest.Encode("path", client.ResourceGroupName),
		"resourceName":              autorest.Encode("path", client.ResourceName),
		"storageClassificationName": autorest.Encode("path", storageClassificationName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications/{storageClassificationName}/replicationStorageClassificationMappings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByReplicationStorageClassificationsSender sends the ListByReplicationStorageClassifications request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationStorageClassificationMappingsClient) ListByReplicationStorageClassificationsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByReplicationStorageClassificationsResponder handles the response to the ListByReplicationStorageClassifications request. The method always
// closes the http.Response Body.
func (client ReplicationStorageClassificationMappingsClient) ListByReplicationStorageClassificationsResponder(resp *http.Response) (result StorageClassificationMappingCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByReplicationStorageClassificationsNextResults retrieves the next set of results, if any.
func (client ReplicationStorageClassificationMappingsClient) listByReplicationStorageClassificationsNextResults(ctx context.Context, lastResults StorageClassificationMappingCollection) (result StorageClassificationMappingCollection, err error) {
	req, err := lastResults.storageClassificationMappingCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationMappingsClient", "listByReplicationStorageClassificationsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByReplicationStorageClassificationsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationMappingsClient", "listByReplicationStorageClassificationsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByReplicationStorageClassificationsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationStorageClassificationMappingsClient", "listByReplicationStorageClassificationsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByReplicationStorageClassificationsComplete enumerates all values, automatically crossing page boundaries as required.
func (client ReplicationStorageClassificationMappingsClient) ListByReplicationStorageClassificationsComplete(ctx context.Context, fabricName string, storageClassificationName string) (result StorageClassificationMappingCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationStorageClassificationMappingsClient.ListByReplicationStorageClassifications")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByReplicationStorageClassifications(ctx, fabricName, storageClassificationName)
	return
}
