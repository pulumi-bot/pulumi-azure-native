// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170101

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupWebService(ctx *pulumi.Context, args *LookupWebServiceArgs, opts ...pulumi.InvokeOption) (*LookupWebServiceResult, error) {
	var rv LookupWebServiceResult
	err := ctx.Invoke("azurerm:machinelearning/v20170101:getWebService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebServiceArgs struct {
	// The name of the web service.
	Name string `pulumi:"name"`
	// The region for which encrypted credential parameters are valid.
	Region *string `pulumi:"region"`
	// Name of the resource group in which the web service is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Instance of an Azure ML web service resource.
type LookupWebServiceResult struct {
	// Specifies the location of the resource.
	Location string `pulumi:"location"`
	// Specifies the name of the resource.
	Name string `pulumi:"name"`
	// Contains the property payload that describes the web service.
	Properties WebServicePropertiesResponse `pulumi:"properties"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the type of the resource.
	Type string `pulumi:"type"`
}