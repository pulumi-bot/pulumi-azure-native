// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupTable(ctx *pulumi.Context, args *LookupTableArgs, opts ...pulumi.InvokeOption) (*LookupTableResult, error) {
	var rv LookupTableResult
	err := ctx.Invoke("azure-nextgen:storage/latest:getTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTableArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A table name must be unique within a storage account and must be between 3 and 63 characters.The name must comprise of only alphanumeric characters and it cannot begin with a numeric character.
	TableName string `pulumi:"tableName"`
}

// Properties of the table, including Id, resource name, resource type.
type LookupTableResult struct {
	// The name of the resource
	Name string `pulumi:"name"`
	// Table name under the specified account
	TableName string `pulumi:"tableName"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}
