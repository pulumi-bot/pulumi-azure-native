// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160627preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupRegistry(ctx *pulumi.Context, args *LookupRegistryArgs, opts ...pulumi.InvokeOption) (*LookupRegistryResult, error) {
	var rv LookupRegistryResult
	err := ctx.Invoke("azurerm:containerregistry/v20160627preview:getRegistry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegistryArgs struct {
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group to which the container registry belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An object that represents a container registry.
type LookupRegistryResult struct {
	// The value that indicates whether the admin user is enabled. This value is false by default.
	AdminUserEnabled *bool `pulumi:"adminUserEnabled"`
	// The creation date of the container registry in ISO8601 format.
	CreationDate string `pulumi:"creationDate"`
	// The location of the resource. This cannot be changed after the resource is created.
	Location string `pulumi:"location"`
	// The URL that can be used to log into the container registry.
	LoginServer string `pulumi:"loginServer"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The properties of the storage account for the container registry. If specified, the storage account must be in the same physical location as the container registry.
	StorageAccount StorageAccountPropertiesResponse `pulumi:"storageAccount"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
}
