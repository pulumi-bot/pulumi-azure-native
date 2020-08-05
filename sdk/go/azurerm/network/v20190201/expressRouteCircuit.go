// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190201

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ExpressRouteCircuit resource
type ExpressRouteCircuit struct {
	pulumi.CustomResourceState

	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the express route circuit.
	Properties ExpressRouteCircuitPropertiesFormatResponseOutput `pulumi:"properties"`
	// The SKU.
	Sku ExpressRouteCircuitSkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewExpressRouteCircuit registers a new resource with the given unique name, arguments, and options.
func NewExpressRouteCircuit(ctx *pulumi.Context,
	name string, args *ExpressRouteCircuitArgs, opts ...pulumi.ResourceOption) (*ExpressRouteCircuit, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ExpressRouteCircuitArgs{}
	}
	var resource ExpressRouteCircuit
	err := ctx.RegisterResource("azurerm:network/v20190201:ExpressRouteCircuit", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExpressRouteCircuit gets an existing ExpressRouteCircuit resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExpressRouteCircuit(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRouteCircuitState, opts ...pulumi.ResourceOption) (*ExpressRouteCircuit, error) {
	var resource ExpressRouteCircuit
	err := ctx.ReadResource("azurerm:network/v20190201:ExpressRouteCircuit", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExpressRouteCircuit resources.
type expressRouteCircuitState struct {
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Properties of the express route circuit.
	Properties *ExpressRouteCircuitPropertiesFormatResponse `pulumi:"properties"`
	// The SKU.
	Sku *ExpressRouteCircuitSkuResponse `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type ExpressRouteCircuitState struct {
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Properties of the express route circuit.
	Properties ExpressRouteCircuitPropertiesFormatResponsePtrInput
	// The SKU.
	Sku ExpressRouteCircuitSkuResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (ExpressRouteCircuitState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitState)(nil)).Elem()
}

type expressRouteCircuitArgs struct {
	// Allow classic operations
	AllowClassicOperations *bool `pulumi:"allowClassicOperations"`
	// The list of authorizations.
	Authorizations []ExpressRouteCircuitAuthorizationType `pulumi:"authorizations"`
	// The bandwidth of the circuit when the circuit is provisioned on an ExpressRoutePort resource.
	BandwidthInGbps *float64 `pulumi:"bandwidthInGbps"`
	// The CircuitProvisioningState state of the resource.
	CircuitProvisioningState *string `pulumi:"circuitProvisioningState"`
	// The reference to the ExpressRoutePort resource when the circuit is provisioned on an ExpressRoutePort resource.
	ExpressRoutePort *SubResource `pulumi:"expressRoutePort"`
	// The GatewayManager Etag.
	GatewayManagerEtag *string `pulumi:"gatewayManagerEtag"`
	// Flag denoting Global reach status.
	GlobalReachEnabled *bool `pulumi:"globalReachEnabled"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the circuit.
	Name string `pulumi:"name"`
	// The list of peerings.
	Peerings []ExpressRouteCircuitPeeringType `pulumi:"peerings"`
	// Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The ServiceKey.
	ServiceKey *string `pulumi:"serviceKey"`
	// The ServiceProviderNotes.
	ServiceProviderNotes *string `pulumi:"serviceProviderNotes"`
	// The ServiceProviderProperties.
	ServiceProviderProperties *ExpressRouteCircuitServiceProviderProperties `pulumi:"serviceProviderProperties"`
	// The ServiceProviderProvisioningState state of the resource.
	ServiceProviderProvisioningState *string `pulumi:"serviceProviderProvisioningState"`
	// The SKU.
	Sku *ExpressRouteCircuitSku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ExpressRouteCircuit resource.
type ExpressRouteCircuitArgs struct {
	// Allow classic operations
	AllowClassicOperations pulumi.BoolPtrInput
	// The list of authorizations.
	Authorizations ExpressRouteCircuitAuthorizationTypeArrayInput
	// The bandwidth of the circuit when the circuit is provisioned on an ExpressRoutePort resource.
	BandwidthInGbps pulumi.Float64PtrInput
	// The CircuitProvisioningState state of the resource.
	CircuitProvisioningState pulumi.StringPtrInput
	// The reference to the ExpressRoutePort resource when the circuit is provisioned on an ExpressRoutePort resource.
	ExpressRoutePort SubResourcePtrInput
	// The GatewayManager Etag.
	GatewayManagerEtag pulumi.StringPtrInput
	// Flag denoting Global reach status.
	GlobalReachEnabled pulumi.BoolPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the circuit.
	Name pulumi.StringInput
	// The list of peerings.
	Peerings ExpressRouteCircuitPeeringTypeArrayInput
	// Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The ServiceKey.
	ServiceKey pulumi.StringPtrInput
	// The ServiceProviderNotes.
	ServiceProviderNotes pulumi.StringPtrInput
	// The ServiceProviderProperties.
	ServiceProviderProperties ExpressRouteCircuitServiceProviderPropertiesPtrInput
	// The ServiceProviderProvisioningState state of the resource.
	ServiceProviderProvisioningState pulumi.StringPtrInput
	// The SKU.
	Sku ExpressRouteCircuitSkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ExpressRouteCircuitArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitArgs)(nil)).Elem()
}
