// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// States and configurations of Cost Analysis.
type ViewByScope struct {
	pulumi.CustomResourceState

	// Show costs accumulated over time.
	Accumulated pulumi.StringPtrOutput `pulumi:"accumulated"`
	// Chart type of the main view in Cost Analysis. Required.
	Chart pulumi.StringPtrOutput `pulumi:"chart"`
	// Date the user created this view.
	CreatedOn pulumi.StringOutput `pulumi:"createdOn"`
	// Has definition for data in this report config.
	Dataset ReportConfigDatasetResponsePtrOutput `pulumi:"dataset"`
	// User input name of the view. Required.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// List of KPIs to show in Cost Analysis UI.
	Kpis KpiPropertiesResponseArrayOutput `pulumi:"kpis"`
	// Metric to use when displaying costs.
	Metric pulumi.StringPtrOutput `pulumi:"metric"`
	// Date when the user last modified this view.
	ModifiedOn pulumi.StringOutput `pulumi:"modifiedOn"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Configuration of 3 sub-views in the Cost Analysis UI.
	Pivots PivotPropertiesResponseArrayOutput `pulumi:"pivots"`
	// Cost Management scope to save the view on. This includes 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, 'providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
	Scope pulumi.StringPtrOutput `pulumi:"scope"`
	// Has time period for pulling data for the report.
	TimePeriod ReportConfigTimePeriodResponsePtrOutput `pulumi:"timePeriod"`
	// The time frame for pulling data for the report. If custom, then a specific time period must be provided.
	Timeframe pulumi.StringOutput `pulumi:"timeframe"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewViewByScope registers a new resource with the given unique name, arguments, and options.
func NewViewByScope(ctx *pulumi.Context,
	name string, args *ViewByScopeArgs, opts ...pulumi.ResourceOption) (*ViewByScope, error) {
	if args == nil || args.Scope == nil {
		return nil, errors.New("missing required argument 'Scope'")
	}
	if args == nil || args.Timeframe == nil {
		return nil, errors.New("missing required argument 'Timeframe'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	if args == nil || args.ViewName == nil {
		return nil, errors.New("missing required argument 'ViewName'")
	}
	if args == nil {
		args = &ViewByScopeArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:costmanagement/v20190401preview:ViewByScope"),
		},
		{
			Type: pulumi.String("azure-nextgen:costmanagement/v20191101:ViewByScope"),
		},
		{
			Type: pulumi.String("azure-nextgen:costmanagement/v20200601:ViewByScope"),
		},
	})
	opts = append(opts, aliases)
	var resource ViewByScope
	err := ctx.RegisterResource("azure-nextgen:costmanagement/latest:ViewByScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetViewByScope gets an existing ViewByScope resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetViewByScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ViewByScopeState, opts ...pulumi.ResourceOption) (*ViewByScope, error) {
	var resource ViewByScope
	err := ctx.ReadResource("azure-nextgen:costmanagement/latest:ViewByScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ViewByScope resources.
type viewByScopeState struct {
	// Show costs accumulated over time.
	Accumulated *string `pulumi:"accumulated"`
	// Chart type of the main view in Cost Analysis. Required.
	Chart *string `pulumi:"chart"`
	// Date the user created this view.
	CreatedOn *string `pulumi:"createdOn"`
	// Has definition for data in this report config.
	Dataset *ReportConfigDatasetResponse `pulumi:"dataset"`
	// User input name of the view. Required.
	DisplayName *string `pulumi:"displayName"`
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag *string `pulumi:"eTag"`
	// List of KPIs to show in Cost Analysis UI.
	Kpis []KpiPropertiesResponse `pulumi:"kpis"`
	// Metric to use when displaying costs.
	Metric *string `pulumi:"metric"`
	// Date when the user last modified this view.
	ModifiedOn *string `pulumi:"modifiedOn"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Configuration of 3 sub-views in the Cost Analysis UI.
	Pivots []PivotPropertiesResponse `pulumi:"pivots"`
	// Cost Management scope to save the view on. This includes 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, 'providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
	Scope *string `pulumi:"scope"`
	// Has time period for pulling data for the report.
	TimePeriod *ReportConfigTimePeriodResponse `pulumi:"timePeriod"`
	// The time frame for pulling data for the report. If custom, then a specific time period must be provided.
	Timeframe *string `pulumi:"timeframe"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type ViewByScopeState struct {
	// Show costs accumulated over time.
	Accumulated pulumi.StringPtrInput
	// Chart type of the main view in Cost Analysis. Required.
	Chart pulumi.StringPtrInput
	// Date the user created this view.
	CreatedOn pulumi.StringPtrInput
	// Has definition for data in this report config.
	Dataset ReportConfigDatasetResponsePtrInput
	// User input name of the view. Required.
	DisplayName pulumi.StringPtrInput
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag pulumi.StringPtrInput
	// List of KPIs to show in Cost Analysis UI.
	Kpis KpiPropertiesResponseArrayInput
	// Metric to use when displaying costs.
	Metric pulumi.StringPtrInput
	// Date when the user last modified this view.
	ModifiedOn pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Configuration of 3 sub-views in the Cost Analysis UI.
	Pivots PivotPropertiesResponseArrayInput
	// Cost Management scope to save the view on. This includes 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, 'providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
	Scope pulumi.StringPtrInput
	// Has time period for pulling data for the report.
	TimePeriod ReportConfigTimePeriodResponsePtrInput
	// The time frame for pulling data for the report. If custom, then a specific time period must be provided.
	Timeframe pulumi.StringPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (ViewByScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*viewByScopeState)(nil)).Elem()
}

type viewByScopeArgs struct {
	// Show costs accumulated over time.
	Accumulated *string `pulumi:"accumulated"`
	// Chart type of the main view in Cost Analysis. Required.
	Chart *string `pulumi:"chart"`
	// Has definition for data in this report config.
	Dataset *ReportConfigDataset `pulumi:"dataset"`
	// User input name of the view. Required.
	DisplayName *string `pulumi:"displayName"`
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag *string `pulumi:"eTag"`
	// List of KPIs to show in Cost Analysis UI.
	Kpis []KpiProperties `pulumi:"kpis"`
	// Metric to use when displaying costs.
	Metric *string `pulumi:"metric"`
	// Configuration of 3 sub-views in the Cost Analysis UI.
	Pivots []PivotProperties `pulumi:"pivots"`
	// Cost Management scope to save the view on. This includes 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, 'providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
	Scope string `pulumi:"scope"`
	// Has time period for pulling data for the report.
	TimePeriod *ReportConfigTimePeriod `pulumi:"timePeriod"`
	// The time frame for pulling data for the report. If custom, then a specific time period must be provided.
	Timeframe string `pulumi:"timeframe"`
	// The type of the report. Usage represents actual usage, forecast represents forecasted data and UsageAndForecast represents both usage and forecasted data. Actual usage and forecasted data can be differentiated based on dates.
	Type string `pulumi:"type"`
	// View name
	ViewName string `pulumi:"viewName"`
}

// The set of arguments for constructing a ViewByScope resource.
type ViewByScopeArgs struct {
	// Show costs accumulated over time.
	Accumulated pulumi.StringPtrInput
	// Chart type of the main view in Cost Analysis. Required.
	Chart pulumi.StringPtrInput
	// Has definition for data in this report config.
	Dataset ReportConfigDatasetPtrInput
	// User input name of the view. Required.
	DisplayName pulumi.StringPtrInput
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag pulumi.StringPtrInput
	// List of KPIs to show in Cost Analysis UI.
	Kpis KpiPropertiesArrayInput
	// Metric to use when displaying costs.
	Metric pulumi.StringPtrInput
	// Configuration of 3 sub-views in the Cost Analysis UI.
	Pivots PivotPropertiesArrayInput
	// Cost Management scope to save the view on. This includes 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, 'providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
	Scope pulumi.StringInput
	// Has time period for pulling data for the report.
	TimePeriod ReportConfigTimePeriodPtrInput
	// The time frame for pulling data for the report. If custom, then a specific time period must be provided.
	Timeframe pulumi.StringInput
	// The type of the report. Usage represents actual usage, forecast represents forecasted data and UsageAndForecast represents both usage and forecasted data. Actual usage and forecasted data can be differentiated based on dates.
	Type pulumi.StringInput
	// View name
	ViewName pulumi.StringInput
}

func (ViewByScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*viewByScopeArgs)(nil)).Elem()
}

type ViewByScopeInput interface {
	pulumi.Input

	ToViewByScopeOutput() ViewByScopeOutput
	ToViewByScopeOutputWithContext(ctx context.Context) ViewByScopeOutput
}

func (ViewByScope) ElementType() reflect.Type {
	return reflect.TypeOf((*ViewByScope)(nil)).Elem()
}

func (i ViewByScope) ToViewByScopeOutput() ViewByScopeOutput {
	return i.ToViewByScopeOutputWithContext(context.Background())
}

func (i ViewByScope) ToViewByScopeOutputWithContext(ctx context.Context) ViewByScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ViewByScopeOutput)
}

type ViewByScopeOutput struct {
	*pulumi.OutputState
}

func (ViewByScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ViewByScopeOutput)(nil)).Elem()
}

func (o ViewByScopeOutput) ToViewByScopeOutput() ViewByScopeOutput {
	return o
}

func (o ViewByScopeOutput) ToViewByScopeOutputWithContext(ctx context.Context) ViewByScopeOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ViewByScopeOutput{})
}
