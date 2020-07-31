// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20140901

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupQueue(ctx *pulumi.Context, args *LookupQueueArgs, opts ...pulumi.InvokeOption) (*LookupQueueResult, error) {
	var rv LookupQueueResult
	err := ctx.Invoke("azurerm:servicebus/v20140901:getQueue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupQueueArgs struct {
	// The queue name.
	Name string `pulumi:"name"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Description of queue Resource.
type LookupQueueResult struct {
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// The Queue Properties definition.
	Properties QueuePropertiesResponse `pulumi:"properties"`
	// Resource type
	Type string `pulumi:"type"`
}