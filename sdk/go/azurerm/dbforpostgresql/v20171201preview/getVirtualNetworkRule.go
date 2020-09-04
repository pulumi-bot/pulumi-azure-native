// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20171201preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupVirtualNetworkRule(ctx *pulumi.Context, args *LookupVirtualNetworkRuleArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkRuleResult, error) {
	var rv LookupVirtualNetworkRuleResult
	err := ctx.Invoke("azurerm:dbforpostgresql/v20171201preview:getVirtualNetworkRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualNetworkRuleArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The name of the virtual network rule.
	VirtualNetworkRuleName string `pulumi:"virtualNetworkRuleName"`
}

// A virtual network rule.
type LookupVirtualNetworkRuleResult struct {
	// Create firewall rule before the virtual network has vnet service endpoint enabled.
	IgnoreMissingVnetServiceEndpoint *bool `pulumi:"ignoreMissingVnetServiceEndpoint"`
	// Resource name.
	Name string `pulumi:"name"`
	// Virtual Network Rule State
	State string `pulumi:"state"`
	// Resource type.
	Type string `pulumi:"type"`
	// The ARM resource id of the virtual network subnet.
	VirtualNetworkSubnetId string `pulumi:"virtualNetworkSubnetId"`
}
