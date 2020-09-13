// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180301

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Class representing a Traffic Manager endpoint.
//
// ## Example Usage
// ### Endpoint-PUT-External-WithCustomHeaders
//
// ```go
// package main
//
// import (
// 	network "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/network/v20180301"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := network.NewEndpoint(ctx, "endpoint", &network.EndpointArgs{
// 			CustomHeaders: network.EndpointPropertiesCustomHeadersArray{
// 				&network.EndpointPropertiesCustomHeadersArgs{
// 					Name:  pulumi.String("header-1"),
// 					Value: pulumi.String("value-1"),
// 				},
// 				&network.EndpointPropertiesCustomHeadersArgs{
// 					Name:  pulumi.String("header-2"),
// 					Value: pulumi.String("value-2"),
// 				},
// 			},
// 			EndpointLocation:  pulumi.String("North Europe"),
// 			EndpointName:      pulumi.String("azsmnet7187"),
// 			EndpointStatus:    pulumi.String("Enabled"),
// 			EndpointType:      pulumi.String("ExternalEndpoints"),
// 			Name:              pulumi.String("azsmnet7187"),
// 			ProfileName:       pulumi.String("azsmnet6386"),
// 			ResourceGroupName: pulumi.String("azuresdkfornetautoresttrafficmanager1421"),
// 			Target:            pulumi.String("foobar.contoso.com"),
// 			Type:              pulumi.String("Microsoft.network/TrafficManagerProfiles/ExternalEndpoints"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Endpoint-PUT-External-WithGeoMapping
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	network "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/network/v20180301"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := network.NewEndpoint(ctx, "endpoint", &network.EndpointArgs{
// 			EndpointName:   pulumi.String(fmt.Sprintf("%v%v%v%v%v", "My", "%", "20external", "%", "20endpoint")),
// 			EndpointStatus: pulumi.String("Enabled"),
// 			EndpointType:   pulumi.String("ExternalEndpoints"),
// 			GeoMapping: pulumi.StringArray{
// 				pulumi.String("GEO-AS"),
// 				pulumi.String("GEO-AF"),
// 			},
// 			Name:              pulumi.String("My external endpoint"),
// 			ProfileName:       pulumi.String("azuresdkfornetautoresttrafficmanager8224"),
// 			ResourceGroupName: pulumi.String("azuresdkfornetautoresttrafficmanager2191"),
// 			Target:            pulumi.String("foobar.contoso.com"),
// 			Type:              pulumi.String("Microsoft.network/TrafficManagerProfiles/ExternalEndpoints"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Endpoint-PUT-External-WithLocation
//
// ```go
// package main
//
// import (
// 	network "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/network/v20180301"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := network.NewEndpoint(ctx, "endpoint", &network.EndpointArgs{
// 			EndpointLocation:  pulumi.String("North Europe"),
// 			EndpointName:      pulumi.String("azsmnet7187"),
// 			EndpointStatus:    pulumi.String("Enabled"),
// 			EndpointType:      pulumi.String("ExternalEndpoints"),
// 			Name:              pulumi.String("azsmnet7187"),
// 			ProfileName:       pulumi.String("azsmnet6386"),
// 			ResourceGroupName: pulumi.String("azuresdkfornetautoresttrafficmanager1421"),
// 			Target:            pulumi.String("foobar.contoso.com"),
// 			Type:              pulumi.String("Microsoft.network/TrafficManagerProfiles/ExternalEndpoints"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Endpoint struct {
	pulumi.CustomResourceState

	// List of custom headers.
	CustomHeaders EndpointPropertiesResponseCustomHeadersArrayOutput `pulumi:"customHeaders"`
	// Specifies the location of the external or nested endpoints when using the ‘Performance’ traffic routing method.
	EndpointLocation pulumi.StringPtrOutput `pulumi:"endpointLocation"`
	// The monitoring status of the endpoint.
	EndpointMonitorStatus pulumi.StringPtrOutput `pulumi:"endpointMonitorStatus"`
	// The status of the endpoint. If the endpoint is Enabled, it is probed for endpoint health and is included in the traffic routing method.
	EndpointStatus pulumi.StringPtrOutput `pulumi:"endpointStatus"`
	// The list of countries/regions mapped to this endpoint when using the ‘Geographic’ traffic routing method. Please consult Traffic Manager Geographic documentation for a full list of accepted values.
	GeoMapping pulumi.StringArrayOutput `pulumi:"geoMapping"`
	// The minimum number of endpoints that must be available in the child profile in order for the parent profile to be considered available. Only applicable to endpoint of type 'NestedEndpoints'.
	MinChildEndpoints pulumi.IntPtrOutput `pulumi:"minChildEndpoints"`
	// The name of the resource
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The priority of this endpoint when using the ‘Priority’ traffic routing method. Possible values are from 1 to 1000, lower values represent higher priority. This is an optional parameter.  If specified, it must be specified on all endpoints, and no two endpoints can share the same priority value.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// The fully-qualified DNS name or IP address of the endpoint. Traffic Manager returns this value in DNS responses to direct traffic to this endpoint.
	Target pulumi.StringPtrOutput `pulumi:"target"`
	// The Azure Resource URI of the of the endpoint. Not applicable to endpoints of type 'ExternalEndpoints'.
	TargetResourceId pulumi.StringPtrOutput `pulumi:"targetResourceId"`
	// The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// The weight of this endpoint when using the 'Weighted' traffic routing method. Possible values are from 1 to 1000.
	Weight pulumi.IntPtrOutput `pulumi:"weight"`
}

// NewEndpoint registers a new resource with the given unique name, arguments, and options.
func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	if args == nil || args.EndpointName == nil {
		return nil, errors.New("missing required argument 'EndpointName'")
	}
	if args == nil || args.EndpointType == nil {
		return nil, errors.New("missing required argument 'EndpointType'")
	}
	if args == nil || args.ProfileName == nil {
		return nil, errors.New("missing required argument 'ProfileName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &EndpointArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:network/latest:Endpoint"),
		},
		{
			Type: pulumi.String("azurerm:network/v20151101:Endpoint"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170301:Endpoint"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170501:Endpoint"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180201:Endpoint"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180401:Endpoint"),
		},
	})
	opts = append(opts, aliases)
	var resource Endpoint
	err := ctx.RegisterResource("azurerm:network/v20180301:Endpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpoint gets an existing Endpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointState, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	var resource Endpoint
	err := ctx.ReadResource("azurerm:network/v20180301:Endpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Endpoint resources.
type endpointState struct {
	// List of custom headers.
	CustomHeaders []EndpointPropertiesResponseCustomHeaders `pulumi:"customHeaders"`
	// Specifies the location of the external or nested endpoints when using the ‘Performance’ traffic routing method.
	EndpointLocation *string `pulumi:"endpointLocation"`
	// The monitoring status of the endpoint.
	EndpointMonitorStatus *string `pulumi:"endpointMonitorStatus"`
	// The status of the endpoint. If the endpoint is Enabled, it is probed for endpoint health and is included in the traffic routing method.
	EndpointStatus *string `pulumi:"endpointStatus"`
	// The list of countries/regions mapped to this endpoint when using the ‘Geographic’ traffic routing method. Please consult Traffic Manager Geographic documentation for a full list of accepted values.
	GeoMapping []string `pulumi:"geoMapping"`
	// The minimum number of endpoints that must be available in the child profile in order for the parent profile to be considered available. Only applicable to endpoint of type 'NestedEndpoints'.
	MinChildEndpoints *int `pulumi:"minChildEndpoints"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The priority of this endpoint when using the ‘Priority’ traffic routing method. Possible values are from 1 to 1000, lower values represent higher priority. This is an optional parameter.  If specified, it must be specified on all endpoints, and no two endpoints can share the same priority value.
	Priority *int `pulumi:"priority"`
	// The fully-qualified DNS name or IP address of the endpoint. Traffic Manager returns this value in DNS responses to direct traffic to this endpoint.
	Target *string `pulumi:"target"`
	// The Azure Resource URI of the of the endpoint. Not applicable to endpoints of type 'ExternalEndpoints'.
	TargetResourceId *string `pulumi:"targetResourceId"`
	// The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
	Type *string `pulumi:"type"`
	// The weight of this endpoint when using the 'Weighted' traffic routing method. Possible values are from 1 to 1000.
	Weight *int `pulumi:"weight"`
}

type EndpointState struct {
	// List of custom headers.
	CustomHeaders EndpointPropertiesResponseCustomHeadersArrayInput
	// Specifies the location of the external or nested endpoints when using the ‘Performance’ traffic routing method.
	EndpointLocation pulumi.StringPtrInput
	// The monitoring status of the endpoint.
	EndpointMonitorStatus pulumi.StringPtrInput
	// The status of the endpoint. If the endpoint is Enabled, it is probed for endpoint health and is included in the traffic routing method.
	EndpointStatus pulumi.StringPtrInput
	// The list of countries/regions mapped to this endpoint when using the ‘Geographic’ traffic routing method. Please consult Traffic Manager Geographic documentation for a full list of accepted values.
	GeoMapping pulumi.StringArrayInput
	// The minimum number of endpoints that must be available in the child profile in order for the parent profile to be considered available. Only applicable to endpoint of type 'NestedEndpoints'.
	MinChildEndpoints pulumi.IntPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The priority of this endpoint when using the ‘Priority’ traffic routing method. Possible values are from 1 to 1000, lower values represent higher priority. This is an optional parameter.  If specified, it must be specified on all endpoints, and no two endpoints can share the same priority value.
	Priority pulumi.IntPtrInput
	// The fully-qualified DNS name or IP address of the endpoint. Traffic Manager returns this value in DNS responses to direct traffic to this endpoint.
	Target pulumi.StringPtrInput
	// The Azure Resource URI of the of the endpoint. Not applicable to endpoints of type 'ExternalEndpoints'.
	TargetResourceId pulumi.StringPtrInput
	// The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
	Type pulumi.StringPtrInput
	// The weight of this endpoint when using the 'Weighted' traffic routing method. Possible values are from 1 to 1000.
	Weight pulumi.IntPtrInput
}

func (EndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointState)(nil)).Elem()
}

type endpointArgs struct {
	// List of custom headers.
	CustomHeaders []EndpointPropertiesCustomHeaders `pulumi:"customHeaders"`
	// Specifies the location of the external or nested endpoints when using the ‘Performance’ traffic routing method.
	EndpointLocation *string `pulumi:"endpointLocation"`
	// The monitoring status of the endpoint.
	EndpointMonitorStatus *string `pulumi:"endpointMonitorStatus"`
	// The name of the Traffic Manager endpoint to be created or updated.
	EndpointName string `pulumi:"endpointName"`
	// The status of the endpoint. If the endpoint is Enabled, it is probed for endpoint health and is included in the traffic routing method.
	EndpointStatus *string `pulumi:"endpointStatus"`
	// The type of the Traffic Manager endpoint to be created or updated.
	EndpointType string `pulumi:"endpointType"`
	// The list of countries/regions mapped to this endpoint when using the ‘Geographic’ traffic routing method. Please consult Traffic Manager Geographic documentation for a full list of accepted values.
	GeoMapping []string `pulumi:"geoMapping"`
	// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{resourceName}
	Id *string `pulumi:"id"`
	// The minimum number of endpoints that must be available in the child profile in order for the parent profile to be considered available. Only applicable to endpoint of type 'NestedEndpoints'.
	MinChildEndpoints *int `pulumi:"minChildEndpoints"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The priority of this endpoint when using the ‘Priority’ traffic routing method. Possible values are from 1 to 1000, lower values represent higher priority. This is an optional parameter.  If specified, it must be specified on all endpoints, and no two endpoints can share the same priority value.
	Priority *int `pulumi:"priority"`
	// The name of the Traffic Manager profile.
	ProfileName string `pulumi:"profileName"`
	// The name of the resource group containing the Traffic Manager endpoint to be created or updated.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The fully-qualified DNS name or IP address of the endpoint. Traffic Manager returns this value in DNS responses to direct traffic to this endpoint.
	Target *string `pulumi:"target"`
	// The Azure Resource URI of the of the endpoint. Not applicable to endpoints of type 'ExternalEndpoints'.
	TargetResourceId *string `pulumi:"targetResourceId"`
	// The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
	Type *string `pulumi:"type"`
	// The weight of this endpoint when using the 'Weighted' traffic routing method. Possible values are from 1 to 1000.
	Weight *int `pulumi:"weight"`
}

// The set of arguments for constructing a Endpoint resource.
type EndpointArgs struct {
	// List of custom headers.
	CustomHeaders EndpointPropertiesCustomHeadersArrayInput
	// Specifies the location of the external or nested endpoints when using the ‘Performance’ traffic routing method.
	EndpointLocation pulumi.StringPtrInput
	// The monitoring status of the endpoint.
	EndpointMonitorStatus pulumi.StringPtrInput
	// The name of the Traffic Manager endpoint to be created or updated.
	EndpointName pulumi.StringInput
	// The status of the endpoint. If the endpoint is Enabled, it is probed for endpoint health and is included in the traffic routing method.
	EndpointStatus pulumi.StringPtrInput
	// The type of the Traffic Manager endpoint to be created or updated.
	EndpointType pulumi.StringInput
	// The list of countries/regions mapped to this endpoint when using the ‘Geographic’ traffic routing method. Please consult Traffic Manager Geographic documentation for a full list of accepted values.
	GeoMapping pulumi.StringArrayInput
	// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{resourceName}
	Id pulumi.StringPtrInput
	// The minimum number of endpoints that must be available in the child profile in order for the parent profile to be considered available. Only applicable to endpoint of type 'NestedEndpoints'.
	MinChildEndpoints pulumi.IntPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The priority of this endpoint when using the ‘Priority’ traffic routing method. Possible values are from 1 to 1000, lower values represent higher priority. This is an optional parameter.  If specified, it must be specified on all endpoints, and no two endpoints can share the same priority value.
	Priority pulumi.IntPtrInput
	// The name of the Traffic Manager profile.
	ProfileName pulumi.StringInput
	// The name of the resource group containing the Traffic Manager endpoint to be created or updated.
	ResourceGroupName pulumi.StringInput
	// The fully-qualified DNS name or IP address of the endpoint. Traffic Manager returns this value in DNS responses to direct traffic to this endpoint.
	Target pulumi.StringPtrInput
	// The Azure Resource URI of the of the endpoint. Not applicable to endpoints of type 'ExternalEndpoints'.
	TargetResourceId pulumi.StringPtrInput
	// The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
	Type pulumi.StringPtrInput
	// The weight of this endpoint when using the 'Weighted' traffic routing method. Possible values are from 1 to 1000.
	Weight pulumi.IntPtrInput
}

func (EndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointArgs)(nil)).Elem()
}
