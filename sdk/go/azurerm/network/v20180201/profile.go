// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180201

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Class representing a Traffic Manager profile.
type Profile struct {
	pulumi.CustomResourceState

	// The Azure Region where the resource lives
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of the Traffic Manager profile.
	Properties ProfilePropertiesResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProfile registers a new resource with the given unique name, arguments, and options.
func NewProfile(ctx *pulumi.Context,
	name string, args *ProfileArgs, opts ...pulumi.ResourceOption) (*Profile, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ProfileArgs{}
	}
	var resource Profile
	err := ctx.RegisterResource("azurerm:network/v20180201:Profile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProfile gets an existing Profile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfileState, opts ...pulumi.ResourceOption) (*Profile, error) {
	var resource Profile
	err := ctx.ReadResource("azurerm:network/v20180201:Profile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Profile resources.
type profileState struct {
	// The Azure Region where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The properties of the Traffic Manager profile.
	Properties *ProfilePropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
	Type *string `pulumi:"type"`
}

type ProfileState struct {
	// The Azure Region where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The properties of the Traffic Manager profile.
	Properties ProfilePropertiesResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
	Type pulumi.StringPtrInput
}

func (ProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*profileState)(nil)).Elem()
}

type profileArgs struct {
	// The DNS settings of the Traffic Manager profile.
	DnsConfig *DnsConfig `pulumi:"dnsConfig"`
	// The list of endpoints in the Traffic Manager profile.
	Endpoints []EndpointType `pulumi:"endpoints"`
	// The Azure Region where the resource lives
	Location *string `pulumi:"location"`
	// The endpoint monitoring settings of the Traffic Manager profile.
	MonitorConfig *MonitorConfig `pulumi:"monitorConfig"`
	// The name of the Traffic Manager profile.
	Name string `pulumi:"name"`
	// The status of the Traffic Manager profile.
	ProfileStatus *string `pulumi:"profileStatus"`
	// The name of the resource group containing the Traffic Manager profile.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The traffic routing method of the Traffic Manager profile.
	TrafficRoutingMethod *string `pulumi:"trafficRoutingMethod"`
	// Indicates whether Traffic View is 'Enabled' or 'Disabled' for the Traffic Manager profile. Null, indicates 'Disabled'. Enabling this feature will increase the cost of the Traffic Manage profile.
	TrafficViewEnrollmentStatus *string `pulumi:"trafficViewEnrollmentStatus"`
}

// The set of arguments for constructing a Profile resource.
type ProfileArgs struct {
	// The DNS settings of the Traffic Manager profile.
	DnsConfig DnsConfigPtrInput
	// The list of endpoints in the Traffic Manager profile.
	Endpoints EndpointTypeArrayInput
	// The Azure Region where the resource lives
	Location pulumi.StringPtrInput
	// The endpoint monitoring settings of the Traffic Manager profile.
	MonitorConfig MonitorConfigPtrInput
	// The name of the Traffic Manager profile.
	Name pulumi.StringInput
	// The status of the Traffic Manager profile.
	ProfileStatus pulumi.StringPtrInput
	// The name of the resource group containing the Traffic Manager profile.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The traffic routing method of the Traffic Manager profile.
	TrafficRoutingMethod pulumi.StringPtrInput
	// Indicates whether Traffic View is 'Enabled' or 'Disabled' for the Traffic Manager profile. Null, indicates 'Disabled'. Enabling this feature will increase the cost of the Traffic Manage profile.
	TrafficViewEnrollmentStatus pulumi.StringPtrInput
}

func (ProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*profileArgs)(nil)).Elem()
}
