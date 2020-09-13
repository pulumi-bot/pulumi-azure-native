// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Trigger resource type.
//
// ## Example Usage
// ### Triggers_Create
//
// ```go
// package main
//
// import (
// 	datafactory "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/datafactory/v20180601"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := datafactory.NewTrigger(ctx, "trigger", &datafactory.TriggerArgs{
// 			FactoryName:       pulumi.String("exampleFactoryName"),
// 			ResourceGroupName: pulumi.String("exampleResourceGroup"),
// 			TriggerName:       pulumi.String("exampleTrigger"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Triggers_Update
//
// ```go
// package main
//
// import (
// 	datafactory "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/datafactory/v20180601"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := datafactory.NewTrigger(ctx, "trigger", &datafactory.TriggerArgs{
// 			FactoryName:       pulumi.String("exampleFactoryName"),
// 			ResourceGroupName: pulumi.String("exampleResourceGroup"),
// 			TriggerName:       pulumi.String("exampleTrigger"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Trigger struct {
	pulumi.CustomResourceState

	// Etag identifies change in the resource.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the trigger.
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTrigger registers a new resource with the given unique name, arguments, and options.
func NewTrigger(ctx *pulumi.Context,
	name string, args *TriggerArgs, opts ...pulumi.ResourceOption) (*Trigger, error) {
	if args == nil || args.FactoryName == nil {
		return nil, errors.New("missing required argument 'FactoryName'")
	}
	if args == nil || args.Properties == nil {
		return nil, errors.New("missing required argument 'Properties'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.TriggerName == nil {
		return nil, errors.New("missing required argument 'TriggerName'")
	}
	if args == nil {
		args = &TriggerArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:datafactory/latest:Trigger"),
		},
		{
			Type: pulumi.String("azurerm:datafactory/v20170901preview:Trigger"),
		},
	})
	opts = append(opts, aliases)
	var resource Trigger
	err := ctx.RegisterResource("azurerm:datafactory/v20180601:Trigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrigger gets an existing Trigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TriggerState, opts ...pulumi.ResourceOption) (*Trigger, error) {
	var resource Trigger
	err := ctx.ReadResource("azurerm:datafactory/v20180601:Trigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Trigger resources.
type triggerState struct {
	// Etag identifies change in the resource.
	Etag *string `pulumi:"etag"`
	// The resource name.
	Name *string `pulumi:"name"`
	// Properties of the trigger.
	Properties interface{} `pulumi:"properties"`
	// The resource type.
	Type *string `pulumi:"type"`
}

type TriggerState struct {
	// Etag identifies change in the resource.
	Etag pulumi.StringPtrInput
	// The resource name.
	Name pulumi.StringPtrInput
	// Properties of the trigger.
	Properties pulumi.Input
	// The resource type.
	Type pulumi.StringPtrInput
}

func (TriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerState)(nil)).Elem()
}

type triggerArgs struct {
	// The factory name.
	FactoryName string `pulumi:"factoryName"`
	// Properties of the trigger.
	Properties interface{} `pulumi:"properties"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The trigger name.
	TriggerName string `pulumi:"triggerName"`
}

// The set of arguments for constructing a Trigger resource.
type TriggerArgs struct {
	// The factory name.
	FactoryName pulumi.StringInput
	// Properties of the trigger.
	Properties pulumi.Input
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The trigger name.
	TriggerName pulumi.StringInput
}

func (TriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerArgs)(nil)).Elem()
}
