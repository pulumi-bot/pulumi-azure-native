// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181001

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A budget resource.
type BudgetByResourceGroupName struct {
	pulumi.CustomResourceState

	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of the budget.
	Properties BudgetPropertiesResponseOutput `pulumi:"properties"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBudgetByResourceGroupName registers a new resource with the given unique name, arguments, and options.
func NewBudgetByResourceGroupName(ctx *pulumi.Context,
	name string, args *BudgetByResourceGroupNameArgs, opts ...pulumi.ResourceOption) (*BudgetByResourceGroupName, error) {
	if args == nil || args.Amount == nil {
		return nil, errors.New("missing required argument 'Amount'")
	}
	if args == nil || args.Category == nil {
		return nil, errors.New("missing required argument 'Category'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.TimeGrain == nil {
		return nil, errors.New("missing required argument 'TimeGrain'")
	}
	if args == nil || args.TimePeriod == nil {
		return nil, errors.New("missing required argument 'TimePeriod'")
	}
	if args == nil {
		args = &BudgetByResourceGroupNameArgs{}
	}
	var resource BudgetByResourceGroupName
	err := ctx.RegisterResource("azurerm:consumption/v20181001:BudgetByResourceGroupName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBudgetByResourceGroupName gets an existing BudgetByResourceGroupName resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBudgetByResourceGroupName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BudgetByResourceGroupNameState, opts ...pulumi.ResourceOption) (*BudgetByResourceGroupName, error) {
	var resource BudgetByResourceGroupName
	err := ctx.ReadResource("azurerm:consumption/v20181001:BudgetByResourceGroupName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BudgetByResourceGroupName resources.
type budgetByResourceGroupNameState struct {
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag *string `pulumi:"eTag"`
	// Resource name.
	Name *string `pulumi:"name"`
	// The properties of the budget.
	Properties *BudgetPropertiesResponse `pulumi:"properties"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type BudgetByResourceGroupNameState struct {
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// The properties of the budget.
	Properties BudgetPropertiesResponsePtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (BudgetByResourceGroupNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*budgetByResourceGroupNameState)(nil)).Elem()
}

type budgetByResourceGroupNameArgs struct {
	// The total amount of cost to track with the budget
	Amount float64 `pulumi:"amount"`
	// The category of the budget, whether the budget tracks cost or usage.
	Category string `pulumi:"category"`
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag *string `pulumi:"eTag"`
	// May be used to filter budgets by resource group, resource, or meter.
	Filters *Filters `pulumi:"filters"`
	// Budget Name.
	Name string `pulumi:"name"`
	// Dictionary of notifications associated with the budget. Budget can have up to five notifications.
	Notifications map[string]Notification `pulumi:"notifications"`
	// Azure Resource Group Name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The time covered by a budget. Tracking of the amount will be reset based on the time grain.
	TimeGrain string `pulumi:"timeGrain"`
	// Has start and end date of the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than three months. Past start date should  be selected within the timegrain period. There are no restrictions on the end date.
	TimePeriod BudgetTimePeriod `pulumi:"timePeriod"`
}

// The set of arguments for constructing a BudgetByResourceGroupName resource.
type BudgetByResourceGroupNameArgs struct {
	// The total amount of cost to track with the budget
	Amount pulumi.Float64Input
	// The category of the budget, whether the budget tracks cost or usage.
	Category pulumi.StringInput
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag pulumi.StringPtrInput
	// May be used to filter budgets by resource group, resource, or meter.
	Filters FiltersPtrInput
	// Budget Name.
	Name pulumi.StringInput
	// Dictionary of notifications associated with the budget. Budget can have up to five notifications.
	Notifications NotificationMapInput
	// Azure Resource Group Name.
	ResourceGroupName pulumi.StringInput
	// The time covered by a budget. Tracking of the amount will be reset based on the time grain.
	TimeGrain pulumi.StringInput
	// Has start and end date of the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than three months. Past start date should  be selected within the timegrain period. There are no restrictions on the end date.
	TimePeriod BudgetTimePeriodInput
}

func (BudgetByResourceGroupNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*budgetByResourceGroupNameArgs)(nil)).Elem()
}
