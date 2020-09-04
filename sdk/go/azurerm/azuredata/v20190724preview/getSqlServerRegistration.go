// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190724preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupSqlServerRegistration(ctx *pulumi.Context, args *LookupSqlServerRegistrationArgs, opts ...pulumi.InvokeOption) (*LookupSqlServerRegistrationResult, error) {
	var rv LookupSqlServerRegistrationResult
	err := ctx.Invoke("azurerm:azuredata/v20190724preview:getSqlServerRegistration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlServerRegistrationArgs struct {
	// Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the SQL Server registration.
	SqlServerRegistrationName string `pulumi:"sqlServerRegistrationName"`
}

// A SQL server registration.
type LookupSqlServerRegistrationResult struct {
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Optional Properties as JSON string
	PropertyBag *string `pulumi:"propertyBag"`
	// Resource Group Name
	ResourceGroup *string `pulumi:"resourceGroup"`
	// Subscription Id
	SubscriptionId *string `pulumi:"subscriptionId"`
	// Read only system data
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
}
