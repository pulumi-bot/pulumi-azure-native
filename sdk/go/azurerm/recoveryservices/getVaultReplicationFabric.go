// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package recoveryservices

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupVaultReplicationFabric(ctx *pulumi.Context, args *LookupVaultReplicationFabricArgs, opts ...pulumi.InvokeOption) (*LookupVaultReplicationFabricResult, error) {
	var rv LookupVaultReplicationFabricResult
	err := ctx.Invoke("azurerm:recoveryservices:getVaultReplicationFabric", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVaultReplicationFabricArgs struct {
	// Fabric name.
	Name string `pulumi:"name"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName string `pulumi:"resourceName"`
}

// Fabric definition.
type LookupVaultReplicationFabricResult struct {
	// Resource Location
	Location *string `pulumi:"location"`
	// Resource Name
	Name string `pulumi:"name"`
	// Fabric related data.
	Properties FabricPropertiesResponse `pulumi:"properties"`
	// Resource Type
	Type string `pulumi:"type"`
}