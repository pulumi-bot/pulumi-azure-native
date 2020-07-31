// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190907

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Class representing an attached database configuration.
type AttachedDatabaseConfiguration struct {
	pulumi.CustomResourceState

	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of the attached database configuration.
	Properties AttachedDatabaseConfigurationPropertiesResponseOutput `pulumi:"properties"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAttachedDatabaseConfiguration registers a new resource with the given unique name, arguments, and options.
func NewAttachedDatabaseConfiguration(ctx *pulumi.Context,
	name string, args *AttachedDatabaseConfigurationArgs, opts ...pulumi.ResourceOption) (*AttachedDatabaseConfiguration, error) {
	if args == nil || args.ClusterName == nil {
		return nil, errors.New("missing required argument 'ClusterName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &AttachedDatabaseConfigurationArgs{}
	}
	var resource AttachedDatabaseConfiguration
	err := ctx.RegisterResource("azurerm:kusto/v20190907:AttachedDatabaseConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttachedDatabaseConfiguration gets an existing AttachedDatabaseConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttachedDatabaseConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttachedDatabaseConfigurationState, opts ...pulumi.ResourceOption) (*AttachedDatabaseConfiguration, error) {
	var resource AttachedDatabaseConfiguration
	err := ctx.ReadResource("azurerm:kusto/v20190907:AttachedDatabaseConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AttachedDatabaseConfiguration resources.
type attachedDatabaseConfigurationState struct {
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The properties of the attached database configuration.
	Properties *AttachedDatabaseConfigurationPropertiesResponse `pulumi:"properties"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type AttachedDatabaseConfigurationState struct {
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The properties of the attached database configuration.
	Properties AttachedDatabaseConfigurationPropertiesResponsePtrInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (AttachedDatabaseConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*attachedDatabaseConfigurationState)(nil)).Elem()
}

type attachedDatabaseConfigurationArgs struct {
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the attached database configuration.
	Name string `pulumi:"name"`
	// The properties of the attached database configuration.
	Properties *AttachedDatabaseConfigurationProperties `pulumi:"properties"`
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a AttachedDatabaseConfiguration resource.
type AttachedDatabaseConfigurationArgs struct {
	// The name of the Kusto cluster.
	ClusterName pulumi.StringInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the attached database configuration.
	Name pulumi.StringInput
	// The properties of the attached database configuration.
	Properties AttachedDatabaseConfigurationPropertiesPtrInput
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName pulumi.StringInput
}

func (AttachedDatabaseConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attachedDatabaseConfigurationArgs)(nil)).Elem()
}