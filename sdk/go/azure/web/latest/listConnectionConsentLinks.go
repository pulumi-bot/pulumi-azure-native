// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Collection of consent links
// Latest API Version: 2016-06-01.
//
// Deprecated: The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:web:listConnectionConsentLinks'.
func ListConnectionConsentLinks(ctx *pulumi.Context, args *ListConnectionConsentLinksArgs, opts ...pulumi.InvokeOption) (*ListConnectionConsentLinksResult, error) {
	var rv ListConnectionConsentLinksResult
	err := ctx.Invoke("azure-native:web/latest:listConnectionConsentLinks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListConnectionConsentLinksArgs struct {
	// Connection name
	ConnectionName string `pulumi:"connectionName"`
	// Collection of resources
	Parameters []ConsentLinkParameterDefinition `pulumi:"parameters"`
	// The resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Subscription Id
	SubscriptionId *string `pulumi:"subscriptionId"`
}

// Collection of consent links
type ListConnectionConsentLinksResult struct {
	// Collection of resources
	Value []ConsentLinkDefinitionResponse `pulumi:"value"`
}
