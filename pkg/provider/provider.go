// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pulumi/pulumi/pkg/resource"
	"github.com/pulumi/pulumi/pkg/resource/plugin"
	"github.com/pulumi/pulumi/pkg/resource/provider"
	rpc "github.com/pulumi/pulumi/sdk/proto/go"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/swag"
	"github.com/golang/glog"
	pbempty "github.com/golang/protobuf/ptypes/empty"
)

const (
	// DefaultBaseURI is the default URI used for the service
	DefaultBaseURI = "https://management.azure.com"

	// DefaultSubscriptionID fixes a subscriptionID for testing
	DefaultSubscriptionID = "0282681f-7a9e-424b-80b2-96babd57a8a1"
	// DefaultResourceGroupName fixes a resoruce group for testing
	DefaultResourceGroupName = "azuretest"
)

type azurermProvider struct {
	host                     *provider.HostClient
	name                     string
	version                  string
	client                   autorest.Client
	containerInstanceSwagger *spec.Swagger
}

func makeProvider(host *provider.HostClient, name, version string) (rpc.ResourceProviderServer, error) {
	// Creating a REST client
	client := autorest.NewClientWithUserAgent("pulumi")
	authorizer, err := auth.NewAuthorizerFromCLI()
	if err != nil {
		return nil, err
	}
	client.Authorizer = authorizer

	// Get and store the swagger specs
	byts, err := swag.LoadFromFileOrHTTP("https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2018-10-01/containerInstance.json")
	if err != nil {
		return nil, err
	}
	swagger := spec.Swagger{}
	err = swagger.UnmarshalJSON(byts)
	if err != nil {
		return nil, err
	}

	// Return the new provider
	return &azurermProvider{
		host:                     host,
		name:                     name,
		version:                  version,
		client:                   client,
		containerInstanceSwagger: &swagger,
	}, nil
}

// Configure configures the resource provider with "globals" that control its behavior.
func (k *azurermProvider) Configure(_ context.Context, req *rpc.ConfigureRequest) (*pbempty.Empty, error) {
	// TODO: Store configuration and initialize clients
	return &pbempty.Empty{}, nil
}

// Invoke dynamically executes a built-in function in the provider.
func (k *azurermProvider) Invoke(context.Context, *rpc.InvokeRequest) (*rpc.InvokeResponse, error) {
	panic("Invoke not implemented")
}

// Check validates that the given property bag is valid for a resource of the given type and returns
// the inputs that should be passed to successive calls to Diff, Create, or Update for this
// resource. As a rule, the provider inputs returned by a call to Check should preserve the original
// representation of the properties as present in the program inputs. Though this rule is not
// required for correctness, violations thereof can negatively impact the end-user experience, as
// the provider inputs are using for detecting and rendering diffs.
func (k *azurermProvider) Check(ctx context.Context, req *rpc.CheckRequest) (*rpc.CheckResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Check(%s)", k.name, urn)
	glog.V(9).Infof("%s executing", label)

	newResInputs := req.GetNews()
	return &rpc.CheckResponse{Inputs: newResInputs, Failures: nil}, nil
}

// Diff checks what impacts a hypothetical update will have on the resource's properties.
func (k *azurermProvider) Diff(ctx context.Context, req *rpc.DiffRequest) (*rpc.DiffResponse, error) {
	panic("Diff not implemented")
	// return &rpc.DiffResponse{
	// 	Changes:             hasChanges,
	// 	Replaces:            replaces,
	// 	Stables:             []string{},
	// 	DeleteBeforeReplace: deleteBeforeReplace,
	// }, nil
}

// Create allocates a new instance of the provided resource and returns its unique ID afterwards.
func (k *azurermProvider) Create(ctx context.Context, req *rpc.CreateRequest) (*rpc.CreateResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Create(%s)", k.name, urn)
	glog.V(9).Infof("%s executing", label)
	inputs, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.properties", label), KeepUnknowns: true, SkipNulls: true,
	})
	if err != nil {
		return nil, err
	}

	bodyProps := inputs.Mappable()
	pathParameters := map[string]interface{}{
		"containerGroupName": autorest.Encode("path", "abcd"),
		"resourceGroupName":  autorest.Encode("path", DefaultResourceGroupName),
		"subscriptionId":     autorest.Encode("path", DefaultSubscriptionID),
	}
	queryParameters := map[string]interface{}{
		"api-version": "2018-10-01",
	}
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(DefaultBaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}", pathParameters),
		autorest.WithJSON(bodyProps),
		autorest.WithQueryParameters(queryParameters))
	prepReq, err := preparer.Prepare((&http.Request{}).WithContext(ctx))
	if err != nil {
		return nil, err
	}

	var resp *http.Response
	resp, err = autorest.SendWithSender(
		k.client,
		prepReq,
		azure.DoRetryWithRegistration(k.client),
	)
	if err != nil {
		return nil, err
	}
	future, err := azure.NewFutureFromResponse(resp)
	if err != nil {
		return nil, err
	}
	err = future.WaitForCompletionRef(ctx, k.client)
	if err != nil {
		return nil, err
	}
	var outputs map[string]interface{}
	fmt.Printf("Responding...\n")
	err = autorest.Respond(
		future.Response(),
		k.client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&outputs),
		autorest.ByClosing())
	if err != nil {
		return nil, err
	}

	outputProperties, err := plugin.MarshalProperties(
		resource.NewPropertyMapFromMap(outputs),
		plugin.MarshalOptions{Label: fmt.Sprintf("%s.autonamedInputs", label), KeepUnknowns: true, SkipNulls: true},
	)
	if err != nil {
		return nil, err
	}
	id, ok := outputs["id"].(string)
	if !ok {
		return nil, fmt.Errorf("failed to receiver 'id' from newly created resource")
	}
	return &rpc.CreateResponse{
		Id:         id,
		Properties: outputProperties,
	}, nil
}

// Read the current live state associated with a resource.
func (k *azurermProvider) Read(ctx context.Context, req *rpc.ReadRequest) (*rpc.ReadResponse, error) {
	panic("Read not implemented")
	// return &rpc.ReadResponse{Id: id, Properties: inputsAndComputed}, nil
}

// Update updates an existing resource with new values.
func (k *azurermProvider) Update(ctx context.Context, req *rpc.UpdateRequest) (*rpc.UpdateResponse, error) {
	panic("Update not implemented")
	// return &rpc.UpdateResponse{Properties: inputsAndComputed}, nil
}

// Delete tears down an existing resource with the given ID.  If it fails, the resource is assumed
// to still exist.
func (k *azurermProvider) Delete(ctx context.Context, req *rpc.DeleteRequest) (*pbempty.Empty, error) {
	// TODO: Actually delete!
	return &pbempty.Empty{}, nil
}

// GetPluginInfo returns generic information about this plugin, like its version.
func (k *azurermProvider) GetPluginInfo(context.Context, *pbempty.Empty) (*rpc.PluginInfo, error) {
	return &rpc.PluginInfo{
		Version: k.version,
	}, nil
}

// Cancel signals the provider to gracefully shut down and abort any ongoing resource operations.
// Operations aborted in this way will return an error (e.g., `Update` and `Create` will either a
// creation error or an initialization error). Since Cancel is advisory and non-blocking, it is up
// to the host to decide how long to wait after Cancel is called before (e.g.)
// hard-closing any gRPC connection.
func (k *azurermProvider) Cancel(context.Context, *pbempty.Empty) (*pbempty.Empty, error) {
	// TODO
	return &pbempty.Empty{}, nil
}
