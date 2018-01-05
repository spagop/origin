package servermanagement

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// GatewayClient is the REST API for Azure Server Management Service.
type GatewayClient struct {
	ManagementClient
}

// NewGatewayClient creates an instance of the GatewayClient client.
func NewGatewayClient(subscriptionID string) GatewayClient {
	return NewGatewayClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewGatewayClientWithBaseURI creates an instance of the GatewayClient client.
func NewGatewayClientWithBaseURI(baseURI string, subscriptionID string) GatewayClient {
	return GatewayClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates or updates a ManagementService gateway. This method may poll for completion. Polling can be canceled
// by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// resourceGroupName is the resource group name uniquely identifies the resource group within the user subscriptionId.
// gatewayName is the gateway name (256 characters maximum). gatewayParameters is parameters supplied to the
// CreateOrUpdate operation.
func (client GatewayClient) Create(resourceGroupName string, gatewayName string, gatewayParameters GatewayParameters, cancel <-chan struct{}) (<-chan GatewayResource, <-chan error) {
	resultChan := make(chan GatewayResource, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9]+`, Chain: nil}}},
		{TargetValue: gatewayName,
			Constraints: []validation.Constraint{{Target: "gatewayName", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "gatewayName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "gatewayName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "servermanagement.GatewayClient", "Create")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result GatewayResource
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreatePreparer(resourceGroupName, gatewayName, gatewayParameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "Create", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "Create", resp, "Failure sending request")
			return
		}

		result, err = client.CreateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "Create", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreatePreparer prepares the Create request.
func (client GatewayClient) CreatePreparer(resourceGroupName string, gatewayName string, gatewayParameters GatewayParameters, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"gatewayName":       autorest.Encode("path", gatewayName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServerManagement/gateways/{gatewayName}", pathParameters),
		autorest.WithJSON(gatewayParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client GatewayClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client GatewayClient) CreateResponder(resp *http.Response) (result GatewayResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a gateway from a resource group.
//
// resourceGroupName is the resource group name uniquely identifies the resource group within the user subscriptionId.
// gatewayName is the gateway name (256 characters maximum).
func (client GatewayClient) Delete(resourceGroupName string, gatewayName string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9]+`, Chain: nil}}},
		{TargetValue: gatewayName,
			Constraints: []validation.Constraint{{Target: "gatewayName", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "gatewayName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "gatewayName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "servermanagement.GatewayClient", "Delete")
	}

	req, err := client.DeletePreparer(resourceGroupName, gatewayName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client GatewayClient) DeletePreparer(resourceGroupName string, gatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"gatewayName":       autorest.Encode("path", gatewayName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServerManagement/gateways/{gatewayName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client GatewayClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client GatewayClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a gateway.
//
// resourceGroupName is the resource group name uniquely identifies the resource group within the user subscriptionId.
// gatewayName is the gateway name (256 characters maximum) expand is gets subscription credentials which uniquely
// identify Microsoft Azure subscription. The subscription ID forms part of the URI for every service call.
func (client GatewayClient) Get(resourceGroupName string, gatewayName string, expand GatewayExpandOption) (result GatewayResource, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9]+`, Chain: nil}}},
		{TargetValue: gatewayName,
			Constraints: []validation.Constraint{{Target: "gatewayName", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "gatewayName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "gatewayName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "servermanagement.GatewayClient", "Get")
	}

	req, err := client.GetPreparer(resourceGroupName, gatewayName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client GatewayClient) GetPreparer(resourceGroupName string, gatewayName string, expand GatewayExpandOption) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"gatewayName":       autorest.Encode("path", gatewayName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(expand)) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServerManagement/gateways/{gatewayName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client GatewayClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client GatewayClient) GetResponder(resp *http.Response) (result GatewayResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetProfile gets a gateway profile. This method may poll for completion. Polling can be canceled by passing the
// cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the resource group name uniquely identifies the resource group within the user subscriptionId.
// gatewayName is the gateway name (256 characters maximum).
func (client GatewayClient) GetProfile(resourceGroupName string, gatewayName string, cancel <-chan struct{}) (<-chan GatewayProfile, <-chan error) {
	resultChan := make(chan GatewayProfile, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9]+`, Chain: nil}}},
		{TargetValue: gatewayName,
			Constraints: []validation.Constraint{{Target: "gatewayName", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "gatewayName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "gatewayName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "servermanagement.GatewayClient", "GetProfile")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result GatewayProfile
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.GetProfilePreparer(resourceGroupName, gatewayName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "GetProfile", nil, "Failure preparing request")
			return
		}

		resp, err := client.GetProfileSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "GetProfile", resp, "Failure sending request")
			return
		}

		result, err = client.GetProfileResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "GetProfile", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// GetProfilePreparer prepares the GetProfile request.
func (client GatewayClient) GetProfilePreparer(resourceGroupName string, gatewayName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"gatewayName":       autorest.Encode("path", gatewayName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServerManagement/gateways/{gatewayName}/profile", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// GetProfileSender sends the GetProfile request. The method will close the
// http.Response Body if it receives an error.
func (client GatewayClient) GetProfileSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// GetProfileResponder handles the response to the GetProfile request. The method always
// closes the http.Response Body.
func (client GatewayClient) GetProfileResponder(resp *http.Response) (result GatewayProfile, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List returns gateways in a subscription.
func (client GatewayClient) List() (result GatewayResources, err error) {
	req, err := client.ListPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client GatewayClient) ListPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ServerManagement/gateways", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client GatewayClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client GatewayClient) ListResponder(resp *http.Response) (result GatewayResources, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client GatewayClient) ListNextResults(lastResults GatewayResources) (result GatewayResources, err error) {
	req, err := lastResults.GatewayResourcesPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client GatewayClient) ListComplete(cancel <-chan struct{}) (<-chan GatewayResource, <-chan error) {
	resultChan := make(chan GatewayResource)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List()
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}

// ListForResourceGroup returns gateways in a resource group.
//
// resourceGroupName is the resource group name uniquely identifies the resource group within the user subscriptionId.
func (client GatewayClient) ListForResourceGroup(resourceGroupName string) (result GatewayResources, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9]+`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "servermanagement.GatewayClient", "ListForResourceGroup")
	}

	req, err := client.ListForResourceGroupPreparer(resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "ListForResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListForResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "ListForResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListForResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "ListForResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListForResourceGroupPreparer prepares the ListForResourceGroup request.
func (client GatewayClient) ListForResourceGroupPreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServerManagement/gateways", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListForResourceGroupSender sends the ListForResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client GatewayClient) ListForResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListForResourceGroupResponder handles the response to the ListForResourceGroup request. The method always
// closes the http.Response Body.
func (client GatewayClient) ListForResourceGroupResponder(resp *http.Response) (result GatewayResources, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListForResourceGroupNextResults retrieves the next set of results, if any.
func (client GatewayClient) ListForResourceGroupNextResults(lastResults GatewayResources) (result GatewayResources, err error) {
	req, err := lastResults.GatewayResourcesPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "ListForResourceGroup", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListForResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "ListForResourceGroup", resp, "Failure sending next results request")
	}

	result, err = client.ListForResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "ListForResourceGroup", resp, "Failure responding to next results request")
	}

	return
}

// ListForResourceGroupComplete gets all elements from the list without paging.
func (client GatewayClient) ListForResourceGroupComplete(resourceGroupName string, cancel <-chan struct{}) (<-chan GatewayResource, <-chan error) {
	resultChan := make(chan GatewayResource)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListForResourceGroup(resourceGroupName)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListForResourceGroupNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}

// RegenerateProfile regenerate a gateway's profile This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the resource group name uniquely identifies the resource group within the user subscriptionId.
// gatewayName is the gateway name (256 characters maximum).
func (client GatewayClient) RegenerateProfile(resourceGroupName string, gatewayName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9]+`, Chain: nil}}},
		{TargetValue: gatewayName,
			Constraints: []validation.Constraint{{Target: "gatewayName", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "gatewayName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "gatewayName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "servermanagement.GatewayClient", "RegenerateProfile")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result autorest.Response
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.RegenerateProfilePreparer(resourceGroupName, gatewayName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "RegenerateProfile", nil, "Failure preparing request")
			return
		}

		resp, err := client.RegenerateProfileSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "RegenerateProfile", resp, "Failure sending request")
			return
		}

		result, err = client.RegenerateProfileResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "RegenerateProfile", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// RegenerateProfilePreparer prepares the RegenerateProfile request.
func (client GatewayClient) RegenerateProfilePreparer(resourceGroupName string, gatewayName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"gatewayName":       autorest.Encode("path", gatewayName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServerManagement/gateways/{gatewayName}/regenerateprofile", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// RegenerateProfileSender sends the RegenerateProfile request. The method will close the
// http.Response Body if it receives an error.
func (client GatewayClient) RegenerateProfileSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// RegenerateProfileResponder handles the response to the RegenerateProfile request. The method always
// closes the http.Response Body.
func (client GatewayClient) RegenerateProfileResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update updates a gateway belonging to a resource group. This method may poll for completion. Polling can be canceled
// by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// resourceGroupName is the resource group name uniquely identifies the resource group within the user subscriptionId.
// gatewayName is the gateway name (256 characters maximum). gatewayParameters is parameters supplied to the Update
// operation.
func (client GatewayClient) Update(resourceGroupName string, gatewayName string, gatewayParameters GatewayParameters, cancel <-chan struct{}) (<-chan GatewayResource, <-chan error) {
	resultChan := make(chan GatewayResource, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9]+`, Chain: nil}}},
		{TargetValue: gatewayName,
			Constraints: []validation.Constraint{{Target: "gatewayName", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "gatewayName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "gatewayName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "servermanagement.GatewayClient", "Update")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result GatewayResource
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.UpdatePreparer(resourceGroupName, gatewayName, gatewayParameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "Update", nil, "Failure preparing request")
			return
		}

		resp, err := client.UpdateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "Update", resp, "Failure sending request")
			return
		}

		result, err = client.UpdateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "Update", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// UpdatePreparer prepares the Update request.
func (client GatewayClient) UpdatePreparer(resourceGroupName string, gatewayName string, gatewayParameters GatewayParameters, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"gatewayName":       autorest.Encode("path", gatewayName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServerManagement/gateways/{gatewayName}", pathParameters),
		autorest.WithJSON(gatewayParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client GatewayClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client GatewayClient) UpdateResponder(resp *http.Response) (result GatewayResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Upgrade upgrades a gateway. This method may poll for completion. Polling can be canceled by passing the cancel
// channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the resource group name uniquely identifies the resource group within the user subscriptionId.
// gatewayName is the gateway name (256 characters maximum).
func (client GatewayClient) Upgrade(resourceGroupName string, gatewayName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9]+`, Chain: nil}}},
		{TargetValue: gatewayName,
			Constraints: []validation.Constraint{{Target: "gatewayName", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "gatewayName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "gatewayName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "servermanagement.GatewayClient", "Upgrade")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result autorest.Response
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.UpgradePreparer(resourceGroupName, gatewayName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "Upgrade", nil, "Failure preparing request")
			return
		}

		resp, err := client.UpgradeSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "Upgrade", resp, "Failure sending request")
			return
		}

		result, err = client.UpgradeResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "servermanagement.GatewayClient", "Upgrade", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// UpgradePreparer prepares the Upgrade request.
func (client GatewayClient) UpgradePreparer(resourceGroupName string, gatewayName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"gatewayName":       autorest.Encode("path", gatewayName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServerManagement/gateways/{gatewayName}/upgradetolatest", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// UpgradeSender sends the Upgrade request. The method will close the
// http.Response Body if it receives an error.
func (client GatewayClient) UpgradeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// UpgradeResponder handles the response to the Upgrade request. The method always
// closes the http.Response Body.
func (client GatewayClient) UpgradeResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
