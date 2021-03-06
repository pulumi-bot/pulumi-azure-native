// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Maintenance configuration record type
// Latest API Version: 2020-04-01.
//
// Deprecated: The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:maintenance:getMaintenanceConfiguration'.
func LookupMaintenanceConfiguration(ctx *pulumi.Context, args *LookupMaintenanceConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupMaintenanceConfigurationResult, error) {
	var rv LookupMaintenanceConfigurationResult
	err := ctx.Invoke("azure-native:maintenance/latest:getMaintenanceConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMaintenanceConfigurationArgs struct {
	// Resource Group Name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource Identifier
	ResourceName string `pulumi:"resourceName"`
}

// Maintenance configuration record type
type LookupMaintenanceConfigurationResult struct {
	// Gets or sets extensionProperties of the maintenanceConfiguration. This is for future use only and would be a set of key value pairs for additional information e.g. whether to follow SDP etc.
	ExtensionProperties map[string]string `pulumi:"extensionProperties"`
	// Fully qualified identifier of the resource
	Id string `pulumi:"id"`
	// Gets or sets location of the resource
	Location *string `pulumi:"location"`
	// Gets or sets maintenanceScope of the configuration. It represent the impact area of the maintenance
	MaintenanceScope *string `pulumi:"maintenanceScope"`
	// Name of the resource
	Name string `pulumi:"name"`
	// Gets or sets namespace of the resource e.g. Microsoft.Maintenance or Microsoft.Sql
	Namespace *string `pulumi:"namespace"`
	// Gets or sets tags of the resource
	Tags map[string]string `pulumi:"tags"`
	// Type of the resource
	Type string `pulumi:"type"`
}
