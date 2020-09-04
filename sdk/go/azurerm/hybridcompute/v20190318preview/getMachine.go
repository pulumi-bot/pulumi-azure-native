// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190318preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupMachine(ctx *pulumi.Context, args *LookupMachineArgs, opts ...pulumi.InvokeOption) (*LookupMachineResult, error) {
	var rv LookupMachineResult
	err := ctx.Invoke("azurerm:hybridcompute/v20190318preview:getMachine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMachineArgs struct {
	// The expand expression to apply on the operation.
	Expand *string `pulumi:"expand"`
	// The name of the hybrid machine.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Describes a hybrid machine.
type LookupMachineResult struct {
	// The hybrid machine agent full version.
	AgentVersion string `pulumi:"agentVersion"`
	// Public Key that the client provides to be used during initial resource onboarding
	ClientPublicKey *string `pulumi:"clientPublicKey"`
	// Specifies the hybrid machine display name.
	DisplayName string `pulumi:"displayName"`
	// Details about the error state.
	ErrorDetails []ErrorDetailResponse `pulumi:"errorDetails"`
	// The time of the last status change.
	LastStatusChange string `pulumi:"lastStatusChange"`
	// Resource location
	Location string `pulumi:"location"`
	// Specifies the hybrid machine FQDN.
	MachineFqdn string `pulumi:"machineFqdn"`
	// Resource name
	Name string `pulumi:"name"`
	// The Operating System running on the hybrid machine.
	OsName string `pulumi:"osName"`
	// Specifies the operating system settings for the hybrid machine.
	OsProfile *OSProfileResponse `pulumi:"osProfile"`
	// The version of Operating System running on the hybrid machine.
	OsVersion string `pulumi:"osVersion"`
	// Resource's Physical Location
	PhysicalLocation *string `pulumi:"physicalLocation"`
	// The identity's principal id.
	PrincipalId string `pulumi:"principalId"`
	// The provisioning state, which only appears in the response.
	ProvisioningState string `pulumi:"provisioningState"`
	// The status of the hybrid machine agent.
	Status string `pulumi:"status"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The identity's tenant id.
	TenantId string `pulumi:"tenantId"`
	// Resource type
	Type string `pulumi:"type"`
	// Specifies the hybrid machine unique ID.
	VmId string `pulumi:"vmId"`
}
