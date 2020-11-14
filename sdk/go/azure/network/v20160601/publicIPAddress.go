// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// PublicIPAddress resource
type PublicIPAddress struct {
	pulumi.CustomResourceState

	// Gets or sets FQDN of the DNS record associated with the public IP address
	DnsSettings PublicIPAddressDnsSettingsResponsePtrOutput `pulumi:"dnsSettings"`
	// Gets a unique read-only string that changes whenever the resource is updated
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Gets or sets the idle timeout of the public IP address
	IdleTimeoutInMinutes pulumi.IntPtrOutput    `pulumi:"idleTimeoutInMinutes"`
	IpAddress            pulumi.StringPtrOutput `pulumi:"ipAddress"`
	// IPConfiguration
	IpConfiguration IPConfigurationResponseOutput `pulumi:"ipConfiguration"`
	// Resource location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets provisioning state of the PublicIP resource Updating/Deleting/Failed
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Gets or sets PublicIP address version (IPv4/IPv6)
	PublicIPAddressVersion pulumi.StringPtrOutput `pulumi:"publicIPAddressVersion"`
	// Gets or sets PublicIP allocation method (Static/Dynamic)
	PublicIPAllocationMethod pulumi.StringPtrOutput `pulumi:"publicIPAllocationMethod"`
	// Gets or sets resource guid property of the PublicIP resource
	ResourceGuid pulumi.StringPtrOutput `pulumi:"resourceGuid"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPublicIPAddress registers a new resource with the given unique name, arguments, and options.
func NewPublicIPAddress(ctx *pulumi.Context,
	name string, args *PublicIPAddressArgs, opts ...pulumi.ResourceOption) (*PublicIPAddress, error) {
	if args == nil || args.PublicIpAddressName == nil {
		return nil, errors.New("missing required argument 'PublicIpAddressName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &PublicIPAddressArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:network/latest:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20150501preview:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20150615:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20160330:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20160901:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20161201:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170301:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170601:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170801:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170901:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20171001:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20171101:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180101:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180201:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180401:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180601:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180701:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180801:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181001:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181101:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181201:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190201:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190401:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190601:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190701:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190801:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190901:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191101:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191201:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200301:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200401:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200501:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200701:PublicIPAddress"),
		},
	})
	opts = append(opts, aliases)
	var resource PublicIPAddress
	err := ctx.RegisterResource("azure-nextgen:network/v20160601:PublicIPAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublicIPAddress gets an existing PublicIPAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicIPAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicIPAddressState, opts ...pulumi.ResourceOption) (*PublicIPAddress, error) {
	var resource PublicIPAddress
	err := ctx.ReadResource("azure-nextgen:network/v20160601:PublicIPAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublicIPAddress resources.
type publicIPAddressState struct {
	// Gets or sets FQDN of the DNS record associated with the public IP address
	DnsSettings *PublicIPAddressDnsSettingsResponse `pulumi:"dnsSettings"`
	// Gets a unique read-only string that changes whenever the resource is updated
	Etag *string `pulumi:"etag"`
	// Gets or sets the idle timeout of the public IP address
	IdleTimeoutInMinutes *int    `pulumi:"idleTimeoutInMinutes"`
	IpAddress            *string `pulumi:"ipAddress"`
	// IPConfiguration
	IpConfiguration *IPConfigurationResponse `pulumi:"ipConfiguration"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Gets provisioning state of the PublicIP resource Updating/Deleting/Failed
	ProvisioningState *string `pulumi:"provisioningState"`
	// Gets or sets PublicIP address version (IPv4/IPv6)
	PublicIPAddressVersion *string `pulumi:"publicIPAddressVersion"`
	// Gets or sets PublicIP allocation method (Static/Dynamic)
	PublicIPAllocationMethod *string `pulumi:"publicIPAllocationMethod"`
	// Gets or sets resource guid property of the PublicIP resource
	ResourceGuid *string `pulumi:"resourceGuid"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type PublicIPAddressState struct {
	// Gets or sets FQDN of the DNS record associated with the public IP address
	DnsSettings PublicIPAddressDnsSettingsResponsePtrInput
	// Gets a unique read-only string that changes whenever the resource is updated
	Etag pulumi.StringPtrInput
	// Gets or sets the idle timeout of the public IP address
	IdleTimeoutInMinutes pulumi.IntPtrInput
	IpAddress            pulumi.StringPtrInput
	// IPConfiguration
	IpConfiguration IPConfigurationResponsePtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Gets provisioning state of the PublicIP resource Updating/Deleting/Failed
	ProvisioningState pulumi.StringPtrInput
	// Gets or sets PublicIP address version (IPv4/IPv6)
	PublicIPAddressVersion pulumi.StringPtrInput
	// Gets or sets PublicIP allocation method (Static/Dynamic)
	PublicIPAllocationMethod pulumi.StringPtrInput
	// Gets or sets resource guid property of the PublicIP resource
	ResourceGuid pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (PublicIPAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIPAddressState)(nil)).Elem()
}

type publicIPAddressArgs struct {
	// Gets or sets FQDN of the DNS record associated with the public IP address
	DnsSettings *PublicIPAddressDnsSettings `pulumi:"dnsSettings"`
	// Gets a unique read-only string that changes whenever the resource is updated
	Etag *string `pulumi:"etag"`
	// Resource Id
	Id *string `pulumi:"id"`
	// Gets or sets the idle timeout of the public IP address
	IdleTimeoutInMinutes *int    `pulumi:"idleTimeoutInMinutes"`
	IpAddress            *string `pulumi:"ipAddress"`
	// Resource location
	Location *string `pulumi:"location"`
	// Gets provisioning state of the PublicIP resource Updating/Deleting/Failed
	ProvisioningState *string `pulumi:"provisioningState"`
	// Gets or sets PublicIP address version (IPv4/IPv6)
	PublicIPAddressVersion *string `pulumi:"publicIPAddressVersion"`
	// Gets or sets PublicIP allocation method (Static/Dynamic)
	PublicIPAllocationMethod *string `pulumi:"publicIPAllocationMethod"`
	// The name of the publicIpAddress.
	PublicIpAddressName string `pulumi:"publicIpAddressName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets resource guid property of the PublicIP resource
	ResourceGuid *string `pulumi:"resourceGuid"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a PublicIPAddress resource.
type PublicIPAddressArgs struct {
	// Gets or sets FQDN of the DNS record associated with the public IP address
	DnsSettings PublicIPAddressDnsSettingsPtrInput
	// Gets a unique read-only string that changes whenever the resource is updated
	Etag pulumi.StringPtrInput
	// Resource Id
	Id pulumi.StringPtrInput
	// Gets or sets the idle timeout of the public IP address
	IdleTimeoutInMinutes pulumi.IntPtrInput
	IpAddress            pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Gets provisioning state of the PublicIP resource Updating/Deleting/Failed
	ProvisioningState pulumi.StringPtrInput
	// Gets or sets PublicIP address version (IPv4/IPv6)
	PublicIPAddressVersion pulumi.StringPtrInput
	// Gets or sets PublicIP allocation method (Static/Dynamic)
	PublicIPAllocationMethod pulumi.StringPtrInput
	// The name of the publicIpAddress.
	PublicIpAddressName pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Gets or sets resource guid property of the PublicIP resource
	ResourceGuid pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (PublicIPAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIPAddressArgs)(nil)).Elem()
}

type PublicIPAddressInput interface {
	pulumi.Input

	ToPublicIPAddressOutput() PublicIPAddressOutput
	ToPublicIPAddressOutputWithContext(ctx context.Context) PublicIPAddressOutput
}

func (PublicIPAddress) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddress)(nil)).Elem()
}

func (i PublicIPAddress) ToPublicIPAddressOutput() PublicIPAddressOutput {
	return i.ToPublicIPAddressOutputWithContext(context.Background())
}

func (i PublicIPAddress) ToPublicIPAddressOutputWithContext(ctx context.Context) PublicIPAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressOutput)
}

type PublicIPAddressOutput struct {
	*pulumi.OutputState
}

func (PublicIPAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressOutput)(nil)).Elem()
}

func (o PublicIPAddressOutput) ToPublicIPAddressOutput() PublicIPAddressOutput {
	return o
}

func (o PublicIPAddressOutput) ToPublicIPAddressOutputWithContext(ctx context.Context) PublicIPAddressOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PublicIPAddressOutput{})
}
