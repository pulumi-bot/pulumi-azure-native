// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Specifies information about the dedicated host group that the dedicated hosts should be assigned to. <br><br> Currently, a dedicated host can only be added to a dedicated host group at creation time. An existing dedicated host cannot be added to another dedicated host group.
type HostGroup struct {
	pulumi.CustomResourceState

	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Dedicated Host Group Properties.
	Properties DedicatedHostGroupPropertiesResponseOutput `pulumi:"properties"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Availability Zone to use for this host group. Only single zone is supported. The zone can be assigned only during creation. If not provided, the group supports all zones in the region. If provided, enforces each host in the group to be in the same zone.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewHostGroup registers a new resource with the given unique name, arguments, and options.
func NewHostGroup(ctx *pulumi.Context,
	name string, args *HostGroupArgs, opts ...pulumi.ResourceOption) (*HostGroup, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &HostGroupArgs{}
	}
	var resource HostGroup
	err := ctx.RegisterResource("azurerm:compute:HostGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHostGroup gets an existing HostGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHostGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostGroupState, opts ...pulumi.ResourceOption) (*HostGroup, error) {
	var resource HostGroup
	err := ctx.ReadResource("azurerm:compute:HostGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HostGroup resources.
type hostGroupState struct {
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Dedicated Host Group Properties.
	Properties *DedicatedHostGroupPropertiesResponse `pulumi:"properties"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
	// Availability Zone to use for this host group. Only single zone is supported. The zone can be assigned only during creation. If not provided, the group supports all zones in the region. If provided, enforces each host in the group to be in the same zone.
	Zones []string `pulumi:"zones"`
}

type HostGroupState struct {
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Dedicated Host Group Properties.
	Properties DedicatedHostGroupPropertiesResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
	// Availability Zone to use for this host group. Only single zone is supported. The zone can be assigned only during creation. If not provided, the group supports all zones in the region. If provided, enforces each host in the group to be in the same zone.
	Zones pulumi.StringArrayInput
}

func (HostGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostGroupState)(nil)).Elem()
}

type hostGroupArgs struct {
	// Resource location
	Location string `pulumi:"location"`
	// The name of the dedicated host group.
	Name string `pulumi:"name"`
	// Dedicated Host Group Properties.
	Properties *DedicatedHostGroupProperties `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Availability Zone to use for this host group. Only single zone is supported. The zone can be assigned only during creation. If not provided, the group supports all zones in the region. If provided, enforces each host in the group to be in the same zone.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a HostGroup resource.
type HostGroupArgs struct {
	// Resource location
	Location pulumi.StringInput
	// The name of the dedicated host group.
	Name pulumi.StringInput
	// Dedicated Host Group Properties.
	Properties DedicatedHostGroupPropertiesPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Availability Zone to use for this host group. Only single zone is supported. The zone can be assigned only during creation. If not provided, the group supports all zones in the region. If provided, enforces each host in the group to be in the same zone.
	Zones pulumi.StringArrayInput
}

func (HostGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostGroupArgs)(nil)).Elem()
}