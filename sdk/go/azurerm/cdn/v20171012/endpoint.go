// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20171012

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// CDN endpoint is the entity within a CDN profile containing configuration information such as origin, protocol, content caching and delivery behavior. The CDN endpoint uses the URL format <endpointname>.azureedge.net.
type Endpoint struct {
	pulumi.CustomResourceState

	// List of content types on which compression applies. The value should be a valid MIME type.
	ContentTypesToCompress pulumi.StringArrayOutput `pulumi:"contentTypesToCompress"`
	// A policy that specifies the delivery rules to be used for an endpoint.
	DeliveryPolicy EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput `pulumi:"deliveryPolicy"`
	// List of rules defining the user's geo access within a CDN endpoint. Each geo filter defines an access rule to a specified path or content, e.g. block APAC for path /pictures/
	GeoFilters GeoFilterResponseArrayOutput `pulumi:"geoFilters"`
	// The host name of the endpoint structured as {endpointName}.{DNSZone}, e.g. contoso.azureedge.net
	HostName pulumi.StringOutput `pulumi:"hostName"`
	// Indicates whether content compression is enabled on CDN. Default value is false. If compression is enabled, content will be served as compressed if user requests for a compressed version. Content won't be compressed on CDN when requested content is smaller than 1 byte or larger than 1 MB.
	IsCompressionEnabled pulumi.BoolPtrOutput `pulumi:"isCompressionEnabled"`
	// Indicates whether HTTP traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
	IsHttpAllowed pulumi.BoolPtrOutput `pulumi:"isHttpAllowed"`
	// Indicates whether HTTPS traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
	IsHttpsAllowed pulumi.BoolPtrOutput `pulumi:"isHttpsAllowed"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies what scenario the customer wants this CDN endpoint to optimize for, e.g. Download, Media services. With this information, CDN can apply scenario driven optimization.
	OptimizationType pulumi.StringPtrOutput `pulumi:"optimizationType"`
	// The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default.
	OriginHostHeader pulumi.StringPtrOutput `pulumi:"originHostHeader"`
	// A directory path on the origin that CDN can use to retrieve content from, e.g. contoso.cloudapp.net/originpath.
	OriginPath pulumi.StringPtrOutput `pulumi:"originPath"`
	// The source of the content being delivered via CDN.
	Origins DeepCreatedOriginResponseArrayOutput `pulumi:"origins"`
	// Path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the origin path.
	ProbePath pulumi.StringPtrOutput `pulumi:"probePath"`
	// Provisioning status of the endpoint.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Defines how CDN caches requests that include query strings. You can ignore any query strings when caching, bypass caching to prevent requests that contain query strings from being cached, or cache every request with a unique URL.
	QueryStringCachingBehavior pulumi.StringPtrOutput `pulumi:"queryStringCachingBehavior"`
	// Resource status of the endpoint.
	ResourceState pulumi.StringOutput `pulumi:"resourceState"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEndpoint registers a new resource with the given unique name, arguments, and options.
func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	if args == nil || args.EndpointName == nil {
		return nil, errors.New("missing required argument 'EndpointName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Origins == nil {
		return nil, errors.New("missing required argument 'Origins'")
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
			Type: pulumi.String("azurerm:cdn/latest:Endpoint"),
		},
		{
			Type: pulumi.String("azurerm:cdn/v20150601:Endpoint"),
		},
		{
			Type: pulumi.String("azurerm:cdn/v20160402:Endpoint"),
		},
		{
			Type: pulumi.String("azurerm:cdn/v20161002:Endpoint"),
		},
		{
			Type: pulumi.String("azurerm:cdn/v20170402:Endpoint"),
		},
		{
			Type: pulumi.String("azurerm:cdn/v20190415:Endpoint"),
		},
		{
			Type: pulumi.String("azurerm:cdn/v20190615:Endpoint"),
		},
		{
			Type: pulumi.String("azurerm:cdn/v20190615preview:Endpoint"),
		},
		{
			Type: pulumi.String("azurerm:cdn/v20191231:Endpoint"),
		},
		{
			Type: pulumi.String("azurerm:cdn/v20200331:Endpoint"),
		},
		{
			Type: pulumi.String("azurerm:cdn/v20200415:Endpoint"),
		},
	})
	opts = append(opts, aliases)
	var resource Endpoint
	err := ctx.RegisterResource("azurerm:cdn/v20171012:Endpoint", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:cdn/v20171012:Endpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Endpoint resources.
type endpointState struct {
	// List of content types on which compression applies. The value should be a valid MIME type.
	ContentTypesToCompress []string `pulumi:"contentTypesToCompress"`
	// A policy that specifies the delivery rules to be used for an endpoint.
	DeliveryPolicy *EndpointPropertiesUpdateParametersResponseDeliveryPolicy `pulumi:"deliveryPolicy"`
	// List of rules defining the user's geo access within a CDN endpoint. Each geo filter defines an access rule to a specified path or content, e.g. block APAC for path /pictures/
	GeoFilters []GeoFilterResponse `pulumi:"geoFilters"`
	// The host name of the endpoint structured as {endpointName}.{DNSZone}, e.g. contoso.azureedge.net
	HostName *string `pulumi:"hostName"`
	// Indicates whether content compression is enabled on CDN. Default value is false. If compression is enabled, content will be served as compressed if user requests for a compressed version. Content won't be compressed on CDN when requested content is smaller than 1 byte or larger than 1 MB.
	IsCompressionEnabled *bool `pulumi:"isCompressionEnabled"`
	// Indicates whether HTTP traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
	IsHttpAllowed *bool `pulumi:"isHttpAllowed"`
	// Indicates whether HTTPS traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
	IsHttpsAllowed *bool `pulumi:"isHttpsAllowed"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Specifies what scenario the customer wants this CDN endpoint to optimize for, e.g. Download, Media services. With this information, CDN can apply scenario driven optimization.
	OptimizationType *string `pulumi:"optimizationType"`
	// The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default.
	OriginHostHeader *string `pulumi:"originHostHeader"`
	// A directory path on the origin that CDN can use to retrieve content from, e.g. contoso.cloudapp.net/originpath.
	OriginPath *string `pulumi:"originPath"`
	// The source of the content being delivered via CDN.
	Origins []DeepCreatedOriginResponse `pulumi:"origins"`
	// Path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the origin path.
	ProbePath *string `pulumi:"probePath"`
	// Provisioning status of the endpoint.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Defines how CDN caches requests that include query strings. You can ignore any query strings when caching, bypass caching to prevent requests that contain query strings from being cached, or cache every request with a unique URL.
	QueryStringCachingBehavior *string `pulumi:"queryStringCachingBehavior"`
	// Resource status of the endpoint.
	ResourceState *string `pulumi:"resourceState"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type EndpointState struct {
	// List of content types on which compression applies. The value should be a valid MIME type.
	ContentTypesToCompress pulumi.StringArrayInput
	// A policy that specifies the delivery rules to be used for an endpoint.
	DeliveryPolicy EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrInput
	// List of rules defining the user's geo access within a CDN endpoint. Each geo filter defines an access rule to a specified path or content, e.g. block APAC for path /pictures/
	GeoFilters GeoFilterResponseArrayInput
	// The host name of the endpoint structured as {endpointName}.{DNSZone}, e.g. contoso.azureedge.net
	HostName pulumi.StringPtrInput
	// Indicates whether content compression is enabled on CDN. Default value is false. If compression is enabled, content will be served as compressed if user requests for a compressed version. Content won't be compressed on CDN when requested content is smaller than 1 byte or larger than 1 MB.
	IsCompressionEnabled pulumi.BoolPtrInput
	// Indicates whether HTTP traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
	IsHttpAllowed pulumi.BoolPtrInput
	// Indicates whether HTTPS traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
	IsHttpsAllowed pulumi.BoolPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Specifies what scenario the customer wants this CDN endpoint to optimize for, e.g. Download, Media services. With this information, CDN can apply scenario driven optimization.
	OptimizationType pulumi.StringPtrInput
	// The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default.
	OriginHostHeader pulumi.StringPtrInput
	// A directory path on the origin that CDN can use to retrieve content from, e.g. contoso.cloudapp.net/originpath.
	OriginPath pulumi.StringPtrInput
	// The source of the content being delivered via CDN.
	Origins DeepCreatedOriginResponseArrayInput
	// Path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the origin path.
	ProbePath pulumi.StringPtrInput
	// Provisioning status of the endpoint.
	ProvisioningState pulumi.StringPtrInput
	// Defines how CDN caches requests that include query strings. You can ignore any query strings when caching, bypass caching to prevent requests that contain query strings from being cached, or cache every request with a unique URL.
	QueryStringCachingBehavior pulumi.StringPtrInput
	// Resource status of the endpoint.
	ResourceState pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (EndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointState)(nil)).Elem()
}

type endpointArgs struct {
	// List of content types on which compression applies. The value should be a valid MIME type.
	ContentTypesToCompress []string `pulumi:"contentTypesToCompress"`
	// A policy that specifies the delivery rules to be used for an endpoint.
	DeliveryPolicy *EndpointPropertiesUpdateParametersDeliveryPolicy `pulumi:"deliveryPolicy"`
	// Name of the endpoint under the profile which is unique globally.
	EndpointName string `pulumi:"endpointName"`
	// List of rules defining the user's geo access within a CDN endpoint. Each geo filter defines an access rule to a specified path or content, e.g. block APAC for path /pictures/
	GeoFilters []GeoFilter `pulumi:"geoFilters"`
	// Indicates whether content compression is enabled on CDN. Default value is false. If compression is enabled, content will be served as compressed if user requests for a compressed version. Content won't be compressed on CDN when requested content is smaller than 1 byte or larger than 1 MB.
	IsCompressionEnabled *bool `pulumi:"isCompressionEnabled"`
	// Indicates whether HTTP traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
	IsHttpAllowed *bool `pulumi:"isHttpAllowed"`
	// Indicates whether HTTPS traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
	IsHttpsAllowed *bool `pulumi:"isHttpsAllowed"`
	// Resource location.
	Location string `pulumi:"location"`
	// Specifies what scenario the customer wants this CDN endpoint to optimize for, e.g. Download, Media services. With this information, CDN can apply scenario driven optimization.
	OptimizationType *string `pulumi:"optimizationType"`
	// The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default.
	OriginHostHeader *string `pulumi:"originHostHeader"`
	// A directory path on the origin that CDN can use to retrieve content from, e.g. contoso.cloudapp.net/originpath.
	OriginPath *string `pulumi:"originPath"`
	// The source of the content being delivered via CDN.
	Origins []DeepCreatedOrigin `pulumi:"origins"`
	// Path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the origin path.
	ProbePath *string `pulumi:"probePath"`
	// Name of the CDN profile which is unique within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Defines how CDN caches requests that include query strings. You can ignore any query strings when caching, bypass caching to prevent requests that contain query strings from being cached, or cache every request with a unique URL.
	QueryStringCachingBehavior *string `pulumi:"queryStringCachingBehavior"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Endpoint resource.
type EndpointArgs struct {
	// List of content types on which compression applies. The value should be a valid MIME type.
	ContentTypesToCompress pulumi.StringArrayInput
	// A policy that specifies the delivery rules to be used for an endpoint.
	DeliveryPolicy EndpointPropertiesUpdateParametersDeliveryPolicyPtrInput
	// Name of the endpoint under the profile which is unique globally.
	EndpointName pulumi.StringInput
	// List of rules defining the user's geo access within a CDN endpoint. Each geo filter defines an access rule to a specified path or content, e.g. block APAC for path /pictures/
	GeoFilters GeoFilterArrayInput
	// Indicates whether content compression is enabled on CDN. Default value is false. If compression is enabled, content will be served as compressed if user requests for a compressed version. Content won't be compressed on CDN when requested content is smaller than 1 byte or larger than 1 MB.
	IsCompressionEnabled pulumi.BoolPtrInput
	// Indicates whether HTTP traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
	IsHttpAllowed pulumi.BoolPtrInput
	// Indicates whether HTTPS traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
	IsHttpsAllowed pulumi.BoolPtrInput
	// Resource location.
	Location pulumi.StringInput
	// Specifies what scenario the customer wants this CDN endpoint to optimize for, e.g. Download, Media services. With this information, CDN can apply scenario driven optimization.
	OptimizationType pulumi.StringPtrInput
	// The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default.
	OriginHostHeader pulumi.StringPtrInput
	// A directory path on the origin that CDN can use to retrieve content from, e.g. contoso.cloudapp.net/originpath.
	OriginPath pulumi.StringPtrInput
	// The source of the content being delivered via CDN.
	Origins DeepCreatedOriginArrayInput
	// Path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the origin path.
	ProbePath pulumi.StringPtrInput
	// Name of the CDN profile which is unique within the resource group.
	ProfileName pulumi.StringInput
	// Defines how CDN caches requests that include query strings. You can ignore any query strings when caching, bypass caching to prevent requests that contain query strings from being cached, or cache every request with a unique URL.
	QueryStringCachingBehavior pulumi.StringPtrInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (EndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointArgs)(nil)).Elem()
}
