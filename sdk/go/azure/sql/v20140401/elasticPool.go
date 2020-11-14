// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20140401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents a database elastic pool.
type ElasticPool struct {
	pulumi.CustomResourceState

	// The creation date of the elastic pool (ISO8601 format).
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// The maximum DTU any one database can consume.
	DatabaseDtuMax pulumi.IntPtrOutput `pulumi:"databaseDtuMax"`
	// The minimum DTU all databases are guaranteed.
	DatabaseDtuMin pulumi.IntPtrOutput `pulumi:"databaseDtuMin"`
	// The total shared DTU for the database elastic pool.
	Dtu pulumi.IntPtrOutput `pulumi:"dtu"`
	// The edition of the elastic pool.
	Edition pulumi.StringPtrOutput `pulumi:"edition"`
	// Kind of elastic pool.  This is metadata used for the Azure portal experience.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The state of the elastic pool.
	State pulumi.StringOutput `pulumi:"state"`
	// Gets storage limit for the database elastic pool in MB.
	StorageMB pulumi.IntPtrOutput `pulumi:"storageMB"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Whether or not this database elastic pool is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant pulumi.BoolPtrOutput `pulumi:"zoneRedundant"`
}

// NewElasticPool registers a new resource with the given unique name, arguments, and options.
func NewElasticPool(ctx *pulumi.Context,
	name string, args *ElasticPoolArgs, opts ...pulumi.ResourceOption) (*ElasticPool, error) {
	if args == nil || args.ElasticPoolName == nil {
		return nil, errors.New("missing required argument 'ElasticPoolName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServerName == nil {
		return nil, errors.New("missing required argument 'ServerName'")
	}
	if args == nil {
		args = &ElasticPoolArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:sql/latest:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20171001preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200801preview:ElasticPool"),
		},
	})
	opts = append(opts, aliases)
	var resource ElasticPool
	err := ctx.RegisterResource("azure-nextgen:sql/v20140401:ElasticPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetElasticPool gets an existing ElasticPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetElasticPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ElasticPoolState, opts ...pulumi.ResourceOption) (*ElasticPool, error) {
	var resource ElasticPool
	err := ctx.ReadResource("azure-nextgen:sql/v20140401:ElasticPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ElasticPool resources.
type elasticPoolState struct {
	// The creation date of the elastic pool (ISO8601 format).
	CreationDate *string `pulumi:"creationDate"`
	// The maximum DTU any one database can consume.
	DatabaseDtuMax *int `pulumi:"databaseDtuMax"`
	// The minimum DTU all databases are guaranteed.
	DatabaseDtuMin *int `pulumi:"databaseDtuMin"`
	// The total shared DTU for the database elastic pool.
	Dtu *int `pulumi:"dtu"`
	// The edition of the elastic pool.
	Edition *string `pulumi:"edition"`
	// Kind of elastic pool.  This is metadata used for the Azure portal experience.
	Kind *string `pulumi:"kind"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// The state of the elastic pool.
	State *string `pulumi:"state"`
	// Gets storage limit for the database elastic pool in MB.
	StorageMB *int `pulumi:"storageMB"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
	// Whether or not this database elastic pool is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

type ElasticPoolState struct {
	// The creation date of the elastic pool (ISO8601 format).
	CreationDate pulumi.StringPtrInput
	// The maximum DTU any one database can consume.
	DatabaseDtuMax pulumi.IntPtrInput
	// The minimum DTU all databases are guaranteed.
	DatabaseDtuMin pulumi.IntPtrInput
	// The total shared DTU for the database elastic pool.
	Dtu pulumi.IntPtrInput
	// The edition of the elastic pool.
	Edition pulumi.StringPtrInput
	// Kind of elastic pool.  This is metadata used for the Azure portal experience.
	Kind pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// The state of the elastic pool.
	State pulumi.StringPtrInput
	// Gets storage limit for the database elastic pool in MB.
	StorageMB pulumi.IntPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
	// Whether or not this database elastic pool is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant pulumi.BoolPtrInput
}

func (ElasticPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticPoolState)(nil)).Elem()
}

type elasticPoolArgs struct {
	// The maximum DTU any one database can consume.
	DatabaseDtuMax *int `pulumi:"databaseDtuMax"`
	// The minimum DTU all databases are guaranteed.
	DatabaseDtuMin *int `pulumi:"databaseDtuMin"`
	// The total shared DTU for the database elastic pool.
	Dtu *int `pulumi:"dtu"`
	// The edition of the elastic pool.
	Edition *string `pulumi:"edition"`
	// The name of the elastic pool to be operated on (updated or created).
	ElasticPoolName string `pulumi:"elasticPoolName"`
	// Resource location.
	Location string `pulumi:"location"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// Gets storage limit for the database elastic pool in MB.
	StorageMB *int `pulumi:"storageMB"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Whether or not this database elastic pool is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

// The set of arguments for constructing a ElasticPool resource.
type ElasticPoolArgs struct {
	// The maximum DTU any one database can consume.
	DatabaseDtuMax pulumi.IntPtrInput
	// The minimum DTU all databases are guaranteed.
	DatabaseDtuMin pulumi.IntPtrInput
	// The total shared DTU for the database elastic pool.
	Dtu pulumi.IntPtrInput
	// The edition of the elastic pool.
	Edition pulumi.StringPtrInput
	// The name of the elastic pool to be operated on (updated or created).
	ElasticPoolName pulumi.StringInput
	// Resource location.
	Location pulumi.StringInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// Gets storage limit for the database elastic pool in MB.
	StorageMB pulumi.IntPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Whether or not this database elastic pool is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant pulumi.BoolPtrInput
}

func (ElasticPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticPoolArgs)(nil)).Elem()
}

type ElasticPoolInput interface {
	pulumi.Input

	ToElasticPoolOutput() ElasticPoolOutput
	ToElasticPoolOutputWithContext(ctx context.Context) ElasticPoolOutput
}

func (ElasticPool) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticPool)(nil)).Elem()
}

func (i ElasticPool) ToElasticPoolOutput() ElasticPoolOutput {
	return i.ToElasticPoolOutputWithContext(context.Background())
}

func (i ElasticPool) ToElasticPoolOutputWithContext(ctx context.Context) ElasticPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolOutput)
}

type ElasticPoolOutput struct {
	*pulumi.OutputState
}

func (ElasticPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticPoolOutput)(nil)).Elem()
}

func (o ElasticPoolOutput) ToElasticPoolOutput() ElasticPoolOutput {
	return o
}

func (o ElasticPoolOutput) ToElasticPoolOutputWithContext(ctx context.Context) ElasticPoolOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ElasticPoolOutput{})
}
