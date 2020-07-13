// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicebus

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupNamespaceQueueAuthorizationRule(ctx *pulumi.Context, args *LookupNamespaceQueueAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceQueueAuthorizationRuleResult, error) {
	var rv LookupNamespaceQueueAuthorizationRuleResult
	err := ctx.Invoke("azurerm:servicebus:getNamespaceQueueAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNamespaceQueueAuthorizationRuleArgs struct {
	// The authorization rule name.
	Name string `pulumi:"name"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// The queue name.
	QueueName string `pulumi:"queueName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Description of a namespace authorization rule.
type LookupNamespaceQueueAuthorizationRuleResult struct {
	// Resource name
	Name string `pulumi:"name"`
	// AuthorizationRule properties.
	Properties SBAuthorizationRuleResponseProperties `pulumi:"properties"`
	// Resource type
	Type string `pulumi:"type"`
}