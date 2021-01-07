// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// NetworkVirtualAppliance Resource.
type NetworkVirtualAppliance struct {
	pulumi.CustomResourceState

	// Address Prefix.
	AddressPrefix pulumi.StringOutput `pulumi:"addressPrefix"`
	// BootStrapConfigurationBlobs storage URLs.
	BootStrapConfigurationBlobs pulumi.StringArrayOutput `pulumi:"bootStrapConfigurationBlobs"`
	// CloudInitConfiguration string in plain text.
	CloudInitConfiguration pulumi.StringPtrOutput `pulumi:"cloudInitConfiguration"`
	// CloudInitConfigurationBlob storage URLs.
	CloudInitConfigurationBlobs pulumi.StringArrayOutput `pulumi:"cloudInitConfigurationBlobs"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The service principal that has read access to cloud-init and config blob.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// List of references to InboundSecurityRules.
	InboundSecurityRules SubResourceResponseArrayOutput `pulumi:"inboundSecurityRules"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Network Virtual Appliance SKU.
	NvaSku VirtualApplianceSkuPropertiesResponsePtrOutput `pulumi:"nvaSku"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// VirtualAppliance ASN.
	VirtualApplianceAsn pulumi.Float64PtrOutput `pulumi:"virtualApplianceAsn"`
	// List of Virtual Appliance Network Interfaces.
	VirtualApplianceNics VirtualApplianceNicPropertiesResponseArrayOutput `pulumi:"virtualApplianceNics"`
	// List of references to VirtualApplianceSite.
	VirtualApplianceSites SubResourceResponseArrayOutput `pulumi:"virtualApplianceSites"`
	// The Virtual Hub where Network Virtual Appliance is being deployed.
	VirtualHub SubResourceResponsePtrOutput `pulumi:"virtualHub"`
}

// NewNetworkVirtualAppliance registers a new resource with the given unique name, arguments, and options.
func NewNetworkVirtualAppliance(ctx *pulumi.Context,
	name string, args *NetworkVirtualApplianceArgs, opts ...pulumi.ResourceOption) (*NetworkVirtualAppliance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkVirtualApplianceName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkVirtualApplianceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:network/latest:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191201:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200301:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200401:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200501:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200701:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200801:NetworkVirtualAppliance"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkVirtualAppliance
	err := ctx.RegisterResource("azure-nextgen:network/v20200601:NetworkVirtualAppliance", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-nextgen:network/v20200601:NetworkVirtualAppliance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkVirtualAppliance resources.
type networkVirtualApplianceState struct {
	// Address Prefix.
	AddressPrefix *string `pulumi:"addressPrefix"`
	// BootStrapConfigurationBlobs storage URLs.
	BootStrapConfigurationBlobs []string `pulumi:"bootStrapConfigurationBlobs"`
	// CloudInitConfiguration string in plain text.
	CloudInitConfiguration *string `pulumi:"cloudInitConfiguration"`
	// CloudInitConfigurationBlob storage URLs.
	CloudInitConfigurationBlobs []string `pulumi:"cloudInitConfigurationBlobs"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// The service principal that has read access to cloud-init and config blob.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// List of references to InboundSecurityRules.
	InboundSecurityRules []SubResourceResponse `pulumi:"inboundSecurityRules"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Network Virtual Appliance SKU.
	NvaSku *VirtualApplianceSkuPropertiesResponse `pulumi:"nvaSku"`
	// The provisioning state of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
	// VirtualAppliance ASN.
	VirtualApplianceAsn *float64 `pulumi:"virtualApplianceAsn"`
	// List of Virtual Appliance Network Interfaces.
	VirtualApplianceNics []VirtualApplianceNicPropertiesResponse `pulumi:"virtualApplianceNics"`
	// List of references to VirtualApplianceSite.
	VirtualApplianceSites []SubResourceResponse `pulumi:"virtualApplianceSites"`
	// The Virtual Hub where Network Virtual Appliance is being deployed.
	VirtualHub *SubResourceResponse `pulumi:"virtualHub"`
}

type NetworkVirtualApplianceState struct {
	// Address Prefix.
	AddressPrefix pulumi.StringPtrInput
	// BootStrapConfigurationBlobs storage URLs.
	BootStrapConfigurationBlobs pulumi.StringArrayInput
	// CloudInitConfiguration string in plain text.
	CloudInitConfiguration pulumi.StringPtrInput
	// CloudInitConfigurationBlob storage URLs.
	CloudInitConfigurationBlobs pulumi.StringArrayInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// The service principal that has read access to cloud-init and config blob.
	Identity ManagedServiceIdentityResponsePtrInput
	// List of references to InboundSecurityRules.
	InboundSecurityRules SubResourceResponseArrayInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Network Virtual Appliance SKU.
	NvaSku VirtualApplianceSkuPropertiesResponsePtrInput
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
	// VirtualAppliance ASN.
	VirtualApplianceAsn pulumi.Float64PtrInput
	// List of Virtual Appliance Network Interfaces.
	VirtualApplianceNics VirtualApplianceNicPropertiesResponseArrayInput
	// List of references to VirtualApplianceSite.
	VirtualApplianceSites SubResourceResponseArrayInput
	// The Virtual Hub where Network Virtual Appliance is being deployed.
	VirtualHub SubResourceResponsePtrInput
}

func (NetworkVirtualApplianceState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkVirtualApplianceState)(nil)).Elem()
}

type networkVirtualApplianceArgs struct {
	// BootStrapConfigurationBlobs storage URLs.
	BootStrapConfigurationBlobs []string `pulumi:"bootStrapConfigurationBlobs"`
	// CloudInitConfiguration string in plain text.
	CloudInitConfiguration *string `pulumi:"cloudInitConfiguration"`
	// CloudInitConfigurationBlob storage URLs.
	CloudInitConfigurationBlobs []string `pulumi:"cloudInitConfigurationBlobs"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The service principal that has read access to cloud-init and config blob.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of Network Virtual Appliance.
	NetworkVirtualApplianceName string `pulumi:"networkVirtualApplianceName"`
	// Network Virtual Appliance SKU.
	NvaSku *VirtualApplianceSkuProperties `pulumi:"nvaSku"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// VirtualAppliance ASN.
	VirtualApplianceAsn *float64 `pulumi:"virtualApplianceAsn"`
	// The Virtual Hub where Network Virtual Appliance is being deployed.
	VirtualHub *SubResource `pulumi:"virtualHub"`
}

// The set of arguments for constructing a NetworkVirtualAppliance resource.
type NetworkVirtualApplianceArgs struct {
	// BootStrapConfigurationBlobs storage URLs.
	BootStrapConfigurationBlobs pulumi.StringArrayInput
	// CloudInitConfiguration string in plain text.
	CloudInitConfiguration pulumi.StringPtrInput
	// CloudInitConfigurationBlob storage URLs.
	CloudInitConfigurationBlobs pulumi.StringArrayInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The service principal that has read access to cloud-init and config blob.
	Identity ManagedServiceIdentityPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of Network Virtual Appliance.
	NetworkVirtualApplianceName pulumi.StringInput
	// Network Virtual Appliance SKU.
	NvaSku VirtualApplianceSkuPropertiesPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// VirtualAppliance ASN.
	VirtualApplianceAsn pulumi.Float64PtrInput
	// The Virtual Hub where Network Virtual Appliance is being deployed.
	VirtualHub SubResourcePtrInput
}

func (NetworkVirtualApplianceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkVirtualApplianceArgs)(nil)).Elem()
}

type NetworkVirtualApplianceInput interface {
	pulumi.Input

	ToNetworkVirtualApplianceOutput() NetworkVirtualApplianceOutput
	ToNetworkVirtualApplianceOutputWithContext(ctx context.Context) NetworkVirtualApplianceOutput
}

func (*NetworkVirtualAppliance) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkVirtualAppliance)(nil))
}

func (i *NetworkVirtualAppliance) ToNetworkVirtualApplianceOutput() NetworkVirtualApplianceOutput {
	return i.ToNetworkVirtualApplianceOutputWithContext(context.Background())
}

func (i *NetworkVirtualAppliance) ToNetworkVirtualApplianceOutputWithContext(ctx context.Context) NetworkVirtualApplianceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkVirtualApplianceOutput)
}

type NetworkVirtualApplianceOutput struct {
	*pulumi.OutputState
}

func (NetworkVirtualApplianceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkVirtualAppliance)(nil))
}

func (o NetworkVirtualApplianceOutput) ToNetworkVirtualApplianceOutput() NetworkVirtualApplianceOutput {
	return o
}

func (o NetworkVirtualApplianceOutput) ToNetworkVirtualApplianceOutputWithContext(ctx context.Context) NetworkVirtualApplianceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NetworkVirtualApplianceOutput{})
}
