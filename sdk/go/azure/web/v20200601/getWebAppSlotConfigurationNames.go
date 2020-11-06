// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupWebAppSlotConfigurationNames(ctx *pulumi.Context, args *LookupWebAppSlotConfigurationNamesArgs, opts ...pulumi.InvokeOption) (*LookupWebAppSlotConfigurationNamesResult, error) {
	var rv LookupWebAppSlotConfigurationNamesResult
	err := ctx.Invoke("azure-nextgen:web/v20200601:getWebAppSlotConfigurationNames", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppSlotConfigurationNamesArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Slot Config names azure resource.
type LookupWebAppSlotConfigurationNamesResult struct {
	// List of application settings names.
	AppSettingNames []string `pulumi:"appSettingNames"`
	// List of external Azure storage account identifiers.
	AzureStorageConfigNames []string `pulumi:"azureStorageConfigNames"`
	// List of connection string names.
	ConnectionStringNames []string `pulumi:"connectionStringNames"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Resource type.
	Type string `pulumi:"type"`
}
