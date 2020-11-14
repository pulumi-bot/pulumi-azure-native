// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupMyWorkbook(ctx *pulumi.Context, args *LookupMyWorkbookArgs, opts ...pulumi.InvokeOption) (*LookupMyWorkbookResult, error) {
	var rv LookupMyWorkbookResult
	err := ctx.Invoke("azure-nextgen:insights/latest:getMyWorkbook", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMyWorkbookArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Application Insights component resource.
	ResourceName string `pulumi:"resourceName"`
}

// An Application Insights private workbook definition.
type LookupMyWorkbookResult struct {
	// Workbook category, as defined by the user at creation time.
	Category string `pulumi:"category"`
	// The user-defined name of the private workbook.
	DisplayName string `pulumi:"displayName"`
	// Identity used for BYOS
	Identity *ManagedIdentityResponse `pulumi:"identity"`
	// The kind of workbook. Choices are user and shared.
	Kind *string `pulumi:"kind"`
	// Resource location
	Location *string `pulumi:"location"`
	// Azure resource name
	Name *string `pulumi:"name"`
	// Configuration of this particular private workbook. Configuration data is a string containing valid JSON
	SerializedData string `pulumi:"serializedData"`
	// Optional resourceId for a source resource.
	SourceId *string `pulumi:"sourceId"`
	// BYOS Storage Account URI
	StorageUri *string `pulumi:"storageUri"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Date and time in UTC of the last modification that was made to this private workbook definition.
	TimeModified string `pulumi:"timeModified"`
	// Azure resource type
	Type *string `pulumi:"type"`
	// Unique user id of the specific user that owns this private workbook.
	UserId string `pulumi:"userId"`
	// This instance's version of the data model. This can change as new features are added that can be marked private workbook.
	Version *string `pulumi:"version"`
}
