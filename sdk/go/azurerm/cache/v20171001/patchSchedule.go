// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20171001

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Response to put/get patch schedules for Redis cache.
//
// ## Example Usage
// ### RedisCachePatchSchedulesCreateOrUpdate
//
// ```go
// package main
//
// import (
// 	cache "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/cache/v20171001"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cache.NewPatchSchedule(ctx, "patchSchedule", &cache.PatchScheduleArgs{
// 			Default:           pulumi.String("default"),
// 			Name:              pulumi.String("cache1"),
// 			ResourceGroupName: pulumi.String("rg1"),
// 			ScheduleEntries: cache.ScheduleEntryArray{
// 				&cache.ScheduleEntryArgs{
// 					DayOfWeek:         pulumi.String("Monday"),
// 					MaintenanceWindow: pulumi.String("PT5H"),
// 					StartHourUtc:      pulumi.Int(12),
// 				},
// 				&cache.ScheduleEntryArgs{
// 					DayOfWeek:    pulumi.String("Tuesday"),
// 					StartHourUtc: pulumi.Int(12),
// 				},
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
type PatchSchedule struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of patch schedules for a Redis cache.
	ScheduleEntries ScheduleEntryResponseArrayOutput `pulumi:"scheduleEntries"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPatchSchedule registers a new resource with the given unique name, arguments, and options.
func NewPatchSchedule(ctx *pulumi.Context,
	name string, args *PatchScheduleArgs, opts ...pulumi.ResourceOption) (*PatchSchedule, error) {
	if args == nil || args.Default == nil {
		return nil, errors.New("missing required argument 'Default'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ScheduleEntries == nil {
		return nil, errors.New("missing required argument 'ScheduleEntries'")
	}
	if args == nil {
		args = &PatchScheduleArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:cache/latest:PatchSchedule"),
		},
		{
			Type: pulumi.String("azurerm:cache/v20180301:PatchSchedule"),
		},
		{
			Type: pulumi.String("azurerm:cache/v20190701:PatchSchedule"),
		},
	})
	opts = append(opts, aliases)
	var resource PatchSchedule
	err := ctx.RegisterResource("azurerm:cache/v20171001:PatchSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPatchSchedule gets an existing PatchSchedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPatchSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PatchScheduleState, opts ...pulumi.ResourceOption) (*PatchSchedule, error) {
	var resource PatchSchedule
	err := ctx.ReadResource("azurerm:cache/v20171001:PatchSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PatchSchedule resources.
type patchScheduleState struct {
	// Resource name.
	Name *string `pulumi:"name"`
	// List of patch schedules for a Redis cache.
	ScheduleEntries []ScheduleEntryResponse `pulumi:"scheduleEntries"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type PatchScheduleState struct {
	// Resource name.
	Name pulumi.StringPtrInput
	// List of patch schedules for a Redis cache.
	ScheduleEntries ScheduleEntryResponseArrayInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (PatchScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*patchScheduleState)(nil)).Elem()
}

type patchScheduleArgs struct {
	// Default string modeled as parameter for auto generation to work correctly.
	Default string `pulumi:"default"`
	// The name of the Redis cache.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// List of patch schedules for a Redis cache.
	ScheduleEntries []ScheduleEntry `pulumi:"scheduleEntries"`
}

// The set of arguments for constructing a PatchSchedule resource.
type PatchScheduleArgs struct {
	// Default string modeled as parameter for auto generation to work correctly.
	Default pulumi.StringInput
	// The name of the Redis cache.
	Name pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// List of patch schedules for a Redis cache.
	ScheduleEntries ScheduleEntryArrayInput
}

func (PatchScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*patchScheduleArgs)(nil)).Elem()
}
