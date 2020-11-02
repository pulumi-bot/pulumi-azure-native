// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("azure-nextgen:operationalinsights/v20201001:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterArgs struct {
	// Name of the Log Analytics Cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The top level Log Analytics cluster resource container.
type LookupClusterResult struct {
	// Configures whether billing will be only on the cluster or each workspace will be billed by its proportional use. This does not change the overall billing, only how it will be distributed. Default value is 'Cluster'
	BillingType *string `pulumi:"billingType"`
	// The ID associated with the cluster.
	ClusterId string `pulumi:"clusterId"`
	// The identity of the resource.
	Identity *IdentityResponse `pulumi:"identity"`
	// Sets whether the cluster will support availability zones. This can be set as true only in regions where Azure Data Explorer support Availability Zones. This Property can not be modified after cluster creation. Default value is 'true' if region supports Availability Zones.
	IsAvailabilityZonesEnabled *bool `pulumi:"isAvailabilityZonesEnabled"`
	// Configures whether cluster will use double encryption. This Property can not be modified after cluster creation. Default value is 'true'
	IsDoubleEncryptionEnabled *bool `pulumi:"isDoubleEncryptionEnabled"`
	// The associated key properties.
	KeyVaultProperties *KeyVaultPropertiesResponse `pulumi:"keyVaultProperties"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state of the cluster.
	ProvisioningState string `pulumi:"provisioningState"`
	// The sku properties.
	Sku *ClusterSkuResponse `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}
