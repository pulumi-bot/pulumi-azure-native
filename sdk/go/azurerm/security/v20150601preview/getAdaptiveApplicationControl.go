// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150601preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupAdaptiveApplicationControl(ctx *pulumi.Context, args *LookupAdaptiveApplicationControlArgs, opts ...pulumi.InvokeOption) (*LookupAdaptiveApplicationControlResult, error) {
	var rv LookupAdaptiveApplicationControlResult
	err := ctx.Invoke("azurerm:security/v20150601preview:getAdaptiveApplicationControl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAdaptiveApplicationControlArgs struct {
	// The location where ASC stores the data of the subscription. can be retrieved from Get locations
	AscLocation string `pulumi:"ascLocation"`
	// Name of an application control VM/server group
	GroupName string `pulumi:"groupName"`
}

type LookupAdaptiveApplicationControlResult struct {
	// The configuration status of the VM/server group or machine or rule on the machine
	ConfigurationStatus *string `pulumi:"configurationStatus"`
	// The application control policy enforcement/protection mode of the VM/server group
	EnforcementMode *string                               `pulumi:"enforcementMode"`
	Issues          []AppWhitelistingIssueSummaryResponse `pulumi:"issues"`
	// Location where the resource is stored
	Location string `pulumi:"location"`
	// Resource name
	Name                string                       `pulumi:"name"`
	PathRecommendations []PathRecommendationResponse `pulumi:"pathRecommendations"`
	// The protection mode of the collection/file types. Exe/Msi/Script are used for Windows, Executable is used for Linux.
	ProtectionMode *ProtectionModeResponse `pulumi:"protectionMode"`
	// The recommendation status of the VM/server group or VM/server
	RecommendationStatus *string `pulumi:"recommendationStatus"`
	// The source type of the VM/server group
	SourceSystem *string `pulumi:"sourceSystem"`
	// Resource type
	Type              string                     `pulumi:"type"`
	VmRecommendations []VmRecommendationResponse `pulumi:"vmRecommendations"`
}
