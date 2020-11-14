// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191210preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupHostPool(ctx *pulumi.Context, args *LookupHostPoolArgs, opts ...pulumi.InvokeOption) (*LookupHostPoolResult, error) {
	var rv LookupHostPoolResult
	err := ctx.Invoke("azure-nextgen:desktopvirtualization/v20191210preview:getHostPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHostPoolArgs struct {
	// The name of the host pool within the specified resource group
	HostPoolName string `pulumi:"hostPoolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Represents a HostPool definition.
type LookupHostPoolResult struct {
	// List of applicationGroup links.
	ApplicationGroupReferences []string `pulumi:"applicationGroupReferences"`
	// Custom rdp property of HostPool.
	CustomRdpProperty *string `pulumi:"customRdpProperty"`
	// Description of HostPool.
	Description *string `pulumi:"description"`
	// Friendly name of HostPool.
	FriendlyName *string `pulumi:"friendlyName"`
	// HostPool type for desktop.
	HostPoolType string `pulumi:"hostPoolType"`
	// The type of the load balancer.
	LoadBalancerType string `pulumi:"loadBalancerType"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The max session limit of HostPool.
	MaxSessionLimit *int `pulumi:"maxSessionLimit"`
	// The name of the resource
	Name string `pulumi:"name"`
	// PersonalDesktopAssignment type for HostPool.
	PersonalDesktopAssignmentType *string `pulumi:"personalDesktopAssignmentType"`
	// The type of preferred application group type, default to Desktop Application Group
	PreferredAppGroupType string `pulumi:"preferredAppGroupType"`
	// The registration info of HostPool.
	RegistrationInfo *RegistrationInfoResponse `pulumi:"registrationInfo"`
	// The ring number of HostPool.
	Ring *int `pulumi:"ring"`
	// Path to keyvault containing ssoContext secret.
	SsoContext *string `pulumi:"ssoContext"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Is validation environment.
	ValidationEnvironment *bool `pulumi:"validationEnvironment"`
	// VM template for sessionhosts configuration within hostpool.
	VmTemplate *string `pulumi:"vmTemplate"`
}
