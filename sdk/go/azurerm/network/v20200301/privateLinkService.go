// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200301

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Private link service resource.
type PrivateLinkService struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the private link service.
	Properties PrivateLinkServicePropertiesResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateLinkService registers a new resource with the given unique name, arguments, and options.
func NewPrivateLinkService(ctx *pulumi.Context,
	name string, args *PrivateLinkServiceArgs, opts ...pulumi.ResourceOption) (*PrivateLinkService, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &PrivateLinkServiceArgs{}
	}
	var resource PrivateLinkService
	err := ctx.RegisterResource("azurerm:network/v20200301:PrivateLinkService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateLinkService gets an existing PrivateLinkService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateLinkService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateLinkServiceState, opts ...pulumi.ResourceOption) (*PrivateLinkService, error) {
	var resource PrivateLinkService
	err := ctx.ReadResource("azurerm:network/v20200301:PrivateLinkService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateLinkService resources.
type privateLinkServiceState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Properties of the private link service.
	Properties *PrivateLinkServicePropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type PrivateLinkServiceState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Properties of the private link service.
	Properties PrivateLinkServicePropertiesResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (PrivateLinkServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkServiceState)(nil)).Elem()
}

type privateLinkServiceArgs struct {
	// The auto-approval list of the private link service.
	AutoApproval map[string]interface{} `pulumi:"autoApproval"`
	// Whether the private link service is enabled for proxy protocol or not.
	EnableProxyProtocol *bool `pulumi:"enableProxyProtocol"`
	// The list of Fqdn.
	Fqdns []string `pulumi:"fqdns"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// An array of private link service IP configurations.
	IpConfigurations []PrivateLinkServiceIpConfiguration `pulumi:"ipConfigurations"`
	// An array of references to the load balancer IP configurations.
	LoadBalancerFrontendIpConfigurations []FrontendIPConfiguration `pulumi:"loadBalancerFrontendIpConfigurations"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the private link service.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The visibility list of the private link service.
	Visibility map[string]interface{} `pulumi:"visibility"`
}

// The set of arguments for constructing a PrivateLinkService resource.
type PrivateLinkServiceArgs struct {
	// The auto-approval list of the private link service.
	AutoApproval pulumi.MapInput
	// Whether the private link service is enabled for proxy protocol or not.
	EnableProxyProtocol pulumi.BoolPtrInput
	// The list of Fqdn.
	Fqdns pulumi.StringArrayInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// An array of private link service IP configurations.
	IpConfigurations PrivateLinkServiceIpConfigurationArrayInput
	// An array of references to the load balancer IP configurations.
	LoadBalancerFrontendIpConfigurations FrontendIPConfigurationArrayInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the private link service.
	Name pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The visibility list of the private link service.
	Visibility pulumi.MapInput
}

func (PrivateLinkServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkServiceArgs)(nil)).Elem()
}
