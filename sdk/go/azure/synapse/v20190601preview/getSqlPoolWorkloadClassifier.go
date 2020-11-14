// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupSqlPoolWorkloadClassifier(ctx *pulumi.Context, args *LookupSqlPoolWorkloadClassifierArgs, opts ...pulumi.InvokeOption) (*LookupSqlPoolWorkloadClassifierResult, error) {
	var rv LookupSqlPoolWorkloadClassifierResult
	err := ctx.Invoke("azure-nextgen:synapse/v20190601preview:getSqlPoolWorkloadClassifier", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlPoolWorkloadClassifierArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// SQL pool name
	SqlPoolName string `pulumi:"sqlPoolName"`
	// The name of the workload classifier.
	WorkloadClassifierName string `pulumi:"workloadClassifierName"`
	// The name of the workload group.
	WorkloadGroupName string `pulumi:"workloadGroupName"`
	// The name of the workspace
	WorkspaceName string `pulumi:"workspaceName"`
}

// Workload classifier operations for a data warehouse
type LookupSqlPoolWorkloadClassifierResult struct {
	// The workload classifier context.
	Context *string `pulumi:"context"`
	// The workload classifier end time for classification.
	EndTime *string `pulumi:"endTime"`
	// The workload classifier importance.
	Importance *string `pulumi:"importance"`
	// The workload classifier label.
	Label *string `pulumi:"label"`
	// The workload classifier member name.
	MemberName string `pulumi:"memberName"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The workload classifier start time for classification.
	StartTime *string `pulumi:"startTime"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}
