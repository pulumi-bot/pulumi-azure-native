// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190401

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Contains information about an Azure Batch account.
//
// ## Example Usage
// ### BatchAccountCreate_BYOS
//
// ```go
// package main
//
// import (
// 	batch "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/batch/v20190401"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := batch.NewBatchAccount(ctx, "batchAccount", &batch.BatchAccountArgs{
// 			AccountName: pulumi.String("sampleacct"),
// 			AutoStorage: &batch.AutoStorageBasePropertiesArgs{
// 				StorageAccountId: pulumi.String("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Storage/storageAccounts/samplestorage"),
// 			},
// 			KeyVaultReference: &batch.KeyVaultReferenceArgs{
// 				Id:  pulumi.String("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.KeyVault/vaults/sample"),
// 				Url: pulumi.String("http://sample.vault.azure.net/"),
// 			},
// 			Location:           pulumi.String("japaneast"),
// 			PoolAllocationMode: pulumi.String("UserSubscription"),
// 			ResourceGroupName:  pulumi.String("default-azurebatch-japaneast"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### BatchAccountCreate_Default
//
// ```go
// package main
//
// import (
// 	batch "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/batch/v20190401"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := batch.NewBatchAccount(ctx, "batchAccount", &batch.BatchAccountArgs{
// 			AccountName: pulumi.String("sampleacct"),
// 			AutoStorage: &batch.AutoStorageBasePropertiesArgs{
// 				StorageAccountId: pulumi.String("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Storage/storageAccounts/samplestorage"),
// 			},
// 			Location:          pulumi.String("japaneast"),
// 			ResourceGroupName: pulumi.String("default-azurebatch-japaneast"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type BatchAccount struct {
	pulumi.CustomResourceState

	// The account endpoint used to interact with the Batch service.
	AccountEndpoint              pulumi.StringOutput `pulumi:"accountEndpoint"`
	ActiveJobAndJobScheduleQuota pulumi.IntOutput    `pulumi:"activeJobAndJobScheduleQuota"`
	// Contains information about the auto-storage account associated with a Batch account.
	AutoStorage AutoStoragePropertiesResponseOutput `pulumi:"autoStorage"`
	// For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription so this value is not returned.
	DedicatedCoreQuota pulumi.IntOutput `pulumi:"dedicatedCoreQuota"`
	// A list of the dedicated core quota per Virtual Machine family for the Batch account. For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription so this value is not returned.
	DedicatedCoreQuotaPerVMFamily VirtualMachineFamilyCoreQuotaResponseArrayOutput `pulumi:"dedicatedCoreQuotaPerVMFamily"`
	// Batch is transitioning its core quota system for dedicated cores to be enforced per Virtual Machine family. During this transitional phase, the dedicated core quota per Virtual Machine family may not yet be enforced. If this flag is false, dedicated core quota is enforced via the old dedicatedCoreQuota property on the account and does not consider Virtual Machine family. If this flag is true, dedicated core quota is enforced via the dedicatedCoreQuotaPerVMFamily property on the account, and the old dedicatedCoreQuota does not apply.
	DedicatedCoreQuotaPerVMFamilyEnforced pulumi.BoolOutput `pulumi:"dedicatedCoreQuotaPerVMFamilyEnforced"`
	// Identifies the Azure key vault associated with a Batch account.
	KeyVaultReference KeyVaultReferenceResponseOutput `pulumi:"keyVaultReference"`
	// The location of the resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription so this value is not returned.
	LowPriorityCoreQuota pulumi.IntOutput `pulumi:"lowPriorityCoreQuota"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The allocation mode for creating pools in the Batch account.
	PoolAllocationMode pulumi.StringOutput `pulumi:"poolAllocationMode"`
	PoolQuota          pulumi.IntOutput    `pulumi:"poolQuota"`
	// The provisioned state of the resource
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBatchAccount registers a new resource with the given unique name, arguments, and options.
func NewBatchAccount(ctx *pulumi.Context,
	name string, args *BatchAccountArgs, opts ...pulumi.ResourceOption) (*BatchAccount, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &BatchAccountArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:batch/latest:BatchAccount"),
		},
		{
			Type: pulumi.String("azurerm:batch/v20151201:BatchAccount"),
		},
		{
			Type: pulumi.String("azurerm:batch/v20170101:BatchAccount"),
		},
		{
			Type: pulumi.String("azurerm:batch/v20170501:BatchAccount"),
		},
		{
			Type: pulumi.String("azurerm:batch/v20170901:BatchAccount"),
		},
		{
			Type: pulumi.String("azurerm:batch/v20181201:BatchAccount"),
		},
		{
			Type: pulumi.String("azurerm:batch/v20190801:BatchAccount"),
		},
		{
			Type: pulumi.String("azurerm:batch/v20200301:BatchAccount"),
		},
		{
			Type: pulumi.String("azurerm:batch/v20200501:BatchAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource BatchAccount
	err := ctx.RegisterResource("azurerm:batch/v20190401:BatchAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBatchAccount gets an existing BatchAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBatchAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BatchAccountState, opts ...pulumi.ResourceOption) (*BatchAccount, error) {
	var resource BatchAccount
	err := ctx.ReadResource("azurerm:batch/v20190401:BatchAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BatchAccount resources.
type batchAccountState struct {
	// The account endpoint used to interact with the Batch service.
	AccountEndpoint              *string `pulumi:"accountEndpoint"`
	ActiveJobAndJobScheduleQuota *int    `pulumi:"activeJobAndJobScheduleQuota"`
	// Contains information about the auto-storage account associated with a Batch account.
	AutoStorage *AutoStoragePropertiesResponse `pulumi:"autoStorage"`
	// For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription so this value is not returned.
	DedicatedCoreQuota *int `pulumi:"dedicatedCoreQuota"`
	// A list of the dedicated core quota per Virtual Machine family for the Batch account. For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription so this value is not returned.
	DedicatedCoreQuotaPerVMFamily []VirtualMachineFamilyCoreQuotaResponse `pulumi:"dedicatedCoreQuotaPerVMFamily"`
	// Batch is transitioning its core quota system for dedicated cores to be enforced per Virtual Machine family. During this transitional phase, the dedicated core quota per Virtual Machine family may not yet be enforced. If this flag is false, dedicated core quota is enforced via the old dedicatedCoreQuota property on the account and does not consider Virtual Machine family. If this flag is true, dedicated core quota is enforced via the dedicatedCoreQuotaPerVMFamily property on the account, and the old dedicatedCoreQuota does not apply.
	DedicatedCoreQuotaPerVMFamilyEnforced *bool `pulumi:"dedicatedCoreQuotaPerVMFamilyEnforced"`
	// Identifies the Azure key vault associated with a Batch account.
	KeyVaultReference *KeyVaultReferenceResponse `pulumi:"keyVaultReference"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription so this value is not returned.
	LowPriorityCoreQuota *int `pulumi:"lowPriorityCoreQuota"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The allocation mode for creating pools in the Batch account.
	PoolAllocationMode *string `pulumi:"poolAllocationMode"`
	PoolQuota          *int    `pulumi:"poolQuota"`
	// The provisioned state of the resource
	ProvisioningState *string `pulumi:"provisioningState"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type BatchAccountState struct {
	// The account endpoint used to interact with the Batch service.
	AccountEndpoint              pulumi.StringPtrInput
	ActiveJobAndJobScheduleQuota pulumi.IntPtrInput
	// Contains information about the auto-storage account associated with a Batch account.
	AutoStorage AutoStoragePropertiesResponsePtrInput
	// For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription so this value is not returned.
	DedicatedCoreQuota pulumi.IntPtrInput
	// A list of the dedicated core quota per Virtual Machine family for the Batch account. For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription so this value is not returned.
	DedicatedCoreQuotaPerVMFamily VirtualMachineFamilyCoreQuotaResponseArrayInput
	// Batch is transitioning its core quota system for dedicated cores to be enforced per Virtual Machine family. During this transitional phase, the dedicated core quota per Virtual Machine family may not yet be enforced. If this flag is false, dedicated core quota is enforced via the old dedicatedCoreQuota property on the account and does not consider Virtual Machine family. If this flag is true, dedicated core quota is enforced via the dedicatedCoreQuotaPerVMFamily property on the account, and the old dedicatedCoreQuota does not apply.
	DedicatedCoreQuotaPerVMFamilyEnforced pulumi.BoolPtrInput
	// Identifies the Azure key vault associated with a Batch account.
	KeyVaultReference KeyVaultReferenceResponsePtrInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription so this value is not returned.
	LowPriorityCoreQuota pulumi.IntPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The allocation mode for creating pools in the Batch account.
	PoolAllocationMode pulumi.StringPtrInput
	PoolQuota          pulumi.IntPtrInput
	// The provisioned state of the resource
	ProvisioningState pulumi.StringPtrInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (BatchAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*batchAccountState)(nil)).Elem()
}

type batchAccountArgs struct {
	// A name for the Batch account which must be unique within the region. Batch account names must be between 3 and 24 characters in length and must use only numbers and lowercase letters. This name is used as part of the DNS name that is used to access the Batch service in the region in which the account is created. For example: http://accountname.region.batch.azure.com/.
	AccountName string `pulumi:"accountName"`
	// The properties related to the auto-storage account.
	AutoStorage *AutoStorageBaseProperties `pulumi:"autoStorage"`
	// A reference to the Azure key vault associated with the Batch account.
	KeyVaultReference *KeyVaultReference `pulumi:"keyVaultReference"`
	// The region in which to create the account.
	Location string `pulumi:"location"`
	// The pool allocation mode also affects how clients may authenticate to the Batch Service API. If the mode is BatchService, clients may authenticate using access keys or Azure Active Directory. If the mode is UserSubscription, clients must use Azure Active Directory. The default is BatchService.
	PoolAllocationMode *string `pulumi:"poolAllocationMode"`
	// The name of the resource group that contains the Batch account.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The user-specified tags associated with the account.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a BatchAccount resource.
type BatchAccountArgs struct {
	// A name for the Batch account which must be unique within the region. Batch account names must be between 3 and 24 characters in length and must use only numbers and lowercase letters. This name is used as part of the DNS name that is used to access the Batch service in the region in which the account is created. For example: http://accountname.region.batch.azure.com/.
	AccountName pulumi.StringInput
	// The properties related to the auto-storage account.
	AutoStorage AutoStorageBasePropertiesPtrInput
	// A reference to the Azure key vault associated with the Batch account.
	KeyVaultReference KeyVaultReferencePtrInput
	// The region in which to create the account.
	Location pulumi.StringInput
	// The pool allocation mode also affects how clients may authenticate to the Batch Service API. If the mode is BatchService, clients may authenticate using access keys or Azure Active Directory. If the mode is UserSubscription, clients must use Azure Active Directory. The default is BatchService.
	PoolAllocationMode pulumi.StringPtrInput
	// The name of the resource group that contains the Batch account.
	ResourceGroupName pulumi.StringInput
	// The user-specified tags associated with the account.
	Tags pulumi.StringMapInput
}

func (BatchAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*batchAccountArgs)(nil)).Elem()
}
