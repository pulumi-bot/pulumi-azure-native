// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180201

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListWebAppHybridConnectionKeys(ctx *pulumi.Context, args *ListWebAppHybridConnectionKeysArgs, opts ...pulumi.InvokeOption) (*ListWebAppHybridConnectionKeysResult, error) {
	var rv ListWebAppHybridConnectionKeysResult
	err := ctx.Invoke("azurerm:web/v20180201:listWebAppHybridConnectionKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppHybridConnectionKeysArgs struct {
	// The relay name for this hybrid connection.
	Name string `pulumi:"name"`
	// The namespace for this hybrid connection.
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Hybrid Connection key contract. This has the send key name and value for a Hybrid Connection.
type ListWebAppHybridConnectionKeysResult struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// HybridConnectionKey resource specific properties
	Properties HybridConnectionKeyResponseProperties `pulumi:"properties"`
	// Resource type.
	Type string `pulumi:"type"`
}