// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150615

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The storage account.
type StorageAccount struct {
	pulumi.CustomResourceState

	// The type of the storage account.
	AccountType pulumi.StringPtrOutput `pulumi:"accountType"`
	// The creation date and time of the storage account in UTC.
	CreationTime pulumi.StringPtrOutput `pulumi:"creationTime"`
	// The custom domain the user assigned to this storage account.
	CustomDomain CustomDomainResponsePtrOutput `pulumi:"customDomain"`
	// The timestamp of the most recent instance of a failover to the secondary location. Only the most recent timestamp is retained. This element is not returned if there has never been a failover instance. Only available if the accountType is Standard_GRS or Standard_RAGRS.
	LastGeoFailoverTime pulumi.StringPtrOutput `pulumi:"lastGeoFailoverTime"`
	// Resource location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The URLs that are used to perform a retrieval of a public blob, queue, or table object. Note that Standard_ZRS and Premium_LRS accounts only return the blob endpoint.
	PrimaryEndpoints EndpointsResponsePtrOutput `pulumi:"primaryEndpoints"`
	// The location of the primary data center for the storage account.
	PrimaryLocation pulumi.StringPtrOutput `pulumi:"primaryLocation"`
	// The status of the storage account at the time the operation was called.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// The URLs that are used to perform a retrieval of a public blob, queue, or table object from the secondary location of the storage account. Only available if the SKU name is Standard_RAGRS.
	SecondaryEndpoints EndpointsResponsePtrOutput `pulumi:"secondaryEndpoints"`
	// The location of the geo-replicated secondary for the storage account. Only available if the accountType is Standard_GRS or Standard_RAGRS.
	SecondaryLocation pulumi.StringPtrOutput `pulumi:"secondaryLocation"`
	// The status indicating whether the primary location of the storage account is available or unavailable.
	StatusOfPrimary pulumi.StringPtrOutput `pulumi:"statusOfPrimary"`
	// The status indicating whether the secondary location of the storage account is available or unavailable. Only available if the SKU name is Standard_GRS or Standard_RAGRS.
	StatusOfSecondary pulumi.StringPtrOutput `pulumi:"statusOfSecondary"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStorageAccount registers a new resource with the given unique name, arguments, and options.
func NewStorageAccount(ctx *pulumi.Context,
	name string, args *StorageAccountArgs, opts ...pulumi.ResourceOption) (*StorageAccount, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.AccountType == nil {
		return nil, errors.New("missing required argument 'AccountType'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &StorageAccountArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:storage/latest:StorageAccount"),
		},
		{
			Type: pulumi.String("azurerm:storage/v20150501preview:StorageAccount"),
		},
		{
			Type: pulumi.String("azurerm:storage/v20160101:StorageAccount"),
		},
		{
			Type: pulumi.String("azurerm:storage/v20160501:StorageAccount"),
		},
		{
			Type: pulumi.String("azurerm:storage/v20161201:StorageAccount"),
		},
		{
			Type: pulumi.String("azurerm:storage/v20170601:StorageAccount"),
		},
		{
			Type: pulumi.String("azurerm:storage/v20171001:StorageAccount"),
		},
		{
			Type: pulumi.String("azurerm:storage/v20180201:StorageAccount"),
		},
		{
			Type: pulumi.String("azurerm:storage/v20180301preview:StorageAccount"),
		},
		{
			Type: pulumi.String("azurerm:storage/v20180701:StorageAccount"),
		},
		{
			Type: pulumi.String("azurerm:storage/v20181101:StorageAccount"),
		},
		{
			Type: pulumi.String("azurerm:storage/v20190401:StorageAccount"),
		},
		{
			Type: pulumi.String("azurerm:storage/v20190601:StorageAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource StorageAccount
	err := ctx.RegisterResource("azurerm:storage/v20150615:StorageAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageAccount gets an existing StorageAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageAccountState, opts ...pulumi.ResourceOption) (*StorageAccount, error) {
	var resource StorageAccount
	err := ctx.ReadResource("azurerm:storage/v20150615:StorageAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageAccount resources.
type storageAccountState struct {
	// The type of the storage account.
	AccountType *string `pulumi:"accountType"`
	// The creation date and time of the storage account in UTC.
	CreationTime *string `pulumi:"creationTime"`
	// The custom domain the user assigned to this storage account.
	CustomDomain *CustomDomainResponse `pulumi:"customDomain"`
	// The timestamp of the most recent instance of a failover to the secondary location. Only the most recent timestamp is retained. This element is not returned if there has never been a failover instance. Only available if the accountType is Standard_GRS or Standard_RAGRS.
	LastGeoFailoverTime *string `pulumi:"lastGeoFailoverTime"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// The URLs that are used to perform a retrieval of a public blob, queue, or table object. Note that Standard_ZRS and Premium_LRS accounts only return the blob endpoint.
	PrimaryEndpoints *EndpointsResponse `pulumi:"primaryEndpoints"`
	// The location of the primary data center for the storage account.
	PrimaryLocation *string `pulumi:"primaryLocation"`
	// The status of the storage account at the time the operation was called.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The URLs that are used to perform a retrieval of a public blob, queue, or table object from the secondary location of the storage account. Only available if the SKU name is Standard_RAGRS.
	SecondaryEndpoints *EndpointsResponse `pulumi:"secondaryEndpoints"`
	// The location of the geo-replicated secondary for the storage account. Only available if the accountType is Standard_GRS or Standard_RAGRS.
	SecondaryLocation *string `pulumi:"secondaryLocation"`
	// The status indicating whether the primary location of the storage account is available or unavailable.
	StatusOfPrimary *string `pulumi:"statusOfPrimary"`
	// The status indicating whether the secondary location of the storage account is available or unavailable. Only available if the SKU name is Standard_GRS or Standard_RAGRS.
	StatusOfSecondary *string `pulumi:"statusOfSecondary"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type StorageAccountState struct {
	// The type of the storage account.
	AccountType pulumi.StringPtrInput
	// The creation date and time of the storage account in UTC.
	CreationTime pulumi.StringPtrInput
	// The custom domain the user assigned to this storage account.
	CustomDomain CustomDomainResponsePtrInput
	// The timestamp of the most recent instance of a failover to the secondary location. Only the most recent timestamp is retained. This element is not returned if there has never been a failover instance. Only available if the accountType is Standard_GRS or Standard_RAGRS.
	LastGeoFailoverTime pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// The URLs that are used to perform a retrieval of a public blob, queue, or table object. Note that Standard_ZRS and Premium_LRS accounts only return the blob endpoint.
	PrimaryEndpoints EndpointsResponsePtrInput
	// The location of the primary data center for the storage account.
	PrimaryLocation pulumi.StringPtrInput
	// The status of the storage account at the time the operation was called.
	ProvisioningState pulumi.StringPtrInput
	// The URLs that are used to perform a retrieval of a public blob, queue, or table object from the secondary location of the storage account. Only available if the SKU name is Standard_RAGRS.
	SecondaryEndpoints EndpointsResponsePtrInput
	// The location of the geo-replicated secondary for the storage account. Only available if the accountType is Standard_GRS or Standard_RAGRS.
	SecondaryLocation pulumi.StringPtrInput
	// The status indicating whether the primary location of the storage account is available or unavailable.
	StatusOfPrimary pulumi.StringPtrInput
	// The status indicating whether the secondary location of the storage account is available or unavailable. Only available if the SKU name is Standard_GRS or Standard_RAGRS.
	StatusOfSecondary pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (StorageAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageAccountState)(nil)).Elem()
}

type storageAccountArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// The sku name. Required for account creation; optional for update. Note that in older versions, sku name was called accountType.
	AccountType string `pulumi:"accountType"`
	// The location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US, East US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is created, but if an identical geo region is specified on update, the request will succeed.
	Location string `pulumi:"location"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A list of key value pairs that describe the resource. These tags can be used for viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key with a length no greater than 128 characters and a value with a length no greater than 256 characters.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a StorageAccount resource.
type StorageAccountArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringInput
	// The sku name. Required for account creation; optional for update. Note that in older versions, sku name was called accountType.
	AccountType pulumi.StringInput
	// The location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US, East US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is created, but if an identical geo region is specified on update, the request will succeed.
	Location pulumi.StringInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// A list of key value pairs that describe the resource. These tags can be used for viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key with a length no greater than 128 characters and a value with a length no greater than 256 characters.
	Tags pulumi.StringMapInput
}

func (StorageAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageAccountArgs)(nil)).Elem()
}
