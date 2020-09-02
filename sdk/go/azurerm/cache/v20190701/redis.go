// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190701

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A single Redis item in List or Get Operation.
type Redis struct {
	pulumi.CustomResourceState

	// The keys of the Redis cache - not set if this object is not the response to Create or Update redis cache
	AccessKeys RedisAccessKeysResponseOutput `pulumi:"accessKeys"`
	// Specifies whether the non-ssl Redis server port (6379) is enabled.
	EnableNonSslPort pulumi.BoolPtrOutput `pulumi:"enableNonSslPort"`
	// Redis host name.
	HostName pulumi.StringOutput `pulumi:"hostName"`
	// List of the Redis instances associated with the cache
	Instances RedisInstanceDetailsResponseArrayOutput `pulumi:"instances"`
	// List of the linked servers associated with the cache
	LinkedServers RedisLinkedServerResponseArrayOutput `pulumi:"linkedServers"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Optional: requires clients to use a specified TLS version (or higher) to connect (e,g, '1.0', '1.1', '1.2')
	MinimumTlsVersion pulumi.StringPtrOutput `pulumi:"minimumTlsVersion"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Redis non-SSL port.
	Port pulumi.IntOutput `pulumi:"port"`
	// Redis instance provisioning status.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// All Redis Settings. Few possible keys: rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value etc.
	RedisConfiguration pulumi.StringMapOutput `pulumi:"redisConfiguration"`
	// Redis version.
	RedisVersion pulumi.StringOutput `pulumi:"redisVersion"`
	// The number of replicas to be created per master.
	ReplicasPerMaster pulumi.IntPtrOutput `pulumi:"replicasPerMaster"`
	// The number of shards to be created on a Premium Cluster Cache.
	ShardCount pulumi.IntPtrOutput `pulumi:"shardCount"`
	// The SKU of the Redis cache to deploy.
	Sku SkuResponseOutput `pulumi:"sku"`
	// Redis SSL port.
	SslPort pulumi.IntOutput `pulumi:"sslPort"`
	// Static IP address. Required when deploying a Redis cache inside an existing Azure Virtual Network.
	StaticIP pulumi.StringPtrOutput `pulumi:"staticIP"`
	// The full resource ID of a subnet in a virtual network to deploy the Redis cache in. Example format: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Microsoft.{Network|ClassicNetwork}/VirtualNetworks/vnet1/subnets/subnet1
	SubnetId pulumi.StringPtrOutput `pulumi:"subnetId"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A dictionary of tenant settings
	TenantSettings pulumi.StringMapOutput `pulumi:"tenantSettings"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// A list of availability zones denoting where the resource needs to come from.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewRedis registers a new resource with the given unique name, arguments, and options.
func NewRedis(ctx *pulumi.Context,
	name string, args *RedisArgs, opts ...pulumi.ResourceOption) (*Redis, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil {
		args = &RedisArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:cache/latest:Redis"),
		},
		{
			Type: pulumi.String("azurerm:cache/v20150801:Redis"),
		},
		{
			Type: pulumi.String("azurerm:cache/v20160401:Redis"),
		},
		{
			Type: pulumi.String("azurerm:cache/v20170201:Redis"),
		},
		{
			Type: pulumi.String("azurerm:cache/v20171001:Redis"),
		},
		{
			Type: pulumi.String("azurerm:cache/v20180301:Redis"),
		},
	})
	opts = append(opts, aliases)
	var resource Redis
	err := ctx.RegisterResource("azurerm:cache/v20190701:Redis", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRedis gets an existing Redis resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRedis(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RedisState, opts ...pulumi.ResourceOption) (*Redis, error) {
	var resource Redis
	err := ctx.ReadResource("azurerm:cache/v20190701:Redis", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Redis resources.
type redisState struct {
	// The keys of the Redis cache - not set if this object is not the response to Create or Update redis cache
	AccessKeys *RedisAccessKeysResponse `pulumi:"accessKeys"`
	// Specifies whether the non-ssl Redis server port (6379) is enabled.
	EnableNonSslPort *bool `pulumi:"enableNonSslPort"`
	// Redis host name.
	HostName *string `pulumi:"hostName"`
	// List of the Redis instances associated with the cache
	Instances []RedisInstanceDetailsResponse `pulumi:"instances"`
	// List of the linked servers associated with the cache
	LinkedServers []RedisLinkedServerResponse `pulumi:"linkedServers"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Optional: requires clients to use a specified TLS version (or higher) to connect (e,g, '1.0', '1.1', '1.2')
	MinimumTlsVersion *string `pulumi:"minimumTlsVersion"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Redis non-SSL port.
	Port *int `pulumi:"port"`
	// Redis instance provisioning status.
	ProvisioningState *string `pulumi:"provisioningState"`
	// All Redis Settings. Few possible keys: rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value etc.
	RedisConfiguration map[string]string `pulumi:"redisConfiguration"`
	// Redis version.
	RedisVersion *string `pulumi:"redisVersion"`
	// The number of replicas to be created per master.
	ReplicasPerMaster *int `pulumi:"replicasPerMaster"`
	// The number of shards to be created on a Premium Cluster Cache.
	ShardCount *int `pulumi:"shardCount"`
	// The SKU of the Redis cache to deploy.
	Sku *SkuResponse `pulumi:"sku"`
	// Redis SSL port.
	SslPort *int `pulumi:"sslPort"`
	// Static IP address. Required when deploying a Redis cache inside an existing Azure Virtual Network.
	StaticIP *string `pulumi:"staticIP"`
	// The full resource ID of a subnet in a virtual network to deploy the Redis cache in. Example format: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Microsoft.{Network|ClassicNetwork}/VirtualNetworks/vnet1/subnets/subnet1
	SubnetId *string `pulumi:"subnetId"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// A dictionary of tenant settings
	TenantSettings map[string]string `pulumi:"tenantSettings"`
	// Resource type.
	Type *string `pulumi:"type"`
	// A list of availability zones denoting where the resource needs to come from.
	Zones []string `pulumi:"zones"`
}

type RedisState struct {
	// The keys of the Redis cache - not set if this object is not the response to Create or Update redis cache
	AccessKeys RedisAccessKeysResponsePtrInput
	// Specifies whether the non-ssl Redis server port (6379) is enabled.
	EnableNonSslPort pulumi.BoolPtrInput
	// Redis host name.
	HostName pulumi.StringPtrInput
	// List of the Redis instances associated with the cache
	Instances RedisInstanceDetailsResponseArrayInput
	// List of the linked servers associated with the cache
	LinkedServers RedisLinkedServerResponseArrayInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Optional: requires clients to use a specified TLS version (or higher) to connect (e,g, '1.0', '1.1', '1.2')
	MinimumTlsVersion pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Redis non-SSL port.
	Port pulumi.IntPtrInput
	// Redis instance provisioning status.
	ProvisioningState pulumi.StringPtrInput
	// All Redis Settings. Few possible keys: rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value etc.
	RedisConfiguration pulumi.StringMapInput
	// Redis version.
	RedisVersion pulumi.StringPtrInput
	// The number of replicas to be created per master.
	ReplicasPerMaster pulumi.IntPtrInput
	// The number of shards to be created on a Premium Cluster Cache.
	ShardCount pulumi.IntPtrInput
	// The SKU of the Redis cache to deploy.
	Sku SkuResponsePtrInput
	// Redis SSL port.
	SslPort pulumi.IntPtrInput
	// Static IP address. Required when deploying a Redis cache inside an existing Azure Virtual Network.
	StaticIP pulumi.StringPtrInput
	// The full resource ID of a subnet in a virtual network to deploy the Redis cache in. Example format: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Microsoft.{Network|ClassicNetwork}/VirtualNetworks/vnet1/subnets/subnet1
	SubnetId pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// A dictionary of tenant settings
	TenantSettings pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
	// A list of availability zones denoting where the resource needs to come from.
	Zones pulumi.StringArrayInput
}

func (RedisState) ElementType() reflect.Type {
	return reflect.TypeOf((*redisState)(nil)).Elem()
}

type redisArgs struct {
	// Specifies whether the non-ssl Redis server port (6379) is enabled.
	EnableNonSslPort *bool `pulumi:"enableNonSslPort"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Optional: requires clients to use a specified TLS version (or higher) to connect (e,g, '1.0', '1.1', '1.2')
	MinimumTlsVersion *string `pulumi:"minimumTlsVersion"`
	// The name of the Redis cache.
	Name string `pulumi:"name"`
	// All Redis Settings. Few possible keys: rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value etc.
	RedisConfiguration map[string]string `pulumi:"redisConfiguration"`
	// The number of replicas to be created per master.
	ReplicasPerMaster *int `pulumi:"replicasPerMaster"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The number of shards to be created on a Premium Cluster Cache.
	ShardCount *int `pulumi:"shardCount"`
	// The SKU of the Redis cache to deploy.
	Sku Sku `pulumi:"sku"`
	// Static IP address. Required when deploying a Redis cache inside an existing Azure Virtual Network.
	StaticIP *string `pulumi:"staticIP"`
	// The full resource ID of a subnet in a virtual network to deploy the Redis cache in. Example format: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Microsoft.{Network|ClassicNetwork}/VirtualNetworks/vnet1/subnets/subnet1
	SubnetId *string `pulumi:"subnetId"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// A dictionary of tenant settings
	TenantSettings map[string]string `pulumi:"tenantSettings"`
	// A list of availability zones denoting where the resource needs to come from.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a Redis resource.
type RedisArgs struct {
	// Specifies whether the non-ssl Redis server port (6379) is enabled.
	EnableNonSslPort pulumi.BoolPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringInput
	// Optional: requires clients to use a specified TLS version (or higher) to connect (e,g, '1.0', '1.1', '1.2')
	MinimumTlsVersion pulumi.StringPtrInput
	// The name of the Redis cache.
	Name pulumi.StringInput
	// All Redis Settings. Few possible keys: rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value etc.
	RedisConfiguration pulumi.StringMapInput
	// The number of replicas to be created per master.
	ReplicasPerMaster pulumi.IntPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The number of shards to be created on a Premium Cluster Cache.
	ShardCount pulumi.IntPtrInput
	// The SKU of the Redis cache to deploy.
	Sku SkuInput
	// Static IP address. Required when deploying a Redis cache inside an existing Azure Virtual Network.
	StaticIP pulumi.StringPtrInput
	// The full resource ID of a subnet in a virtual network to deploy the Redis cache in. Example format: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Microsoft.{Network|ClassicNetwork}/VirtualNetworks/vnet1/subnets/subnet1
	SubnetId pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// A dictionary of tenant settings
	TenantSettings pulumi.StringMapInput
	// A list of availability zones denoting where the resource needs to come from.
	Zones pulumi.StringArrayInput
}

func (RedisArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*redisArgs)(nil)).Elem()
}
