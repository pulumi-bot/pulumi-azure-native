// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Describes the RedisEnterprise cluster
type RedisEnterprise struct {
	pulumi.CustomResourceState

	// DNS name of the cluster endpoint
	HostName pulumi.StringOutput `pulumi:"hostName"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The minimum TLS version for the cluster to support, e.g. '1.2'
	MinimumTlsVersion pulumi.StringPtrOutput `pulumi:"minimumTlsVersion"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// List of private endpoint connections associated with the specified RedisEnterprise cluster
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Current provisioning status of the cluster
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Version of redis the cluster supports, e.g. '6'
	RedisVersion pulumi.StringOutput `pulumi:"redisVersion"`
	// Current resource status of the cluster
	ResourceState pulumi.StringOutput `pulumi:"resourceState"`
	// The SKU to create, which affects price, performance, and features.
	Sku SkuResponseOutput `pulumi:"sku"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The zones where this cluster will be deployed.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewRedisEnterprise registers a new resource with the given unique name, arguments, and options.
func NewRedisEnterprise(ctx *pulumi.Context,
	name string, args *RedisEnterpriseArgs, opts ...pulumi.ResourceOption) (*RedisEnterprise, error) {
	if args == nil || args.ClusterName == nil {
		return nil, errors.New("missing required argument 'ClusterName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil {
		args = &RedisEnterpriseArgs{}
	}
	var resource RedisEnterprise
	err := ctx.RegisterResource("azure-nextgen:cache/v20201001preview:RedisEnterprise", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRedisEnterprise gets an existing RedisEnterprise resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRedisEnterprise(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RedisEnterpriseState, opts ...pulumi.ResourceOption) (*RedisEnterprise, error) {
	var resource RedisEnterprise
	err := ctx.ReadResource("azure-nextgen:cache/v20201001preview:RedisEnterprise", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RedisEnterprise resources.
type redisEnterpriseState struct {
	// DNS name of the cluster endpoint
	HostName *string `pulumi:"hostName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The minimum TLS version for the cluster to support, e.g. '1.2'
	MinimumTlsVersion *string `pulumi:"minimumTlsVersion"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// List of private endpoint connections associated with the specified RedisEnterprise cluster
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// Current provisioning status of the cluster
	ProvisioningState *string `pulumi:"provisioningState"`
	// Version of redis the cluster supports, e.g. '6'
	RedisVersion *string `pulumi:"redisVersion"`
	// Current resource status of the cluster
	ResourceState *string `pulumi:"resourceState"`
	// The SKU to create, which affects price, performance, and features.
	Sku *SkuResponse `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `pulumi:"type"`
	// The zones where this cluster will be deployed.
	Zones []string `pulumi:"zones"`
}

type RedisEnterpriseState struct {
	// DNS name of the cluster endpoint
	HostName pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The minimum TLS version for the cluster to support, e.g. '1.2'
	MinimumTlsVersion pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// List of private endpoint connections associated with the specified RedisEnterprise cluster
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayInput
	// Current provisioning status of the cluster
	ProvisioningState pulumi.StringPtrInput
	// Version of redis the cluster supports, e.g. '6'
	RedisVersion pulumi.StringPtrInput
	// Current resource status of the cluster
	ResourceState pulumi.StringPtrInput
	// The SKU to create, which affects price, performance, and features.
	Sku SkuResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringPtrInput
	// The zones where this cluster will be deployed.
	Zones pulumi.StringArrayInput
}

func (RedisEnterpriseState) ElementType() reflect.Type {
	return reflect.TypeOf((*redisEnterpriseState)(nil)).Elem()
}

type redisEnterpriseArgs struct {
	// The name of the RedisEnterprise cluster.
	ClusterName string `pulumi:"clusterName"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The minimum TLS version for the cluster to support, e.g. '1.2'
	MinimumTlsVersion *string `pulumi:"minimumTlsVersion"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU to create, which affects price, performance, and features.
	Sku Sku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The zones where this cluster will be deployed.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a RedisEnterprise resource.
type RedisEnterpriseArgs struct {
	// The name of the RedisEnterprise cluster.
	ClusterName pulumi.StringInput
	// The geo-location where the resource lives
	Location pulumi.StringInput
	// The minimum TLS version for the cluster to support, e.g. '1.2'
	MinimumTlsVersion pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The SKU to create, which affects price, performance, and features.
	Sku SkuInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The zones where this cluster will be deployed.
	Zones pulumi.StringArrayInput
}

func (RedisEnterpriseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*redisEnterpriseArgs)(nil)).Elem()
}
