// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicebus

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupNamespaceTopicSubscriptionRule(ctx *pulumi.Context, args *LookupNamespaceTopicSubscriptionRuleArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceTopicSubscriptionRuleResult, error) {
	var rv LookupNamespaceTopicSubscriptionRuleResult
	err := ctx.Invoke("azurerm:servicebus:getNamespaceTopicSubscriptionRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNamespaceTopicSubscriptionRuleArgs struct {
	// The rule name.
	Name string `pulumi:"name"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The subscription name.
	SubscriptionName string `pulumi:"subscriptionName"`
	// The topic name.
	TopicName string `pulumi:"topicName"`
}

// Description of Rule Resource.
type LookupNamespaceTopicSubscriptionRuleResult struct {
	// Resource name
	Name string `pulumi:"name"`
	// Properties of Rule resource
	Properties RulepropertiesResponse `pulumi:"properties"`
	// Resource type
	Type string `pulumi:"type"`
}