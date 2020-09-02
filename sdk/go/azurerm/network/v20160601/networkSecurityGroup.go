// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// NetworkSecurityGroup resource
type NetworkSecurityGroup struct {
	pulumi.CustomResourceState

	// Gets or default security rules of network security group
	DefaultSecurityRules SecurityRuleResponseArrayOutput `pulumi:"defaultSecurityRules"`
	// Gets a unique read-only string that changes whenever the resource is updated
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Resource location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets collection of references to Network Interfaces
	NetworkInterfaces NetworkInterfaceResponseArrayOutput `pulumi:"networkInterfaces"`
	// Gets provisioning state of the PublicIP resource Updating/Deleting/Failed
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Gets or sets resource guid property of the network security group resource
	ResourceGuid pulumi.StringPtrOutput `pulumi:"resourceGuid"`
	// Gets or sets security rules of network security group
	SecurityRules SecurityRuleResponseArrayOutput `pulumi:"securityRules"`
	// Gets collection of references to subnets
	Subnets SubnetResponseArrayOutput `pulumi:"subnets"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNetworkSecurityGroup registers a new resource with the given unique name, arguments, and options.
func NewNetworkSecurityGroup(ctx *pulumi.Context,
	name string, args *NetworkSecurityGroupArgs, opts ...pulumi.ResourceOption) (*NetworkSecurityGroup, error) {
	if args == nil || args.NetworkSecurityGroupName == nil {
		return nil, errors.New("missing required argument 'NetworkSecurityGroupName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &NetworkSecurityGroupArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:network/latest:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20150501preview:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20150615:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20160330:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20160901:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20161201:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170301:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170601:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170801:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170901:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20171001:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20171101:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180101:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180201:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180401:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180601:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180701:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180801:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181001:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181101:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181201:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190201:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190401:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190601:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190701:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190801:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190901:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191101:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191201:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200301:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200401:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200501:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200601:NetworkSecurityGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkSecurityGroup
	err := ctx.RegisterResource("azurerm:network/v20160601:NetworkSecurityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkSecurityGroup gets an existing NetworkSecurityGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkSecurityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkSecurityGroupState, opts ...pulumi.ResourceOption) (*NetworkSecurityGroup, error) {
	var resource NetworkSecurityGroup
	err := ctx.ReadResource("azurerm:network/v20160601:NetworkSecurityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkSecurityGroup resources.
type networkSecurityGroupState struct {
	// Gets or default security rules of network security group
	DefaultSecurityRules []SecurityRuleResponse `pulumi:"defaultSecurityRules"`
	// Gets a unique read-only string that changes whenever the resource is updated
	Etag *string `pulumi:"etag"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Gets collection of references to Network Interfaces
	NetworkInterfaces []NetworkInterfaceResponse `pulumi:"networkInterfaces"`
	// Gets provisioning state of the PublicIP resource Updating/Deleting/Failed
	ProvisioningState *string `pulumi:"provisioningState"`
	// Gets or sets resource guid property of the network security group resource
	ResourceGuid *string `pulumi:"resourceGuid"`
	// Gets or sets security rules of network security group
	SecurityRules []SecurityRuleResponse `pulumi:"securityRules"`
	// Gets collection of references to subnets
	Subnets []SubnetResponse `pulumi:"subnets"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type NetworkSecurityGroupState struct {
	// Gets or default security rules of network security group
	DefaultSecurityRules SecurityRuleResponseArrayInput
	// Gets a unique read-only string that changes whenever the resource is updated
	Etag pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Gets collection of references to Network Interfaces
	NetworkInterfaces NetworkInterfaceResponseArrayInput
	// Gets provisioning state of the PublicIP resource Updating/Deleting/Failed
	ProvisioningState pulumi.StringPtrInput
	// Gets or sets resource guid property of the network security group resource
	ResourceGuid pulumi.StringPtrInput
	// Gets or sets security rules of network security group
	SecurityRules SecurityRuleResponseArrayInput
	// Gets collection of references to subnets
	Subnets SubnetResponseArrayInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (NetworkSecurityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkSecurityGroupState)(nil)).Elem()
}

type networkSecurityGroupArgs struct {
	// Gets or default security rules of network security group
	DefaultSecurityRules []SecurityRuleType `pulumi:"defaultSecurityRules"`
	// Gets a unique read-only string that changes whenever the resource is updated
	Etag *string `pulumi:"etag"`
	// Resource Id
	Id *string `pulumi:"id"`
	// Resource location
	Location *string `pulumi:"location"`
	// The name of the network security group.
	NetworkSecurityGroupName string `pulumi:"networkSecurityGroupName"`
	// Gets provisioning state of the PublicIP resource Updating/Deleting/Failed
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets resource guid property of the network security group resource
	ResourceGuid *string `pulumi:"resourceGuid"`
	// Gets or sets security rules of network security group
	SecurityRules []SecurityRuleType `pulumi:"securityRules"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkSecurityGroup resource.
type NetworkSecurityGroupArgs struct {
	// Gets or default security rules of network security group
	DefaultSecurityRules SecurityRuleTypeArrayInput
	// Gets a unique read-only string that changes whenever the resource is updated
	Etag pulumi.StringPtrInput
	// Resource Id
	Id pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// The name of the network security group.
	NetworkSecurityGroupName pulumi.StringInput
	// Gets provisioning state of the PublicIP resource Updating/Deleting/Failed
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Gets or sets resource guid property of the network security group resource
	ResourceGuid pulumi.StringPtrInput
	// Gets or sets security rules of network security group
	SecurityRules SecurityRuleTypeArrayInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (NetworkSecurityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkSecurityGroupArgs)(nil)).Elem()
}
