// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Public IP prefix resource.
type PublicIPPrefix struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Public IP prefix properties.
	Properties PublicIPPrefixPropertiesFormatResponseOutput `pulumi:"properties"`
	// The public IP prefix SKU.
	Sku PublicIPPrefixSkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewPublicIPPrefix registers a new resource with the given unique name, arguments, and options.
func NewPublicIPPrefix(ctx *pulumi.Context,
	name string, args *PublicIPPrefixArgs, opts ...pulumi.ResourceOption) (*PublicIPPrefix, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &PublicIPPrefixArgs{}
	}
	var resource PublicIPPrefix
	err := ctx.RegisterResource("azurerm:network/v20180801:PublicIPPrefix", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublicIPPrefix gets an existing PublicIPPrefix resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicIPPrefix(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicIPPrefixState, opts ...pulumi.ResourceOption) (*PublicIPPrefix, error) {
	var resource PublicIPPrefix
	err := ctx.ReadResource("azurerm:network/v20180801:PublicIPPrefix", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublicIPPrefix resources.
type publicIPPrefixState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Public IP prefix properties.
	Properties *PublicIPPrefixPropertiesFormatResponse `pulumi:"properties"`
	// The public IP prefix SKU.
	Sku *PublicIPPrefixSkuResponse `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
	// A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones []string `pulumi:"zones"`
}

type PublicIPPrefixState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Public IP prefix properties.
	Properties PublicIPPrefixPropertiesFormatResponsePtrInput
	// The public IP prefix SKU.
	Sku PublicIPPrefixSkuResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
	// A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones pulumi.StringArrayInput
}

func (PublicIPPrefixState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIPPrefixState)(nil)).Elem()
}

type publicIPPrefixArgs struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The allocated Prefix
	IpPrefix *string `pulumi:"ipPrefix"`
	// The list of tags associated with the public IP prefix.
	IpTags []IpTag `pulumi:"ipTags"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the public IP prefix.
	Name string `pulumi:"name"`
	// The Length of the Public IP Prefix.
	PrefixLength *int `pulumi:"prefixLength"`
	// The provisioning state of the Public IP prefix resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The public IP address version. Possible values are: 'IPv4' and 'IPv6'.
	PublicIPAddressVersion *string `pulumi:"publicIPAddressVersion"`
	// The list of all referenced PublicIPAddresses
	PublicIPAddresses []ReferencedPublicIpAddress `pulumi:"publicIPAddresses"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource GUID property of the public IP prefix resource.
	ResourceGuid *string `pulumi:"resourceGuid"`
	// The public IP prefix SKU.
	Sku *PublicIPPrefixSku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a PublicIPPrefix resource.
type PublicIPPrefixArgs struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The allocated Prefix
	IpPrefix pulumi.StringPtrInput
	// The list of tags associated with the public IP prefix.
	IpTags IpTagArrayInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the public IP prefix.
	Name pulumi.StringInput
	// The Length of the Public IP Prefix.
	PrefixLength pulumi.IntPtrInput
	// The provisioning state of the Public IP prefix resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringPtrInput
	// The public IP address version. Possible values are: 'IPv4' and 'IPv6'.
	PublicIPAddressVersion pulumi.StringPtrInput
	// The list of all referenced PublicIPAddresses
	PublicIPAddresses ReferencedPublicIpAddressArrayInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The resource GUID property of the public IP prefix resource.
	ResourceGuid pulumi.StringPtrInput
	// The public IP prefix SKU.
	Sku PublicIPPrefixSkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones pulumi.StringArrayInput
}

func (PublicIPPrefixArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIPPrefixArgs)(nil)).Elem()
}
