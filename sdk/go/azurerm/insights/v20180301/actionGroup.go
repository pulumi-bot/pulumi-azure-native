// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180301

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An action group resource.
//
// ## Example Usage
// ### Create or update an action group
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	insights "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/insights/v20180301"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := insights.NewActionGroup(ctx, "actionGroup", &insights.ActionGroupArgs{
// 			ActionGroupName: pulumi.String("SampleActionGroup"),
// 			AutomationRunbookReceivers: insights.AutomationRunbookReceiverArray{
// 				&insights.AutomationRunbookReceiverArgs{
// 					AutomationAccountId: pulumi.String("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/runbookTest/providers/Microsoft.Automation/automationAccounts/runbooktest"),
// 					IsGlobalRunbook:     pulumi.Bool(false),
// 					Name:                pulumi.String("testRunbook"),
// 					RunbookName:         pulumi.String("Sample runbook"),
// 					ServiceUri:          pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v", "https://s13events.azure-automation.net/webhooks?token=iimE", "%", "2fD19Eg", "%", "2bvDy22yUMecIZY6Uiz", "%", "2bHfuQ67r8r1wY", "%", "2fI", "%", "3d")),
// 					WebhookResourceId:   pulumi.String("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/runbookTest/providers/Microsoft.Automation/automationAccounts/runbooktest/webhooks/Alert1510184037084"),
// 				},
// 			},
// 			AzureAppPushReceivers: insights.AzureAppPushReceiverArray{
// 				&insights.AzureAppPushReceiverArgs{
// 					EmailAddress: pulumi.String("johndoe@email.com"),
// 					Name:         pulumi.String("Sample azureAppPush"),
// 				},
// 			},
// 			AzureFunctionReceivers: insights.AzureFunctionReceiverArray{
// 				&insights.AzureFunctionReceiverArgs{
// 					FunctionAppResourceId: pulumi.String("/subscriptions/5def922a-3ed4-49c1-b9fd-05ec533819a3/resourceGroups/aznsTest/providers/Microsoft.Web/sites/testFunctionApp"),
// 					FunctionName:          pulumi.String("HttpTriggerCSharp1"),
// 					HttpTriggerUrl:        pulumi.String("https://testfunctionapp.azurewebsites.net/api/HttpTriggerCSharp1?code=4CopFfiXqUQC8dvIM7F53J7tIU3Gy9QQIG/vKAXMe2avhHqK3/jVYw=="),
// 					Name:                  pulumi.String("Sample azureFunction"),
// 				},
// 			},
// 			EmailReceivers: insights.EmailReceiverArray{
// 				&insights.EmailReceiverArgs{
// 					EmailAddress: pulumi.String("johndoe@email.com"),
// 					Name:         pulumi.String("John Doe's email"),
// 				},
// 				&insights.EmailReceiverArgs{
// 					EmailAddress: pulumi.String("janesmith@email.com"),
// 					Name:         pulumi.String("Jane Smith's email"),
// 				},
// 			},
// 			Enabled:        pulumi.Bool(true),
// 			GroupShortName: pulumi.String("sample"),
// 			ItsmReceivers: insights.ItsmReceiverArray{
// 				&insights.ItsmReceiverArgs{
// 					ConnectionId:        pulumi.String("a3b9076c-ce8e-434e-85b4-aff10cb3c8f1"),
// 					Name:                pulumi.String("Sample itsm"),
// 					Region:              pulumi.String("westcentralus"),
// 					TicketConfiguration: pulumi.String("{\"PayloadRevision\":0,\"WorkItemType\":\"Incident\",\"UseTemplate\":false,\"WorkItemData\":\"{}\",\"CreateOneWIPerCI\":false}"),
// 					WorkspaceId:         pulumi.String("5def922a-3ed4-49c1-b9fd-05ec533819a3|55dfd1f8-7e59-4f89-bf56-4c82f5ace23c"),
// 				},
// 			},
// 			Location: pulumi.String("Global"),
// 			LogicAppReceivers: insights.LogicAppReceiverArray{
// 				&insights.LogicAppReceiverArgs{
// 					CallbackUrl: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v", "https://prod-27.northcentralus.logic.azure.com/workflows/68e572e818e5457ba898763b7db90877/triggers/manual/paths/invoke/azns/test?api-version=2016-10-01&sp=", "%", "2Ftriggers", "%", "2Fmanual", "%", "2Frun&sv=1.0&sig=Abpsb72UYJxPPvmDo937uzofupO5r_vIeWEx7KVHo7w")),
// 					Name:        pulumi.String("Sample logicApp"),
// 					ResourceId:  pulumi.String("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/LogicApp/providers/Microsoft.Logic/workflows/testLogicApp"),
// 				},
// 			},
// 			ResourceGroupName: pulumi.String("Default-NotificationRules"),
// 			SmsReceivers: insights.SmsReceiverArray{
// 				&insights.SmsReceiverArgs{
// 					CountryCode: pulumi.String("1"),
// 					Name:        pulumi.String("John Doe's mobile"),
// 					PhoneNumber: pulumi.String("1234567890"),
// 				},
// 				&insights.SmsReceiverArgs{
// 					CountryCode: pulumi.String("1"),
// 					Name:        pulumi.String("Jane Smith's mobile"),
// 					PhoneNumber: pulumi.String("0987654321"),
// 				},
// 			},
// 			Tags: nil,
// 			VoiceReceivers: insights.VoiceReceiverArray{
// 				&insights.VoiceReceiverArgs{
// 					CountryCode: pulumi.String("1"),
// 					Name:        pulumi.String("Sample voice"),
// 					PhoneNumber: pulumi.String("1234567890"),
// 				},
// 			},
// 			WebhookReceivers: insights.WebhookReceiverArray{
// 				&insights.WebhookReceiverArgs{
// 					Name:       pulumi.String("Sample webhook"),
// 					ServiceUri: pulumi.String("http://www.example.com/webhook"),
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
type ActionGroup struct {
	pulumi.CustomResourceState

	// The list of AutomationRunbook receivers that are part of this action group.
	AutomationRunbookReceivers AutomationRunbookReceiverResponseArrayOutput `pulumi:"automationRunbookReceivers"`
	// The list of AzureAppPush receivers that are part of this action group.
	AzureAppPushReceivers AzureAppPushReceiverResponseArrayOutput `pulumi:"azureAppPushReceivers"`
	// The list of azure function receivers that are part of this action group.
	AzureFunctionReceivers AzureFunctionReceiverResponseArrayOutput `pulumi:"azureFunctionReceivers"`
	// The list of email receivers that are part of this action group.
	EmailReceivers EmailReceiverResponseArrayOutput `pulumi:"emailReceivers"`
	// Indicates whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The short name of the action group. This will be used in SMS messages.
	GroupShortName pulumi.StringOutput `pulumi:"groupShortName"`
	// The list of ITSM receivers that are part of this action group.
	ItsmReceivers ItsmReceiverResponseArrayOutput `pulumi:"itsmReceivers"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// The list of logic app receivers that are part of this action group.
	LogicAppReceivers LogicAppReceiverResponseArrayOutput `pulumi:"logicAppReceivers"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of SMS receivers that are part of this action group.
	SmsReceivers SmsReceiverResponseArrayOutput `pulumi:"smsReceivers"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// The list of voice receivers that are part of this action group.
	VoiceReceivers VoiceReceiverResponseArrayOutput `pulumi:"voiceReceivers"`
	// The list of webhook receivers that are part of this action group.
	WebhookReceivers WebhookReceiverResponseArrayOutput `pulumi:"webhookReceivers"`
}

// NewActionGroup registers a new resource with the given unique name, arguments, and options.
func NewActionGroup(ctx *pulumi.Context,
	name string, args *ActionGroupArgs, opts ...pulumi.ResourceOption) (*ActionGroup, error) {
	if args == nil || args.ActionGroupName == nil {
		return nil, errors.New("missing required argument 'ActionGroupName'")
	}
	if args == nil || args.Enabled == nil {
		return nil, errors.New("missing required argument 'Enabled'")
	}
	if args == nil || args.GroupShortName == nil {
		return nil, errors.New("missing required argument 'GroupShortName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ActionGroupArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:insights/latest:ActionGroup"),
		},
		{
			Type: pulumi.String("azurerm:insights/v20170401:ActionGroup"),
		},
		{
			Type: pulumi.String("azurerm:insights/v20180901:ActionGroup"),
		},
		{
			Type: pulumi.String("azurerm:insights/v20190301:ActionGroup"),
		},
		{
			Type: pulumi.String("azurerm:insights/v20190601:ActionGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource ActionGroup
	err := ctx.RegisterResource("azurerm:insights/v20180301:ActionGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActionGroup gets an existing ActionGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActionGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActionGroupState, opts ...pulumi.ResourceOption) (*ActionGroup, error) {
	var resource ActionGroup
	err := ctx.ReadResource("azurerm:insights/v20180301:ActionGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActionGroup resources.
type actionGroupState struct {
	// The list of AutomationRunbook receivers that are part of this action group.
	AutomationRunbookReceivers []AutomationRunbookReceiverResponse `pulumi:"automationRunbookReceivers"`
	// The list of AzureAppPush receivers that are part of this action group.
	AzureAppPushReceivers []AzureAppPushReceiverResponse `pulumi:"azureAppPushReceivers"`
	// The list of azure function receivers that are part of this action group.
	AzureFunctionReceivers []AzureFunctionReceiverResponse `pulumi:"azureFunctionReceivers"`
	// The list of email receivers that are part of this action group.
	EmailReceivers []EmailReceiverResponse `pulumi:"emailReceivers"`
	// Indicates whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications.
	Enabled *bool `pulumi:"enabled"`
	// The short name of the action group. This will be used in SMS messages.
	GroupShortName *string `pulumi:"groupShortName"`
	// The list of ITSM receivers that are part of this action group.
	ItsmReceivers []ItsmReceiverResponse `pulumi:"itsmReceivers"`
	// Resource location
	Location *string `pulumi:"location"`
	// The list of logic app receivers that are part of this action group.
	LogicAppReceivers []LogicAppReceiverResponse `pulumi:"logicAppReceivers"`
	// Azure resource name
	Name *string `pulumi:"name"`
	// The list of SMS receivers that are part of this action group.
	SmsReceivers []SmsReceiverResponse `pulumi:"smsReceivers"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Azure resource type
	Type *string `pulumi:"type"`
	// The list of voice receivers that are part of this action group.
	VoiceReceivers []VoiceReceiverResponse `pulumi:"voiceReceivers"`
	// The list of webhook receivers that are part of this action group.
	WebhookReceivers []WebhookReceiverResponse `pulumi:"webhookReceivers"`
}

type ActionGroupState struct {
	// The list of AutomationRunbook receivers that are part of this action group.
	AutomationRunbookReceivers AutomationRunbookReceiverResponseArrayInput
	// The list of AzureAppPush receivers that are part of this action group.
	AzureAppPushReceivers AzureAppPushReceiverResponseArrayInput
	// The list of azure function receivers that are part of this action group.
	AzureFunctionReceivers AzureFunctionReceiverResponseArrayInput
	// The list of email receivers that are part of this action group.
	EmailReceivers EmailReceiverResponseArrayInput
	// Indicates whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications.
	Enabled pulumi.BoolPtrInput
	// The short name of the action group. This will be used in SMS messages.
	GroupShortName pulumi.StringPtrInput
	// The list of ITSM receivers that are part of this action group.
	ItsmReceivers ItsmReceiverResponseArrayInput
	// Resource location
	Location pulumi.StringPtrInput
	// The list of logic app receivers that are part of this action group.
	LogicAppReceivers LogicAppReceiverResponseArrayInput
	// Azure resource name
	Name pulumi.StringPtrInput
	// The list of SMS receivers that are part of this action group.
	SmsReceivers SmsReceiverResponseArrayInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Azure resource type
	Type pulumi.StringPtrInput
	// The list of voice receivers that are part of this action group.
	VoiceReceivers VoiceReceiverResponseArrayInput
	// The list of webhook receivers that are part of this action group.
	WebhookReceivers WebhookReceiverResponseArrayInput
}

func (ActionGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*actionGroupState)(nil)).Elem()
}

type actionGroupArgs struct {
	// The name of the action group.
	ActionGroupName string `pulumi:"actionGroupName"`
	// The list of AutomationRunbook receivers that are part of this action group.
	AutomationRunbookReceivers []AutomationRunbookReceiver `pulumi:"automationRunbookReceivers"`
	// The list of AzureAppPush receivers that are part of this action group.
	AzureAppPushReceivers []AzureAppPushReceiver `pulumi:"azureAppPushReceivers"`
	// The list of azure function receivers that are part of this action group.
	AzureFunctionReceivers []AzureFunctionReceiver `pulumi:"azureFunctionReceivers"`
	// The list of email receivers that are part of this action group.
	EmailReceivers []EmailReceiver `pulumi:"emailReceivers"`
	// Indicates whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications.
	Enabled bool `pulumi:"enabled"`
	// The short name of the action group. This will be used in SMS messages.
	GroupShortName string `pulumi:"groupShortName"`
	// The list of ITSM receivers that are part of this action group.
	ItsmReceivers []ItsmReceiver `pulumi:"itsmReceivers"`
	// Resource location
	Location string `pulumi:"location"`
	// The list of logic app receivers that are part of this action group.
	LogicAppReceivers []LogicAppReceiver `pulumi:"logicAppReceivers"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The list of SMS receivers that are part of this action group.
	SmsReceivers []SmsReceiver `pulumi:"smsReceivers"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The list of voice receivers that are part of this action group.
	VoiceReceivers []VoiceReceiver `pulumi:"voiceReceivers"`
	// The list of webhook receivers that are part of this action group.
	WebhookReceivers []WebhookReceiver `pulumi:"webhookReceivers"`
}

// The set of arguments for constructing a ActionGroup resource.
type ActionGroupArgs struct {
	// The name of the action group.
	ActionGroupName pulumi.StringInput
	// The list of AutomationRunbook receivers that are part of this action group.
	AutomationRunbookReceivers AutomationRunbookReceiverArrayInput
	// The list of AzureAppPush receivers that are part of this action group.
	AzureAppPushReceivers AzureAppPushReceiverArrayInput
	// The list of azure function receivers that are part of this action group.
	AzureFunctionReceivers AzureFunctionReceiverArrayInput
	// The list of email receivers that are part of this action group.
	EmailReceivers EmailReceiverArrayInput
	// Indicates whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications.
	Enabled pulumi.BoolInput
	// The short name of the action group. This will be used in SMS messages.
	GroupShortName pulumi.StringInput
	// The list of ITSM receivers that are part of this action group.
	ItsmReceivers ItsmReceiverArrayInput
	// Resource location
	Location pulumi.StringInput
	// The list of logic app receivers that are part of this action group.
	LogicAppReceivers LogicAppReceiverArrayInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The list of SMS receivers that are part of this action group.
	SmsReceivers SmsReceiverArrayInput
	// Resource tags
	Tags pulumi.StringMapInput
	// The list of voice receivers that are part of this action group.
	VoiceReceivers VoiceReceiverArrayInput
	// The list of webhook receivers that are part of this action group.
	WebhookReceivers WebhookReceiverArrayInput
}

func (ActionGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*actionGroupArgs)(nil)).Elem()
}
