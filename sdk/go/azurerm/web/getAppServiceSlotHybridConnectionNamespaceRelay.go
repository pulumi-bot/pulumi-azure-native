// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupAppServiceSlotHybridConnectionNamespaceRelay(ctx *pulumi.Context, args *LookupAppServiceSlotHybridConnectionNamespaceRelayArgs, opts ...pulumi.InvokeOption) (*LookupAppServiceSlotHybridConnectionNamespaceRelayResult, error) {
	var rv LookupAppServiceSlotHybridConnectionNamespaceRelayResult
	err := ctx.Invoke("azurerm:web:getAppServiceSlotHybridConnectionNamespaceRelay", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAppServiceSlotHybridConnectionNamespaceRelayArgs struct {
	// The relay name for this hybrid connection.
	Name string `pulumi:"name"`
	// The namespace for this hybrid connection.
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the slot for the web app.
	Slot string `pulumi:"slot"`
}

// Hybrid Connection contract. This is used to configure a Hybrid Connection.
type LookupAppServiceSlotHybridConnectionNamespaceRelayResult struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// HybridConnection resource specific properties
	Properties HybridConnectionResponseProperties `pulumi:"properties"`
	// Resource type.
	Type string `pulumi:"type"`
}
