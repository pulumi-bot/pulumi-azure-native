// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dbformysql

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A virtual network rule.
type ServerVirtualNetworkRule struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource properties.
	Properties VirtualNetworkRulePropertiesResponseOutput `pulumi:"properties"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewServerVirtualNetworkRule registers a new resource with the given unique name, arguments, and options.
func NewServerVirtualNetworkRule(ctx *pulumi.Context,
	name string, args *ServerVirtualNetworkRuleArgs, opts ...pulumi.ResourceOption) (*ServerVirtualNetworkRule, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServerName == nil {
		return nil, errors.New("missing required argument 'ServerName'")
	}
	if args == nil {
		args = &ServerVirtualNetworkRuleArgs{}
	}
	var resource ServerVirtualNetworkRule
	err := ctx.RegisterResource("azurerm:dbformysql:ServerVirtualNetworkRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerVirtualNetworkRule gets an existing ServerVirtualNetworkRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerVirtualNetworkRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerVirtualNetworkRuleState, opts ...pulumi.ResourceOption) (*ServerVirtualNetworkRule, error) {
	var resource ServerVirtualNetworkRule
	err := ctx.ReadResource("azurerm:dbformysql:ServerVirtualNetworkRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerVirtualNetworkRule resources.
type serverVirtualNetworkRuleState struct {
	// The name of the resource
	Name *string `pulumi:"name"`
	// Resource properties.
	Properties *VirtualNetworkRulePropertiesResponse `pulumi:"properties"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type ServerVirtualNetworkRuleState struct {
	// The name of the resource
	Name pulumi.StringPtrInput
	// Resource properties.
	Properties VirtualNetworkRulePropertiesResponsePtrInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (ServerVirtualNetworkRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverVirtualNetworkRuleState)(nil)).Elem()
}

type serverVirtualNetworkRuleArgs struct {
	// The name of the virtual network rule.
	Name string `pulumi:"name"`
	// Resource properties.
	Properties *VirtualNetworkRuleProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// The set of arguments for constructing a ServerVirtualNetworkRule resource.
type ServerVirtualNetworkRuleArgs struct {
	// The name of the virtual network rule.
	Name pulumi.StringInput
	// Resource properties.
	Properties VirtualNetworkRulePropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
}

func (ServerVirtualNetworkRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverVirtualNetworkRuleArgs)(nil)).Elem()
}