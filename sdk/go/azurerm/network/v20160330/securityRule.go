// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160330

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Network security rule
type SecurityRule struct {
	pulumi.CustomResourceState

	// Gets or sets network traffic is allowed or denied. Possible values are 'Allow' and 'Deny'
	Access pulumi.StringOutput `pulumi:"access"`
	// Gets or sets a description for this rule. Restricted to 140 chars.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Gets or sets destination address prefix. CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
	DestinationAddressPrefix pulumi.StringOutput `pulumi:"destinationAddressPrefix"`
	// Gets or sets Destination Port or Range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	DestinationPortRange pulumi.StringPtrOutput `pulumi:"destinationPortRange"`
	// Gets or sets the direction of the rule.InBound or Outbound. The direction specifies if rule will be evaluated on incoming or outgoing traffic.
	Direction pulumi.StringOutput `pulumi:"direction"`
	// A unique read-only string that changes whenever the resource is updated
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Gets name of the resource that is unique within a resource group. This name can be used to access the resource
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Gets or sets the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// Gets or sets Network protocol this rule applies to. Can be Tcp, Udp or All(*).
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Gets or sets source address prefix. CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from.
	SourceAddressPrefix pulumi.StringOutput `pulumi:"sourceAddressPrefix"`
	// Gets or sets Source Port or Range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	SourcePortRange pulumi.StringPtrOutput `pulumi:"sourcePortRange"`
}

// NewSecurityRule registers a new resource with the given unique name, arguments, and options.
func NewSecurityRule(ctx *pulumi.Context,
	name string, args *SecurityRuleArgs, opts ...pulumi.ResourceOption) (*SecurityRule, error) {
	if args == nil || args.Access == nil {
		return nil, errors.New("missing required argument 'Access'")
	}
	if args == nil || args.DestinationAddressPrefix == nil {
		return nil, errors.New("missing required argument 'DestinationAddressPrefix'")
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
	if args == nil || args.SourceAddressPrefix == nil {
		return nil, errors.New("missing required argument 'SourceAddressPrefix'")
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
			Type: pulumi.String("azurerm:network/v20200301:SecurityRule"),
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
	err := ctx.RegisterResource("azurerm:network/v20160330:SecurityRule", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:network/v20160330:SecurityRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityRule resources.
type securityRuleState struct {
	// Gets or sets network traffic is allowed or denied. Possible values are 'Allow' and 'Deny'
	Access *string `pulumi:"access"`
	// Gets or sets a description for this rule. Restricted to 140 chars.
	Description *string `pulumi:"description"`
	// Gets or sets destination address prefix. CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
	DestinationAddressPrefix *string `pulumi:"destinationAddressPrefix"`
	// Gets or sets Destination Port or Range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	DestinationPortRange *string `pulumi:"destinationPortRange"`
	// Gets or sets the direction of the rule.InBound or Outbound. The direction specifies if rule will be evaluated on incoming or outgoing traffic.
	Direction *string `pulumi:"direction"`
	// A unique read-only string that changes whenever the resource is updated
	Etag *string `pulumi:"etag"`
	// Gets name of the resource that is unique within a resource group. This name can be used to access the resource
	Name *string `pulumi:"name"`
	// Gets or sets the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
	Priority *int `pulumi:"priority"`
	// Gets or sets Network protocol this rule applies to. Can be Tcp, Udp or All(*).
	Protocol *string `pulumi:"protocol"`
	// Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
	ProvisioningState *string `pulumi:"provisioningState"`
	// Gets or sets source address prefix. CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from.
	SourceAddressPrefix *string `pulumi:"sourceAddressPrefix"`
	// Gets or sets Source Port or Range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	SourcePortRange *string `pulumi:"sourcePortRange"`
}

type SecurityRuleState struct {
	// Gets or sets network traffic is allowed or denied. Possible values are 'Allow' and 'Deny'
	Access pulumi.StringPtrInput
	// Gets or sets a description for this rule. Restricted to 140 chars.
	Description pulumi.StringPtrInput
	// Gets or sets destination address prefix. CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
	DestinationAddressPrefix pulumi.StringPtrInput
	// Gets or sets Destination Port or Range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	DestinationPortRange pulumi.StringPtrInput
	// Gets or sets the direction of the rule.InBound or Outbound. The direction specifies if rule will be evaluated on incoming or outgoing traffic.
	Direction pulumi.StringPtrInput
	// A unique read-only string that changes whenever the resource is updated
	Etag pulumi.StringPtrInput
	// Gets name of the resource that is unique within a resource group. This name can be used to access the resource
	Name pulumi.StringPtrInput
	// Gets or sets the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
	Priority pulumi.IntPtrInput
	// Gets or sets Network protocol this rule applies to. Can be Tcp, Udp or All(*).
	Protocol pulumi.StringPtrInput
	// Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
	ProvisioningState pulumi.StringPtrInput
	// Gets or sets source address prefix. CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from.
	SourceAddressPrefix pulumi.StringPtrInput
	// Gets or sets Source Port or Range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	SourcePortRange pulumi.StringPtrInput
}

func (SecurityRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityRuleState)(nil)).Elem()
}

type securityRuleArgs struct {
	// Gets or sets network traffic is allowed or denied. Possible values are 'Allow' and 'Deny'
	Access string `pulumi:"access"`
	// Gets or sets a description for this rule. Restricted to 140 chars.
	Description *string `pulumi:"description"`
	// Gets or sets destination address prefix. CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
	DestinationAddressPrefix string `pulumi:"destinationAddressPrefix"`
	// Gets or sets Destination Port or Range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	DestinationPortRange *string `pulumi:"destinationPortRange"`
	// Gets or sets the direction of the rule.InBound or Outbound. The direction specifies if rule will be evaluated on incoming or outgoing traffic.
	Direction string `pulumi:"direction"`
	// A unique read-only string that changes whenever the resource is updated
	Etag *string `pulumi:"etag"`
	// Resource Id
	Id *string `pulumi:"id"`
	// Gets name of the resource that is unique within a resource group. This name can be used to access the resource
	Name *string `pulumi:"name"`
	// The name of the network security group.
	NetworkSecurityGroupName string `pulumi:"networkSecurityGroupName"`
	// Gets or sets the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
	Priority *int `pulumi:"priority"`
	// Gets or sets Network protocol this rule applies to. Can be Tcp, Udp or All(*).
	Protocol string `pulumi:"protocol"`
	// Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the security rule.
	SecurityRuleName string `pulumi:"securityRuleName"`
	// Gets or sets source address prefix. CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from.
	SourceAddressPrefix string `pulumi:"sourceAddressPrefix"`
	// Gets or sets Source Port or Range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	SourcePortRange *string `pulumi:"sourcePortRange"`
}

// The set of arguments for constructing a SecurityRule resource.
type SecurityRuleArgs struct {
	// Gets or sets network traffic is allowed or denied. Possible values are 'Allow' and 'Deny'
	Access pulumi.StringInput
	// Gets or sets a description for this rule. Restricted to 140 chars.
	Description pulumi.StringPtrInput
	// Gets or sets destination address prefix. CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
	DestinationAddressPrefix pulumi.StringInput
	// Gets or sets Destination Port or Range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	DestinationPortRange pulumi.StringPtrInput
	// Gets or sets the direction of the rule.InBound or Outbound. The direction specifies if rule will be evaluated on incoming or outgoing traffic.
	Direction pulumi.StringInput
	// A unique read-only string that changes whenever the resource is updated
	Etag pulumi.StringPtrInput
	// Resource Id
	Id pulumi.StringPtrInput
	// Gets name of the resource that is unique within a resource group. This name can be used to access the resource
	Name pulumi.StringPtrInput
	// The name of the network security group.
	NetworkSecurityGroupName pulumi.StringInput
	// Gets or sets the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
	Priority pulumi.IntPtrInput
	// Gets or sets Network protocol this rule applies to. Can be Tcp, Udp or All(*).
	Protocol pulumi.StringInput
	// Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the security rule.
	SecurityRuleName pulumi.StringInput
	// Gets or sets source address prefix. CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from.
	SourceAddressPrefix pulumi.StringInput
	// Gets or sets Source Port or Range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	SourcePortRange pulumi.StringPtrInput
}

func (SecurityRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityRuleArgs)(nil)).Elem()
}
