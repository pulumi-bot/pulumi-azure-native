// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupTagAtScope(ctx *pulumi.Context, args *LookupTagAtScopeArgs, opts ...pulumi.InvokeOption) (*LookupTagAtScopeResult, error) {
	var rv LookupTagAtScopeResult
	err := ctx.Invoke("azurerm:resources/v20200601:getTagAtScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTagAtScopeArgs struct {
	// The resource scope.
	Scope string `pulumi:"scope"`
}

// Wrapper resource for tags API requests and responses.
type LookupTagAtScopeResult struct {
	// The name of the tags wrapper resource.
	Name string `pulumi:"name"`
	// The set of tags.
	Properties TagsResponse `pulumi:"properties"`
	// The type of the tags wrapper resource.
	Type string `pulumi:"type"`
}
