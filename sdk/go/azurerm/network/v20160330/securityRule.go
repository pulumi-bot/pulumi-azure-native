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

	// A unique read-only string that changes whenever the resource is updated
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Gets name of the resource that is unique within a resource group. This name can be used to access the resource
	Name       pulumi.StringPtrOutput                     `pulumi:"name"`
	Properties SecurityRulePropertiesFormatResponseOutput `pulumi:"properties"`
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
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
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
	if args == nil || args.SourceAddressPrefix == nil {
		return nil, errors.New("missing required argument 'SourceAddressPrefix'")
	}
	if args == nil {
		args = &SecurityRuleArgs{}
	}
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
	// A unique read-only string that changes whenever the resource is updated
	Etag *string `pulumi:"etag"`
	// Gets name of the resource that is unique within a resource group. This name can be used to access the resource
	Name       *string                               `pulumi:"name"`
	Properties *SecurityRulePropertiesFormatResponse `pulumi:"properties"`
}

type SecurityRuleState struct {
	// A unique read-only string that changes whenever the resource is updated
	Etag pulumi.StringPtrInput
	// Gets name of the resource that is unique within a resource group. This name can be used to access the resource
	Name       pulumi.StringPtrInput
	Properties SecurityRulePropertiesFormatResponsePtrInput
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
	// The name of the security rule.
	Name string `pulumi:"name"`
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
	// The name of the security rule.
	Name pulumi.StringInput
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
	// Gets or sets source address prefix. CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from.
	SourceAddressPrefix pulumi.StringInput
	// Gets or sets Source Port or Range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	SourcePortRange pulumi.StringPtrInput
}

func (SecurityRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityRuleArgs)(nil)).Elem()
}
