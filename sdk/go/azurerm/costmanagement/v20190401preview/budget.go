// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190401preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A budget resource.
type Budget struct {
	pulumi.CustomResourceState

	// The total amount of cost to track with the budget
	Amount pulumi.Float64Output `pulumi:"amount"`
	// The category of the budget, whether the budget tracks cost or usage.
	Category pulumi.StringOutput `pulumi:"category"`
	// The current amount of cost which is being tracked for a budget.
	CurrentSpend CurrentSpendResponseOutput `pulumi:"currentSpend"`
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// May be used to filter budgets.
	Filter ReportConfigFilterResponsePtrOutput `pulumi:"filter"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Dictionary of notifications associated with the budget. Budget can have up to five notifications.
	Notifications NotificationResponseMapOutput `pulumi:"notifications"`
	// The time covered by a budget. Tracking of the amount will be reset based on the time grain.
	TimeGrain pulumi.StringOutput `pulumi:"timeGrain"`
	// Has start and end date of the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than three months. Past start date should  be selected within the timegrain period. There are no restrictions on the end date.
	TimePeriod BudgetTimePeriodResponseOutput `pulumi:"timePeriod"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBudget registers a new resource with the given unique name, arguments, and options.
func NewBudget(ctx *pulumi.Context,
	name string, args *BudgetArgs, opts ...pulumi.ResourceOption) (*Budget, error) {
	if args == nil || args.Amount == nil {
		return nil, errors.New("missing required argument 'Amount'")
	}
	if args == nil || args.BudgetName == nil {
		return nil, errors.New("missing required argument 'BudgetName'")
	}
	if args == nil || args.Category == nil {
		return nil, errors.New("missing required argument 'Category'")
	}
	if args == nil || args.Scope == nil {
		return nil, errors.New("missing required argument 'Scope'")
	}
	if args == nil || args.TimeGrain == nil {
		return nil, errors.New("missing required argument 'TimeGrain'")
	}
	if args == nil || args.TimePeriod == nil {
		return nil, errors.New("missing required argument 'TimePeriod'")
	}
	if args == nil {
		args = &BudgetArgs{}
	}
	var resource Budget
	err := ctx.RegisterResource("azurerm:costmanagement/v20190401preview:Budget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBudget gets an existing Budget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBudget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BudgetState, opts ...pulumi.ResourceOption) (*Budget, error) {
	var resource Budget
	err := ctx.ReadResource("azurerm:costmanagement/v20190401preview:Budget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Budget resources.
type budgetState struct {
	// The total amount of cost to track with the budget
	Amount *float64 `pulumi:"amount"`
	// The category of the budget, whether the budget tracks cost or usage.
	Category *string `pulumi:"category"`
	// The current amount of cost which is being tracked for a budget.
	CurrentSpend *CurrentSpendResponse `pulumi:"currentSpend"`
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag *string `pulumi:"eTag"`
	// May be used to filter budgets.
	Filter *ReportConfigFilterResponse `pulumi:"filter"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Dictionary of notifications associated with the budget. Budget can have up to five notifications.
	Notifications map[string]NotificationResponse `pulumi:"notifications"`
	// The time covered by a budget. Tracking of the amount will be reset based on the time grain.
	TimeGrain *string `pulumi:"timeGrain"`
	// Has start and end date of the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than three months. Past start date should  be selected within the timegrain period. There are no restrictions on the end date.
	TimePeriod *BudgetTimePeriodResponse `pulumi:"timePeriod"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type BudgetState struct {
	// The total amount of cost to track with the budget
	Amount pulumi.Float64PtrInput
	// The category of the budget, whether the budget tracks cost or usage.
	Category pulumi.StringPtrInput
	// The current amount of cost which is being tracked for a budget.
	CurrentSpend CurrentSpendResponsePtrInput
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag pulumi.StringPtrInput
	// May be used to filter budgets.
	Filter ReportConfigFilterResponsePtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Dictionary of notifications associated with the budget. Budget can have up to five notifications.
	Notifications NotificationResponseMapInput
	// The time covered by a budget. Tracking of the amount will be reset based on the time grain.
	TimeGrain pulumi.StringPtrInput
	// Has start and end date of the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than three months. Past start date should  be selected within the timegrain period. There are no restrictions on the end date.
	TimePeriod BudgetTimePeriodResponsePtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (BudgetState) ElementType() reflect.Type {
	return reflect.TypeOf((*budgetState)(nil)).Elem()
}

type budgetArgs struct {
	// The total amount of cost to track with the budget
	Amount float64 `pulumi:"amount"`
	// Budget Name.
	BudgetName string `pulumi:"budgetName"`
	// The category of the budget, whether the budget tracks cost or usage.
	Category string `pulumi:"category"`
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag *string `pulumi:"eTag"`
	// May be used to filter budgets.
	Filter *ReportConfigFilter `pulumi:"filter"`
	// Dictionary of notifications associated with the budget. Budget can have up to five notifications.
	Notifications map[string]Notification `pulumi:"notifications"`
	// The scope associated with budget operations. This includes '/subscriptions/{subscriptionId}/' for subscription scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, '/providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for invoiceSection scope.
	Scope string `pulumi:"scope"`
	// The time covered by a budget. Tracking of the amount will be reset based on the time grain.
	TimeGrain string `pulumi:"timeGrain"`
	// Has start and end date of the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than three months. Past start date should  be selected within the timegrain period. There are no restrictions on the end date.
	TimePeriod BudgetTimePeriod `pulumi:"timePeriod"`
}

// The set of arguments for constructing a Budget resource.
type BudgetArgs struct {
	// The total amount of cost to track with the budget
	Amount pulumi.Float64Input
	// Budget Name.
	BudgetName pulumi.StringInput
	// The category of the budget, whether the budget tracks cost or usage.
	Category pulumi.StringInput
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag pulumi.StringPtrInput
	// May be used to filter budgets.
	Filter ReportConfigFilterPtrInput
	// Dictionary of notifications associated with the budget. Budget can have up to five notifications.
	Notifications NotificationMapInput
	// The scope associated with budget operations. This includes '/subscriptions/{subscriptionId}/' for subscription scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, '/providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for invoiceSection scope.
	Scope pulumi.StringInput
	// The time covered by a budget. Tracking of the amount will be reset based on the time grain.
	TimeGrain pulumi.StringInput
	// Has start and end date of the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than three months. Past start date should  be selected within the timegrain period. There are no restrictions on the end date.
	TimePeriod BudgetTimePeriodInput
}

func (BudgetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*budgetArgs)(nil)).Elem()
}
