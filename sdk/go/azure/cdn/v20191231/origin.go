// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191231

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// CDN origin is the source of the content being delivered via CDN. When the edge nodes represented by an endpoint do not have the requested content cached, they attempt to fetch it from one or more of the configured origins.
type Origin struct {
	pulumi.CustomResourceState

	// Origin is enabled for load balancing or not
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The address of the origin. Domain names, IPv4 addresses, and IPv6 addresses are supported.This should be unique across all origins in an endpoint.
	HostName pulumi.StringOutput `pulumi:"hostName"`
	// The value of the HTTP port. Must be between 1 and 65535.
	HttpPort pulumi.IntPtrOutput `pulumi:"httpPort"`
	// The value of the HTTPS port. Must be between 1 and 65535.
	HttpsPort pulumi.IntPtrOutput `pulumi:"httpsPort"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default. If endpoint uses multiple origins for load balancing, then the host header at endpoint is ignored and this one is considered.
	OriginHostHeader pulumi.StringPtrOutput `pulumi:"originHostHeader"`
	// Priority of origin in given origin group for load balancing. Higher priorities will not be used for load balancing if any lower priority origin is healthy.Must be between 1 and 5
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// Provisioning status of the origin.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource status of the origin.
	ResourceState pulumi.StringOutput `pulumi:"resourceState"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Weight of the origin in given origin group for load balancing. Must be between 1 and 1000
	Weight pulumi.IntPtrOutput `pulumi:"weight"`
}

// NewOrigin registers a new resource with the given unique name, arguments, and options.
func NewOrigin(ctx *pulumi.Context,
	name string, args *OriginArgs, opts ...pulumi.ResourceOption) (*Origin, error) {
	if args == nil || args.EndpointName == nil {
		return nil, errors.New("missing required argument 'EndpointName'")
	}
	if args == nil || args.HostName == nil {
		return nil, errors.New("missing required argument 'HostName'")
	}
	if args == nil || args.OriginName == nil {
		return nil, errors.New("missing required argument 'OriginName'")
	}
	if args == nil || args.ProfileName == nil {
		return nil, errors.New("missing required argument 'ProfileName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &OriginArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:cdn/latest:Origin"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn/v20150601:Origin"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn/v20160402:Origin"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn/v20200331:Origin"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn/v20200415:Origin"),
		},
	})
	opts = append(opts, aliases)
	var resource Origin
	err := ctx.RegisterResource("azure-nextgen:cdn/v20191231:Origin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrigin gets an existing Origin resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrigin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OriginState, opts ...pulumi.ResourceOption) (*Origin, error) {
	var resource Origin
	err := ctx.ReadResource("azure-nextgen:cdn/v20191231:Origin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Origin resources.
type originState struct {
	// Origin is enabled for load balancing or not
	Enabled *bool `pulumi:"enabled"`
	// The address of the origin. Domain names, IPv4 addresses, and IPv6 addresses are supported.This should be unique across all origins in an endpoint.
	HostName *string `pulumi:"hostName"`
	// The value of the HTTP port. Must be between 1 and 65535.
	HttpPort *int `pulumi:"httpPort"`
	// The value of the HTTPS port. Must be between 1 and 65535.
	HttpsPort *int `pulumi:"httpsPort"`
	// Resource name.
	Name *string `pulumi:"name"`
	// The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default. If endpoint uses multiple origins for load balancing, then the host header at endpoint is ignored and this one is considered.
	OriginHostHeader *string `pulumi:"originHostHeader"`
	// Priority of origin in given origin group for load balancing. Higher priorities will not be used for load balancing if any lower priority origin is healthy.Must be between 1 and 5
	Priority *int `pulumi:"priority"`
	// Provisioning status of the origin.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Resource status of the origin.
	ResourceState *string `pulumi:"resourceState"`
	// Resource type.
	Type *string `pulumi:"type"`
	// Weight of the origin in given origin group for load balancing. Must be between 1 and 1000
	Weight *int `pulumi:"weight"`
}

type OriginState struct {
	// Origin is enabled for load balancing or not
	Enabled pulumi.BoolPtrInput
	// The address of the origin. Domain names, IPv4 addresses, and IPv6 addresses are supported.This should be unique across all origins in an endpoint.
	HostName pulumi.StringPtrInput
	// The value of the HTTP port. Must be between 1 and 65535.
	HttpPort pulumi.IntPtrInput
	// The value of the HTTPS port. Must be between 1 and 65535.
	HttpsPort pulumi.IntPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default. If endpoint uses multiple origins for load balancing, then the host header at endpoint is ignored and this one is considered.
	OriginHostHeader pulumi.StringPtrInput
	// Priority of origin in given origin group for load balancing. Higher priorities will not be used for load balancing if any lower priority origin is healthy.Must be between 1 and 5
	Priority pulumi.IntPtrInput
	// Provisioning status of the origin.
	ProvisioningState pulumi.StringPtrInput
	// Resource status of the origin.
	ResourceState pulumi.StringPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
	// Weight of the origin in given origin group for load balancing. Must be between 1 and 1000
	Weight pulumi.IntPtrInput
}

func (OriginState) ElementType() reflect.Type {
	return reflect.TypeOf((*originState)(nil)).Elem()
}

type originArgs struct {
	// Origin is enabled for load balancing or not
	Enabled *bool `pulumi:"enabled"`
	// Name of the endpoint under the profile which is unique globally.
	EndpointName string `pulumi:"endpointName"`
	// The address of the origin. Domain names, IPv4 addresses, and IPv6 addresses are supported.This should be unique across all origins in an endpoint.
	HostName string `pulumi:"hostName"`
	// The value of the HTTP port. Must be between 1 and 65535.
	HttpPort *int `pulumi:"httpPort"`
	// The value of the HTTPS port. Must be between 1 and 65535.
	HttpsPort *int `pulumi:"httpsPort"`
	// The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default. If endpoint uses multiple origins for load balancing, then the host header at endpoint is ignored and this one is considered.
	OriginHostHeader *string `pulumi:"originHostHeader"`
	// Name of the origin that is unique within the endpoint.
	OriginName string `pulumi:"originName"`
	// Priority of origin in given origin group for load balancing. Higher priorities will not be used for load balancing if any lower priority origin is healthy.Must be between 1 and 5
	Priority *int `pulumi:"priority"`
	// Name of the CDN profile which is unique within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Weight of the origin in given origin group for load balancing. Must be between 1 and 1000
	Weight *int `pulumi:"weight"`
}

// The set of arguments for constructing a Origin resource.
type OriginArgs struct {
	// Origin is enabled for load balancing or not
	Enabled pulumi.BoolPtrInput
	// Name of the endpoint under the profile which is unique globally.
	EndpointName pulumi.StringInput
	// The address of the origin. Domain names, IPv4 addresses, and IPv6 addresses are supported.This should be unique across all origins in an endpoint.
	HostName pulumi.StringInput
	// The value of the HTTP port. Must be between 1 and 65535.
	HttpPort pulumi.IntPtrInput
	// The value of the HTTPS port. Must be between 1 and 65535.
	HttpsPort pulumi.IntPtrInput
	// The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default. If endpoint uses multiple origins for load balancing, then the host header at endpoint is ignored and this one is considered.
	OriginHostHeader pulumi.StringPtrInput
	// Name of the origin that is unique within the endpoint.
	OriginName pulumi.StringInput
	// Priority of origin in given origin group for load balancing. Higher priorities will not be used for load balancing if any lower priority origin is healthy.Must be between 1 and 5
	Priority pulumi.IntPtrInput
	// Name of the CDN profile which is unique within the resource group.
	ProfileName pulumi.StringInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// Weight of the origin in given origin group for load balancing. Must be between 1 and 1000
	Weight pulumi.IntPtrInput
}

func (OriginArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*originArgs)(nil)).Elem()
}

type OriginInput interface {
	pulumi.Input

	ToOriginOutput() OriginOutput
	ToOriginOutputWithContext(ctx context.Context) OriginOutput
}

func (Origin) ElementType() reflect.Type {
	return reflect.TypeOf((*Origin)(nil)).Elem()
}

func (i Origin) ToOriginOutput() OriginOutput {
	return i.ToOriginOutputWithContext(context.Background())
}

func (i Origin) ToOriginOutputWithContext(ctx context.Context) OriginOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OriginOutput)
}

type OriginOutput struct {
	*pulumi.OutputState
}

func (OriginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OriginOutput)(nil)).Elem()
}

func (o OriginOutput) ToOriginOutput() OriginOutput {
	return o
}

func (o OriginOutput) ToOriginOutputWithContext(ctx context.Context) OriginOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OriginOutput{})
}
