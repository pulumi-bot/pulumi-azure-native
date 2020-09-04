// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170601preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An object that represents a replication for a container registry.
type Replication struct {
	pulumi.CustomResourceState

	// The location of the resource. This cannot be changed after the resource is created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the replication at the time the operation was called.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The status of the replication at the time the operation was called.
	Status StatusResponseOutput `pulumi:"status"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewReplication registers a new resource with the given unique name, arguments, and options.
func NewReplication(ctx *pulumi.Context,
	name string, args *ReplicationArgs, opts ...pulumi.ResourceOption) (*Replication, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.RegistryName == nil {
		return nil, errors.New("missing required argument 'RegistryName'")
	}
	if args == nil || args.ReplicationName == nil {
		return nil, errors.New("missing required argument 'ReplicationName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ReplicationArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:containerregistry/latest:Replication"),
		},
		{
			Type: pulumi.String("azurerm:containerregistry/v20171001:Replication"),
		},
		{
			Type: pulumi.String("azurerm:containerregistry/v20190501:Replication"),
		},
		{
			Type: pulumi.String("azurerm:containerregistry/v20191201preview:Replication"),
		},
	})
	opts = append(opts, aliases)
	var resource Replication
	err := ctx.RegisterResource("azurerm:containerregistry/v20170601preview:Replication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplication gets an existing Replication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationState, opts ...pulumi.ResourceOption) (*Replication, error) {
	var resource Replication
	err := ctx.ReadResource("azurerm:containerregistry/v20170601preview:Replication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Replication resources.
type replicationState struct {
	// The location of the resource. This cannot be changed after the resource is created.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The provisioning state of the replication at the time the operation was called.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The status of the replication at the time the operation was called.
	Status *StatusResponse `pulumi:"status"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type ReplicationState struct {
	// The location of the resource. This cannot be changed after the resource is created.
	Location pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The provisioning state of the replication at the time the operation was called.
	ProvisioningState pulumi.StringPtrInput
	// The status of the replication at the time the operation was called.
	Status StatusResponsePtrInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (ReplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationState)(nil)).Elem()
}

type replicationArgs struct {
	// The location of the resource. This cannot be changed after the resource is created.
	Location string `pulumi:"location"`
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the replication.
	ReplicationName string `pulumi:"replicationName"`
	// The name of the resource group to which the container registry belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Replication resource.
type ReplicationArgs struct {
	// The location of the resource. This cannot be changed after the resource is created.
	Location pulumi.StringInput
	// The name of the container registry.
	RegistryName pulumi.StringInput
	// The name of the replication.
	ReplicationName pulumi.StringInput
	// The name of the resource group to which the container registry belongs.
	ResourceGroupName pulumi.StringInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
}

func (ReplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationArgs)(nil)).Elem()
}
