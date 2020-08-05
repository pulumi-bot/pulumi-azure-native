// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Inbound NAT rule of the load balancer.
type InboundNatRule struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Properties of load balancer inbound nat rule.
	Properties InboundNatRulePropertiesFormatResponseOutput `pulumi:"properties"`
}

// NewInboundNatRule registers a new resource with the given unique name, arguments, and options.
func NewInboundNatRule(ctx *pulumi.Context,
	name string, args *InboundNatRuleArgs, opts ...pulumi.ResourceOption) (*InboundNatRule, error) {
	if args == nil || args.LoadBalancerName == nil {
		return nil, errors.New("missing required argument 'LoadBalancerName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &InboundNatRuleArgs{}
	}
	var resource InboundNatRule
	err := ctx.RegisterResource("azurerm:network/v20180801:InboundNatRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInboundNatRule gets an existing InboundNatRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInboundNatRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InboundNatRuleState, opts ...pulumi.ResourceOption) (*InboundNatRule, error) {
	var resource InboundNatRule
	err := ctx.ReadResource("azurerm:network/v20180801:InboundNatRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InboundNatRule resources.
type inboundNatRuleState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Properties of load balancer inbound nat rule.
	Properties *InboundNatRulePropertiesFormatResponse `pulumi:"properties"`
}

type InboundNatRuleState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// Properties of load balancer inbound nat rule.
	Properties InboundNatRulePropertiesFormatResponsePtrInput
}

func (InboundNatRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*inboundNatRuleState)(nil)).Elem()
}

type inboundNatRuleArgs struct {
	// The port used for the internal endpoint. Acceptable values range from 1 to 65535.
	BackendPort *int `pulumi:"backendPort"`
	// Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server. This setting can't be changed after you create the endpoint.
	EnableFloatingIP *bool `pulumi:"enableFloatingIP"`
	// Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.
	EnableTcpReset *bool `pulumi:"enableTcpReset"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// A reference to frontend IP addresses.
	FrontendIPConfiguration *SubResource `pulumi:"frontendIPConfiguration"`
	// The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values range from 1 to 65534.
	FrontendPort *int `pulumi:"frontendPort"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to TCP.
	IdleTimeoutInMinutes *int `pulumi:"idleTimeoutInMinutes"`
	// The name of the load balancer.
	LoadBalancerName string `pulumi:"loadBalancerName"`
	// The name of the inbound nat rule.
	Name string `pulumi:"name"`
	// The transport protocol for the endpoint. Possible values are 'Udp' or 'Tcp' or 'All'.
	Protocol *string `pulumi:"protocol"`
	// Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a InboundNatRule resource.
type InboundNatRuleArgs struct {
	// The port used for the internal endpoint. Acceptable values range from 1 to 65535.
	BackendPort pulumi.IntPtrInput
	// Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server. This setting can't be changed after you create the endpoint.
	EnableFloatingIP pulumi.BoolPtrInput
	// Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.
	EnableTcpReset pulumi.BoolPtrInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// A reference to frontend IP addresses.
	FrontendIPConfiguration SubResourcePtrInput
	// The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values range from 1 to 65534.
	FrontendPort pulumi.IntPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to TCP.
	IdleTimeoutInMinutes pulumi.IntPtrInput
	// The name of the load balancer.
	LoadBalancerName pulumi.StringInput
	// The name of the inbound nat rule.
	Name pulumi.StringInput
	// The transport protocol for the endpoint. Possible values are 'Udp' or 'Tcp' or 'All'.
	Protocol pulumi.StringPtrInput
	// Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (InboundNatRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inboundNatRuleArgs)(nil)).Elem()
}
