// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20171201

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupConfiguration(ctx *pulumi.Context, args *LookupConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationResult, error) {
	var rv LookupConfigurationResult
	err := ctx.Invoke("azure-nextgen:dbforpostgresql/v20171201:getConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationArgs struct {
	// The name of the server configuration.
	ConfigurationName string `pulumi:"configurationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// Represents a Configuration.
type LookupConfigurationResult struct {
	// Allowed values of the configuration.
	AllowedValues string `pulumi:"allowedValues"`
	// Data type of the configuration.
	DataType string `pulumi:"dataType"`
	// Default value of the configuration.
	DefaultValue string `pulumi:"defaultValue"`
	// Description of the configuration.
	Description string `pulumi:"description"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Source of the configuration.
	Source *string `pulumi:"source"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Value of the configuration.
	Value *string `pulumi:"value"`
}
