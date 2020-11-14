// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupServer(ctx *pulumi.Context, args *LookupServerArgs, opts ...pulumi.InvokeOption) (*LookupServerResult, error) {
	var rv LookupServerResult
	err := ctx.Invoke("azure-nextgen:dbformariadb/v20180601:getServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// Represents a server.
type LookupServerResult struct {
	// The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
	AdministratorLogin *string `pulumi:"administratorLogin"`
	// Earliest restore point creation time (ISO8601 format)
	EarliestRestoreDate *string `pulumi:"earliestRestoreDate"`
	// The fully qualified domain name of a server.
	FullyQualifiedDomainName *string `pulumi:"fullyQualifiedDomainName"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The master server id of a replica server.
	MasterServerId *string `pulumi:"masterServerId"`
	// The name of the resource
	Name string `pulumi:"name"`
	// List of private endpoint connections on a server
	PrivateEndpointConnections []ServerPrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// Whether or not public network access is allowed for this server. Value is optional but if passed in, must be 'Enabled' or 'Disabled'
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The maximum number of replicas that a master server can have.
	ReplicaCapacity *int `pulumi:"replicaCapacity"`
	// The replication role of the server.
	ReplicationRole *string `pulumi:"replicationRole"`
	// The SKU (pricing tier) of the server.
	Sku *SkuResponse `pulumi:"sku"`
	// Enable ssl enforcement or not when connect to server.
	SslEnforcement *string `pulumi:"sslEnforcement"`
	// Storage profile of a server.
	StorageProfile *StorageProfileResponse `pulumi:"storageProfile"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// A state of a server that is visible to user.
	UserVisibleState *string `pulumi:"userVisibleState"`
	// Server version.
	Version *string `pulumi:"version"`
}
