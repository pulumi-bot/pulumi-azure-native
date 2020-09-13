// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Cache details.
//
// ## Example Usage
// ### ApiManagementCreateCache
//
// ```go
// package main
//
// import (
// 	apimanagement "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/apimanagement/v20191201preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apimanagement.NewCache(ctx, "cache", &apimanagement.CacheArgs{
// 			CacheId:           pulumi.String("westindia"),
// 			ConnectionString:  pulumi.String("contoso5.redis.cache.windows.net,ssl=true,password=..."),
// 			Description:       pulumi.String("Redis cache instances in West India"),
// 			ResourceGroupName: pulumi.String("rg1"),
// 			ResourceId:        pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/Redis/contoso5"),
// 			ServiceName:       pulumi.String("apimService1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Cache struct {
	pulumi.CustomResourceState

	// Runtime connection string to cache
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	// Cache description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Original uri of entity in external system cache points to
	ResourceId pulumi.StringPtrOutput `pulumi:"resourceId"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCache registers a new resource with the given unique name, arguments, and options.
func NewCache(ctx *pulumi.Context,
	name string, args *CacheArgs, opts ...pulumi.ResourceOption) (*Cache, error) {
	if args == nil || args.CacheId == nil {
		return nil, errors.New("missing required argument 'CacheId'")
	}
	if args == nil || args.ConnectionString == nil {
		return nil, errors.New("missing required argument 'ConnectionString'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil {
		args = &CacheArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:apimanagement/latest:Cache"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20180601preview:Cache"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20190101:Cache"),
		},
		{
			Type: pulumi.String("azurerm:apimanagement/v20191201:Cache"),
		},
	})
	opts = append(opts, aliases)
	var resource Cache
	err := ctx.RegisterResource("azurerm:apimanagement/v20191201preview:Cache", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:apimanagement/v20191201preview:Cache", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cache resources.
type cacheState struct {
	// Runtime connection string to cache
	ConnectionString *string `pulumi:"connectionString"`
	// Cache description
	Description *string `pulumi:"description"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Original uri of entity in external system cache points to
	ResourceId *string `pulumi:"resourceId"`
	// Resource type for API Management resource.
	Type *string `pulumi:"type"`
}

type CacheState struct {
	// Runtime connection string to cache
	ConnectionString pulumi.StringPtrInput
	// Cache description
	Description pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Original uri of entity in external system cache points to
	ResourceId pulumi.StringPtrInput
	// Resource type for API Management resource.
	Type pulumi.StringPtrInput
}

func (CacheState) ElementType() reflect.Type {
	return reflect.TypeOf((*cacheState)(nil)).Elem()
}

type cacheArgs struct {
	// Identifier of the Cache entity. Cache identifier (should be either 'default' or valid Azure region identifier).
	CacheId string `pulumi:"cacheId"`
	// Runtime connection string to cache
	ConnectionString string `pulumi:"connectionString"`
	// Cache description
	Description *string `pulumi:"description"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Original uri of entity in external system cache points to
	ResourceId *string `pulumi:"resourceId"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a Cache resource.
type CacheArgs struct {
	// Identifier of the Cache entity. Cache identifier (should be either 'default' or valid Azure region identifier).
	CacheId pulumi.StringInput
	// Runtime connection string to cache
	ConnectionString pulumi.StringInput
	// Cache description
	Description pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Original uri of entity in external system cache points to
	ResourceId pulumi.StringPtrInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (CacheArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cacheArgs)(nil)).Elem()
}
