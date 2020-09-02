// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A data set mapping data transfer object.
type DataSetMapping struct {
	pulumi.CustomResourceState

	// Kind of data set mapping.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the azure resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Type of the azure resource
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDataSetMapping registers a new resource with the given unique name, arguments, and options.
func NewDataSetMapping(ctx *pulumi.Context,
	name string, args *DataSetMappingArgs, opts ...pulumi.ResourceOption) (*DataSetMapping, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.DataSetMappingName == nil {
		return nil, errors.New("missing required argument 'DataSetMappingName'")
	}
	if args == nil || args.Kind == nil {
		return nil, errors.New("missing required argument 'Kind'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ShareSubscriptionName == nil {
		return nil, errors.New("missing required argument 'ShareSubscriptionName'")
	}
	if args == nil {
		args = &DataSetMappingArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:datashare/v20191101:DataSetMapping"),
		},
	})
	opts = append(opts, aliases)
	var resource DataSetMapping
	err := ctx.RegisterResource("azurerm:datashare/latest:DataSetMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataSetMapping gets an existing DataSetMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataSetMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataSetMappingState, opts ...pulumi.ResourceOption) (*DataSetMapping, error) {
	var resource DataSetMapping
	err := ctx.ReadResource("azurerm:datashare/latest:DataSetMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataSetMapping resources.
type dataSetMappingState struct {
	// Kind of data set mapping.
	Kind *string `pulumi:"kind"`
	// Name of the azure resource
	Name *string `pulumi:"name"`
	// Type of the azure resource
	Type *string `pulumi:"type"`
}

type DataSetMappingState struct {
	// Kind of data set mapping.
	Kind pulumi.StringPtrInput
	// Name of the azure resource
	Name pulumi.StringPtrInput
	// Type of the azure resource
	Type pulumi.StringPtrInput
}

func (DataSetMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSetMappingState)(nil)).Elem()
}

type dataSetMappingArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The name of the data set mapping to be created.
	DataSetMappingName string `pulumi:"dataSetMappingName"`
	// Kind of data set mapping.
	Kind string `pulumi:"kind"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the share subscription which will hold the data set sink.
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}

// The set of arguments for constructing a DataSetMapping resource.
type DataSetMappingArgs struct {
	// The name of the share account.
	AccountName pulumi.StringInput
	// The name of the data set mapping to be created.
	DataSetMappingName pulumi.StringInput
	// Kind of data set mapping.
	Kind pulumi.StringInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The name of the share subscription which will hold the data set sink.
	ShareSubscriptionName pulumi.StringInput
}

func (DataSetMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSetMappingArgs)(nil)).Elem()
}