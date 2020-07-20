// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Definition of the schedule.
type AutomationAccountSchedule struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets or sets the properties of the schedule.
	Properties SchedulePropertiesResponseOutput `pulumi:"properties"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAutomationAccountSchedule registers a new resource with the given unique name, arguments, and options.
func NewAutomationAccountSchedule(ctx *pulumi.Context,
	name string, args *AutomationAccountScheduleArgs, opts ...pulumi.ResourceOption) (*AutomationAccountSchedule, error) {
	if args == nil || args.AutomationAccountName == nil {
		return nil, errors.New("missing required argument 'AutomationAccountName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.Properties == nil {
		return nil, errors.New("missing required argument 'Properties'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &AutomationAccountScheduleArgs{}
	}
	var resource AutomationAccountSchedule
	err := ctx.RegisterResource("azurerm:automation:AutomationAccountSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutomationAccountSchedule gets an existing AutomationAccountSchedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutomationAccountSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutomationAccountScheduleState, opts ...pulumi.ResourceOption) (*AutomationAccountSchedule, error) {
	var resource AutomationAccountSchedule
	err := ctx.ReadResource("azurerm:automation:AutomationAccountSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AutomationAccountSchedule resources.
type automationAccountScheduleState struct {
	// The name of the resource
	Name *string `pulumi:"name"`
	// Gets or sets the properties of the schedule.
	Properties *SchedulePropertiesResponse `pulumi:"properties"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type AutomationAccountScheduleState struct {
	// The name of the resource
	Name pulumi.StringPtrInput
	// Gets or sets the properties of the schedule.
	Properties SchedulePropertiesResponsePtrInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (AutomationAccountScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*automationAccountScheduleState)(nil)).Elem()
}

type automationAccountScheduleArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The schedule name.
	Name string `pulumi:"name"`
	// Gets or sets the list of schedule properties.
	Properties ScheduleCreateOrUpdateProperties `pulumi:"properties"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a AutomationAccountSchedule resource.
type AutomationAccountScheduleArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput
	// The schedule name.
	Name pulumi.StringInput
	// Gets or sets the list of schedule properties.
	Properties ScheduleCreateOrUpdatePropertiesInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
}

func (AutomationAccountScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*automationAccountScheduleArgs)(nil)).Elem()
}
