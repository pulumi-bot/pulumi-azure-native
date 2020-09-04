// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180701preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListRemediationDeploymentsAtResource(ctx *pulumi.Context, args *ListRemediationDeploymentsAtResourceArgs, opts ...pulumi.InvokeOption) (*ListRemediationDeploymentsAtResourceResult, error) {
	var rv ListRemediationDeploymentsAtResourceResult
	err := ctx.Invoke("azurerm:policyinsights/v20180701preview:listRemediationDeploymentsAtResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRemediationDeploymentsAtResourceArgs struct {
	// The name of the remediation.
	RemediationName string `pulumi:"remediationName"`
	// Resource ID.
	ResourceId string `pulumi:"resourceId"`
	// Maximum number of records to return.
	Top *int `pulumi:"top"`
}

// List of deployments for a remediation.
type ListRemediationDeploymentsAtResourceResult struct {
	// The URL to get the next set of results.
	NextLink string `pulumi:"nextLink"`
	// Array of deployments for the remediation.
	Value []RemediationDeploymentResponse `pulumi:"value"`
}
