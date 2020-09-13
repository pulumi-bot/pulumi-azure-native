// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Properties that define an Analytics item that is associated to an Application Insights component.
//
// ## Example Usage
// ### AnalyticsItemPut
//
// ```go
// package main
//
// import (
// 	insights "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/insights/latest"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := insights.NewAnalyticsItem(ctx, "analyticsItem", &insights.AnalyticsItemArgs{
// 			Content:           pulumi.String("let newExceptionsTimeRange = 1d;\nlet timeRangeToCheckBefore = 7d;\nexceptions\n| where timestamp < ago(timeRangeToCheckBefore)\n| summarize count() by problemId\n| join kind= rightanti (\nexceptions\n| where timestamp >= ago(newExceptionsTimeRange)\n| extend stack = tostring(details[0].rawStack)\n| summarize count(), dcount(user_AuthenticatedId), min(timestamp), max(timestamp), any(stack) by problemId  \n) on problemId \n| order by  count_ desc\n"),
// 			Name:              pulumi.String("Exceptions - New in the last 24 hours"),
// 			ResourceGroupName: pulumi.String("my-resource-group"),
// 			ResourceName:      pulumi.String("my-component"),
// 			Scope:             pulumi.String("shared"),
// 			ScopePath:         pulumi.String("analyticsItems"),
// 			Type:              pulumi.String("query"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type AnalyticsItem struct {
	pulumi.CustomResourceState

	// The content of this item
	Content pulumi.StringPtrOutput `pulumi:"content"`
	// The user-defined name of the item.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// A set of properties that can be defined in the context of a specific item type. Each type may have its own properties.
	Properties ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput `pulumi:"properties"`
	// Enum indicating if this item definition is owned by a specific user or is shared between all users with access to the Application Insights component.
	Scope pulumi.StringPtrOutput `pulumi:"scope"`
	// Date and time in UTC when this item was created.
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// Date and time in UTC of the last modification that was made to this item.
	TimeModified pulumi.StringOutput `pulumi:"timeModified"`
	// Enum indicating the type of the Analytics item.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// This instance's version of the data model. This can change as new features are added.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewAnalyticsItem registers a new resource with the given unique name, arguments, and options.
func NewAnalyticsItem(ctx *pulumi.Context,
	name string, args *AnalyticsItemArgs, opts ...pulumi.ResourceOption) (*AnalyticsItem, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ResourceName == nil {
		return nil, errors.New("missing required argument 'ResourceName'")
	}
	if args == nil || args.ScopePath == nil {
		return nil, errors.New("missing required argument 'ScopePath'")
	}
	if args == nil {
		args = &AnalyticsItemArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:insights/v20150501:AnalyticsItem"),
		},
	})
	opts = append(opts, aliases)
	var resource AnalyticsItem
	err := ctx.RegisterResource("azurerm:insights/latest:AnalyticsItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnalyticsItem gets an existing AnalyticsItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnalyticsItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnalyticsItemState, opts ...pulumi.ResourceOption) (*AnalyticsItem, error) {
	var resource AnalyticsItem
	err := ctx.ReadResource("azurerm:insights/latest:AnalyticsItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AnalyticsItem resources.
type analyticsItemState struct {
	// The content of this item
	Content *string `pulumi:"content"`
	// The user-defined name of the item.
	Name *string `pulumi:"name"`
	// A set of properties that can be defined in the context of a specific item type. Each type may have its own properties.
	Properties *ApplicationInsightsComponentAnalyticsItemPropertiesResponse `pulumi:"properties"`
	// Enum indicating if this item definition is owned by a specific user or is shared between all users with access to the Application Insights component.
	Scope *string `pulumi:"scope"`
	// Date and time in UTC when this item was created.
	TimeCreated *string `pulumi:"timeCreated"`
	// Date and time in UTC of the last modification that was made to this item.
	TimeModified *string `pulumi:"timeModified"`
	// Enum indicating the type of the Analytics item.
	Type *string `pulumi:"type"`
	// This instance's version of the data model. This can change as new features are added.
	Version *string `pulumi:"version"`
}

type AnalyticsItemState struct {
	// The content of this item
	Content pulumi.StringPtrInput
	// The user-defined name of the item.
	Name pulumi.StringPtrInput
	// A set of properties that can be defined in the context of a specific item type. Each type may have its own properties.
	Properties ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrInput
	// Enum indicating if this item definition is owned by a specific user or is shared between all users with access to the Application Insights component.
	Scope pulumi.StringPtrInput
	// Date and time in UTC when this item was created.
	TimeCreated pulumi.StringPtrInput
	// Date and time in UTC of the last modification that was made to this item.
	TimeModified pulumi.StringPtrInput
	// Enum indicating the type of the Analytics item.
	Type pulumi.StringPtrInput
	// This instance's version of the data model. This can change as new features are added.
	Version pulumi.StringPtrInput
}

func (AnalyticsItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*analyticsItemState)(nil)).Elem()
}

type analyticsItemArgs struct {
	// The content of this item
	Content *string `pulumi:"content"`
	// Internally assigned unique id of the item definition.
	Id *string `pulumi:"id"`
	// The user-defined name of the item.
	Name *string `pulumi:"name"`
	// Flag indicating whether or not to force save an item. This allows overriding an item if it already exists.
	OverrideItem *bool `pulumi:"overrideItem"`
	// A set of properties that can be defined in the context of a specific item type. Each type may have its own properties.
	Properties *ApplicationInsightsComponentAnalyticsItemProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Application Insights component resource.
	ResourceName string `pulumi:"resourceName"`
	// Enum indicating if this item definition is owned by a specific user or is shared between all users with access to the Application Insights component.
	Scope *string `pulumi:"scope"`
	// Enum indicating if this item definition is owned by a specific user or is shared between all users with access to the Application Insights component.
	ScopePath string `pulumi:"scopePath"`
	// Enum indicating the type of the Analytics item.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a AnalyticsItem resource.
type AnalyticsItemArgs struct {
	// The content of this item
	Content pulumi.StringPtrInput
	// Internally assigned unique id of the item definition.
	Id pulumi.StringPtrInput
	// The user-defined name of the item.
	Name pulumi.StringPtrInput
	// Flag indicating whether or not to force save an item. This allows overriding an item if it already exists.
	OverrideItem pulumi.BoolPtrInput
	// A set of properties that can be defined in the context of a specific item type. Each type may have its own properties.
	Properties ApplicationInsightsComponentAnalyticsItemPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the Application Insights component resource.
	ResourceName pulumi.StringInput
	// Enum indicating if this item definition is owned by a specific user or is shared between all users with access to the Application Insights component.
	Scope pulumi.StringPtrInput
	// Enum indicating if this item definition is owned by a specific user or is shared between all users with access to the Application Insights component.
	ScopePath pulumi.StringInput
	// Enum indicating the type of the Analytics item.
	Type pulumi.StringPtrInput
}

func (AnalyticsItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*analyticsItemArgs)(nil)).Elem()
}
