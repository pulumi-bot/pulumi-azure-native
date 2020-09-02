// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupApiOperation(ctx *pulumi.Context, args *LookupApiOperationArgs, opts ...pulumi.InvokeOption) (*LookupApiOperationResult, error) {
	var rv LookupApiOperationResult
	err := ctx.Invoke("azurerm:apimanagement/v20191201preview:getApiOperation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiOperationArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId string `pulumi:"apiId"`
	// Operation identifier within an API. Must be unique in the current API Management service instance.
	OperationId string `pulumi:"operationId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Api Operation details.
type LookupApiOperationResult struct {
	// Description of the operation. May include HTML formatting tags.
	Description *string `pulumi:"description"`
	// Operation Name.
	DisplayName string `pulumi:"displayName"`
	// A Valid HTTP Operation Method. Typical Http Methods like GET, PUT, POST but not limited by only them.
	Method string `pulumi:"method"`
	// Resource name.
	Name string `pulumi:"name"`
	// Operation Policies
	Policies *string `pulumi:"policies"`
	// An entity containing request details.
	Request *RequestContractResponse `pulumi:"request"`
	// Array of Operation responses.
	Responses []ResponseContractResponse `pulumi:"responses"`
	// Collection of URL template parameters.
	TemplateParameters []ParameterContractResponse `pulumi:"templateParameters"`
	// Resource type for API Management resource.
	Type string `pulumi:"type"`
	// Relative URL template identifying the target resource for this operation. May include parameters. Example: /customers/{cid}/orders/{oid}/?date={date}
	UrlTemplate string `pulumi:"urlTemplate"`
}
