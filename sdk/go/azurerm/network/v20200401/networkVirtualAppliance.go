// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// NetworkVirtualAppliance Resource.
//
// ## Example Usage
// ### Create NetworkVirtualAppliance
//
// ```go
// package main
//
// import (
// 	network "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/network/v20200401"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := network.NewNetworkVirtualAppliance(ctx, "networkVirtualAppliance", &network.NetworkVirtualApplianceArgs{
// 			BootStrapConfigurationBlob: pulumi.StringArray{
// 				pulumi.String("https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrbootstrapconfig"),
// 			},
// 			CloudInitConfigurationBlob: pulumi.StringArray{
// 				pulumi.String("https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrcloudinitconfig"),
// 			},
// 			Identity: &network.ManagedServiceIdentityArgs{
// 				Type: pulumi.String("UserAssigned"),
// 			},
// 			Location:                    pulumi.String("West US"),
// 			NetworkVirtualApplianceName: pulumi.String("nva"),
// 			ResourceGroupName:           pulumi.String("rg1"),
// 			Sku: &network.VirtualApplianceSkuPropertiesArgs{
// 				BundledScaleUnit:   pulumi.String("1"),
// 				MarketPlaceVersion: pulumi.String("12.1"),
// 				Vendor:             pulumi.String("Cisco SDWAN"),
// 			},
// 			Tags: pulumi.StringMap{
// 				"key1": pulumi.String("value1"),
// 			},
// 			VirtualApplianceAsn: pulumi.Int(10000),
// 			VirtualHub: &network.SubResourceArgs{
// 				Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type NetworkVirtualAppliance struct {
	pulumi.CustomResourceState

	// BootStrapConfigurationBlob storage URLs.
	BootStrapConfigurationBlob pulumi.StringArrayOutput `pulumi:"bootStrapConfigurationBlob"`
	// CloudInitConfigurationBlob storage URLs.
	CloudInitConfigurationBlob pulumi.StringArrayOutput `pulumi:"cloudInitConfigurationBlob"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The service principal that has read access to cloud-init and config blob.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Network Virtual Appliance SKU.
	Sku VirtualApplianceSkuPropertiesResponsePtrOutput `pulumi:"sku"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// VirtualAppliance ASN.
	VirtualApplianceAsn pulumi.IntPtrOutput `pulumi:"virtualApplianceAsn"`
	// List of Virtual Appliance Network Interfaces.
	VirtualApplianceNics VirtualApplianceNicPropertiesResponseArrayOutput `pulumi:"virtualApplianceNics"`
	// The Virtual Hub where Network Virtual Appliance is being deployed.
	VirtualHub SubResourceResponsePtrOutput `pulumi:"virtualHub"`
}

// NewNetworkVirtualAppliance registers a new resource with the given unique name, arguments, and options.
func NewNetworkVirtualAppliance(ctx *pulumi.Context,
	name string, args *NetworkVirtualApplianceArgs, opts ...pulumi.ResourceOption) (*NetworkVirtualAppliance, error) {
	if args == nil || args.NetworkVirtualApplianceName == nil {
		return nil, errors.New("missing required argument 'NetworkVirtualApplianceName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &NetworkVirtualApplianceArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:network/latest:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191201:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200301:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200501:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200601:NetworkVirtualAppliance"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkVirtualAppliance
	err := ctx.RegisterResource("azurerm:network/v20200401:NetworkVirtualAppliance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkVirtualAppliance gets an existing NetworkVirtualAppliance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkVirtualAppliance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkVirtualApplianceState, opts ...pulumi.ResourceOption) (*NetworkVirtualAppliance, error) {
	var resource NetworkVirtualAppliance
	err := ctx.ReadResource("azurerm:network/v20200401:NetworkVirtualAppliance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkVirtualAppliance resources.
type networkVirtualApplianceState struct {
	// BootStrapConfigurationBlob storage URLs.
	BootStrapConfigurationBlob []string `pulumi:"bootStrapConfigurationBlob"`
	// CloudInitConfigurationBlob storage URLs.
	CloudInitConfigurationBlob []string `pulumi:"cloudInitConfigurationBlob"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// The service principal that has read access to cloud-init and config blob.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Network Virtual Appliance SKU.
	Sku *VirtualApplianceSkuPropertiesResponse `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
	// VirtualAppliance ASN.
	VirtualApplianceAsn *int `pulumi:"virtualApplianceAsn"`
	// List of Virtual Appliance Network Interfaces.
	VirtualApplianceNics []VirtualApplianceNicPropertiesResponse `pulumi:"virtualApplianceNics"`
	// The Virtual Hub where Network Virtual Appliance is being deployed.
	VirtualHub *SubResourceResponse `pulumi:"virtualHub"`
}

type NetworkVirtualApplianceState struct {
	// BootStrapConfigurationBlob storage URLs.
	BootStrapConfigurationBlob pulumi.StringArrayInput
	// CloudInitConfigurationBlob storage URLs.
	CloudInitConfigurationBlob pulumi.StringArrayInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// The service principal that has read access to cloud-init and config blob.
	Identity ManagedServiceIdentityResponsePtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringPtrInput
	// Network Virtual Appliance SKU.
	Sku VirtualApplianceSkuPropertiesResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
	// VirtualAppliance ASN.
	VirtualApplianceAsn pulumi.IntPtrInput
	// List of Virtual Appliance Network Interfaces.
	VirtualApplianceNics VirtualApplianceNicPropertiesResponseArrayInput
	// The Virtual Hub where Network Virtual Appliance is being deployed.
	VirtualHub SubResourceResponsePtrInput
}

func (NetworkVirtualApplianceState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkVirtualApplianceState)(nil)).Elem()
}

type networkVirtualApplianceArgs struct {
	// BootStrapConfigurationBlob storage URLs.
	BootStrapConfigurationBlob []string `pulumi:"bootStrapConfigurationBlob"`
	// CloudInitConfigurationBlob storage URLs.
	CloudInitConfigurationBlob []string `pulumi:"cloudInitConfigurationBlob"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The service principal that has read access to cloud-init and config blob.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of Network Virtual Appliance.
	NetworkVirtualApplianceName string `pulumi:"networkVirtualApplianceName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Network Virtual Appliance SKU.
	Sku *VirtualApplianceSkuProperties `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// VirtualAppliance ASN.
	VirtualApplianceAsn *int `pulumi:"virtualApplianceAsn"`
	// The Virtual Hub where Network Virtual Appliance is being deployed.
	VirtualHub *SubResource `pulumi:"virtualHub"`
}

// The set of arguments for constructing a NetworkVirtualAppliance resource.
type NetworkVirtualApplianceArgs struct {
	// BootStrapConfigurationBlob storage URLs.
	BootStrapConfigurationBlob pulumi.StringArrayInput
	// CloudInitConfigurationBlob storage URLs.
	CloudInitConfigurationBlob pulumi.StringArrayInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The service principal that has read access to cloud-init and config blob.
	Identity ManagedServiceIdentityPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of Network Virtual Appliance.
	NetworkVirtualApplianceName pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Network Virtual Appliance SKU.
	Sku VirtualApplianceSkuPropertiesPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// VirtualAppliance ASN.
	VirtualApplianceAsn pulumi.IntPtrInput
	// The Virtual Hub where Network Virtual Appliance is being deployed.
	VirtualHub SubResourcePtrInput
}

func (NetworkVirtualApplianceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkVirtualApplianceArgs)(nil)).Elem()
}
