// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150831preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupUserAssignedIdentity(ctx *pulumi.Context, args *LookupUserAssignedIdentityArgs, opts ...pulumi.InvokeOption) (*LookupUserAssignedIdentityResult, error) {
	var rv LookupUserAssignedIdentityResult
	err := ctx.Invoke("azurerm:managedidentity/v20150831preview:getUserAssignedIdentity", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserAssignedIdentityArgs struct {
	// The name of the Resource Group to which the identity belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the identity resource.
	ResourceName string `pulumi:"resourceName"`
}

// Describes an identity resource.
type LookupUserAssignedIdentityResult struct {
	// The id of the app associated with the identity. This is a random generated UUID by MSI.
	ClientId string `pulumi:"clientId"`
	//  The ManagedServiceIdentity DataPlane URL that can be queried to obtain the identity credentials.
	ClientSecretUrl string `pulumi:"clientSecretUrl"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The id of the service principal object associated with the created identity.
	PrincipalId string `pulumi:"principalId"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The id of the tenant which the identity belongs to.
	TenantId string `pulumi:"tenantId"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
}
