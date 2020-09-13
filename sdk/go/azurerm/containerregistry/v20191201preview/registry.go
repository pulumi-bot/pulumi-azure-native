// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An object that represents a container registry.
//
// ## Example Usage
// ### RegistryCreate
//
// ```go
// package main
//
// import (
// 	containerregistry "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/containerregistry/v20191201preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := containerregistry.NewRegistry(ctx, "registry", &containerregistry.RegistryArgs{
// 			AdminUserEnabled:  pulumi.Bool(true),
// 			Location:          pulumi.String("westus"),
// 			RegistryName:      pulumi.String("myRegistry"),
// 			ResourceGroupName: pulumi.String("myResourceGroup"),
// 			Sku: &containerregistry.SkuArgs{
// 				Name: pulumi.String("Standard"),
// 			},
// 			Tags: pulumi.StringMap{
// 				"key": pulumi.String("value"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Registry struct {
	pulumi.CustomResourceState

	// The value that indicates whether the admin user is enabled.
	AdminUserEnabled pulumi.BoolPtrOutput `pulumi:"adminUserEnabled"`
	// The creation date of the container registry in ISO8601 format.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// Enable a single data endpoint per region for serving data.
	DataEndpointEnabled pulumi.BoolPtrOutput `pulumi:"dataEndpointEnabled"`
	// List of host names that will serve data when dataEndpointEnabled is true.
	DataEndpointHostNames pulumi.StringArrayOutput `pulumi:"dataEndpointHostNames"`
	// The encryption settings of container registry.
	Encryption EncryptionPropertyResponsePtrOutput `pulumi:"encryption"`
	// The identity of the container registry.
	Identity IdentityPropertiesResponsePtrOutput `pulumi:"identity"`
	// The location of the resource. This cannot be changed after the resource is created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The URL that can be used to log into the container registry.
	LoginServer pulumi.StringOutput `pulumi:"loginServer"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The network rule set for a container registry.
	NetworkRuleSet NetworkRuleSetResponsePtrOutput `pulumi:"networkRuleSet"`
	// The policies for a container registry.
	Policies PoliciesResponsePtrOutput `pulumi:"policies"`
	// List of private endpoint connections for a container registry.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// The provisioning state of the container registry at the time the operation was called.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Whether or not public network access is allowed for the container registry.
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// The SKU of the container registry.
	Sku SkuResponseOutput `pulumi:"sku"`
	// The status of the container registry at the time the operation was called.
	Status StatusResponseOutput `pulumi:"status"`
	// The properties of the storage account for the container registry. Only applicable to Classic SKU.
	StorageAccount StorageAccountPropertiesResponsePtrOutput `pulumi:"storageAccount"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRegistry registers a new resource with the given unique name, arguments, and options.
func NewRegistry(ctx *pulumi.Context,
	name string, args *RegistryArgs, opts ...pulumi.ResourceOption) (*Registry, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.RegistryName == nil {
		return nil, errors.New("missing required argument 'RegistryName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil {
		args = &RegistryArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:containerregistry/latest:Registry"),
		},
		{
			Type: pulumi.String("azurerm:containerregistry/v20160627preview:Registry"),
		},
		{
			Type: pulumi.String("azurerm:containerregistry/v20170301:Registry"),
		},
		{
			Type: pulumi.String("azurerm:containerregistry/v20170601preview:Registry"),
		},
		{
			Type: pulumi.String("azurerm:containerregistry/v20171001:Registry"),
		},
		{
			Type: pulumi.String("azurerm:containerregistry/v20190501:Registry"),
		},
	})
	opts = append(opts, aliases)
	var resource Registry
	err := ctx.RegisterResource("azurerm:containerregistry/v20191201preview:Registry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistry gets an existing Registry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryState, opts ...pulumi.ResourceOption) (*Registry, error) {
	var resource Registry
	err := ctx.ReadResource("azurerm:containerregistry/v20191201preview:Registry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Registry resources.
type registryState struct {
	// The value that indicates whether the admin user is enabled.
	AdminUserEnabled *bool `pulumi:"adminUserEnabled"`
	// The creation date of the container registry in ISO8601 format.
	CreationDate *string `pulumi:"creationDate"`
	// Enable a single data endpoint per region for serving data.
	DataEndpointEnabled *bool `pulumi:"dataEndpointEnabled"`
	// List of host names that will serve data when dataEndpointEnabled is true.
	DataEndpointHostNames []string `pulumi:"dataEndpointHostNames"`
	// The encryption settings of container registry.
	Encryption *EncryptionPropertyResponse `pulumi:"encryption"`
	// The identity of the container registry.
	Identity *IdentityPropertiesResponse `pulumi:"identity"`
	// The location of the resource. This cannot be changed after the resource is created.
	Location *string `pulumi:"location"`
	// The URL that can be used to log into the container registry.
	LoginServer *string `pulumi:"loginServer"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The network rule set for a container registry.
	NetworkRuleSet *NetworkRuleSetResponse `pulumi:"networkRuleSet"`
	// The policies for a container registry.
	Policies *PoliciesResponse `pulumi:"policies"`
	// List of private endpoint connections for a container registry.
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// The provisioning state of the container registry at the time the operation was called.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Whether or not public network access is allowed for the container registry.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The SKU of the container registry.
	Sku *SkuResponse `pulumi:"sku"`
	// The status of the container registry at the time the operation was called.
	Status *StatusResponse `pulumi:"status"`
	// The properties of the storage account for the container registry. Only applicable to Classic SKU.
	StorageAccount *StorageAccountPropertiesResponse `pulumi:"storageAccount"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type RegistryState struct {
	// The value that indicates whether the admin user is enabled.
	AdminUserEnabled pulumi.BoolPtrInput
	// The creation date of the container registry in ISO8601 format.
	CreationDate pulumi.StringPtrInput
	// Enable a single data endpoint per region for serving data.
	DataEndpointEnabled pulumi.BoolPtrInput
	// List of host names that will serve data when dataEndpointEnabled is true.
	DataEndpointHostNames pulumi.StringArrayInput
	// The encryption settings of container registry.
	Encryption EncryptionPropertyResponsePtrInput
	// The identity of the container registry.
	Identity IdentityPropertiesResponsePtrInput
	// The location of the resource. This cannot be changed after the resource is created.
	Location pulumi.StringPtrInput
	// The URL that can be used to log into the container registry.
	LoginServer pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The network rule set for a container registry.
	NetworkRuleSet NetworkRuleSetResponsePtrInput
	// The policies for a container registry.
	Policies PoliciesResponsePtrInput
	// List of private endpoint connections for a container registry.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayInput
	// The provisioning state of the container registry at the time the operation was called.
	ProvisioningState pulumi.StringPtrInput
	// Whether or not public network access is allowed for the container registry.
	PublicNetworkAccess pulumi.StringPtrInput
	// The SKU of the container registry.
	Sku SkuResponsePtrInput
	// The status of the container registry at the time the operation was called.
	Status StatusResponsePtrInput
	// The properties of the storage account for the container registry. Only applicable to Classic SKU.
	StorageAccount StorageAccountPropertiesResponsePtrInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (RegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryState)(nil)).Elem()
}

type registryArgs struct {
	// The value that indicates whether the admin user is enabled.
	AdminUserEnabled *bool `pulumi:"adminUserEnabled"`
	// Enable a single data endpoint per region for serving data.
	DataEndpointEnabled *bool `pulumi:"dataEndpointEnabled"`
	// The encryption settings of container registry.
	Encryption *EncryptionProperty `pulumi:"encryption"`
	// The identity of the container registry.
	Identity *IdentityProperties `pulumi:"identity"`
	// The location of the resource. This cannot be changed after the resource is created.
	Location string `pulumi:"location"`
	// The network rule set for a container registry.
	NetworkRuleSet *NetworkRuleSet `pulumi:"networkRuleSet"`
	// The policies for a container registry.
	Policies *Policies `pulumi:"policies"`
	// Whether or not public network access is allowed for the container registry.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group to which the container registry belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU of the container registry.
	Sku Sku `pulumi:"sku"`
	// The properties of the storage account for the container registry. Only applicable to Classic SKU.
	StorageAccount *StorageAccountProperties `pulumi:"storageAccount"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Registry resource.
type RegistryArgs struct {
	// The value that indicates whether the admin user is enabled.
	AdminUserEnabled pulumi.BoolPtrInput
	// Enable a single data endpoint per region for serving data.
	DataEndpointEnabled pulumi.BoolPtrInput
	// The encryption settings of container registry.
	Encryption EncryptionPropertyPtrInput
	// The identity of the container registry.
	Identity IdentityPropertiesPtrInput
	// The location of the resource. This cannot be changed after the resource is created.
	Location pulumi.StringInput
	// The network rule set for a container registry.
	NetworkRuleSet NetworkRuleSetPtrInput
	// The policies for a container registry.
	Policies PoliciesPtrInput
	// Whether or not public network access is allowed for the container registry.
	PublicNetworkAccess pulumi.StringPtrInput
	// The name of the container registry.
	RegistryName pulumi.StringInput
	// The name of the resource group to which the container registry belongs.
	ResourceGroupName pulumi.StringInput
	// The SKU of the container registry.
	Sku SkuInput
	// The properties of the storage account for the container registry. Only applicable to Classic SKU.
	StorageAccount StorageAccountPropertiesPtrInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
}

func (RegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryArgs)(nil)).Elem()
}
