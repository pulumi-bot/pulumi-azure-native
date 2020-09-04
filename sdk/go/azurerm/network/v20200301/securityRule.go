// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200301

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Network security rule.
type SecurityRule struct {
	pulumi.CustomResourceState

	// The network traffic is allowed or denied.
	Access pulumi.StringOutput `pulumi:"access"`
	// A description for this rule. Restricted to 140 chars.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The destination address prefix. CIDR or destination IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
	DestinationAddressPrefix pulumi.StringPtrOutput `pulumi:"destinationAddressPrefix"`
	// The destination address prefixes. CIDR or destination IP ranges.
	DestinationAddressPrefixes pulumi.StringArrayOutput `pulumi:"destinationAddressPrefixes"`
	// The application security group specified as destination.
	DestinationApplicationSecurityGroups ApplicationSecurityGroupResponseArrayOutput `pulumi:"destinationApplicationSecurityGroups"`
	// The destination port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	DestinationPortRange pulumi.StringPtrOutput `pulumi:"destinationPortRange"`
	// The destination port ranges.
	DestinationPortRanges pulumi.StringArrayOutput `pulumi:"destinationPortRanges"`
	// The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic.
	Direction pulumi.StringOutput `pulumi:"direction"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// Network protocol this rule applies to.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The provisioning state of the security rule resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from.
	SourceAddressPrefix pulumi.StringPtrOutput `pulumi:"sourceAddressPrefix"`
	// The CIDR or source IP ranges.
	SourceAddressPrefixes pulumi.StringArrayOutput `pulumi:"sourceAddressPrefixes"`
	// The application security group specified as source.
	SourceApplicationSecurityGroups ApplicationSecurityGroupResponseArrayOutput `pulumi:"sourceApplicationSecurityGroups"`
	// The source port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	SourcePortRange pulumi.StringPtrOutput `pulumi:"sourcePortRange"`
	// The source port ranges.
	SourcePortRanges pulumi.StringArrayOutput `pulumi:"sourcePortRanges"`
}

// NewSecurityRule registers a new resource with the given unique name, arguments, and options.
func NewSecurityRule(ctx *pulumi.Context,
	name string, args *SecurityRuleArgs, opts ...pulumi.ResourceOption) (*SecurityRule, error) {
	if args == nil || args.Access == nil {
		return nil, errors.New("missing required argument 'Access'")
	}
	if args == nil || args.Direction == nil {
		return nil, errors.New("missing required argument 'Direction'")
	}
	if args == nil || args.NetworkSecurityGroupName == nil {
		return nil, errors.New("missing required argument 'NetworkSecurityGroupName'")
	}
	if args == nil || args.Protocol == nil {
		return nil, errors.New("missing required argument 'Protocol'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SecurityRuleName == nil {
		return nil, errors.New("missing required argument 'SecurityRuleName'")
	}
	if args == nil {
		args = &SecurityRuleArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:network/latest:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20150501preview:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20150615:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20160330:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20160601:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20160901:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20161201:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170301:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170601:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170801:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170901:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20171001:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20171101:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180101:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180201:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180401:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180601:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180701:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180801:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181001:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181101:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181201:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190201:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190401:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190601:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190701:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190801:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190901:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191101:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191201:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200401:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200501:SecurityRule"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200601:SecurityRule"),
		},
	})
	opts = append(opts, aliases)
	var resource SecurityRule
	err := ctx.RegisterResource("azurerm:network/v20200301:SecurityRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityRule gets an existing SecurityRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityRuleState, opts ...pulumi.ResourceOption) (*SecurityRule, error) {
	var resource SecurityRule
	err := ctx.ReadResource("azurerm:network/v20200301:SecurityRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityRule resources.
type securityRuleState struct {
	// The network traffic is allowed or denied.
	Access *string `pulumi:"access"`
	// A description for this rule. Restricted to 140 chars.
	Description *string `pulumi:"description"`
	// The destination address prefix. CIDR or destination IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
	DestinationAddressPrefix *string `pulumi:"destinationAddressPrefix"`
	// The destination address prefixes. CIDR or destination IP ranges.
	DestinationAddressPrefixes []string `pulumi:"destinationAddressPrefixes"`
	// The application security group specified as destination.
	DestinationApplicationSecurityGroups []ApplicationSecurityGroupResponse `pulumi:"destinationApplicationSecurityGroups"`
	// The destination port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	DestinationPortRange *string `pulumi:"destinationPortRange"`
	// The destination port ranges.
	DestinationPortRanges []string `pulumi:"destinationPortRanges"`
	// The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic.
	Direction *string `pulumi:"direction"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
	Priority *int `pulumi:"priority"`
	// Network protocol this rule applies to.
	Protocol *string `pulumi:"protocol"`
	// The provisioning state of the security rule resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from.
	SourceAddressPrefix *string `pulumi:"sourceAddressPrefix"`
	// The CIDR or source IP ranges.
	SourceAddressPrefixes []string `pulumi:"sourceAddressPrefixes"`
	// The application security group specified as source.
	SourceApplicationSecurityGroups []ApplicationSecurityGroupResponse `pulumi:"sourceApplicationSecurityGroups"`
	// The source port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	SourcePortRange *string `pulumi:"sourcePortRange"`
	// The source port ranges.
	SourcePortRanges []string `pulumi:"sourcePortRanges"`
}

type SecurityRuleState struct {
	// The network traffic is allowed or denied.
	Access pulumi.StringPtrInput
	// A description for this rule. Restricted to 140 chars.
	Description pulumi.StringPtrInput
	// The destination address prefix. CIDR or destination IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
	DestinationAddressPrefix pulumi.StringPtrInput
	// The destination address prefixes. CIDR or destination IP ranges.
	DestinationAddressPrefixes pulumi.StringArrayInput
	// The application security group specified as destination.
	DestinationApplicationSecurityGroups ApplicationSecurityGroupResponseArrayInput
	// The destination port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	DestinationPortRange pulumi.StringPtrInput
	// The destination port ranges.
	DestinationPortRanges pulumi.StringArrayInput
	// The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic.
	Direction pulumi.StringPtrInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
	Priority pulumi.IntPtrInput
	// Network protocol this rule applies to.
	Protocol pulumi.StringPtrInput
	// The provisioning state of the security rule resource.
	ProvisioningState pulumi.StringPtrInput
	// The CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from.
	SourceAddressPrefix pulumi.StringPtrInput
	// The CIDR or source IP ranges.
	SourceAddressPrefixes pulumi.StringArrayInput
	// The application security group specified as source.
	SourceApplicationSecurityGroups ApplicationSecurityGroupResponseArrayInput
	// The source port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	SourcePortRange pulumi.StringPtrInput
	// The source port ranges.
	SourcePortRanges pulumi.StringArrayInput
}

func (SecurityRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityRuleState)(nil)).Elem()
}

type securityRuleArgs struct {
	// The network traffic is allowed or denied.
	Access string `pulumi:"access"`
	// A description for this rule. Restricted to 140 chars.
	Description *string `pulumi:"description"`
	// The destination address prefix. CIDR or destination IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
	DestinationAddressPrefix *string `pulumi:"destinationAddressPrefix"`
	// The destination address prefixes. CIDR or destination IP ranges.
	DestinationAddressPrefixes []string `pulumi:"destinationAddressPrefixes"`
	// The application security group specified as destination.
	DestinationApplicationSecurityGroups []ApplicationSecurityGroupType `pulumi:"destinationApplicationSecurityGroups"`
	// The destination port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	DestinationPortRange *string `pulumi:"destinationPortRange"`
	// The destination port ranges.
	DestinationPortRanges []string `pulumi:"destinationPortRanges"`
	// The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic.
	Direction string `pulumi:"direction"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The name of the network security group.
	NetworkSecurityGroupName string `pulumi:"networkSecurityGroupName"`
	// The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
	Priority *int `pulumi:"priority"`
	// Network protocol this rule applies to.
	Protocol string `pulumi:"protocol"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the security rule.
	SecurityRuleName string `pulumi:"securityRuleName"`
	// The CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from.
	SourceAddressPrefix *string `pulumi:"sourceAddressPrefix"`
	// The CIDR or source IP ranges.
	SourceAddressPrefixes []string `pulumi:"sourceAddressPrefixes"`
	// The application security group specified as source.
	SourceApplicationSecurityGroups []ApplicationSecurityGroupType `pulumi:"sourceApplicationSecurityGroups"`
	// The source port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	SourcePortRange *string `pulumi:"sourcePortRange"`
	// The source port ranges.
	SourcePortRanges []string `pulumi:"sourcePortRanges"`
}

// The set of arguments for constructing a SecurityRule resource.
type SecurityRuleArgs struct {
	// The network traffic is allowed or denied.
	Access pulumi.StringInput
	// A description for this rule. Restricted to 140 chars.
	Description pulumi.StringPtrInput
	// The destination address prefix. CIDR or destination IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
	DestinationAddressPrefix pulumi.StringPtrInput
	// The destination address prefixes. CIDR or destination IP ranges.
	DestinationAddressPrefixes pulumi.StringArrayInput
	// The application security group specified as destination.
	DestinationApplicationSecurityGroups ApplicationSecurityGroupTypeArrayInput
	// The destination port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	DestinationPortRange pulumi.StringPtrInput
	// The destination port ranges.
	DestinationPortRanges pulumi.StringArrayInput
	// The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic.
	Direction pulumi.StringInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// The name of the network security group.
	NetworkSecurityGroupName pulumi.StringInput
	// The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
	Priority pulumi.IntPtrInput
	// Network protocol this rule applies to.
	Protocol pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the security rule.
	SecurityRuleName pulumi.StringInput
	// The CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from.
	SourceAddressPrefix pulumi.StringPtrInput
	// The CIDR or source IP ranges.
	SourceAddressPrefixes pulumi.StringArrayInput
	// The application security group specified as source.
	SourceApplicationSecurityGroups ApplicationSecurityGroupTypeArrayInput
	// The source port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	SourcePortRange pulumi.StringPtrInput
	// The source port ranges.
	SourcePortRanges pulumi.StringArrayInput
}

func (SecurityRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityRuleArgs)(nil)).Elem()
}
