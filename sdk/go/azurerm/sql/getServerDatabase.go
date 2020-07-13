// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupServerDatabase(ctx *pulumi.Context, args *LookupServerDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupServerDatabaseResult, error) {
	var rv LookupServerDatabaseResult
	err := ctx.Invoke("azurerm:sql:getServerDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerDatabaseArgs struct {
	// The name of the database to be retrieved.
	Name string `pulumi:"name"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// Represents a database.
type LookupServerDatabaseResult struct {
	// Kind of database.  This is metadata used for the Azure portal experience.
	Kind string `pulumi:"kind"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The properties representing the resource.
	Properties DatabasePropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}