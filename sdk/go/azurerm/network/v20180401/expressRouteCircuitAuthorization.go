// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180401

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Authorization in an ExpressRouteCircuit resource.
type ExpressRouteCircuitAuthorization struct {
	pulumi.CustomResourceState

	// The authorization key.
	AuthorizationKey pulumi.StringPtrOutput `pulumi:"authorizationKey"`
	// AuthorizationUseStatus. Possible values are: 'Available' and 'InUse'.
	AuthorizationUseStatus pulumi.StringPtrOutput `pulumi:"authorizationUseStatus"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
}

// NewExpressRouteCircuitAuthorization registers a new resource with the given unique name, arguments, and options.
func NewExpressRouteCircuitAuthorization(ctx *pulumi.Context,
	name string, args *ExpressRouteCircuitAuthorizationArgs, opts ...pulumi.ResourceOption) (*ExpressRouteCircuitAuthorization, error) {
	if args == nil || args.AuthorizationName == nil {
		return nil, errors.New("missing required argument 'AuthorizationName'")
	}
	if args == nil || args.CircuitName == nil {
		return nil, errors.New("missing required argument 'CircuitName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ExpressRouteCircuitAuthorizationArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:network/latest:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20150501preview:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20150615:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20160330:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20160601:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20160901:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20161201:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170301:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170601:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170801:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170901:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20171001:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20171101:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180101:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180201:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180601:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180701:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180801:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181001:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181101:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181201:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190201:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190401:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190601:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190701:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190801:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190901:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191101:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191201:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200301:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200401:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200501:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200601:ExpressRouteCircuitAuthorization"),
		},
	})
	opts = append(opts, aliases)
	var resource ExpressRouteCircuitAuthorization
	err := ctx.RegisterResource("azurerm:network/v20180401:ExpressRouteCircuitAuthorization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExpressRouteCircuitAuthorization gets an existing ExpressRouteCircuitAuthorization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExpressRouteCircuitAuthorization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRouteCircuitAuthorizationState, opts ...pulumi.ResourceOption) (*ExpressRouteCircuitAuthorization, error) {
	var resource ExpressRouteCircuitAuthorization
	err := ctx.ReadResource("azurerm:network/v20180401:ExpressRouteCircuitAuthorization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExpressRouteCircuitAuthorization resources.
type expressRouteCircuitAuthorizationState struct {
	// The authorization key.
	AuthorizationKey *string `pulumi:"authorizationKey"`
	// AuthorizationUseStatus. Possible values are: 'Available' and 'InUse'.
	AuthorizationUseStatus *string `pulumi:"authorizationUseStatus"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState *string `pulumi:"provisioningState"`
}

type ExpressRouteCircuitAuthorizationState struct {
	// The authorization key.
	AuthorizationKey pulumi.StringPtrInput
	// AuthorizationUseStatus. Possible values are: 'Available' and 'InUse'.
	AuthorizationUseStatus pulumi.StringPtrInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringPtrInput
}

func (ExpressRouteCircuitAuthorizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitAuthorizationState)(nil)).Elem()
}

type expressRouteCircuitAuthorizationArgs struct {
	// The authorization key.
	AuthorizationKey *string `pulumi:"authorizationKey"`
	// The name of the authorization.
	AuthorizationName string `pulumi:"authorizationName"`
	// AuthorizationUseStatus. Possible values are: 'Available' and 'InUse'.
	AuthorizationUseStatus *string `pulumi:"authorizationUseStatus"`
	// The name of the express route circuit.
	CircuitName string `pulumi:"circuitName"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ExpressRouteCircuitAuthorization resource.
type ExpressRouteCircuitAuthorizationArgs struct {
	// The authorization key.
	AuthorizationKey pulumi.StringPtrInput
	// The name of the authorization.
	AuthorizationName pulumi.StringInput
	// AuthorizationUseStatus. Possible values are: 'Available' and 'InUse'.
	AuthorizationUseStatus pulumi.StringPtrInput
	// The name of the express route circuit.
	CircuitName pulumi.StringInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (ExpressRouteCircuitAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitAuthorizationArgs)(nil)).Elem()
}
