// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Azure Firewall resource.
type AzureFirewall struct {
	pulumi.CustomResourceState

	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the azure firewall.
	Properties AzureFirewallPropertiesFormatResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// A list of availability zones denoting where the resource needs to come from.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewAzureFirewall registers a new resource with the given unique name, arguments, and options.
func NewAzureFirewall(ctx *pulumi.Context,
	name string, args *AzureFirewallArgs, opts ...pulumi.ResourceOption) (*AzureFirewall, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &AzureFirewallArgs{}
	}
	var resource AzureFirewall
	err := ctx.RegisterResource("azurerm:network/v20190601:AzureFirewall", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAzureFirewall gets an existing AzureFirewall resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAzureFirewall(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzureFirewallState, opts ...pulumi.ResourceOption) (*AzureFirewall, error) {
	var resource AzureFirewall
	err := ctx.ReadResource("azurerm:network/v20190601:AzureFirewall", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AzureFirewall resources.
type azureFirewallState struct {
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Properties of the azure firewall.
	Properties *AzureFirewallPropertiesFormatResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
	// A list of availability zones denoting where the resource needs to come from.
	Zones []string `pulumi:"zones"`
}

type AzureFirewallState struct {
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Properties of the azure firewall.
	Properties AzureFirewallPropertiesFormatResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
	// A list of availability zones denoting where the resource needs to come from.
	Zones pulumi.StringArrayInput
}

func (AzureFirewallState) ElementType() reflect.Type {
	return reflect.TypeOf((*azureFirewallState)(nil)).Elem()
}

type azureFirewallArgs struct {
	// Collection of application rule collections used by Azure Firewall.
	ApplicationRuleCollections []AzureFirewallApplicationRuleCollection `pulumi:"applicationRuleCollections"`
	// The firewallPolicy associated with this azure firewall.
	FirewallPolicy *SubResource `pulumi:"firewallPolicy"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// IP configuration of the Azure Firewall resource.
	IpConfigurations []AzureFirewallIPConfiguration `pulumi:"ipConfigurations"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the Azure Firewall.
	Name string `pulumi:"name"`
	// Collection of NAT rule collections used by Azure Firewall.
	NatRuleCollections []AzureFirewallNatRuleCollection `pulumi:"natRuleCollections"`
	// Collection of network rule collections used by Azure Firewall.
	NetworkRuleCollections []AzureFirewallNetworkRuleCollection `pulumi:"networkRuleCollections"`
	// The provisioning state of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The operation mode for Threat Intelligence.
	ThreatIntelMode *string `pulumi:"threatIntelMode"`
	// The virtualHub to which the firewall belongs.
	VirtualHub *SubResource `pulumi:"virtualHub"`
	// A list of availability zones denoting where the resource needs to come from.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a AzureFirewall resource.
type AzureFirewallArgs struct {
	// Collection of application rule collections used by Azure Firewall.
	ApplicationRuleCollections AzureFirewallApplicationRuleCollectionArrayInput
	// The firewallPolicy associated with this azure firewall.
	FirewallPolicy SubResourcePtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// IP configuration of the Azure Firewall resource.
	IpConfigurations AzureFirewallIPConfigurationArrayInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the Azure Firewall.
	Name pulumi.StringInput
	// Collection of NAT rule collections used by Azure Firewall.
	NatRuleCollections AzureFirewallNatRuleCollectionArrayInput
	// Collection of network rule collections used by Azure Firewall.
	NetworkRuleCollections AzureFirewallNetworkRuleCollectionArrayInput
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The operation mode for Threat Intelligence.
	ThreatIntelMode pulumi.StringPtrInput
	// The virtualHub to which the firewall belongs.
	VirtualHub SubResourcePtrInput
	// A list of availability zones denoting where the resource needs to come from.
	Zones pulumi.StringArrayInput
}

func (AzureFirewallArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azureFirewallArgs)(nil)).Elem()
}
