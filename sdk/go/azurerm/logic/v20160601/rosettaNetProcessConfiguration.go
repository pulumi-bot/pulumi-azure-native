// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The integration account RosettaNet process configuration.
//
// ## Example Usage
// ### Create or update an RosettaNetProcessConfiguration
//
// ```go
// package main
//
// import (
// 	logic "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/logic/v20160601"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := logic.NewRosettaNetProcessConfiguration(ctx, "rosettaNetProcessConfiguration", &logic.RosettaNetProcessConfigurationArgs{
// 			ActivitySettings: &logic.RosettaNetPipActivitySettingsArgs{
// 				AcknowledgmentOfReceiptSettings: &logic.RosettaNetPipAcknowledgmentOfReceiptSettingsArgs{
// 					IsNonRepudiationRequired:   pulumi.Bool(false),
// 					TimeToAcknowledgeInSeconds: pulumi.Int(600),
// 				},
// 				ActivityBehavior: &logic.RosettaNetPipActivityBehaviorArgs{
// 					ActionType:                       pulumi.String("DoubleAction"),
// 					IsAuthorizationRequired:          pulumi.Bool(false),
// 					IsSecuredTransportRequired:       pulumi.Bool(false),
// 					NonRepudiationOfOriginAndContent: pulumi.Bool(false),
// 					PersistentConfidentialityScope:   pulumi.String("None"),
// 					ResponseType:                     pulumi.String("Async"),
// 					RetryCount:                       pulumi.Int(2),
// 					TimeToPerformInSeconds:           pulumi.Int(7200),
// 				},
// 				ActivityType: pulumi.String("RequestResponse"),
// 			},
// 			Description: pulumi.String("Test description"),
// 			InitiatorRoleSettings: &logic.RosettaNetPipRoleSettingsArgs{
// 				Action: pulumi.String("Purchase Order Request"),
// 				BusinessDocument: &logic.RosettaNetPipBusinessDocumentArgs{
// 					Description: pulumi.String("A request to accept a purchase order for fulfillment.."),
// 					Name:        pulumi.String("Purchase Order Request"),
// 					Version:     pulumi.String("V02.02.00"),
// 				},
// 				Description:           pulumi.String("This partner role creates a demand for a product or service."),
// 				Role:                  pulumi.String("Buyer"),
// 				RoleType:              pulumi.String("Functional"),
// 				Service:               pulumi.String("Buyer Service"),
// 				ServiceClassification: pulumi.String("Business Service"),
// 			},
// 			IntegrationAccountName: pulumi.String("testia123"),
// 			ProcessCode:            pulumi.String("3A4"),
// 			ProcessName:            pulumi.String("Request Purchase Order"),
// 			ProcessVersion:         pulumi.String("V02.02.00"),
// 			ResourceGroupName:      pulumi.String("testrg123"),
// 			ResponderRoleSettings: &logic.RosettaNetPipRoleSettingsArgs{
// 				Action: pulumi.String("Purchase Order Confirmation Action"),
// 				BusinessDocument: &logic.RosettaNetPipBusinessDocumentArgs{
// 					Description: pulumi.String("Formally confirms the status of line item(s) in a Purchase Order. A Purchase Order line item may have one of the following states: accepted, rejected, or pending."),
// 					Name:        pulumi.String("Purchase Order Confirmation"),
// 					Version:     pulumi.String("V02.02.00"),
// 				},
// 				Description:           pulumi.String("An organization that sells products to partners in the supply chain."),
// 				Role:                  pulumi.String("Seller"),
// 				RoleType:              pulumi.String("Organizational"),
// 				Service:               pulumi.String("Seller Service"),
// 				ServiceClassification: pulumi.String("Business Service"),
// 			},
// 			RosettaNetProcessConfigurationName: pulumi.String("3A4"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type RosettaNetProcessConfiguration struct {
	pulumi.CustomResourceState

	// The RosettaNet process configuration activity settings.
	ActivitySettings RosettaNetPipActivitySettingsResponseOutput `pulumi:"activitySettings"`
	// The changed time.
	ChangedTime pulumi.StringOutput `pulumi:"changedTime"`
	// The created time.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// The integration account RosettaNet ProcessConfiguration properties.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The RosettaNet initiator role settings.
	InitiatorRoleSettings RosettaNetPipRoleSettingsResponseOutput `pulumi:"initiatorRoleSettings"`
	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The metadata.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// Gets the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The integration account RosettaNet process code.
	ProcessCode pulumi.StringOutput `pulumi:"processCode"`
	// The integration account RosettaNet process name.
	ProcessName pulumi.StringOutput `pulumi:"processName"`
	// The integration account RosettaNet process version.
	ProcessVersion pulumi.StringOutput `pulumi:"processVersion"`
	// The RosettaNet responder role settings.
	ResponderRoleSettings RosettaNetPipRoleSettingsResponseOutput `pulumi:"responderRoleSettings"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Gets the resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRosettaNetProcessConfiguration registers a new resource with the given unique name, arguments, and options.
func NewRosettaNetProcessConfiguration(ctx *pulumi.Context,
	name string, args *RosettaNetProcessConfigurationArgs, opts ...pulumi.ResourceOption) (*RosettaNetProcessConfiguration, error) {
	if args == nil || args.ActivitySettings == nil {
		return nil, errors.New("missing required argument 'ActivitySettings'")
	}
	if args == nil || args.InitiatorRoleSettings == nil {
		return nil, errors.New("missing required argument 'InitiatorRoleSettings'")
	}
	if args == nil || args.IntegrationAccountName == nil {
		return nil, errors.New("missing required argument 'IntegrationAccountName'")
	}
	if args == nil || args.ProcessCode == nil {
		return nil, errors.New("missing required argument 'ProcessCode'")
	}
	if args == nil || args.ProcessName == nil {
		return nil, errors.New("missing required argument 'ProcessName'")
	}
	if args == nil || args.ProcessVersion == nil {
		return nil, errors.New("missing required argument 'ProcessVersion'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ResponderRoleSettings == nil {
		return nil, errors.New("missing required argument 'ResponderRoleSettings'")
	}
	if args == nil || args.RosettaNetProcessConfigurationName == nil {
		return nil, errors.New("missing required argument 'RosettaNetProcessConfigurationName'")
	}
	if args == nil {
		args = &RosettaNetProcessConfigurationArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:logic/latest:RosettaNetProcessConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource RosettaNetProcessConfiguration
	err := ctx.RegisterResource("azurerm:logic/v20160601:RosettaNetProcessConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRosettaNetProcessConfiguration gets an existing RosettaNetProcessConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRosettaNetProcessConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RosettaNetProcessConfigurationState, opts ...pulumi.ResourceOption) (*RosettaNetProcessConfiguration, error) {
	var resource RosettaNetProcessConfiguration
	err := ctx.ReadResource("azurerm:logic/v20160601:RosettaNetProcessConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RosettaNetProcessConfiguration resources.
type rosettaNetProcessConfigurationState struct {
	// The RosettaNet process configuration activity settings.
	ActivitySettings *RosettaNetPipActivitySettingsResponse `pulumi:"activitySettings"`
	// The changed time.
	ChangedTime *string `pulumi:"changedTime"`
	// The created time.
	CreatedTime *string `pulumi:"createdTime"`
	// The integration account RosettaNet ProcessConfiguration properties.
	Description *string `pulumi:"description"`
	// The RosettaNet initiator role settings.
	InitiatorRoleSettings *RosettaNetPipRoleSettingsResponse `pulumi:"initiatorRoleSettings"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The metadata.
	Metadata map[string]string `pulumi:"metadata"`
	// Gets the resource name.
	Name *string `pulumi:"name"`
	// The integration account RosettaNet process code.
	ProcessCode *string `pulumi:"processCode"`
	// The integration account RosettaNet process name.
	ProcessName *string `pulumi:"processName"`
	// The integration account RosettaNet process version.
	ProcessVersion *string `pulumi:"processVersion"`
	// The RosettaNet responder role settings.
	ResponderRoleSettings *RosettaNetPipRoleSettingsResponse `pulumi:"responderRoleSettings"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets the resource type.
	Type *string `pulumi:"type"`
}

type RosettaNetProcessConfigurationState struct {
	// The RosettaNet process configuration activity settings.
	ActivitySettings RosettaNetPipActivitySettingsResponsePtrInput
	// The changed time.
	ChangedTime pulumi.StringPtrInput
	// The created time.
	CreatedTime pulumi.StringPtrInput
	// The integration account RosettaNet ProcessConfiguration properties.
	Description pulumi.StringPtrInput
	// The RosettaNet initiator role settings.
	InitiatorRoleSettings RosettaNetPipRoleSettingsResponsePtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The metadata.
	Metadata pulumi.StringMapInput
	// Gets the resource name.
	Name pulumi.StringPtrInput
	// The integration account RosettaNet process code.
	ProcessCode pulumi.StringPtrInput
	// The integration account RosettaNet process name.
	ProcessName pulumi.StringPtrInput
	// The integration account RosettaNet process version.
	ProcessVersion pulumi.StringPtrInput
	// The RosettaNet responder role settings.
	ResponderRoleSettings RosettaNetPipRoleSettingsResponsePtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// Gets the resource type.
	Type pulumi.StringPtrInput
}

func (RosettaNetProcessConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*rosettaNetProcessConfigurationState)(nil)).Elem()
}

type rosettaNetProcessConfigurationArgs struct {
	// The RosettaNet process configuration activity settings.
	ActivitySettings RosettaNetPipActivitySettings `pulumi:"activitySettings"`
	// The integration account RosettaNet ProcessConfiguration properties.
	Description *string `pulumi:"description"`
	// The RosettaNet initiator role settings.
	InitiatorRoleSettings RosettaNetPipRoleSettings `pulumi:"initiatorRoleSettings"`
	// The integration account name.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The metadata.
	Metadata map[string]string `pulumi:"metadata"`
	// The integration account RosettaNet process code.
	ProcessCode string `pulumi:"processCode"`
	// The integration account RosettaNet process name.
	ProcessName string `pulumi:"processName"`
	// The integration account RosettaNet process version.
	ProcessVersion string `pulumi:"processVersion"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The RosettaNet responder role settings.
	ResponderRoleSettings RosettaNetPipRoleSettings `pulumi:"responderRoleSettings"`
	// The integration account RosettaNet ProcessConfiguration name.
	RosettaNetProcessConfigurationName string `pulumi:"rosettaNetProcessConfigurationName"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a RosettaNetProcessConfiguration resource.
type RosettaNetProcessConfigurationArgs struct {
	// The RosettaNet process configuration activity settings.
	ActivitySettings RosettaNetPipActivitySettingsInput
	// The integration account RosettaNet ProcessConfiguration properties.
	Description pulumi.StringPtrInput
	// The RosettaNet initiator role settings.
	InitiatorRoleSettings RosettaNetPipRoleSettingsInput
	// The integration account name.
	IntegrationAccountName pulumi.StringInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The metadata.
	Metadata pulumi.StringMapInput
	// The integration account RosettaNet process code.
	ProcessCode pulumi.StringInput
	// The integration account RosettaNet process name.
	ProcessName pulumi.StringInput
	// The integration account RosettaNet process version.
	ProcessVersion pulumi.StringInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The RosettaNet responder role settings.
	ResponderRoleSettings RosettaNetPipRoleSettingsInput
	// The integration account RosettaNet ProcessConfiguration name.
	RosettaNetProcessConfigurationName pulumi.StringInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (RosettaNetProcessConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rosettaNetProcessConfigurationArgs)(nil)).Elem()
}
