// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package media

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupMediaServiceTransform(ctx *pulumi.Context, args *LookupMediaServiceTransformArgs, opts ...pulumi.InvokeOption) (*LookupMediaServiceTransformResult, error) {
	var rv LookupMediaServiceTransformResult
	err := ctx.Invoke("azurerm:media:getMediaServiceTransform", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMediaServiceTransformArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The Transform name.
	Name string `pulumi:"name"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A Transform encapsulates the rules or instructions for generating desired outputs from input media, such as by transcoding or by extracting insights. After the Transform is created, it can be applied to input media by creating Jobs.
type LookupMediaServiceTransformResult struct {
	// The name of the resource
	Name string `pulumi:"name"`
	// The resource properties.
	Properties TransformPropertiesResponse `pulumi:"properties"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
}