// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150501

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupWebTest(ctx *pulumi.Context, args *LookupWebTestArgs, opts ...pulumi.InvokeOption) (*LookupWebTestResult, error) {
	var rv LookupWebTestResult
	err := ctx.Invoke("azurerm:insights/v20150501:getWebTest", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebTestArgs struct {
	// The name of the Application Insights webtest resource.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An Application Insights web test definition.
type LookupWebTestResult struct {
	// The kind of web test that this web test watches. Choices are ping and multistep.
	Kind *string `pulumi:"kind"`
	// Resource location
	Location string `pulumi:"location"`
	// Azure resource name
	Name string `pulumi:"name"`
	// Metadata describing a web test for an Azure resource.
	Properties WebTestPropertiesResponse `pulumi:"properties"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Azure resource type
	Type string `pulumi:"type"`
}