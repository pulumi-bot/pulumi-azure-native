// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190701

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Capacity pool resource
//
// ## Example Usage
// ### Pools_CreateOrUpdate
//
// ```go
// package main
//
// import (
// 	netapp "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/netapp/v20190701"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := netapp.NewPool(ctx, "pool", &netapp.PoolArgs{
// 			AccountName:       pulumi.String("accountName"),
// 			Location:          pulumi.String("eastus"),
// 			PoolName:          pulumi.String("poolName"),
// 			ResourceGroupName: pulumi.String("resourceGroup"),
// 			ServiceLevel:      pulumi.String("Premium"),
// 			Size:              pulumi.Int(4398046511104),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
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
	Size pulumi.IntOutput `pulumi:"size"`
	// Resource tags
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPool registers a new resource with the given unique name, arguments, and options.
func NewPool(ctx *pulumi.Context,
	name string, args *PoolArgs, opts ...pulumi.ResourceOption) (*Pool, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.PoolName == nil {
		return nil, errors.New("missing required argument 'PoolName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceLevel == nil {
		return nil, errors.New("missing required argument 'ServiceLevel'")
	}
	if args == nil || args.Size == nil {
		return nil, errors.New("missing required argument 'Size'")
	}
	if args == nil {
		args = &PoolArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:netapp/latest:Pool"),
		},
		{
			Type: pulumi.String("azurerm:netapp/v20170815:Pool"),
		},
		{
			Type: pulumi.String("azurerm:netapp/v20190501:Pool"),
		},
		{
			Type: pulumi.String("azurerm:netapp/v20190601:Pool"),
		},
		{
			Type: pulumi.String("azurerm:netapp/v20190801:Pool"),
		},
		{
			Type: pulumi.String("azurerm:netapp/v20191001:Pool"),
		},
		{
			Type: pulumi.String("azurerm:netapp/v20191101:Pool"),
		},
		{
			Type: pulumi.String("azurerm:netapp/v20200201:Pool"),
		},
		{
			Type: pulumi.String("azurerm:netapp/v20200601:Pool"),
		},
	})
	opts = append(opts, aliases)
	var resource Pool
	err := ctx.RegisterResource("azurerm:netapp/v20190701:Pool", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:netapp/v20190701:Pool", name, id, state, &resource, opts...)
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
	Size *int `pulumi:"size"`
	// Resource tags
	Tags map[string]interface{} `pulumi:"tags"`
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
	Size pulumi.IntPtrInput
	// Resource tags
	Tags pulumi.MapInput
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
	Size int `pulumi:"size"`
	// Resource tags
	Tags map[string]interface{} `pulumi:"tags"`
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
	Size pulumi.IntInput
	// Resource tags
	Tags pulumi.MapInput
}

func (PoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*poolArgs)(nil)).Elem()
}
