// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170701privatepreview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupApp(ctx *pulumi.Context, args *LookupAppArgs, opts ...pulumi.InvokeOption) (*LookupAppResult, error) {
	var rv LookupAppResult
	err := ctx.Invoke("azurerm:iotcentral/v20170701privatepreview:getApp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAppArgs struct {
	// The name of the resource group that contains the IoT Central application.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The ARM resource name of the IoT Central application.
	ResourceName string `pulumi:"resourceName"`
}

// The IoT Central application.
type LookupAppResult struct {
	// The ID of the application.
	ApplicationId string `pulumi:"applicationId"`
	// The display name of the application.
	DisplayName *string `pulumi:"displayName"`
	// The resource location.
	Location string `pulumi:"location"`
	// The ARM resource name.
	Name string `pulumi:"name"`
	// A valid instance SKU.
	Sku AppSkuInfoResponse `pulumi:"sku"`
	// The subdomain of the application.
	Subdomain *string `pulumi:"subdomain"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the application template, which is a blueprint that defines the characteristics and behaviors of an application. Optional; if not specified, defaults to a blank blueprint and allows the application to be defined from scratch.
	Template *string `pulumi:"template"`
	// The resource type.
	Type string `pulumi:"type"`
}
