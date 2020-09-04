// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150501preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupFailoverGroup(ctx *pulumi.Context, args *LookupFailoverGroupArgs, opts ...pulumi.InvokeOption) (*LookupFailoverGroupResult, error) {
	var rv LookupFailoverGroupResult
	err := ctx.Invoke("azurerm:sql/v20150501preview:getFailoverGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFailoverGroupArgs struct {
	// The name of the failover group.
	FailoverGroupName string `pulumi:"failoverGroupName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server containing the failover group.
	ServerName string `pulumi:"serverName"`
}

// A failover group.
type LookupFailoverGroupResult struct {
	// List of databases in the failover group.
	Databases []string `pulumi:"databases"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// List of partner server information for the failover group.
	PartnerServers []PartnerInfoResponse `pulumi:"partnerServers"`
	// Read-only endpoint of the failover group instance.
	ReadOnlyEndpoint *FailoverGroupReadOnlyEndpointResponse `pulumi:"readOnlyEndpoint"`
	// Read-write endpoint of the failover group instance.
	ReadWriteEndpoint FailoverGroupReadWriteEndpointResponse `pulumi:"readWriteEndpoint"`
	// Local replication role of the failover group instance.
	ReplicationRole string `pulumi:"replicationRole"`
	// Replication state of the failover group instance.
	ReplicationState string `pulumi:"replicationState"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}
