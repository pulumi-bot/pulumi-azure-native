// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storsimple

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupManagerStorageDomain(ctx *pulumi.Context, args *LookupManagerStorageDomainArgs, opts ...pulumi.InvokeOption) (*LookupManagerStorageDomainResult, error) {
	var rv LookupManagerStorageDomainResult
	err := ctx.Invoke("azurerm:storsimple:getManagerStorageDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagerStorageDomainArgs struct {
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The storage domain name.
	Name string `pulumi:"name"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The storage domain.
type LookupManagerStorageDomainResult struct {
	// The name.
	Name string `pulumi:"name"`
	// The properties.
	Properties StorageDomainPropertiesResponse `pulumi:"properties"`
	// The type.
	Type string `pulumi:"type"`
}