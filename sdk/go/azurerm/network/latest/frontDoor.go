// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Front Door represents a collection of backend endpoints to route traffic to along with rules that specify how traffic is sent there.
//
// ## Example Usage
// ### Create or update specific Front Door
//
// ```go
// package main
//
// import (
// 	network "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/network/latest"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := network.NewFrontDoor(ctx, "frontDoor", &network.FrontDoorArgs{
// 			BackendPools: network.BackendPoolArray{
// 				&network.BackendPoolArgs{
// 					Name: pulumi.String("backendPool1"),
// 				},
// 			},
// 			BackendPoolsSettings: &network.BackendPoolsSettingsArgs{
// 				EnforceCertificateNameCheck: pulumi.String("Enabled"),
// 				SendRecvTimeoutSeconds:      pulumi.Int(60),
// 			},
// 			EnabledState:  pulumi.String("Enabled"),
// 			FrontDoorName: pulumi.String("frontDoor1"),
// 			FrontendEndpoints: network.FrontendEndpointArray{
// 				&network.FrontendEndpointArgs{
// 					Name: pulumi.String("frontendEndpoint1"),
// 				},
// 				&network.FrontendEndpointArgs{
// 					Name: pulumi.String("default"),
// 				},
// 			},
// 			HealthProbeSettings: network.HealthProbeSettingsModelArray{
// 				&network.HealthProbeSettingsModelArgs{
// 					Name: pulumi.String("healthProbeSettings1"),
// 				},
// 			},
// 			LoadBalancingSettings: network.LoadBalancingSettingsModelArray{
// 				&network.LoadBalancingSettingsModelArgs{
// 					Name: pulumi.String("loadBalancingSettings1"),
// 				},
// 			},
// 			Location:          pulumi.String("westus"),
// 			ResourceGroupName: pulumi.String("rg1"),
// 			RoutingRules: network.RoutingRuleArray{
// 				&network.RoutingRuleArgs{
// 					Name: pulumi.String("routingRule1"),
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"tag1": pulumi.String("value1"),
// 				"tag2": pulumi.String("value2"),
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
type FrontDoor struct {
	pulumi.CustomResourceState

	// Backend pools available to routing rules.
	BackendPools BackendPoolResponseArrayOutput `pulumi:"backendPools"`
	// Settings for all backendPools
	BackendPoolsSettings BackendPoolsSettingsResponsePtrOutput `pulumi:"backendPoolsSettings"`
	// The host that each frontendEndpoint must CNAME to.
	Cname pulumi.StringOutput `pulumi:"cname"`
	// Operational status of the Front Door load balancer. Permitted values are 'Enabled' or 'Disabled'
	EnabledState pulumi.StringPtrOutput `pulumi:"enabledState"`
	// A friendly name for the frontDoor
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	// The Id of the frontdoor.
	FrontdoorId pulumi.StringOutput `pulumi:"frontdoorId"`
	// Frontend endpoints available to routing rules.
	FrontendEndpoints FrontendEndpointResponseArrayOutput `pulumi:"frontendEndpoints"`
	// Health probe settings associated with this Front Door instance.
	HealthProbeSettings HealthProbeSettingsModelResponseArrayOutput `pulumi:"healthProbeSettings"`
	// Load balancing settings associated with this Front Door instance.
	LoadBalancingSettings LoadBalancingSettingsModelResponseArrayOutput `pulumi:"loadBalancingSettings"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the Front Door.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource status of the Front Door.
	ResourceState pulumi.StringPtrOutput `pulumi:"resourceState"`
	// Routing rules associated with this Front Door.
	RoutingRules RoutingRuleResponseArrayOutput `pulumi:"routingRules"`
	// Rules Engine Configurations available to routing rules.
	RulesEngines RulesEngineResponseArrayOutput `pulumi:"rulesEngines"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFrontDoor registers a new resource with the given unique name, arguments, and options.
func NewFrontDoor(ctx *pulumi.Context,
	name string, args *FrontDoorArgs, opts ...pulumi.ResourceOption) (*FrontDoor, error) {
	if args == nil || args.FrontDoorName == nil {
		return nil, errors.New("missing required argument 'FrontDoorName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &FrontDoorArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:network/v20180801:FrontDoor"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190401:FrontDoor"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190501:FrontDoor"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200101:FrontDoor"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200401:FrontDoor"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200501:FrontDoor"),
		},
	})
	opts = append(opts, aliases)
	var resource FrontDoor
	err := ctx.RegisterResource("azurerm:network/latest:FrontDoor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFrontDoor gets an existing FrontDoor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFrontDoor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FrontDoorState, opts ...pulumi.ResourceOption) (*FrontDoor, error) {
	var resource FrontDoor
	err := ctx.ReadResource("azurerm:network/latest:FrontDoor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FrontDoor resources.
type frontDoorState struct {
	// Backend pools available to routing rules.
	BackendPools []BackendPoolResponse `pulumi:"backendPools"`
	// Settings for all backendPools
	BackendPoolsSettings *BackendPoolsSettingsResponse `pulumi:"backendPoolsSettings"`
	// The host that each frontendEndpoint must CNAME to.
	Cname *string `pulumi:"cname"`
	// Operational status of the Front Door load balancer. Permitted values are 'Enabled' or 'Disabled'
	EnabledState *string `pulumi:"enabledState"`
	// A friendly name for the frontDoor
	FriendlyName *string `pulumi:"friendlyName"`
	// The Id of the frontdoor.
	FrontdoorId *string `pulumi:"frontdoorId"`
	// Frontend endpoints available to routing rules.
	FrontendEndpoints []FrontendEndpointResponse `pulumi:"frontendEndpoints"`
	// Health probe settings associated with this Front Door instance.
	HealthProbeSettings []HealthProbeSettingsModelResponse `pulumi:"healthProbeSettings"`
	// Load balancing settings associated with this Front Door instance.
	LoadBalancingSettings []LoadBalancingSettingsModelResponse `pulumi:"loadBalancingSettings"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Provisioning state of the Front Door.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Resource status of the Front Door.
	ResourceState *string `pulumi:"resourceState"`
	// Routing rules associated with this Front Door.
	RoutingRules []RoutingRuleResponse `pulumi:"routingRules"`
	// Rules Engine Configurations available to routing rules.
	RulesEngines []RulesEngineResponse `pulumi:"rulesEngines"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type FrontDoorState struct {
	// Backend pools available to routing rules.
	BackendPools BackendPoolResponseArrayInput
	// Settings for all backendPools
	BackendPoolsSettings BackendPoolsSettingsResponsePtrInput
	// The host that each frontendEndpoint must CNAME to.
	Cname pulumi.StringPtrInput
	// Operational status of the Front Door load balancer. Permitted values are 'Enabled' or 'Disabled'
	EnabledState pulumi.StringPtrInput
	// A friendly name for the frontDoor
	FriendlyName pulumi.StringPtrInput
	// The Id of the frontdoor.
	FrontdoorId pulumi.StringPtrInput
	// Frontend endpoints available to routing rules.
	FrontendEndpoints FrontendEndpointResponseArrayInput
	// Health probe settings associated with this Front Door instance.
	HealthProbeSettings HealthProbeSettingsModelResponseArrayInput
	// Load balancing settings associated with this Front Door instance.
	LoadBalancingSettings LoadBalancingSettingsModelResponseArrayInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Provisioning state of the Front Door.
	ProvisioningState pulumi.StringPtrInput
	// Resource status of the Front Door.
	ResourceState pulumi.StringPtrInput
	// Routing rules associated with this Front Door.
	RoutingRules RoutingRuleResponseArrayInput
	// Rules Engine Configurations available to routing rules.
	RulesEngines RulesEngineResponseArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (FrontDoorState) ElementType() reflect.Type {
	return reflect.TypeOf((*frontDoorState)(nil)).Elem()
}

type frontDoorArgs struct {
	// Backend pools available to routing rules.
	BackendPools []BackendPool `pulumi:"backendPools"`
	// Settings for all backendPools
	BackendPoolsSettings *BackendPoolsSettings `pulumi:"backendPoolsSettings"`
	// Operational status of the Front Door load balancer. Permitted values are 'Enabled' or 'Disabled'
	EnabledState *string `pulumi:"enabledState"`
	// A friendly name for the frontDoor
	FriendlyName *string `pulumi:"friendlyName"`
	// Name of the Front Door which is globally unique.
	FrontDoorName string `pulumi:"frontDoorName"`
	// Frontend endpoints available to routing rules.
	FrontendEndpoints []FrontendEndpoint `pulumi:"frontendEndpoints"`
	// Health probe settings associated with this Front Door instance.
	HealthProbeSettings []HealthProbeSettingsModel `pulumi:"healthProbeSettings"`
	// Load balancing settings associated with this Front Door instance.
	LoadBalancingSettings []LoadBalancingSettingsModel `pulumi:"loadBalancingSettings"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource status of the Front Door.
	ResourceState *string `pulumi:"resourceState"`
	// Routing rules associated with this Front Door.
	RoutingRules []RoutingRule `pulumi:"routingRules"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a FrontDoor resource.
type FrontDoorArgs struct {
	// Backend pools available to routing rules.
	BackendPools BackendPoolArrayInput
	// Settings for all backendPools
	BackendPoolsSettings BackendPoolsSettingsPtrInput
	// Operational status of the Front Door load balancer. Permitted values are 'Enabled' or 'Disabled'
	EnabledState pulumi.StringPtrInput
	// A friendly name for the frontDoor
	FriendlyName pulumi.StringPtrInput
	// Name of the Front Door which is globally unique.
	FrontDoorName pulumi.StringInput
	// Frontend endpoints available to routing rules.
	FrontendEndpoints FrontendEndpointArrayInput
	// Health probe settings associated with this Front Door instance.
	HealthProbeSettings HealthProbeSettingsModelArrayInput
	// Load balancing settings associated with this Front Door instance.
	LoadBalancingSettings LoadBalancingSettingsModelArrayInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// Resource status of the Front Door.
	ResourceState pulumi.StringPtrInput
	// Routing rules associated with this Front Door.
	RoutingRules RoutingRuleArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (FrontDoorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*frontDoorArgs)(nil)).Elem()
}
