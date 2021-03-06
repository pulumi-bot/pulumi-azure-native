// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Defines the security Configuration
// API Version: 2021-02-01-preview.
func LookupSecurityConfiguration(ctx *pulumi.Context, args *LookupSecurityConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupSecurityConfigurationResult, error) {
	var rv LookupSecurityConfigurationResult
	err := ctx.Invoke("azure-native:network:getSecurityConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecurityConfigurationArgs struct {
	// The name of the network manager security Configuration.
	ConfigurationName string `pulumi:"configurationName"`
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Defines the security Configuration
type LookupSecurityConfigurationResult struct {
	// Groups for configuration
	AppliesToGroups []NetworkManagerSecurityGroupItemResponse `pulumi:"appliesToGroups"`
	// Flag if need to delete existing network security groups.
	DeleteExistingNSGs *bool `pulumi:"deleteExistingNSGs"`
	// A description of the security Configuration.
	Description *string `pulumi:"description"`
	// A display name of the security Configuration.
	DisplayName *string `pulumi:"displayName"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// The provisioning state of the scope assignment resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Security Type.
	SecurityType *string `pulumi:"securityType"`
	// The system metadata related to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type.
	Type string `pulumi:"type"`
}
