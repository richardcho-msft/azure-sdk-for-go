//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery

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

// ReplicationStorageClassificationsClient contains the methods for the ReplicationStorageClassifications group.
// Don't use this type directly, use NewReplicationStorageClassificationsClient() instead.
type ReplicationStorageClassificationsClient struct {
	host              string
	resourceName      string
	resourceGroupName string
	subscriptionID    string
	pl                runtime.Pipeline
}

// NewReplicationStorageClassificationsClient creates a new instance of ReplicationStorageClassificationsClient with the specified values.
// resourceName - The name of the recovery services vault.
// resourceGroupName - The name of the resource group where the recovery services vault is present.
// subscriptionID - The subscription Id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewReplicationStorageClassificationsClient(resourceName string, resourceGroupName string, subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ReplicationStorageClassificationsClient, error) {
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
	client := &ReplicationStorageClassificationsClient{
		resourceName:      resourceName,
		resourceGroupName: resourceGroupName,
		subscriptionID:    subscriptionID,
		host:              ep,
		pl:                pl,
	}
	return client, nil
}

// Get - Gets the details of the specified storage classification.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-01
// fabricName - Fabric name.
// storageClassificationName - Storage classification name.
// options - ReplicationStorageClassificationsClientGetOptions contains the optional parameters for the ReplicationStorageClassificationsClient.Get
// method.
func (client *ReplicationStorageClassificationsClient) Get(ctx context.Context, fabricName string, storageClassificationName string, options *ReplicationStorageClassificationsClientGetOptions) (ReplicationStorageClassificationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, fabricName, storageClassificationName, options)
	if err != nil {
		return ReplicationStorageClassificationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ReplicationStorageClassificationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ReplicationStorageClassificationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ReplicationStorageClassificationsClient) getCreateRequest(ctx context.Context, fabricName string, storageClassificationName string, options *ReplicationStorageClassificationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications/{storageClassificationName}"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if storageClassificationName == "" {
		return nil, errors.New("parameter storageClassificationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageClassificationName}", url.PathEscape(storageClassificationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReplicationStorageClassificationsClient) getHandleResponse(resp *http.Response) (ReplicationStorageClassificationsClientGetResponse, error) {
	result := ReplicationStorageClassificationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageClassification); err != nil {
		return ReplicationStorageClassificationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the storage classifications in the vault.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-01
// options - ReplicationStorageClassificationsClientListOptions contains the optional parameters for the ReplicationStorageClassificationsClient.List
// method.
func (client *ReplicationStorageClassificationsClient) NewListPager(options *ReplicationStorageClassificationsClientListOptions) *runtime.Pager[ReplicationStorageClassificationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReplicationStorageClassificationsClientListResponse]{
		More: func(page ReplicationStorageClassificationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReplicationStorageClassificationsClientListResponse) (ReplicationStorageClassificationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReplicationStorageClassificationsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ReplicationStorageClassificationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReplicationStorageClassificationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ReplicationStorageClassificationsClient) listCreateRequest(ctx context.Context, options *ReplicationStorageClassificationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationStorageClassifications"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReplicationStorageClassificationsClient) listHandleResponse(resp *http.Response) (ReplicationStorageClassificationsClientListResponse, error) {
	result := ReplicationStorageClassificationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageClassificationCollection); err != nil {
		return ReplicationStorageClassificationsClientListResponse{}, err
	}
	return result, nil
}

// NewListByReplicationFabricsPager - Lists the storage classifications available in the specified fabric.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-01
// fabricName - Site name of interest.
// options - ReplicationStorageClassificationsClientListByReplicationFabricsOptions contains the optional parameters for the
// ReplicationStorageClassificationsClient.ListByReplicationFabrics method.
func (client *ReplicationStorageClassificationsClient) NewListByReplicationFabricsPager(fabricName string, options *ReplicationStorageClassificationsClientListByReplicationFabricsOptions) *runtime.Pager[ReplicationStorageClassificationsClientListByReplicationFabricsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReplicationStorageClassificationsClientListByReplicationFabricsResponse]{
		More: func(page ReplicationStorageClassificationsClientListByReplicationFabricsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReplicationStorageClassificationsClientListByReplicationFabricsResponse) (ReplicationStorageClassificationsClientListByReplicationFabricsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByReplicationFabricsCreateRequest(ctx, fabricName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReplicationStorageClassificationsClientListByReplicationFabricsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ReplicationStorageClassificationsClientListByReplicationFabricsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReplicationStorageClassificationsClientListByReplicationFabricsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByReplicationFabricsHandleResponse(resp)
		},
	})
}

// listByReplicationFabricsCreateRequest creates the ListByReplicationFabrics request.
func (client *ReplicationStorageClassificationsClient) listByReplicationFabricsCreateRequest(ctx context.Context, fabricName string, options *ReplicationStorageClassificationsClientListByReplicationFabricsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByReplicationFabricsHandleResponse handles the ListByReplicationFabrics response.
func (client *ReplicationStorageClassificationsClient) listByReplicationFabricsHandleResponse(resp *http.Response) (ReplicationStorageClassificationsClientListByReplicationFabricsResponse, error) {
	result := ReplicationStorageClassificationsClientListByReplicationFabricsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageClassificationCollection); err != nil {
		return ReplicationStorageClassificationsClientListByReplicationFabricsResponse{}, err
	}
	return result, nil
}
