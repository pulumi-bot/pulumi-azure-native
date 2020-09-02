// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupEventHubAuthorizationRule(ctx *pulumi.Context, args *LookupEventHubAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupEventHubAuthorizationRuleResult, error) {
	var rv LookupEventHubAuthorizationRuleResult
	err := ctx.Invoke("azurerm:eventhub/latest:getEventHubAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventHubAuthorizationRuleArgs struct {
	// The authorization rule name.
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	// The Event Hub name
	EventHubName string `pulumi:"eventHubName"`
	// The Namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Single item in a List or Get AuthorizationRule operation
type LookupEventHubAuthorizationRuleResult struct {
	// Resource name.
	Name string `pulumi:"name"`
	// The rights associated with the rule.
	Rights []string `pulumi:"rights"`
	// Resource type.
	Type string `pulumi:"type"`
}