// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180715preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	var rv LookupServiceResult
	err := ctx.Invoke("azurerm:datamigration/v20180715preview:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceArgs struct {
	// Name of the resource group
	GroupName string `pulumi:"groupName"`
	// Name of the service
	ServiceName string `pulumi:"serviceName"`
}

// A Database Migration Service resource
type LookupServiceResult struct {
	// HTTP strong entity tag value. Ignored if submitted
	Etag *string `pulumi:"etag"`
	// The resource kind. Only 'vm' (the default) is supported.
	Kind *string `pulumi:"kind"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The resource's provisioning state
	ProvisioningState string `pulumi:"provisioningState"`
	// The public key of the service, used to encrypt secrets sent to the service
	PublicKey *string `pulumi:"publicKey"`
	// Service SKU
	Sku *ServiceSkuResponse `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// The ID of the Microsoft.Network/virtualNetworks/subnets resource to which the service should be joined
	VirtualSubnetId string `pulumi:"virtualSubnetId"`
}
