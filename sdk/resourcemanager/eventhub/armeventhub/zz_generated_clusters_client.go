//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventhub

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ClustersClient contains the methods for the Clusters group.
// Don't use this type directly, use NewClustersClient() instead.
type ClustersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewClustersClient creates a new instance of ClustersClient with the specified values.
// subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewClustersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClustersClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ClustersClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates an instance of an Event Hubs Cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// resourceGroupName - Name of the resource group within the azure subscription.
// clusterName - The name of the Event Hubs Cluster.
// parameters - Parameters for creating a eventhub cluster resource.
// options - ClustersClientBeginCreateOrUpdateOptions contains the optional parameters for the ClustersClient.BeginCreateOrUpdate
// method.
func (client *ClustersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, parameters Cluster, options *ClustersClientBeginCreateOrUpdateOptions) (*runtime.Poller[ClustersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, clusterName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ClustersClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ClustersClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates an instance of an Event Hubs Cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
func (client *ClustersClient) createOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, parameters Cluster, options *ClustersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, clusterName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ClustersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, parameters Cluster, options *ClustersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/clusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes an existing Event Hubs Cluster. This operation is idempotent.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// resourceGroupName - Name of the resource group within the azure subscription.
// clusterName - The name of the Event Hubs Cluster.
// options - ClustersClientBeginDeleteOptions contains the optional parameters for the ClustersClient.BeginDelete method.
func (client *ClustersClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, options *ClustersClientBeginDeleteOptions) (*runtime.Poller[ClustersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ClustersClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ClustersClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes an existing Event Hubs Cluster. This operation is idempotent.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
func (client *ClustersClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, options *ClustersClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ClustersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *ClustersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/clusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the resource description of the specified Event Hubs Cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// resourceGroupName - Name of the resource group within the azure subscription.
// clusterName - The name of the Event Hubs Cluster.
// options - ClustersClientGetOptions contains the optional parameters for the ClustersClient.Get method.
func (client *ClustersClient) Get(ctx context.Context, resourceGroupName string, clusterName string, options *ClustersClientGetOptions) (ClustersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return ClustersClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClustersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClustersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ClustersClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *ClustersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/clusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ClustersClient) getHandleResponse(resp *http.Response) (ClustersClientGetResponse, error) {
	result := ClustersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Cluster); err != nil {
		return ClustersClientGetResponse{}, err
	}
	return result, nil
}

// ListAvailableClusterRegion - List the quantity of available pre-provisioned Event Hubs Clusters, indexed by Azure region.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// options - ClustersClientListAvailableClusterRegionOptions contains the optional parameters for the ClustersClient.ListAvailableClusterRegion
// method.
func (client *ClustersClient) ListAvailableClusterRegion(ctx context.Context, options *ClustersClientListAvailableClusterRegionOptions) (ClustersClientListAvailableClusterRegionResponse, error) {
	req, err := client.listAvailableClusterRegionCreateRequest(ctx, options)
	if err != nil {
		return ClustersClientListAvailableClusterRegionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClustersClientListAvailableClusterRegionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClustersClientListAvailableClusterRegionResponse{}, runtime.NewResponseError(resp)
	}
	return client.listAvailableClusterRegionHandleResponse(resp)
}

// listAvailableClusterRegionCreateRequest creates the ListAvailableClusterRegion request.
func (client *ClustersClient) listAvailableClusterRegionCreateRequest(ctx context.Context, options *ClustersClientListAvailableClusterRegionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.EventHub/availableClusterRegions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAvailableClusterRegionHandleResponse handles the ListAvailableClusterRegion response.
func (client *ClustersClient) listAvailableClusterRegionHandleResponse(resp *http.Response) (ClustersClientListAvailableClusterRegionResponse, error) {
	result := ClustersClientListAvailableClusterRegionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailableClustersList); err != nil {
		return ClustersClientListAvailableClusterRegionResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists the available Event Hubs Clusters within an ARM resource group
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// resourceGroupName - Name of the resource group within the azure subscription.
// options - ClustersClientListByResourceGroupOptions contains the optional parameters for the ClustersClient.ListByResourceGroup
// method.
func (client *ClustersClient) NewListByResourceGroupPager(resourceGroupName string, options *ClustersClientListByResourceGroupOptions) *runtime.Pager[ClustersClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClustersClientListByResourceGroupResponse]{
		More: func(page ClustersClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClustersClientListByResourceGroupResponse) (ClustersClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ClustersClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ClustersClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ClustersClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ClustersClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ClustersClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/clusters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ClustersClient) listByResourceGroupHandleResponse(resp *http.Response) (ClustersClientListByResourceGroupResponse, error) {
	result := ClustersClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterListResult); err != nil {
		return ClustersClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists the available Event Hubs Clusters within an ARM resource group
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// options - ClustersClientListBySubscriptionOptions contains the optional parameters for the ClustersClient.ListBySubscription
// method.
func (client *ClustersClient) NewListBySubscriptionPager(options *ClustersClientListBySubscriptionOptions) *runtime.Pager[ClustersClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClustersClientListBySubscriptionResponse]{
		More: func(page ClustersClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClustersClientListBySubscriptionResponse) (ClustersClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ClustersClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ClustersClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ClustersClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ClustersClient) listBySubscriptionCreateRequest(ctx context.Context, options *ClustersClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.EventHub/clusters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ClustersClient) listBySubscriptionHandleResponse(resp *http.Response) (ClustersClientListBySubscriptionResponse, error) {
	result := ClustersClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterListResult); err != nil {
		return ClustersClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// ListNamespaces - List all Event Hubs Namespace IDs in an Event Hubs Dedicated Cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// resourceGroupName - Name of the resource group within the azure subscription.
// clusterName - The name of the Event Hubs Cluster.
// options - ClustersClientListNamespacesOptions contains the optional parameters for the ClustersClient.ListNamespaces method.
func (client *ClustersClient) ListNamespaces(ctx context.Context, resourceGroupName string, clusterName string, options *ClustersClientListNamespacesOptions) (ClustersClientListNamespacesResponse, error) {
	req, err := client.listNamespacesCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return ClustersClientListNamespacesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClustersClientListNamespacesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClustersClientListNamespacesResponse{}, runtime.NewResponseError(resp)
	}
	return client.listNamespacesHandleResponse(resp)
}

// listNamespacesCreateRequest creates the ListNamespaces request.
func (client *ClustersClient) listNamespacesCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *ClustersClientListNamespacesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/clusters/{clusterName}/namespaces"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listNamespacesHandleResponse handles the ListNamespaces response.
func (client *ClustersClient) listNamespacesHandleResponse(resp *http.Response) (ClustersClientListNamespacesResponse, error) {
	result := ClustersClientListNamespacesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EHNamespaceIDListResult); err != nil {
		return ClustersClientListNamespacesResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Modifies mutable properties on the Event Hubs Cluster. This operation is idempotent.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// resourceGroupName - Name of the resource group within the azure subscription.
// clusterName - The name of the Event Hubs Cluster.
// parameters - The properties of the Event Hubs Cluster which should be updated.
// options - ClustersClientBeginUpdateOptions contains the optional parameters for the ClustersClient.BeginUpdate method.
func (client *ClustersClient) BeginUpdate(ctx context.Context, resourceGroupName string, clusterName string, parameters Cluster, options *ClustersClientBeginUpdateOptions) (*runtime.Poller[ClustersClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, clusterName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ClustersClientUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ClustersClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Modifies mutable properties on the Event Hubs Cluster. This operation is idempotent.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
func (client *ClustersClient) update(ctx context.Context, resourceGroupName string, clusterName string, parameters Cluster, options *ClustersClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, clusterName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ClustersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, parameters Cluster, options *ClustersClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/clusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}
