// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Front Door represents a collection of backend endpoints to route traffic to along with rules that specify how traffic is sent there.
type FrontDoor struct {
	pulumi.CustomResourceState

	// Backend pools available to routing rules.
	BackendPools BackendPoolResponseArrayOutput `pulumi:"backendPools"`
	// The host that each frontendEndpoint must CNAME to.
	Cname pulumi.StringOutput `pulumi:"cname"`
	// Operational status of the Front Door load balancer. Permitted values are 'Enabled' or 'Disabled'
	EnabledState pulumi.StringPtrOutput `pulumi:"enabledState"`
	// A friendly name for the frontDoor
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
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
			Type: pulumi.String("azurerm:network/latest:FrontDoor"),
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
	err := ctx.RegisterResource("azurerm:network/v20180801:FrontDoor", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:network/v20180801:FrontDoor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FrontDoor resources.
type frontDoorState struct {
	// Backend pools available to routing rules.
	BackendPools []BackendPoolResponse `pulumi:"backendPools"`
	// The host that each frontendEndpoint must CNAME to.
	Cname *string `pulumi:"cname"`
	// Operational status of the Front Door load balancer. Permitted values are 'Enabled' or 'Disabled'
	EnabledState *string `pulumi:"enabledState"`
	// A friendly name for the frontDoor
	FriendlyName *string `pulumi:"friendlyName"`
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
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type FrontDoorState struct {
	// Backend pools available to routing rules.
	BackendPools BackendPoolResponseArrayInput
	// The host that each frontendEndpoint must CNAME to.
	Cname pulumi.StringPtrInput
	// Operational status of the Front Door load balancer. Permitted values are 'Enabled' or 'Disabled'
	EnabledState pulumi.StringPtrInput
	// A friendly name for the frontDoor
	FriendlyName pulumi.StringPtrInput
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
