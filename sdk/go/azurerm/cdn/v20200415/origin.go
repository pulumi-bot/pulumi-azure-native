// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200415

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// CDN origin is the source of the content being delivered via CDN. When the edge nodes represented by an endpoint do not have the requested content cached, they attempt to fetch it from one or more of the configured origins.
//
// ## Example Usage
// ### Origins_Create
//
// ```go
// package main
//
// import (
// 	cdn "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/cdn/v20200415"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cdn.NewOrigin(ctx, "origin", &cdn.OriginArgs{
// 			Enabled:                    pulumi.Bool(true),
// 			EndpointName:               pulumi.String("endpoint1"),
// 			HostName:                   pulumi.String("www.someDomain.net"),
// 			HttpPort:                   pulumi.Int(80),
// 			HttpsPort:                  pulumi.Int(443),
// 			OriginHostHeader:           pulumi.String("www.someDomain.net"),
// 			OriginName:                 pulumi.String("www-someDomain-net"),
// 			Priority:                   pulumi.Int(1),
// 			PrivateLinkApprovalMessage: pulumi.String("Please approve the connection request for this Private Link"),
// 			PrivateLinkLocation:        pulumi.String("eastus"),
// 			PrivateLinkResourceId:      pulumi.String("/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.Network/privateLinkServices/pls1"),
// 			ProfileName:                pulumi.String("profile1"),
// 			ResourceGroupName:          pulumi.String("RG"),
// 			Weight:                     pulumi.Int(50),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
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
	// The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default. This overrides the host header defined at Endpoint
	OriginHostHeader pulumi.StringPtrOutput `pulumi:"originHostHeader"`
	// Priority of origin in given origin group for load balancing. Higher priorities will not be used for load balancing if any lower priority origin is healthy.Must be between 1 and 5
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// The approval status for the connection to the Private Link
	PrivateEndpointStatus pulumi.StringOutput `pulumi:"privateEndpointStatus"`
	// The Alias of the Private Link resource. Populating this optional field indicates that this origin is 'Private'
	PrivateLinkAlias pulumi.StringPtrOutput `pulumi:"privateLinkAlias"`
	// A custom message to be included in the approval request to connect to the Private Link.
	PrivateLinkApprovalMessage pulumi.StringPtrOutput `pulumi:"privateLinkApprovalMessage"`
	// The location of the Private Link resource. Required only if 'privateLinkResourceId' is populated
	PrivateLinkLocation pulumi.StringPtrOutput `pulumi:"privateLinkLocation"`
	// The Resource Id of the Private Link resource. Populating this optional field indicates that this backend is 'Private'
	PrivateLinkResourceId pulumi.StringPtrOutput `pulumi:"privateLinkResourceId"`
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
			Type: pulumi.String("azurerm:cdn/latest:Origin"),
		},
		{
			Type: pulumi.String("azurerm:cdn/v20150601:Origin"),
		},
		{
			Type: pulumi.String("azurerm:cdn/v20160402:Origin"),
		},
		{
			Type: pulumi.String("azurerm:cdn/v20191231:Origin"),
		},
		{
			Type: pulumi.String("azurerm:cdn/v20200331:Origin"),
		},
	})
	opts = append(opts, aliases)
	var resource Origin
	err := ctx.RegisterResource("azurerm:cdn/v20200415:Origin", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:cdn/v20200415:Origin", name, id, state, &resource, opts...)
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
	// The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default. This overrides the host header defined at Endpoint
	OriginHostHeader *string `pulumi:"originHostHeader"`
	// Priority of origin in given origin group for load balancing. Higher priorities will not be used for load balancing if any lower priority origin is healthy.Must be between 1 and 5
	Priority *int `pulumi:"priority"`
	// The approval status for the connection to the Private Link
	PrivateEndpointStatus *string `pulumi:"privateEndpointStatus"`
	// The Alias of the Private Link resource. Populating this optional field indicates that this origin is 'Private'
	PrivateLinkAlias *string `pulumi:"privateLinkAlias"`
	// A custom message to be included in the approval request to connect to the Private Link.
	PrivateLinkApprovalMessage *string `pulumi:"privateLinkApprovalMessage"`
	// The location of the Private Link resource. Required only if 'privateLinkResourceId' is populated
	PrivateLinkLocation *string `pulumi:"privateLinkLocation"`
	// The Resource Id of the Private Link resource. Populating this optional field indicates that this backend is 'Private'
	PrivateLinkResourceId *string `pulumi:"privateLinkResourceId"`
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
	// The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default. This overrides the host header defined at Endpoint
	OriginHostHeader pulumi.StringPtrInput
	// Priority of origin in given origin group for load balancing. Higher priorities will not be used for load balancing if any lower priority origin is healthy.Must be between 1 and 5
	Priority pulumi.IntPtrInput
	// The approval status for the connection to the Private Link
	PrivateEndpointStatus pulumi.StringPtrInput
	// The Alias of the Private Link resource. Populating this optional field indicates that this origin is 'Private'
	PrivateLinkAlias pulumi.StringPtrInput
	// A custom message to be included in the approval request to connect to the Private Link.
	PrivateLinkApprovalMessage pulumi.StringPtrInput
	// The location of the Private Link resource. Required only if 'privateLinkResourceId' is populated
	PrivateLinkLocation pulumi.StringPtrInput
	// The Resource Id of the Private Link resource. Populating this optional field indicates that this backend is 'Private'
	PrivateLinkResourceId pulumi.StringPtrInput
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
	// The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default. This overrides the host header defined at Endpoint
	OriginHostHeader *string `pulumi:"originHostHeader"`
	// Name of the origin that is unique within the endpoint.
	OriginName string `pulumi:"originName"`
	// Priority of origin in given origin group for load balancing. Higher priorities will not be used for load balancing if any lower priority origin is healthy.Must be between 1 and 5
	Priority *int `pulumi:"priority"`
	// The Alias of the Private Link resource. Populating this optional field indicates that this origin is 'Private'
	PrivateLinkAlias *string `pulumi:"privateLinkAlias"`
	// A custom message to be included in the approval request to connect to the Private Link.
	PrivateLinkApprovalMessage *string `pulumi:"privateLinkApprovalMessage"`
	// The location of the Private Link resource. Required only if 'privateLinkResourceId' is populated
	PrivateLinkLocation *string `pulumi:"privateLinkLocation"`
	// The Resource Id of the Private Link resource. Populating this optional field indicates that this backend is 'Private'
	PrivateLinkResourceId *string `pulumi:"privateLinkResourceId"`
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
	// The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default. This overrides the host header defined at Endpoint
	OriginHostHeader pulumi.StringPtrInput
	// Name of the origin that is unique within the endpoint.
	OriginName pulumi.StringInput
	// Priority of origin in given origin group for load balancing. Higher priorities will not be used for load balancing if any lower priority origin is healthy.Must be between 1 and 5
	Priority pulumi.IntPtrInput
	// The Alias of the Private Link resource. Populating this optional field indicates that this origin is 'Private'
	PrivateLinkAlias pulumi.StringPtrInput
	// A custom message to be included in the approval request to connect to the Private Link.
	PrivateLinkApprovalMessage pulumi.StringPtrInput
	// The location of the Private Link resource. Required only if 'privateLinkResourceId' is populated
	PrivateLinkLocation pulumi.StringPtrInput
	// The Resource Id of the Private Link resource. Populating this optional field indicates that this backend is 'Private'
	PrivateLinkResourceId pulumi.StringPtrInput
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
