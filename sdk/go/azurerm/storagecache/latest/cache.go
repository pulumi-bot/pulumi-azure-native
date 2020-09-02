// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Cache instance. Follows Azure Resource Manager standards: https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/resource-api-reference.md
type Cache struct {
	pulumi.CustomResourceState

	// The size of this Cache, in GB.
	CacheSizeGB pulumi.IntPtrOutput `pulumi:"cacheSizeGB"`
	// Specifies encryption settings of the cache.
	EncryptionSettings CacheEncryptionSettingsResponsePtrOutput `pulumi:"encryptionSettings"`
	// Health of the Cache.
	Health CacheHealthResponseOutput `pulumi:"health"`
	// The identity of the cache, if configured.
	Identity CacheIdentityResponsePtrOutput `pulumi:"identity"`
	// Region name string.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Array of IP addresses that can be used by clients mounting this Cache.
	MountAddresses pulumi.StringArrayOutput `pulumi:"mountAddresses"`
	// Name of Cache.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies network settings of the cache.
	NetworkSettings CacheNetworkSettingsResponsePtrOutput `pulumi:"networkSettings"`
	// ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Specifies security settings of the cache.
	SecuritySettings CacheSecuritySettingsResponsePtrOutput `pulumi:"securitySettings"`
	// SKU for the Cache.
	Sku CacheResponseSkuPtrOutput `pulumi:"sku"`
	// Subnet used for the Cache.
	Subnet pulumi.StringPtrOutput `pulumi:"subnet"`
	// ARM tags as name/value pairs.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Type of the Cache; Microsoft.StorageCache/Cache
	Type pulumi.StringOutput `pulumi:"type"`
	// Upgrade status of the Cache.
	UpgradeStatus CacheUpgradeStatusResponsePtrOutput `pulumi:"upgradeStatus"`
}

// NewCache registers a new resource with the given unique name, arguments, and options.
func NewCache(ctx *pulumi.Context,
	name string, args *CacheArgs, opts ...pulumi.ResourceOption) (*Cache, error) {
	if args == nil || args.CacheName == nil {
		return nil, errors.New("missing required argument 'CacheName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &CacheArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:storagecache/v20190801preview:Cache"),
		},
		{
			Type: pulumi.String("azurerm:storagecache/v20191101:Cache"),
		},
		{
			Type: pulumi.String("azurerm:storagecache/v20200301:Cache"),
		},
	})
	opts = append(opts, aliases)
	var resource Cache
	err := ctx.RegisterResource("azurerm:storagecache/latest:Cache", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCache gets an existing Cache resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCache(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CacheState, opts ...pulumi.ResourceOption) (*Cache, error) {
	var resource Cache
	err := ctx.ReadResource("azurerm:storagecache/latest:Cache", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cache resources.
type cacheState struct {
	// The size of this Cache, in GB.
	CacheSizeGB *int `pulumi:"cacheSizeGB"`
	// Specifies encryption settings of the cache.
	EncryptionSettings *CacheEncryptionSettingsResponse `pulumi:"encryptionSettings"`
	// Health of the Cache.
	Health *CacheHealthResponse `pulumi:"health"`
	// The identity of the cache, if configured.
	Identity *CacheIdentityResponse `pulumi:"identity"`
	// Region name string.
	Location *string `pulumi:"location"`
	// Array of IP addresses that can be used by clients mounting this Cache.
	MountAddresses []string `pulumi:"mountAddresses"`
	// Name of Cache.
	Name *string `pulumi:"name"`
	// Specifies network settings of the cache.
	NetworkSettings *CacheNetworkSettingsResponse `pulumi:"networkSettings"`
	// ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
	ProvisioningState *string `pulumi:"provisioningState"`
	// Specifies security settings of the cache.
	SecuritySettings *CacheSecuritySettingsResponse `pulumi:"securitySettings"`
	// SKU for the Cache.
	Sku *CacheResponseSku `pulumi:"sku"`
	// Subnet used for the Cache.
	Subnet *string `pulumi:"subnet"`
	// ARM tags as name/value pairs.
	Tags map[string]interface{} `pulumi:"tags"`
	// Type of the Cache; Microsoft.StorageCache/Cache
	Type *string `pulumi:"type"`
	// Upgrade status of the Cache.
	UpgradeStatus *CacheUpgradeStatusResponse `pulumi:"upgradeStatus"`
}

type CacheState struct {
	// The size of this Cache, in GB.
	CacheSizeGB pulumi.IntPtrInput
	// Specifies encryption settings of the cache.
	EncryptionSettings CacheEncryptionSettingsResponsePtrInput
	// Health of the Cache.
	Health CacheHealthResponsePtrInput
	// The identity of the cache, if configured.
	Identity CacheIdentityResponsePtrInput
	// Region name string.
	Location pulumi.StringPtrInput
	// Array of IP addresses that can be used by clients mounting this Cache.
	MountAddresses pulumi.StringArrayInput
	// Name of Cache.
	Name pulumi.StringPtrInput
	// Specifies network settings of the cache.
	NetworkSettings CacheNetworkSettingsResponsePtrInput
	// ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
	ProvisioningState pulumi.StringPtrInput
	// Specifies security settings of the cache.
	SecuritySettings CacheSecuritySettingsResponsePtrInput
	// SKU for the Cache.
	Sku CacheResponseSkuPtrInput
	// Subnet used for the Cache.
	Subnet pulumi.StringPtrInput
	// ARM tags as name/value pairs.
	Tags pulumi.MapInput
	// Type of the Cache; Microsoft.StorageCache/Cache
	Type pulumi.StringPtrInput
	// Upgrade status of the Cache.
	UpgradeStatus CacheUpgradeStatusResponsePtrInput
}

func (CacheState) ElementType() reflect.Type {
	return reflect.TypeOf((*cacheState)(nil)).Elem()
}

type cacheArgs struct {
	// Name of Cache. Length of name must be not greater than 80 and chars must be in list of [-0-9a-zA-Z_] char class.
	CacheName string `pulumi:"cacheName"`
	// The size of this Cache, in GB.
	CacheSizeGB *int `pulumi:"cacheSizeGB"`
	// Specifies encryption settings of the cache.
	EncryptionSettings *CacheEncryptionSettings `pulumi:"encryptionSettings"`
	// The identity of the cache, if configured.
	Identity *CacheIdentity `pulumi:"identity"`
	// Region name string.
	Location *string `pulumi:"location"`
	// Specifies network settings of the cache.
	NetworkSettings *CacheNetworkSettings `pulumi:"networkSettings"`
	// ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
	ProvisioningState *string `pulumi:"provisioningState"`
	// Target resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies security settings of the cache.
	SecuritySettings *CacheSecuritySettings `pulumi:"securitySettings"`
	// SKU for the Cache.
	Sku *CacheSku `pulumi:"sku"`
	// Subnet used for the Cache.
	Subnet *string `pulumi:"subnet"`
	// ARM tags as name/value pairs.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Cache resource.
type CacheArgs struct {
	// Name of Cache. Length of name must be not greater than 80 and chars must be in list of [-0-9a-zA-Z_] char class.
	CacheName pulumi.StringInput
	// The size of this Cache, in GB.
	CacheSizeGB pulumi.IntPtrInput
	// Specifies encryption settings of the cache.
	EncryptionSettings CacheEncryptionSettingsPtrInput
	// The identity of the cache, if configured.
	Identity CacheIdentityPtrInput
	// Region name string.
	Location pulumi.StringPtrInput
	// Specifies network settings of the cache.
	NetworkSettings CacheNetworkSettingsPtrInput
	// ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
	ProvisioningState pulumi.StringPtrInput
	// Target resource group.
	ResourceGroupName pulumi.StringInput
	// Specifies security settings of the cache.
	SecuritySettings CacheSecuritySettingsPtrInput
	// SKU for the Cache.
	Sku CacheSkuPtrInput
	// Subnet used for the Cache.
	Subnet pulumi.StringPtrInput
	// ARM tags as name/value pairs.
	Tags pulumi.MapInput
}

func (CacheArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cacheArgs)(nil)).Elem()
}
