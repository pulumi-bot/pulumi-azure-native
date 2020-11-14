// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The DataManager resource.
type DataManager struct {
	pulumi.CustomResourceState

	// Etag of the Resource.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US, East
	// US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is created, but if an identical geo
	// region is specified on update the request will succeed.
	Location pulumi.StringOutput `pulumi:"location"`
	// The Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The sku type.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource
	// (across resource groups).
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDataManager registers a new resource with the given unique name, arguments, and options.
func NewDataManager(ctx *pulumi.Context,
	name string, args *DataManagerArgs, opts ...pulumi.ResourceOption) (*DataManager, error) {
	if args == nil || args.DataManagerName == nil {
		return nil, errors.New("missing required argument 'DataManagerName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &DataManagerArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:hybriddata/latest:DataManager"),
		},
		{
			Type: pulumi.String("azure-nextgen:hybriddata/v20190601:DataManager"),
		},
	})
	opts = append(opts, aliases)
	var resource DataManager
	err := ctx.RegisterResource("azure-nextgen:hybriddata/v20160601:DataManager", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataManager gets an existing DataManager resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataManager(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataManagerState, opts ...pulumi.ResourceOption) (*DataManager, error) {
	var resource DataManager
	err := ctx.ReadResource("azure-nextgen:hybriddata/v20160601:DataManager", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataManager resources.
type dataManagerState struct {
	// Etag of the Resource.
	Etag *string `pulumi:"etag"`
	// The location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US, East
	// US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is created, but if an identical geo
	// region is specified on update the request will succeed.
	Location *string `pulumi:"location"`
	// The Resource Name.
	Name *string `pulumi:"name"`
	// The sku type.
	Sku *SkuResponse `pulumi:"sku"`
	// The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource
	// (across resource groups).
	Tags map[string]string `pulumi:"tags"`
	// The Resource type.
	Type *string `pulumi:"type"`
}

type DataManagerState struct {
	// Etag of the Resource.
	Etag pulumi.StringPtrInput
	// The location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US, East
	// US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is created, but if an identical geo
	// region is specified on update the request will succeed.
	Location pulumi.StringPtrInput
	// The Resource Name.
	Name pulumi.StringPtrInput
	// The sku type.
	Sku SkuResponsePtrInput
	// The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource
	// (across resource groups).
	Tags pulumi.StringMapInput
	// The Resource type.
	Type pulumi.StringPtrInput
}

func (DataManagerState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataManagerState)(nil)).Elem()
}

type dataManagerArgs struct {
	// The name of the DataManager Resource within the specified resource group. DataManager names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
	DataManagerName string `pulumi:"dataManagerName"`
	// Etag of the Resource.
	Etag *string `pulumi:"etag"`
	// The location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US, East
	// US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is created, but if an identical geo
	// region is specified on update the request will succeed.
	Location string `pulumi:"location"`
	// The Resource Group Name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The sku type.
	Sku *Sku `pulumi:"sku"`
	// The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource
	// (across resource groups).
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DataManager resource.
type DataManagerArgs struct {
	// The name of the DataManager Resource within the specified resource group. DataManager names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
	DataManagerName pulumi.StringInput
	// Etag of the Resource.
	Etag pulumi.StringPtrInput
	// The location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US, East
	// US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is created, but if an identical geo
	// region is specified on update the request will succeed.
	Location pulumi.StringInput
	// The Resource Group Name
	ResourceGroupName pulumi.StringInput
	// The sku type.
	Sku SkuPtrInput
	// The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource
	// (across resource groups).
	Tags pulumi.StringMapInput
}

func (DataManagerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataManagerArgs)(nil)).Elem()
}

type DataManagerInput interface {
	pulumi.Input

	ToDataManagerOutput() DataManagerOutput
	ToDataManagerOutputWithContext(ctx context.Context) DataManagerOutput
}

func (DataManager) ElementType() reflect.Type {
	return reflect.TypeOf((*DataManager)(nil)).Elem()
}

func (i DataManager) ToDataManagerOutput() DataManagerOutput {
	return i.ToDataManagerOutputWithContext(context.Background())
}

func (i DataManager) ToDataManagerOutputWithContext(ctx context.Context) DataManagerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataManagerOutput)
}

type DataManagerOutput struct {
	*pulumi.OutputState
}

func (DataManagerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataManagerOutput)(nil)).Elem()
}

func (o DataManagerOutput) ToDataManagerOutput() DataManagerOutput {
	return o
}

func (o DataManagerOutput) ToDataManagerOutputWithContext(ctx context.Context) DataManagerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DataManagerOutput{})
}
