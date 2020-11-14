// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupTransparentDataEncryption(ctx *pulumi.Context, args *LookupTransparentDataEncryptionArgs, opts ...pulumi.InvokeOption) (*LookupTransparentDataEncryptionResult, error) {
	var rv LookupTransparentDataEncryptionResult
	err := ctx.Invoke("azure-nextgen:sql/latest:getTransparentDataEncryption", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTransparentDataEncryptionArgs struct {
	// The name of the database for which the transparent data encryption applies.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The name of the transparent data encryption configuration.
	TransparentDataEncryptionName string `pulumi:"transparentDataEncryptionName"`
}

// Represents a database transparent data encryption configuration.
type LookupTransparentDataEncryptionResult struct {
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The status of the database transparent data encryption.
	Status *string `pulumi:"status"`
	// Resource type.
	Type string `pulumi:"type"`
}
