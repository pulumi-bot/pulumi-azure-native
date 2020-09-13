// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200501preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The scheduled query rule resource.
//
// ## Example Usage
// ### Create or update a scheduled query rule for Single Resource
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	insights "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/insights/v20200501preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := insights.NewScheduledQueryRule(ctx, "scheduledQueryRule", &insights.ScheduledQueryRuleArgs{
// 			Actions: insights.ActionArray{
// 				&insights.ActionArgs{
// 					ActionGroupId: pulumi.String("/subscriptions/1cf177ed-1330-4692-80ea-fd3d7783b147/resourcegroups/sqrapi/providers/microsoft.insights/actiongroups/myactiongroup"),
// 					WebHookProperties: pulumi.StringMap{
// 						"key11": pulumi.String("value11"),
// 						"key12": pulumi.String("value12"),
// 					},
// 				},
// 			},
// 			Criteria: &insights.ScheduledQueryRuleCriteriaArgs{
// 				AllOf: insights.ConditionArray{
// 					&insights.ConditionArgs{
// 						Dimensions: insights.DimensionArray{
// 							&insights.DimensionArgs{
// 								Name:     pulumi.String("ComputerIp"),
// 								Operator: pulumi.String("Exclude"),
// 								Values: pulumi.StringArray{
// 									pulumi.String("192.168.1.1"),
// 								},
// 							},
// 							&insights.DimensionArgs{
// 								Name:     pulumi.String("OSType"),
// 								Operator: pulumi.String("Include"),
// 								Values: pulumi.StringArray{
// 									pulumi.String("*"),
// 								},
// 							},
// 						},
// 						FailingPeriods: &insights.ConditionFailingPeriodsArgs{
// 							MinFailingPeriodsToAlert:  pulumi.Float64(1),
// 							NumberOfEvaluationPeriods: pulumi.Float64(1),
// 						},
// 						MetricMeasureColumn: pulumi.String(fmt.Sprintf("%v%v", "%", " Processor Time")),
// 						Operator:            pulumi.String("GreaterThan"),
// 						Query:               pulumi.String("Perf | where ObjectName == \"Processor\""),
// 						ResourceIdColumn:    pulumi.String("resourceId"),
// 						Threshold:           pulumi.Float64(70),
// 						TimeAggregation:     pulumi.String("Average"),
// 					},
// 				},
// 			},
// 			Description:         pulumi.String("Performance rule"),
// 			Enabled:             pulumi.Bool(true),
// 			EvaluationFrequency: pulumi.String("PT5M"),
// 			Location:            pulumi.String("eastus"),
// 			MuteActionsDuration: pulumi.String("PT30M"),
// 			ResourceGroupName:   pulumi.String("QueryResourceGroupName"),
// 			RuleName:            pulumi.String("perf"),
// 			Scopes: pulumi.StringArray{
// 				pulumi.String("/subscriptions/aaf177ed-1330-a9f2-80ea-fd3d7783b147/resourceGroups/scopeResourceGroup1/providers/Microsoft.Compute/virtualMachines/vm1"),
// 			},
// 			Severity:   pulumi.Float64(4),
// 			WindowSize: pulumi.String("PT10M"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Create or update a scheduled query rule on Resource group(s)
//
// ```go
// package main
//
// import (
// 	insights "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/insights/v20200501preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := insights.NewScheduledQueryRule(ctx, "scheduledQueryRule", &insights.ScheduledQueryRuleArgs{
// 			Actions: insights.ActionArray{
// 				&insights.ActionArgs{
// 					ActionGroupId: pulumi.String("/subscriptions/1cf177ed-1330-4692-80ea-fd3d7783b147/resourcegroups/sqrapi/providers/microsoft.insights/actiongroups/myactiongroup"),
// 					WebHookProperties: pulumi.StringMap{
// 						"key11": pulumi.String("value11"),
// 						"key12": pulumi.String("value12"),
// 					},
// 				},
// 			},
// 			Criteria: &insights.ScheduledQueryRuleCriteriaArgs{
// 				AllOf: insights.ConditionArray{
// 					&insights.ConditionArgs{
// 						Dimensions: insights.DimensionArray{},
// 						FailingPeriods: &insights.ConditionFailingPeriodsArgs{
// 							MinFailingPeriodsToAlert:  pulumi.Float64(1),
// 							NumberOfEvaluationPeriods: pulumi.Float64(1),
// 						},
// 						Operator:        pulumi.String("GreaterThan"),
// 						Query:           pulumi.String("Heartbeat"),
// 						Threshold:       pulumi.Float64(360),
// 						TimeAggregation: pulumi.String("Count"),
// 					},
// 				},
// 			},
// 			Description:         pulumi.String("Health check rule"),
// 			Enabled:             pulumi.Bool(true),
// 			EvaluationFrequency: pulumi.String("PT5M"),
// 			Location:            pulumi.String("eastus"),
// 			MuteActionsDuration: pulumi.String("PT30M"),
// 			ResourceGroupName:   pulumi.String("QueryResourceGroupName"),
// 			RuleName:            pulumi.String("heartbeat"),
// 			Scopes: pulumi.StringArray{
// 				pulumi.String("/subscriptions/aaf177ed-1330-a9f2-80ea-fd3d7783b147/resourceGroups/scopeResourceGroup1"),
// 			},
// 			Severity: pulumi.Float64(4),
// 			TargetResourceTypes: pulumi.StringArray{
// 				pulumi.String("Microsoft.Compute/virtualMachines"),
// 			},
// 			WindowSize: pulumi.String("PT10M"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Create or update a scheduled query rule on Subscription
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	insights "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/insights/v20200501preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := insights.NewScheduledQueryRule(ctx, "scheduledQueryRule", &insights.ScheduledQueryRuleArgs{
// 			Actions: insights.ActionArray{
// 				&insights.ActionArgs{
// 					ActionGroupId: pulumi.String("/subscriptions/1cf177ed-1330-4692-80ea-fd3d7783b147/resourcegroups/sqrapi/providers/microsoft.insights/actiongroups/myactiongroup"),
// 					WebHookProperties: pulumi.StringMap{
// 						"key11": pulumi.String("value11"),
// 						"key12": pulumi.String("value12"),
// 					},
// 				},
// 			},
// 			Criteria: &insights.ScheduledQueryRuleCriteriaArgs{
// 				AllOf: insights.ConditionArray{
// 					&insights.ConditionArgs{
// 						Dimensions: insights.DimensionArray{
// 							&insights.DimensionArgs{
// 								Name:     pulumi.String("ComputerIp"),
// 								Operator: pulumi.String("Exclude"),
// 								Values: pulumi.StringArray{
// 									pulumi.String("192.168.1.1"),
// 								},
// 							},
// 							&insights.DimensionArgs{
// 								Name:     pulumi.String("OSType"),
// 								Operator: pulumi.String("Include"),
// 								Values: pulumi.StringArray{
// 									pulumi.String("*"),
// 								},
// 							},
// 						},
// 						FailingPeriods: &insights.ConditionFailingPeriodsArgs{
// 							MinFailingPeriodsToAlert:  pulumi.Float64(1),
// 							NumberOfEvaluationPeriods: pulumi.Float64(1),
// 						},
// 						MetricMeasureColumn: pulumi.String(fmt.Sprintf("%v%v", "%", " Processor Time")),
// 						Operator:            pulumi.String("GreaterThan"),
// 						Query:               pulumi.String("Perf | where ObjectName == \"Processor\""),
// 						ResourceIdColumn:    pulumi.String("resourceId"),
// 						Threshold:           pulumi.Float64(70),
// 						TimeAggregation:     pulumi.String("Average"),
// 					},
// 				},
// 			},
// 			Description:         pulumi.String("Performance rule"),
// 			Enabled:             pulumi.Bool(true),
// 			EvaluationFrequency: pulumi.String("PT5M"),
// 			Location:            pulumi.String("eastus"),
// 			MuteActionsDuration: pulumi.String("PT30M"),
// 			ResourceGroupName:   pulumi.String("QueryResourceGroupName"),
// 			RuleName:            pulumi.String("perf"),
// 			Scopes: pulumi.StringArray{
// 				pulumi.String("/subscriptions/aaf177ed-1330-a9f2-80ea-fd3d7783b147"),
// 			},
// 			Severity: pulumi.Float64(4),
// 			TargetResourceTypes: pulumi.StringArray{
// 				pulumi.String("Microsoft.Compute/virtualMachines"),
// 			},
// 			WindowSize: pulumi.String("PT10M"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type ScheduledQueryRule struct {
	pulumi.CustomResourceState

	Actions ActionResponseArrayOutput `pulumi:"actions"`
	// The rule criteria that defines the conditions of the scheduled query rule.
	Criteria ScheduledQueryRuleCriteriaResponsePtrOutput `pulumi:"criteria"`
	// The description of the scheduled query rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The flag which indicates whether this scheduled query rule is enabled. Value should be true or false
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// How often the scheduled query rule is evaluated represented in ISO 8601 duration format.
	EvaluationFrequency pulumi.StringPtrOutput `pulumi:"evaluationFrequency"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Mute actions for the chosen period of time (in ISO 8601 duration format) after the alert is fired.
	MuteActionsDuration pulumi.StringPtrOutput `pulumi:"muteActionsDuration"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of resource id's that this scheduled query rule is scoped to.
	Scopes pulumi.StringArrayOutput `pulumi:"scopes"`
	// Severity of the alert. Should be an integer between [0-4]. Value of 0 is severest
	Severity pulumi.Float64PtrOutput `pulumi:"severity"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// List of resource type of the target resource(s) on which the alert is created/updated. For example if the scope is a resource group and targetResourceTypes is Microsoft.Compute/virtualMachines, then a different alert will be fired for each virtual machine in the resource group which meet the alert criteria
	TargetResourceTypes pulumi.StringArrayOutput `pulumi:"targetResourceTypes"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
	// The period of time (in ISO 8601 duration format) on which the Alert query will be executed (bin size).
	WindowSize pulumi.StringPtrOutput `pulumi:"windowSize"`
}

// NewScheduledQueryRule registers a new resource with the given unique name, arguments, and options.
func NewScheduledQueryRule(ctx *pulumi.Context,
	name string, args *ScheduledQueryRuleArgs, opts ...pulumi.ResourceOption) (*ScheduledQueryRule, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.RuleName == nil {
		return nil, errors.New("missing required argument 'RuleName'")
	}
	if args == nil {
		args = &ScheduledQueryRuleArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:insights/latest:ScheduledQueryRule"),
		},
		{
			Type: pulumi.String("azurerm:insights/v20180416:ScheduledQueryRule"),
		},
	})
	opts = append(opts, aliases)
	var resource ScheduledQueryRule
	err := ctx.RegisterResource("azurerm:insights/v20200501preview:ScheduledQueryRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScheduledQueryRule gets an existing ScheduledQueryRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScheduledQueryRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduledQueryRuleState, opts ...pulumi.ResourceOption) (*ScheduledQueryRule, error) {
	var resource ScheduledQueryRule
	err := ctx.ReadResource("azurerm:insights/v20200501preview:ScheduledQueryRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScheduledQueryRule resources.
type scheduledQueryRuleState struct {
	Actions []ActionResponse `pulumi:"actions"`
	// The rule criteria that defines the conditions of the scheduled query rule.
	Criteria *ScheduledQueryRuleCriteriaResponse `pulumi:"criteria"`
	// The description of the scheduled query rule.
	Description *string `pulumi:"description"`
	// The flag which indicates whether this scheduled query rule is enabled. Value should be true or false
	Enabled *bool `pulumi:"enabled"`
	// How often the scheduled query rule is evaluated represented in ISO 8601 duration format.
	EvaluationFrequency *string `pulumi:"evaluationFrequency"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Mute actions for the chosen period of time (in ISO 8601 duration format) after the alert is fired.
	MuteActionsDuration *string `pulumi:"muteActionsDuration"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The list of resource id's that this scheduled query rule is scoped to.
	Scopes []string `pulumi:"scopes"`
	// Severity of the alert. Should be an integer between [0-4]. Value of 0 is severest
	Severity *float64 `pulumi:"severity"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// List of resource type of the target resource(s) on which the alert is created/updated. For example if the scope is a resource group and targetResourceTypes is Microsoft.Compute/virtualMachines, then a different alert will be fired for each virtual machine in the resource group which meet the alert criteria
	TargetResourceTypes []string `pulumi:"targetResourceTypes"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
	// The period of time (in ISO 8601 duration format) on which the Alert query will be executed (bin size).
	WindowSize *string `pulumi:"windowSize"`
}

type ScheduledQueryRuleState struct {
	Actions ActionResponseArrayInput
	// The rule criteria that defines the conditions of the scheduled query rule.
	Criteria ScheduledQueryRuleCriteriaResponsePtrInput
	// The description of the scheduled query rule.
	Description pulumi.StringPtrInput
	// The flag which indicates whether this scheduled query rule is enabled. Value should be true or false
	Enabled pulumi.BoolPtrInput
	// How often the scheduled query rule is evaluated represented in ISO 8601 duration format.
	EvaluationFrequency pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Mute actions for the chosen period of time (in ISO 8601 duration format) after the alert is fired.
	MuteActionsDuration pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The list of resource id's that this scheduled query rule is scoped to.
	Scopes pulumi.StringArrayInput
	// Severity of the alert. Should be an integer between [0-4]. Value of 0 is severest
	Severity pulumi.Float64PtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// List of resource type of the target resource(s) on which the alert is created/updated. For example if the scope is a resource group and targetResourceTypes is Microsoft.Compute/virtualMachines, then a different alert will be fired for each virtual machine in the resource group which meet the alert criteria
	TargetResourceTypes pulumi.StringArrayInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
	// The period of time (in ISO 8601 duration format) on which the Alert query will be executed (bin size).
	WindowSize pulumi.StringPtrInput
}

func (ScheduledQueryRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledQueryRuleState)(nil)).Elem()
}

type scheduledQueryRuleArgs struct {
	Actions []Action `pulumi:"actions"`
	// The rule criteria that defines the conditions of the scheduled query rule.
	Criteria *ScheduledQueryRuleCriteria `pulumi:"criteria"`
	// The description of the scheduled query rule.
	Description *string `pulumi:"description"`
	// The flag which indicates whether this scheduled query rule is enabled. Value should be true or false
	Enabled *bool `pulumi:"enabled"`
	// How often the scheduled query rule is evaluated represented in ISO 8601 duration format.
	EvaluationFrequency *string `pulumi:"evaluationFrequency"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Mute actions for the chosen period of time (in ISO 8601 duration format) after the alert is fired.
	MuteActionsDuration *string `pulumi:"muteActionsDuration"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the rule.
	RuleName string `pulumi:"ruleName"`
	// The list of resource id's that this scheduled query rule is scoped to.
	Scopes []string `pulumi:"scopes"`
	// Severity of the alert. Should be an integer between [0-4]. Value of 0 is severest
	Severity *float64 `pulumi:"severity"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// List of resource type of the target resource(s) on which the alert is created/updated. For example if the scope is a resource group and targetResourceTypes is Microsoft.Compute/virtualMachines, then a different alert will be fired for each virtual machine in the resource group which meet the alert criteria
	TargetResourceTypes []string `pulumi:"targetResourceTypes"`
	// The period of time (in ISO 8601 duration format) on which the Alert query will be executed (bin size).
	WindowSize *string `pulumi:"windowSize"`
}

// The set of arguments for constructing a ScheduledQueryRule resource.
type ScheduledQueryRuleArgs struct {
	Actions ActionArrayInput
	// The rule criteria that defines the conditions of the scheduled query rule.
	Criteria ScheduledQueryRuleCriteriaPtrInput
	// The description of the scheduled query rule.
	Description pulumi.StringPtrInput
	// The flag which indicates whether this scheduled query rule is enabled. Value should be true or false
	Enabled pulumi.BoolPtrInput
	// How often the scheduled query rule is evaluated represented in ISO 8601 duration format.
	EvaluationFrequency pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringInput
	// Mute actions for the chosen period of time (in ISO 8601 duration format) after the alert is fired.
	MuteActionsDuration pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the rule.
	RuleName pulumi.StringInput
	// The list of resource id's that this scheduled query rule is scoped to.
	Scopes pulumi.StringArrayInput
	// Severity of the alert. Should be an integer between [0-4]. Value of 0 is severest
	Severity pulumi.Float64PtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// List of resource type of the target resource(s) on which the alert is created/updated. For example if the scope is a resource group and targetResourceTypes is Microsoft.Compute/virtualMachines, then a different alert will be fired for each virtual machine in the resource group which meet the alert criteria
	TargetResourceTypes pulumi.StringArrayInput
	// The period of time (in ISO 8601 duration format) on which the Alert query will be executed (bin size).
	WindowSize pulumi.StringPtrInput
}

func (ScheduledQueryRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledQueryRuleArgs)(nil)).Elem()
}
