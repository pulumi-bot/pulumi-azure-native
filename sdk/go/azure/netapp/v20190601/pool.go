// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Capacity pool resource
type Pool struct {
	pulumi.CustomResourceState

	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// UUID v4 used to identify the Pool
	PoolId pulumi.StringOutput `pulumi:"poolId"`
	// Azure lifecycle management
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The service level of the file system
	ServiceLevel pulumi.StringOutput `pulumi:"serviceLevel"`
	// Provisioned size of the pool (in bytes). Allowed values are in 4TiB chunks (value must be multiply of 4398046511104).
	Size pulumi.Float64Output `pulumi:"size"`
	// Resource tags
	Tags pulumi.AnyOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPool registers a new resource with the given unique name, arguments, and options.
func NewPool(ctx *pulumi.Context,
	name string, args *PoolArgs, opts ...pulumi.ResourceOption) (*Pool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.PoolName == nil {
		return nil, errors.New("invalid value for required argument 'PoolName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceLevel == nil {
		return nil, errors.New("invalid value for required argument 'ServiceLevel'")
	}
	if args.Size == nil {
		return nil, errors.New("invalid value for required argument 'Size'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:netapp/latest:Pool"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20170815:Pool"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20190501:Pool"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20190701:Pool"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20190801:Pool"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20191001:Pool"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20191101:Pool"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200201:Pool"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200301:Pool"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200501:Pool"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200601:Pool"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200701:Pool"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200801:Pool"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200901:Pool"),
		},
	})
	opts = append(opts, aliases)
	var resource Pool
	err := ctx.RegisterResource("azure-nextgen:netapp/v20190601:Pool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPool gets an existing Pool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PoolState, opts ...pulumi.ResourceOption) (*Pool, error) {
	var resource Pool
	err := ctx.ReadResource("azure-nextgen:netapp/v20190601:Pool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pool resources.
type poolState struct {
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// UUID v4 used to identify the Pool
	PoolId *string `pulumi:"poolId"`
	// Azure lifecycle management
	ProvisioningState *string `pulumi:"provisioningState"`
	// The service level of the file system
	ServiceLevel *string `pulumi:"serviceLevel"`
	// Provisioned size of the pool (in bytes). Allowed values are in 4TiB chunks (value must be multiply of 4398046511104).
	Size *float64 `pulumi:"size"`
	// Resource tags
	Tags interface{} `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type PoolState struct {
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// UUID v4 used to identify the Pool
	PoolId pulumi.StringPtrInput
	// Azure lifecycle management
	ProvisioningState pulumi.StringPtrInput
	// The service level of the file system
	ServiceLevel pulumi.StringPtrInput
	// Provisioned size of the pool (in bytes). Allowed values are in 4TiB chunks (value must be multiply of 4398046511104).
	Size pulumi.Float64PtrInput
	// Resource tags
	Tags pulumi.Input
	// Resource type
	Type pulumi.StringPtrInput
}

func (PoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*poolState)(nil)).Elem()
}

type poolArgs struct {
	// The name of the NetApp account
	AccountName string `pulumi:"accountName"`
	// Resource location
	Location string `pulumi:"location"`
	// The name of the capacity pool
	PoolName string `pulumi:"poolName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The service level of the file system
	ServiceLevel string `pulumi:"serviceLevel"`
	// Provisioned size of the pool (in bytes). Allowed values are in 4TiB chunks (value must be multiply of 4398046511104).
	Size float64 `pulumi:"size"`
	// Resource tags
	Tags interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Pool resource.
type PoolArgs struct {
	// The name of the NetApp account
	AccountName pulumi.StringInput
	// Resource location
	Location pulumi.StringInput
	// The name of the capacity pool
	PoolName pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The service level of the file system
	ServiceLevel pulumi.StringInput
	// Provisioned size of the pool (in bytes). Allowed values are in 4TiB chunks (value must be multiply of 4398046511104).
	Size pulumi.Float64Input
	// Resource tags
	Tags pulumi.Input
}

func (PoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*poolArgs)(nil)).Elem()
}

type PoolInput interface {
	pulumi.Input

	ToPoolOutput() PoolOutput
	ToPoolOutputWithContext(ctx context.Context) PoolOutput
}

func (*Pool) ElementType() reflect.Type {
	return reflect.TypeOf((*Pool)(nil))
}

func (i *Pool) ToPoolOutput() PoolOutput {
	return i.ToPoolOutputWithContext(context.Background())
}

func (i *Pool) ToPoolOutputWithContext(ctx context.Context) PoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolOutput)
}

type PoolOutput struct {
	*pulumi.OutputState
}

func (PoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Pool)(nil))
}

func (o PoolOutput) ToPoolOutput() PoolOutput {
	return o
}

func (o PoolOutput) ToPoolOutputWithContext(ctx context.Context) PoolOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PoolOutput{})
}
