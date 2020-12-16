// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20201005privatepreview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupFirewallRule(ctx *pulumi.Context, args *LookupFirewallRuleArgs, opts ...pulumi.InvokeOption) (*LookupFirewallRuleResult, error) {
	var rv LookupFirewallRuleResult
	err := ctx.Invoke("azure-nextgen:dbforpostgresql/v20201005privatepreview:getFirewallRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFirewallRuleArgs struct {
	// The name of the server group firewall rule.
	FirewallRuleName string `pulumi:"firewallRuleName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server group.
	ServerGroupName string `pulumi:"serverGroupName"`
}

// Represents a server group firewall rule.
type LookupFirewallRuleResult struct {
	// The end IP address of the server group firewall rule. Must be IPv4 format.
	EndIpAddress string `pulumi:"endIpAddress"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The start IP address of the server group firewall rule. Must be IPv4 format.
	StartIpAddress string `pulumi:"startIpAddress"`
	// The system metadata relating to this resource
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}
