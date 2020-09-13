// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20171103preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// HANA instance info on Azure (ARM properties and HANA properties)
//
// ## Example Usage
// ### Get properties of a HANA instance
//
// ```go
// package main
//
// import (
// 	hanaonazure "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/hanaonazure/v20171103preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := hanaonazure.NewHanaInstance(ctx, "hanaInstance", &hanaonazure.HanaInstanceArgs{
// 			HanaInstanceName: pulumi.String("myHanaInstance"),
// 			Location:         pulumi.String("westus"),
// 			NetworkProfile: &hanaonazure.NetworkProfileArgs{
// 				NetworkInterfaces: hanaonazure.IpAddressArray{
// 					&hanaonazure.IpAddressArgs{
// 						IpAddress: pulumi.String("100.100.100.100"),
// 					},
// 				},
// 			},
// 			OsProfile: &hanaonazure.OSProfileArgs{
// 				ComputerName: pulumi.String("myComputerName"),
// 				SshPublicKey: pulumi.String("AAAAB3NzaC1yc2EAAAABJQAAAQB/nAmOjTmezNUDKYvEeIRf2YnwM9/uUG1d0BYsc8/tRtx+RGi7N2lUbp728MXGwdnL9od4cItzky/zVdLZE2cycOa18xBK9cOWmcKS0A8FYBxEQWJ/q9YVUgZbFKfYGaGQxsER+A0w/fX8ALuk78ktP31K69LcQgxIsl7rNzxsoOQKJ/CIxOGMMxczYTiEoLvQhapFQMs3FL96didKr/QbrfB1WT6s3838SEaXfgZvLef1YB2xmfhbT9OXFE3FXvh2UPBfN+ffE7iiayQf/2XR+8j4N4bW30DiPtOQLGUrH1y5X/rpNZNlWW2+jGIxqZtgWg7lTy3mXy5x836Sj/6L"),
// 			},
// 			PartnerNodeId:     pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.HanaOnAzure/hanaInstances/myHanaInstance2"),
// 			ResourceGroupName: pulumi.String("myResourceGroup"),
// 			StorageProfile:    nil,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type HanaInstance struct {
	pulumi.CustomResourceState

	// Specifies the HANA instance unique ID.
	HanaInstanceId pulumi.StringOutput `pulumi:"hanaInstanceId"`
	// Specifies the hardware settings for the HANA instance.
	HardwareProfile HardwareProfileResponsePtrOutput `pulumi:"hardwareProfile"`
	// Hardware revision of a HANA instance
	HwRevision pulumi.StringOutput `pulumi:"hwRevision"`
	// Resource location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the network settings for the HANA instance.
	NetworkProfile NetworkProfileResponsePtrOutput `pulumi:"networkProfile"`
	// Specifies the operating system settings for the HANA instance.
	OsProfile OSProfileResponsePtrOutput `pulumi:"osProfile"`
	// ARM ID of another HanaInstance that will share a network with this HanaInstance
	PartnerNodeId pulumi.StringPtrOutput `pulumi:"partnerNodeId"`
	// Resource power state
	PowerState pulumi.StringOutput `pulumi:"powerState"`
	// State of provisioning of the HanaInstance
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource proximity placement group
	ProximityPlacementGroup pulumi.StringOutput `pulumi:"proximityPlacementGroup"`
	// Specifies the storage settings for the HANA instance disks.
	StorageProfile StorageProfileResponsePtrOutput `pulumi:"storageProfile"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewHanaInstance registers a new resource with the given unique name, arguments, and options.
func NewHanaInstance(ctx *pulumi.Context,
	name string, args *HanaInstanceArgs, opts ...pulumi.ResourceOption) (*HanaInstance, error) {
	if args == nil || args.HanaInstanceName == nil {
		return nil, errors.New("missing required argument 'HanaInstanceName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &HanaInstanceArgs{}
	}
	var resource HanaInstance
	err := ctx.RegisterResource("azurerm:hanaonazure/v20171103preview:HanaInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHanaInstance gets an existing HanaInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHanaInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HanaInstanceState, opts ...pulumi.ResourceOption) (*HanaInstance, error) {
	var resource HanaInstance
	err := ctx.ReadResource("azurerm:hanaonazure/v20171103preview:HanaInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HanaInstance resources.
type hanaInstanceState struct {
	// Specifies the HANA instance unique ID.
	HanaInstanceId *string `pulumi:"hanaInstanceId"`
	// Specifies the hardware settings for the HANA instance.
	HardwareProfile *HardwareProfileResponse `pulumi:"hardwareProfile"`
	// Hardware revision of a HANA instance
	HwRevision *string `pulumi:"hwRevision"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Specifies the network settings for the HANA instance.
	NetworkProfile *NetworkProfileResponse `pulumi:"networkProfile"`
	// Specifies the operating system settings for the HANA instance.
	OsProfile *OSProfileResponse `pulumi:"osProfile"`
	// ARM ID of another HanaInstance that will share a network with this HanaInstance
	PartnerNodeId *string `pulumi:"partnerNodeId"`
	// Resource power state
	PowerState *string `pulumi:"powerState"`
	// State of provisioning of the HanaInstance
	ProvisioningState *string `pulumi:"provisioningState"`
	// Resource proximity placement group
	ProximityPlacementGroup *string `pulumi:"proximityPlacementGroup"`
	// Specifies the storage settings for the HANA instance disks.
	StorageProfile *StorageProfileResponse `pulumi:"storageProfile"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type HanaInstanceState struct {
	// Specifies the HANA instance unique ID.
	HanaInstanceId pulumi.StringPtrInput
	// Specifies the hardware settings for the HANA instance.
	HardwareProfile HardwareProfileResponsePtrInput
	// Hardware revision of a HANA instance
	HwRevision pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Specifies the network settings for the HANA instance.
	NetworkProfile NetworkProfileResponsePtrInput
	// Specifies the operating system settings for the HANA instance.
	OsProfile OSProfileResponsePtrInput
	// ARM ID of another HanaInstance that will share a network with this HanaInstance
	PartnerNodeId pulumi.StringPtrInput
	// Resource power state
	PowerState pulumi.StringPtrInput
	// State of provisioning of the HanaInstance
	ProvisioningState pulumi.StringPtrInput
	// Resource proximity placement group
	ProximityPlacementGroup pulumi.StringPtrInput
	// Specifies the storage settings for the HANA instance disks.
	StorageProfile StorageProfileResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (HanaInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*hanaInstanceState)(nil)).Elem()
}

type hanaInstanceArgs struct {
	// Name of the SAP HANA on Azure instance.
	HanaInstanceName string `pulumi:"hanaInstanceName"`
	// Resource location
	Location *string `pulumi:"location"`
	// Specifies the network settings for the HANA instance.
	NetworkProfile *NetworkProfile `pulumi:"networkProfile"`
	// Specifies the operating system settings for the HANA instance.
	OsProfile *OSProfile `pulumi:"osProfile"`
	// ARM ID of another HanaInstance that will share a network with this HanaInstance
	PartnerNodeId *string `pulumi:"partnerNodeId"`
	// Name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the storage settings for the HANA instance disks.
	StorageProfile *StorageProfile `pulumi:"storageProfile"`
}

// The set of arguments for constructing a HanaInstance resource.
type HanaInstanceArgs struct {
	// Name of the SAP HANA on Azure instance.
	HanaInstanceName pulumi.StringInput
	// Resource location
	Location pulumi.StringPtrInput
	// Specifies the network settings for the HANA instance.
	NetworkProfile NetworkProfilePtrInput
	// Specifies the operating system settings for the HANA instance.
	OsProfile OSProfilePtrInput
	// ARM ID of another HanaInstance that will share a network with this HanaInstance
	PartnerNodeId pulumi.StringPtrInput
	// Name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Specifies the storage settings for the HANA instance disks.
	StorageProfile StorageProfilePtrInput
}

func (HanaInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hanaInstanceArgs)(nil)).Elem()
}
