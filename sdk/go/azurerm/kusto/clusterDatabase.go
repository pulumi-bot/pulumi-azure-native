// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kusto

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Class representing a Kusto database.
type ClusterDatabase struct {
	pulumi.CustomResourceState

	// Kind of the database
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewClusterDatabase registers a new resource with the given unique name, arguments, and options.
func NewClusterDatabase(ctx *pulumi.Context,
	name string, args *ClusterDatabaseArgs, opts ...pulumi.ResourceOption) (*ClusterDatabase, error) {
	if args == nil || args.ClusterName == nil {
		return nil, errors.New("missing required argument 'ClusterName'")
	}
	if args == nil || args.Kind == nil {
		return nil, errors.New("missing required argument 'Kind'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ClusterDatabaseArgs{}
	}
	var resource ClusterDatabase
	err := ctx.RegisterResource("azurerm:kusto:ClusterDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterDatabase gets an existing ClusterDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterDatabaseState, opts ...pulumi.ResourceOption) (*ClusterDatabase, error) {
	var resource ClusterDatabase
	err := ctx.ReadResource("azurerm:kusto:ClusterDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterDatabase resources.
type clusterDatabaseState struct {
	// Kind of the database
	Kind *string `pulumi:"kind"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type ClusterDatabaseState struct {
	// Kind of the database
	Kind pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (ClusterDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterDatabaseState)(nil)).Elem()
}

type clusterDatabaseArgs struct {
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// Kind of the database
	Kind string `pulumi:"kind"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the database in the Kusto cluster.
	Name string `pulumi:"name"`
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ClusterDatabase resource.
type ClusterDatabaseArgs struct {
	// The name of the Kusto cluster.
	ClusterName pulumi.StringInput
	// Kind of the database
	Kind pulumi.StringInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the database in the Kusto cluster.
	Name pulumi.StringInput
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName pulumi.StringInput
}

func (ClusterDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterDatabaseArgs)(nil)).Elem()
}