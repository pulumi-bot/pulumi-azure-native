// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20201201preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	var rv LookupAccountResult
	err := ctx.Invoke("azure-nextgen:purview/v20201201preview:getAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccountArgs struct {
	// The name of the account.
	AccountName string `pulumi:"accountName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Account resource
type LookupAccountResult struct {
	// Cloud connectors.
	// External cloud identifier used as part of scanning configuration.
	CloudConnectors *CloudConnectorsResponse `pulumi:"cloudConnectors"`
	// Gets the time at which the entity was created.
	CreatedAt string `pulumi:"createdAt"`
	// Gets the creator of the entity.
	CreatedBy string `pulumi:"createdBy"`
	// Gets the creators of the entity's object id.
	CreatedByObjectId string `pulumi:"createdByObjectId"`
	// The URIs that are the public endpoints of the account.
	Endpoints AccountPropertiesResponseEndpoints `pulumi:"endpoints"`
	// Gets or sets the friendly name.
	FriendlyName string `pulumi:"friendlyName"`
	// Gets or sets the identifier.
	Id string `pulumi:"id"`
	// Identity Info on the tracked resource
	Identity *IdentityResponse `pulumi:"identity"`
	// Gets or sets the location.
	Location *string `pulumi:"location"`
	// Gets the resource identifiers of the managed resources.
	ManagedResources AccountPropertiesResponseManagedResources `pulumi:"managedResources"`
	// Gets or sets the name.
	Name string `pulumi:"name"`
	// Get the network ACLs.
	NetworkAcls *NetworkAclsResponse `pulumi:"networkAcls"`
	// Gets the private endpoint connections information.
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// Gets or sets the state of the provisioning.
	ProvisioningState string `pulumi:"provisioningState"`
	// Gets or sets the Sku.
	Sku *AccountSkuResponse `pulumi:"sku"`
	// Tags on the azure resource.
	Tags map[string]string `pulumi:"tags"`
	// Gets or sets the type.
	Type string `pulumi:"type"`
}
