// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190701

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Specifies information about the dedicated host group that the dedicated hosts should be assigned to. <br><br> Currently, a dedicated host can only be added to a dedicated host group at creation time. An existing dedicated host cannot be added to another dedicated host group.
//
// ## Example Usage
// ### Create or update a dedicated host group.
//
// ```go
// package main
//
// import (
// 	compute "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/compute/v20190701"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewDedicatedHostGroup(ctx, "dedicatedHostGroup", &compute.DedicatedHostGroupArgs{
// 			HostGroupName:            pulumi.String("myDedicatedHostGroup"),
// 			Location:                 pulumi.String("westus"),
// 			PlatformFaultDomainCount: pulumi.Int(3),
// 			ResourceGroupName:        pulumi.String("myResourceGroup"),
// 			Tags: pulumi.StringMap{
// 				"department": pulumi.String("finance"),
// 			},
// 			Zones: pulumi.StringArray{
// 				pulumi.String("1"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type DedicatedHostGroup struct {
	pulumi.CustomResourceState

	// A list of references to all dedicated hosts in the dedicated host group.
	Hosts SubResourceReadOnlyResponseArrayOutput `pulumi:"hosts"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Number of fault domains that the host group can span.
	PlatformFaultDomainCount pulumi.IntOutput `pulumi:"platformFaultDomainCount"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Availability Zone to use for this host group. Only single zone is supported. The zone can be assigned only during creation. If not provided, the group supports all zones in the region. If provided, enforces each host in the group to be in the same zone.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewDedicatedHostGroup registers a new resource with the given unique name, arguments, and options.
func NewDedicatedHostGroup(ctx *pulumi.Context,
	name string, args *DedicatedHostGroupArgs, opts ...pulumi.ResourceOption) (*DedicatedHostGroup, error) {
	if args == nil || args.HostGroupName == nil {
		return nil, errors.New("missing required argument 'HostGroupName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.PlatformFaultDomainCount == nil {
		return nil, errors.New("missing required argument 'PlatformFaultDomainCount'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &DedicatedHostGroupArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:compute/latest:DedicatedHostGroup"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20190301:DedicatedHostGroup"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20191201:DedicatedHostGroup"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20200601:DedicatedHostGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource DedicatedHostGroup
	err := ctx.RegisterResource("azurerm:compute/v20190701:DedicatedHostGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDedicatedHostGroup gets an existing DedicatedHostGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDedicatedHostGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DedicatedHostGroupState, opts ...pulumi.ResourceOption) (*DedicatedHostGroup, error) {
	var resource DedicatedHostGroup
	err := ctx.ReadResource("azurerm:compute/v20190701:DedicatedHostGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DedicatedHostGroup resources.
type dedicatedHostGroupState struct {
	// A list of references to all dedicated hosts in the dedicated host group.
	Hosts []SubResourceReadOnlyResponse `pulumi:"hosts"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Number of fault domains that the host group can span.
	PlatformFaultDomainCount *int `pulumi:"platformFaultDomainCount"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
	// Availability Zone to use for this host group. Only single zone is supported. The zone can be assigned only during creation. If not provided, the group supports all zones in the region. If provided, enforces each host in the group to be in the same zone.
	Zones []string `pulumi:"zones"`
}

type DedicatedHostGroupState struct {
	// A list of references to all dedicated hosts in the dedicated host group.
	Hosts SubResourceReadOnlyResponseArrayInput
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Number of fault domains that the host group can span.
	PlatformFaultDomainCount pulumi.IntPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
	// Availability Zone to use for this host group. Only single zone is supported. The zone can be assigned only during creation. If not provided, the group supports all zones in the region. If provided, enforces each host in the group to be in the same zone.
	Zones pulumi.StringArrayInput
}

func (DedicatedHostGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedHostGroupState)(nil)).Elem()
}

type dedicatedHostGroupArgs struct {
	// The name of the dedicated host group.
	HostGroupName string `pulumi:"hostGroupName"`
	// Resource location
	Location string `pulumi:"location"`
	// Number of fault domains that the host group can span.
	PlatformFaultDomainCount int `pulumi:"platformFaultDomainCount"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Availability Zone to use for this host group. Only single zone is supported. The zone can be assigned only during creation. If not provided, the group supports all zones in the region. If provided, enforces each host in the group to be in the same zone.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a DedicatedHostGroup resource.
type DedicatedHostGroupArgs struct {
	// The name of the dedicated host group.
	HostGroupName pulumi.StringInput
	// Resource location
	Location pulumi.StringInput
	// Number of fault domains that the host group can span.
	PlatformFaultDomainCount pulumi.IntInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Availability Zone to use for this host group. Only single zone is supported. The zone can be assigned only during creation. If not provided, the group supports all zones in the region. If provided, enforces each host in the group to be in the same zone.
	Zones pulumi.StringArrayInput
}

func (DedicatedHostGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedHostGroupArgs)(nil)).Elem()
}
