// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupHostingEnvironment(ctx *pulumi.Context, args *LookupHostingEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupHostingEnvironmentResult, error) {
	var rv LookupHostingEnvironmentResult
	err := ctx.Invoke("azurerm:web:getHostingEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHostingEnvironmentArgs struct {
	// Name of the App Service Environment.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// App Service Environment ARM resource.
type LookupHostingEnvironmentResult struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Location.
	Location string `pulumi:"location"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Core resource properties
	Properties AppServiceEnvironmentResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}