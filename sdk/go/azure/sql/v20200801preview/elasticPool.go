// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An elastic pool.
type ElasticPool struct {
	pulumi.CustomResourceState

	// The creation date of the elastic pool (ISO8601 format).
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// Kind of elastic pool. This is metadata used for the Azure portal experience.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The license type to apply for this elastic pool.
	LicenseType pulumi.StringPtrOutput `pulumi:"licenseType"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Maintenance configuration id assigned to the elastic pool. This configuration defines the period when the maintenance updates will will occur.
	MaintenanceConfigurationId pulumi.StringPtrOutput `pulumi:"maintenanceConfigurationId"`
	// The storage limit for the database elastic pool in bytes.
	MaxSizeBytes pulumi.Float64PtrOutput `pulumi:"maxSizeBytes"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The per database settings for the elastic pool.
	PerDatabaseSettings ElasticPoolPerDatabaseSettingsResponsePtrOutput `pulumi:"perDatabaseSettings"`
	// The elastic pool SKU.
	//
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or the following command:
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// The state of the elastic pool.
	State pulumi.StringOutput `pulumi:"state"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Whether or not this elastic pool is zone redundant, which means the replicas of this elastic pool will be spread across multiple availability zones.
	ZoneRedundant pulumi.BoolPtrOutput `pulumi:"zoneRedundant"`
}

// NewElasticPool registers a new resource with the given unique name, arguments, and options.
func NewElasticPool(ctx *pulumi.Context,
	name string, args *ElasticPoolArgs, opts ...pulumi.ResourceOption) (*ElasticPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ElasticPoolName == nil {
		return nil, errors.New("invalid value for required argument 'ElasticPoolName'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:sql/latest:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20140401:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20171001preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200202preview:ElasticPool"),
		},
	})
	opts = append(opts, aliases)
	var resource ElasticPool
	err := ctx.RegisterResource("azure-nextgen:sql/v20200801preview:ElasticPool", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-nextgen:sql/v20200801preview:ElasticPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ElasticPool resources.
type elasticPoolState struct {
	// The creation date of the elastic pool (ISO8601 format).
	CreationDate *string `pulumi:"creationDate"`
	// Kind of elastic pool. This is metadata used for the Azure portal experience.
	Kind *string `pulumi:"kind"`
	// The license type to apply for this elastic pool.
	LicenseType *string `pulumi:"licenseType"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Maintenance configuration id assigned to the elastic pool. This configuration defines the period when the maintenance updates will will occur.
	MaintenanceConfigurationId *string `pulumi:"maintenanceConfigurationId"`
	// The storage limit for the database elastic pool in bytes.
	MaxSizeBytes *float64 `pulumi:"maxSizeBytes"`
	// Resource name.
	Name *string `pulumi:"name"`
	// The per database settings for the elastic pool.
	PerDatabaseSettings *ElasticPoolPerDatabaseSettingsResponse `pulumi:"perDatabaseSettings"`
	// The elastic pool SKU.
	//
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or the following command:
	Sku *SkuResponse `pulumi:"sku"`
	// The state of the elastic pool.
	State *string `pulumi:"state"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
	// Whether or not this elastic pool is zone redundant, which means the replicas of this elastic pool will be spread across multiple availability zones.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

type ElasticPoolState struct {
	// The creation date of the elastic pool (ISO8601 format).
	CreationDate pulumi.StringPtrInput
	// Kind of elastic pool. This is metadata used for the Azure portal experience.
	Kind pulumi.StringPtrInput
	// The license type to apply for this elastic pool.
	LicenseType pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Maintenance configuration id assigned to the elastic pool. This configuration defines the period when the maintenance updates will will occur.
	MaintenanceConfigurationId pulumi.StringPtrInput
	// The storage limit for the database elastic pool in bytes.
	MaxSizeBytes pulumi.Float64PtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// The per database settings for the elastic pool.
	PerDatabaseSettings ElasticPoolPerDatabaseSettingsResponsePtrInput
	// The elastic pool SKU.
	//
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or the following command:
	Sku SkuResponsePtrInput
	// The state of the elastic pool.
	State pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
	// Whether or not this elastic pool is zone redundant, which means the replicas of this elastic pool will be spread across multiple availability zones.
	ZoneRedundant pulumi.BoolPtrInput
}

func (ElasticPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticPoolState)(nil)).Elem()
}

type elasticPoolArgs struct {
	// The name of the elastic pool.
	ElasticPoolName string `pulumi:"elasticPoolName"`
	// The license type to apply for this elastic pool.
	LicenseType *string `pulumi:"licenseType"`
	// Resource location.
	Location string `pulumi:"location"`
	// Maintenance configuration id assigned to the elastic pool. This configuration defines the period when the maintenance updates will will occur.
	MaintenanceConfigurationId *string `pulumi:"maintenanceConfigurationId"`
	// The storage limit for the database elastic pool in bytes.
	MaxSizeBytes *float64 `pulumi:"maxSizeBytes"`
	// The per database settings for the elastic pool.
	PerDatabaseSettings *ElasticPoolPerDatabaseSettings `pulumi:"perDatabaseSettings"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The elastic pool SKU.
	//
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or the following command:
	Sku *Sku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Whether or not this elastic pool is zone redundant, which means the replicas of this elastic pool will be spread across multiple availability zones.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

// The set of arguments for constructing a ElasticPool resource.
type ElasticPoolArgs struct {
	// The name of the elastic pool.
	ElasticPoolName pulumi.StringInput
	// The license type to apply for this elastic pool.
	LicenseType pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringInput
	// Maintenance configuration id assigned to the elastic pool. This configuration defines the period when the maintenance updates will will occur.
	MaintenanceConfigurationId pulumi.StringPtrInput
	// The storage limit for the database elastic pool in bytes.
	MaxSizeBytes pulumi.Float64PtrInput
	// The per database settings for the elastic pool.
	PerDatabaseSettings ElasticPoolPerDatabaseSettingsPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The elastic pool SKU.
	//
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or the following command:
	Sku SkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Whether or not this elastic pool is zone redundant, which means the replicas of this elastic pool will be spread across multiple availability zones.
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

func (*ElasticPool) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticPool)(nil))
}

func (i *ElasticPool) ToElasticPoolOutput() ElasticPoolOutput {
	return i.ToElasticPoolOutputWithContext(context.Background())
}

func (i *ElasticPool) ToElasticPoolOutputWithContext(ctx context.Context) ElasticPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolOutput)
}

type ElasticPoolOutput struct {
	*pulumi.OutputState
}

func (ElasticPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticPool)(nil))
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
