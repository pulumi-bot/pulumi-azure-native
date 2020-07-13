// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupServiceApiOperation(ctx *pulumi.Context, args *LookupServiceApiOperationArgs, opts ...pulumi.InvokeOption) (*LookupServiceApiOperationResult, error) {
	var rv LookupServiceApiOperationResult
	err := ctx.Invoke("azurerm:apimanagement:getServiceApiOperation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceApiOperationArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId string `pulumi:"apiId"`
	// Operation identifier within an API. Must be unique in the current API Management service instance.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Api Operation details.
type LookupServiceApiOperationResult struct {
	// Resource name.
	Name string `pulumi:"name"`
	// Properties of the Operation Contract.
	Properties OperationContractPropertiesResponse `pulumi:"properties"`
	// Resource type for API Management resource.
	Type string `pulumi:"type"`
}