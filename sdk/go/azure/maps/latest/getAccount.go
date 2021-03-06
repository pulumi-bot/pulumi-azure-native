// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An Azure resource which represents access to a suite of Maps REST APIs.
// Latest API Version: 2018-05-01.
//
// Deprecated: The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:maps:getAccount'.
func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	var rv LookupAccountResult
	err := ctx.Invoke("azure-native:maps/latest:getAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccountArgs struct {
	// The name of the Maps Account.
	AccountName string `pulumi:"accountName"`
	// The name of the Azure Resource Group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An Azure resource which represents access to a suite of Maps REST APIs.
type LookupAccountResult struct {
	// The fully qualified Maps Account resource identifier.
	Id string `pulumi:"id"`
	// The location of the resource.
	Location string `pulumi:"location"`
	// The name of the Maps Account, which is unique within a Resource Group.
	Name string `pulumi:"name"`
	// The map account properties.
	Properties MapsAccountPropertiesResponse `pulumi:"properties"`
	// The SKU of this account.
	Sku SkuResponse `pulumi:"sku"`
	// Gets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags map[string]string `pulumi:"tags"`
	// Azure resource type.
	Type string `pulumi:"type"`
}
