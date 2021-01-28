// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupRegisteredServer(ctx *pulumi.Context, args *LookupRegisteredServerArgs, opts ...pulumi.InvokeOption) (*LookupRegisteredServerResult, error) {
	var rv LookupRegisteredServerResult
	err := ctx.Invoke("azure-nextgen:storagesync/latest:getRegisteredServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegisteredServerArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// GUID identifying the on-premises server.
	ServerId string `pulumi:"serverId"`
	// Name of Storage Sync Service resource.
	StorageSyncServiceName string `pulumi:"storageSyncServiceName"`
}

// Registered Server resource.
type LookupRegisteredServerResult struct {
	// Registered Server Agent Version
	AgentVersion *string `pulumi:"agentVersion"`
	// Registered Server clusterId
	ClusterId *string `pulumi:"clusterId"`
	// Registered Server clusterName
	ClusterName *string `pulumi:"clusterName"`
	// Resource discoveryEndpointUri
	DiscoveryEndpointUri *string `pulumi:"discoveryEndpointUri"`
	// Friendly Name
	FriendlyName *string `pulumi:"friendlyName"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Registered Server last heart beat
	LastHeartBeat *string `pulumi:"lastHeartBeat"`
	// Resource Last Operation Name
	LastOperationName *string `pulumi:"lastOperationName"`
	// Registered Server lastWorkflowId
	LastWorkflowId *string `pulumi:"lastWorkflowId"`
	// Management Endpoint Uri
	ManagementEndpointUri *string `pulumi:"managementEndpointUri"`
	// Monitoring Configuration
	MonitoringConfiguration *string `pulumi:"monitoringConfiguration"`
	// Telemetry Endpoint Uri
	MonitoringEndpointUri *string `pulumi:"monitoringEndpointUri"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Registered Server Provisioning State
	ProvisioningState *string `pulumi:"provisioningState"`
	// Resource Location
	ResourceLocation *string `pulumi:"resourceLocation"`
	// Registered Server Certificate
	ServerCertificate *string `pulumi:"serverCertificate"`
	// Registered Server serverId
	ServerId *string `pulumi:"serverId"`
	// Registered Server Management Error Code
	ServerManagementErrorCode *int `pulumi:"serverManagementErrorCode"`
	// Registered Server OS Version
	ServerOSVersion *string `pulumi:"serverOSVersion"`
	// Registered Server serverRole
	ServerRole *string `pulumi:"serverRole"`
	// Service Location
	ServiceLocation *string `pulumi:"serviceLocation"`
	// Registered Server storageSyncServiceUid
	StorageSyncServiceUid *string `pulumi:"storageSyncServiceUid"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}
