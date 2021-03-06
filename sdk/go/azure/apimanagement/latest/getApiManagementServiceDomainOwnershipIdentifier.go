// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Response of the GetDomainOwnershipIdentifier operation.
// Latest API Version: 2020-12-01.
//
// Deprecated: The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:apimanagement:getApiManagementServiceDomainOwnershipIdentifier'.
func GetApiManagementServiceDomainOwnershipIdentifier(ctx *pulumi.Context, args *GetApiManagementServiceDomainOwnershipIdentifierArgs, opts ...pulumi.InvokeOption) (*GetApiManagementServiceDomainOwnershipIdentifierResult, error) {
	var rv GetApiManagementServiceDomainOwnershipIdentifierResult
	err := ctx.Invoke("azure-native:apimanagement/latest:getApiManagementServiceDomainOwnershipIdentifier", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetApiManagementServiceDomainOwnershipIdentifierArgs struct {
}

// Response of the GetDomainOwnershipIdentifier operation.
type GetApiManagementServiceDomainOwnershipIdentifierResult struct {
	// The domain ownership identifier value.
	DomainOwnershipIdentifier string `pulumi:"domainOwnershipIdentifier"`
}
