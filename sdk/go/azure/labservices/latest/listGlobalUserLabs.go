// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Lists the labs owned by a user
// Latest API Version: 2018-10-15.
//
// Deprecated: The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:labservices:listGlobalUserLabs'.
func ListGlobalUserLabs(ctx *pulumi.Context, args *ListGlobalUserLabsArgs, opts ...pulumi.InvokeOption) (*ListGlobalUserLabsResult, error) {
	var rv ListGlobalUserLabsResult
	err := ctx.Invoke("azure-native:labservices/latest:listGlobalUserLabs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListGlobalUserLabsArgs struct {
	// The name of the user.
	UserName string `pulumi:"userName"`
}

// Lists the labs owned by a user
type ListGlobalUserLabsResult struct {
	// List of all the labs
	Labs []LabDetailsResponse `pulumi:"labs"`
}
