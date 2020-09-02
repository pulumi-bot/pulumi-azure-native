// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170501preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	var rv LookupAccountResult
	err := ctx.Invoke("azurerm:machinelearningexperimentation/v20170501preview:getAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccountArgs struct {
	// The name of the machine learning team account.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group to which the machine learning team account belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An object that represents a machine learning team account.
type LookupAccountResult struct {
	// The immutable id associated with this team account.
	AccountId string `pulumi:"accountId"`
	// The creation date of the machine learning team account in ISO8601 format.
	CreationDate string `pulumi:"creationDate"`
	// The description of this workspace.
	Description *string `pulumi:"description"`
	// The uri for this machine learning team account.
	DiscoveryUri string `pulumi:"discoveryUri"`
	// The friendly name for this workspace. This will be the workspace name in the arm id when the workspace object gets created
	FriendlyName *string `pulumi:"friendlyName"`
	// The fully qualified arm id of the user key vault.
	KeyVaultId string `pulumi:"keyVaultId"`
	// The location of the resource. This cannot be changed after the resource is created.
	Location string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The current deployment state of team account resource. The provisioningState is to indicate states for resource provisioning.
	ProvisioningState string `pulumi:"provisioningState"`
	// The no of users/seats who can access this team account. This property defines the charge on the team account.
	Seats *string `pulumi:"seats"`
	// The properties of the storage account for the machine learning team account.
	StorageAccount StorageAccountPropertiesResponse `pulumi:"storageAccount"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
	// The fully qualified arm id of the vso account to be used for this team account.
	VsoAccountId string `pulumi:"vsoAccountId"`
}
