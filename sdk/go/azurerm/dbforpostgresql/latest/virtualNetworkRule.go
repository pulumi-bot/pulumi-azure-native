// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A virtual network rule.
type VirtualNetworkRule struct {
	pulumi.CustomResourceState

	// Create firewall rule before the virtual network has vnet service endpoint enabled.
	IgnoreMissingVnetServiceEndpoint pulumi.BoolPtrOutput `pulumi:"ignoreMissingVnetServiceEndpoint"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Virtual Network Rule State
	State pulumi.StringOutput `pulumi:"state"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
	// The ARM resource id of the virtual network subnet.
	VirtualNetworkSubnetId pulumi.StringOutput `pulumi:"virtualNetworkSubnetId"`
}

// NewVirtualNetworkRule registers a new resource with the given unique name, arguments, and options.
func NewVirtualNetworkRule(ctx *pulumi.Context,
	name string, args *VirtualNetworkRuleArgs, opts ...pulumi.ResourceOption) (*VirtualNetworkRule, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServerName == nil {
		return nil, errors.New("missing required argument 'ServerName'")
	}
	if args == nil || args.VirtualNetworkRuleName == nil {
		return nil, errors.New("missing required argument 'VirtualNetworkRuleName'")
	}
	if args == nil || args.VirtualNetworkSubnetId == nil {
		return nil, errors.New("missing required argument 'VirtualNetworkSubnetId'")
	}
	if args == nil {
		args = &VirtualNetworkRuleArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:dbforpostgresql/v20171201:VirtualNetworkRule"),
		},
		{
			Type: pulumi.String("azurerm:dbforpostgresql/v20171201preview:VirtualNetworkRule"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualNetworkRule
	err := ctx.RegisterResource("azurerm:dbforpostgresql/latest:VirtualNetworkRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualNetworkRule gets an existing VirtualNetworkRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNetworkRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkRuleState, opts ...pulumi.ResourceOption) (*VirtualNetworkRule, error) {
	var resource VirtualNetworkRule
	err := ctx.ReadResource("azurerm:dbforpostgresql/latest:VirtualNetworkRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualNetworkRule resources.
type virtualNetworkRuleState struct {
	// Create firewall rule before the virtual network has vnet service endpoint enabled.
	IgnoreMissingVnetServiceEndpoint *bool `pulumi:"ignoreMissingVnetServiceEndpoint"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// Virtual Network Rule State
	State *string `pulumi:"state"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
	// The ARM resource id of the virtual network subnet.
	VirtualNetworkSubnetId *string `pulumi:"virtualNetworkSubnetId"`
}

type VirtualNetworkRuleState struct {
	// Create firewall rule before the virtual network has vnet service endpoint enabled.
	IgnoreMissingVnetServiceEndpoint pulumi.BoolPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// Virtual Network Rule State
	State pulumi.StringPtrInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
	// The ARM resource id of the virtual network subnet.
	VirtualNetworkSubnetId pulumi.StringPtrInput
}

func (VirtualNetworkRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkRuleState)(nil)).Elem()
}

type virtualNetworkRuleArgs struct {
	// Create firewall rule before the virtual network has vnet service endpoint enabled.
	IgnoreMissingVnetServiceEndpoint *bool `pulumi:"ignoreMissingVnetServiceEndpoint"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The name of the virtual network rule.
	VirtualNetworkRuleName string `pulumi:"virtualNetworkRuleName"`
	// The ARM resource id of the virtual network subnet.
	VirtualNetworkSubnetId string `pulumi:"virtualNetworkSubnetId"`
}

// The set of arguments for constructing a VirtualNetworkRule resource.
type VirtualNetworkRuleArgs struct {
	// Create firewall rule before the virtual network has vnet service endpoint enabled.
	IgnoreMissingVnetServiceEndpoint pulumi.BoolPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The name of the virtual network rule.
	VirtualNetworkRuleName pulumi.StringInput
	// The ARM resource id of the virtual network subnet.
	VirtualNetworkSubnetId pulumi.StringInput
}

func (VirtualNetworkRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkRuleArgs)(nil)).Elem()
}
