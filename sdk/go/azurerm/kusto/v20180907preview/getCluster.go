// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180907preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("azurerm:kusto/v20180907preview:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterArgs struct {
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Class representing a Kusto cluster.
type LookupClusterResult struct {
	// The cluster data ingestion URI.
	DataIngestionUri string `pulumi:"dataIngestionUri"`
	// An ETag of the resource created.
	Etag string `pulumi:"etag"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioned state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The SKU of the cluster.
	Sku AzureSkuResponse `pulumi:"sku"`
	// The state of the resource.
	State string `pulumi:"state"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The cluster's external tenants.
	TrustedExternalTenants []TrustedExternalTenantResponse `pulumi:"trustedExternalTenants"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
	// The cluster URI.
	Uri string `pulumi:"uri"`
}
