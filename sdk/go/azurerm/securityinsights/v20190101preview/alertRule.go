// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190101preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Alert rule.
//
// ## Example Usage
// ### Creates or updates a Fusion alert rule.
//
// ```go
// package main
//
// import (
// 	securityinsights "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/securityinsights/v20190101preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := securityinsights.NewAlertRule(ctx, "alertRule", &securityinsights.AlertRuleArgs{
// 			Etag:                                pulumi.String("3d00c3ca-0000-0100-0000-5d42d5010000"),
// 			Kind:                                pulumi.String("Fusion"),
// 			OperationalInsightsResourceProvider: pulumi.String("Microsoft.OperationalInsights"),
// 			ResourceGroupName:                   pulumi.String("myRg"),
// 			RuleId:                              pulumi.String("myFirstFusionRule"),
// 			WorkspaceName:                       pulumi.String("myWorkspace"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Creates or updates a MicrosoftSecurityIncidentCreation rule.
//
// ```go
// package main
//
// import (
// 	securityinsights "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/securityinsights/v20190101preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := securityinsights.NewAlertRule(ctx, "alertRule", &securityinsights.AlertRuleArgs{
// 			Etag:                                pulumi.String("\"260097e0-0000-0d00-0000-5d6fa88f0000\""),
// 			Kind:                                pulumi.String("MicrosoftSecurityIncidentCreation"),
// 			OperationalInsightsResourceProvider: pulumi.String("Microsoft.OperationalInsights"),
// 			ResourceGroupName:                   pulumi.String("myRg"),
// 			RuleId:                              pulumi.String("microsoftSecurityIncidentCreationRuleExample"),
// 			WorkspaceName:                       pulumi.String("myWorkspace"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Creates or updates a Scheduled alert rule.
//
// ```go
// package main
//
// import (
// 	securityinsights "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/securityinsights/v20190101preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := securityinsights.NewAlertRule(ctx, "alertRule", &securityinsights.AlertRuleArgs{
// 			Etag:                                pulumi.String("\"0300bf09-0000-0000-0000-5c37296e0000\""),
// 			Kind:                                pulumi.String("Scheduled"),
// 			OperationalInsightsResourceProvider: pulumi.String("Microsoft.OperationalInsights"),
// 			ResourceGroupName:                   pulumi.String("myRg"),
// 			RuleId:                              pulumi.String("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
// 			WorkspaceName:                       pulumi.String("myWorkspace"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type AlertRule struct {
	pulumi.CustomResourceState

	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the alert rule
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAlertRule registers a new resource with the given unique name, arguments, and options.
func NewAlertRule(ctx *pulumi.Context,
	name string, args *AlertRuleArgs, opts ...pulumi.ResourceOption) (*AlertRule, error) {
	if args == nil || args.Kind == nil {
		return nil, errors.New("missing required argument 'Kind'")
	}
	if args == nil || args.OperationalInsightsResourceProvider == nil {
		return nil, errors.New("missing required argument 'OperationalInsightsResourceProvider'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.RuleId == nil {
		return nil, errors.New("missing required argument 'RuleId'")
	}
	if args == nil || args.WorkspaceName == nil {
		return nil, errors.New("missing required argument 'WorkspaceName'")
	}
	if args == nil {
		args = &AlertRuleArgs{}
	}
	var resource AlertRule
	err := ctx.RegisterResource("azurerm:securityinsights/v20190101preview:AlertRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlertRule gets an existing AlertRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlertRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlertRuleState, opts ...pulumi.ResourceOption) (*AlertRule, error) {
	var resource AlertRule
	err := ctx.ReadResource("azurerm:securityinsights/v20190101preview:AlertRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AlertRule resources.
type alertRuleState struct {
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// The kind of the alert rule
	Kind *string `pulumi:"kind"`
	// Azure resource name
	Name *string `pulumi:"name"`
	// Azure resource type
	Type *string `pulumi:"type"`
}

type AlertRuleState struct {
	// Etag of the azure resource
	Etag pulumi.StringPtrInput
	// The kind of the alert rule
	Kind pulumi.StringPtrInput
	// Azure resource name
	Name pulumi.StringPtrInput
	// Azure resource type
	Type pulumi.StringPtrInput
}

func (AlertRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertRuleState)(nil)).Elem()
}

type alertRuleArgs struct {
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// The kind of the alert rule
	Kind string `pulumi:"kind"`
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Alert rule ID
	RuleId string `pulumi:"ruleId"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a AlertRule resource.
type AlertRuleArgs struct {
	// Etag of the azure resource
	Etag pulumi.StringPtrInput
	// The kind of the alert rule
	Kind pulumi.StringInput
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider pulumi.StringInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Alert rule ID
	RuleId pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (AlertRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alertRuleArgs)(nil)).Elem()
}
