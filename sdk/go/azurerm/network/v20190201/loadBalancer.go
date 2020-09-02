// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190201

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// LoadBalancer resource
type LoadBalancer struct {
	pulumi.CustomResourceState

	// Collection of backend address pools used by a load balancer
	BackendAddressPools BackendAddressPoolResponseArrayOutput `pulumi:"backendAddressPools"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Object representing the frontend IPs to be used for the load balancer
	FrontendIPConfigurations FrontendIPConfigurationResponseArrayOutput `pulumi:"frontendIPConfigurations"`
	// Defines an external port range for inbound NAT to a single backend port on NICs associated with a load balancer. Inbound NAT rules are created automatically for each NIC associated with the Load Balancer using an external port from this range. Defining an Inbound NAT pool on your Load Balancer is mutually exclusive with defining inbound Nat rules. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an inbound NAT pool. They have to reference individual inbound NAT rules.
	InboundNatPools InboundNatPoolResponseArrayOutput `pulumi:"inboundNatPools"`
	// Collection of inbound NAT Rules used by a load balancer. Defining inbound NAT rules on your load balancer is mutually exclusive with defining an inbound NAT pool. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an Inbound NAT pool. They have to reference individual inbound NAT rules.
	InboundNatRules InboundNatRuleResponseArrayOutput `pulumi:"inboundNatRules"`
	// Object collection representing the load balancing rules Gets the provisioning
	LoadBalancingRules LoadBalancingRuleResponseArrayOutput `pulumi:"loadBalancingRules"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The outbound rules.
	OutboundRules OutboundRuleResponseArrayOutput `pulumi:"outboundRules"`
	// Collection of probe objects used in the load balancer
	Probes ProbeResponseArrayOutput `pulumi:"probes"`
	// Gets the provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// The resource GUID property of the load balancer resource.
	ResourceGuid pulumi.StringPtrOutput `pulumi:"resourceGuid"`
	// The load balancer SKU.
	Sku LoadBalancerSkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewLoadBalancer registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancer(ctx *pulumi.Context,
	name string, args *LoadBalancerArgs, opts ...pulumi.ResourceOption) (*LoadBalancer, error) {
	if args == nil || args.LoadBalancerName == nil {
		return nil, errors.New("missing required argument 'LoadBalancerName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &LoadBalancerArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:network/latest:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20150501preview:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20150615:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20160330:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20160601:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20160901:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20161201:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170301:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170601:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170801:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170901:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20171001:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20171101:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180101:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180201:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180401:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180601:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180701:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180801:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181001:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181101:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181201:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190401:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190601:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190701:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190801:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190901:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191101:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191201:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200301:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200401:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200501:LoadBalancer"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200601:LoadBalancer"),
		},
	})
	opts = append(opts, aliases)
	var resource LoadBalancer
	err := ctx.RegisterResource("azurerm:network/v20190201:LoadBalancer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadBalancer gets an existing LoadBalancer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadBalancer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadBalancerState, opts ...pulumi.ResourceOption) (*LoadBalancer, error) {
	var resource LoadBalancer
	err := ctx.ReadResource("azurerm:network/v20190201:LoadBalancer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoadBalancer resources.
type loadBalancerState struct {
	// Collection of backend address pools used by a load balancer
	BackendAddressPools []BackendAddressPoolResponse `pulumi:"backendAddressPools"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Object representing the frontend IPs to be used for the load balancer
	FrontendIPConfigurations []FrontendIPConfigurationResponse `pulumi:"frontendIPConfigurations"`
	// Defines an external port range for inbound NAT to a single backend port on NICs associated with a load balancer. Inbound NAT rules are created automatically for each NIC associated with the Load Balancer using an external port from this range. Defining an Inbound NAT pool on your Load Balancer is mutually exclusive with defining inbound Nat rules. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an inbound NAT pool. They have to reference individual inbound NAT rules.
	InboundNatPools []InboundNatPoolResponse `pulumi:"inboundNatPools"`
	// Collection of inbound NAT Rules used by a load balancer. Defining inbound NAT rules on your load balancer is mutually exclusive with defining an inbound NAT pool. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an Inbound NAT pool. They have to reference individual inbound NAT rules.
	InboundNatRules []InboundNatRuleResponse `pulumi:"inboundNatRules"`
	// Object collection representing the load balancing rules Gets the provisioning
	LoadBalancingRules []LoadBalancingRuleResponse `pulumi:"loadBalancingRules"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// The outbound rules.
	OutboundRules []OutboundRuleResponse `pulumi:"outboundRules"`
	// Collection of probe objects used in the load balancer
	Probes []ProbeResponse `pulumi:"probes"`
	// Gets the provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The resource GUID property of the load balancer resource.
	ResourceGuid *string `pulumi:"resourceGuid"`
	// The load balancer SKU.
	Sku *LoadBalancerSkuResponse `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type LoadBalancerState struct {
	// Collection of backend address pools used by a load balancer
	BackendAddressPools BackendAddressPoolResponseArrayInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Object representing the frontend IPs to be used for the load balancer
	FrontendIPConfigurations FrontendIPConfigurationResponseArrayInput
	// Defines an external port range for inbound NAT to a single backend port on NICs associated with a load balancer. Inbound NAT rules are created automatically for each NIC associated with the Load Balancer using an external port from this range. Defining an Inbound NAT pool on your Load Balancer is mutually exclusive with defining inbound Nat rules. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an inbound NAT pool. They have to reference individual inbound NAT rules.
	InboundNatPools InboundNatPoolResponseArrayInput
	// Collection of inbound NAT Rules used by a load balancer. Defining inbound NAT rules on your load balancer is mutually exclusive with defining an inbound NAT pool. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an Inbound NAT pool. They have to reference individual inbound NAT rules.
	InboundNatRules InboundNatRuleResponseArrayInput
	// Object collection representing the load balancing rules Gets the provisioning
	LoadBalancingRules LoadBalancingRuleResponseArrayInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// The outbound rules.
	OutboundRules OutboundRuleResponseArrayInput
	// Collection of probe objects used in the load balancer
	Probes ProbeResponseArrayInput
	// Gets the provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringPtrInput
	// The resource GUID property of the load balancer resource.
	ResourceGuid pulumi.StringPtrInput
	// The load balancer SKU.
	Sku LoadBalancerSkuResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (LoadBalancerState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerState)(nil)).Elem()
}

type loadBalancerArgs struct {
	// Collection of backend address pools used by a load balancer
	BackendAddressPools []BackendAddressPool `pulumi:"backendAddressPools"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Object representing the frontend IPs to be used for the load balancer
	FrontendIPConfigurations []FrontendIPConfiguration `pulumi:"frontendIPConfigurations"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Defines an external port range for inbound NAT to a single backend port on NICs associated with a load balancer. Inbound NAT rules are created automatically for each NIC associated with the Load Balancer using an external port from this range. Defining an Inbound NAT pool on your Load Balancer is mutually exclusive with defining inbound Nat rules. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an inbound NAT pool. They have to reference individual inbound NAT rules.
	InboundNatPools []InboundNatPool `pulumi:"inboundNatPools"`
	// Collection of inbound NAT Rules used by a load balancer. Defining inbound NAT rules on your load balancer is mutually exclusive with defining an inbound NAT pool. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an Inbound NAT pool. They have to reference individual inbound NAT rules.
	InboundNatRules []InboundNatRuleType `pulumi:"inboundNatRules"`
	// The name of the load balancer.
	LoadBalancerName string `pulumi:"loadBalancerName"`
	// Object collection representing the load balancing rules Gets the provisioning
	LoadBalancingRules []LoadBalancingRule `pulumi:"loadBalancingRules"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The outbound rules.
	OutboundRules []OutboundRule `pulumi:"outboundRules"`
	// Collection of probe objects used in the load balancer
	Probes []Probe `pulumi:"probes"`
	// Gets the provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource GUID property of the load balancer resource.
	ResourceGuid *string `pulumi:"resourceGuid"`
	// The load balancer SKU.
	Sku *LoadBalancerSku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a LoadBalancer resource.
type LoadBalancerArgs struct {
	// Collection of backend address pools used by a load balancer
	BackendAddressPools BackendAddressPoolArrayInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Object representing the frontend IPs to be used for the load balancer
	FrontendIPConfigurations FrontendIPConfigurationArrayInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Defines an external port range for inbound NAT to a single backend port on NICs associated with a load balancer. Inbound NAT rules are created automatically for each NIC associated with the Load Balancer using an external port from this range. Defining an Inbound NAT pool on your Load Balancer is mutually exclusive with defining inbound Nat rules. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an inbound NAT pool. They have to reference individual inbound NAT rules.
	InboundNatPools InboundNatPoolArrayInput
	// Collection of inbound NAT Rules used by a load balancer. Defining inbound NAT rules on your load balancer is mutually exclusive with defining an inbound NAT pool. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an Inbound NAT pool. They have to reference individual inbound NAT rules.
	InboundNatRules InboundNatRuleTypeArrayInput
	// The name of the load balancer.
	LoadBalancerName pulumi.StringInput
	// Object collection representing the load balancing rules Gets the provisioning
	LoadBalancingRules LoadBalancingRuleArrayInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The outbound rules.
	OutboundRules OutboundRuleArrayInput
	// Collection of probe objects used in the load balancer
	Probes ProbeArrayInput
	// Gets the provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The resource GUID property of the load balancer resource.
	ResourceGuid pulumi.StringPtrInput
	// The load balancer SKU.
	Sku LoadBalancerSkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (LoadBalancerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerArgs)(nil)).Elem()
}
