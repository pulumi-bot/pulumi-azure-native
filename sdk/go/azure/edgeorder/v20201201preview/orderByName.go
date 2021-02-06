// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20201201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents order contract
type OrderByName struct {
	pulumi.CustomResourceState

	// Represents shipping and return address for order
	AddressDetails AddressDetailsResponseOutput `pulumi:"addressDetails"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Id of the order collection to which order belongs to
	OrderCollectionId pulumi.StringOutput `pulumi:"orderCollectionId"`
	// Represents order details.
	OrderDetails OrderDetailsResponseOutput `pulumi:"orderDetails"`
	// Start time of order
	StartTime pulumi.StringOutput `pulumi:"startTime"`
	// Represents resource creation and update time
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewOrderByName registers a new resource with the given unique name, arguments, and options.
func NewOrderByName(ctx *pulumi.Context,
	name string, args *OrderByNameArgs, opts ...pulumi.ResourceOption) (*OrderByName, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AddressDetails == nil {
		return nil, errors.New("invalid value for required argument 'AddressDetails'")
	}
	if args.OrderDetails == nil {
		return nil, errors.New("invalid value for required argument 'OrderDetails'")
	}
	if args.OrderName == nil {
		return nil, errors.New("invalid value for required argument 'OrderName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource OrderByName
	err := ctx.RegisterResource("azure-nextgen:edgeorder/v20201201preview:OrderByName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrderByName gets an existing OrderByName resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrderByName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrderByNameState, opts ...pulumi.ResourceOption) (*OrderByName, error) {
	var resource OrderByName
	err := ctx.ReadResource("azure-nextgen:edgeorder/v20201201preview:OrderByName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrderByName resources.
type orderByNameState struct {
	// Represents shipping and return address for order
	AddressDetails *AddressDetailsResponse `pulumi:"addressDetails"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// Id of the order collection to which order belongs to
	OrderCollectionId *string `pulumi:"orderCollectionId"`
	// Represents order details.
	OrderDetails *OrderDetailsResponse `pulumi:"orderDetails"`
	// Start time of order
	StartTime *string `pulumi:"startTime"`
	// Represents resource creation and update time
	SystemData *SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `pulumi:"type"`
}

type OrderByNameState struct {
	// Represents shipping and return address for order
	AddressDetails AddressDetailsResponsePtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// Id of the order collection to which order belongs to
	OrderCollectionId pulumi.StringPtrInput
	// Represents order details.
	OrderDetails OrderDetailsResponsePtrInput
	// Start time of order
	StartTime pulumi.StringPtrInput
	// Represents resource creation and update time
	SystemData SystemDataResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringPtrInput
}

func (OrderByNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*orderByNameState)(nil)).Elem()
}

type orderByNameArgs struct {
	// Represents shipping and return address for order
	AddressDetails AddressDetails `pulumi:"addressDetails"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Represents order details.
	OrderDetails OrderDetails `pulumi:"orderDetails"`
	// The name of the order
	OrderName string `pulumi:"orderName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a OrderByName resource.
type OrderByNameArgs struct {
	// Represents shipping and return address for order
	AddressDetails AddressDetailsInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Represents order details.
	OrderDetails OrderDetailsInput
	// The name of the order
	OrderName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (OrderByNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*orderByNameArgs)(nil)).Elem()
}

type OrderByNameInput interface {
	pulumi.Input

	ToOrderByNameOutput() OrderByNameOutput
	ToOrderByNameOutputWithContext(ctx context.Context) OrderByNameOutput
}

func (*OrderByName) ElementType() reflect.Type {
	return reflect.TypeOf((*OrderByName)(nil))
}

func (i *OrderByName) ToOrderByNameOutput() OrderByNameOutput {
	return i.ToOrderByNameOutputWithContext(context.Background())
}

func (i *OrderByName) ToOrderByNameOutputWithContext(ctx context.Context) OrderByNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrderByNameOutput)
}

type OrderByNameOutput struct {
	*pulumi.OutputState
}

func (OrderByNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrderByName)(nil))
}

func (o OrderByNameOutput) ToOrderByNameOutput() OrderByNameOutput {
	return o
}

func (o OrderByNameOutput) ToOrderByNameOutputWithContext(ctx context.Context) OrderByNameOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OrderByNameOutput{})
}
