// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Latest API Version: 2021-01-01.
//
// Deprecated: The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:machinelearningservices:listNotebookKeys'.
func ListNotebookKeys(ctx *pulumi.Context, args *ListNotebookKeysArgs, opts ...pulumi.InvokeOption) (*ListNotebookKeysResult, error) {
	var rv ListNotebookKeysResult
	err := ctx.Invoke("azure-native:machinelearningservices/latest:listNotebookKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListNotebookKeysArgs struct {
	// Name of the resource group in which workspace is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

type ListNotebookKeysResult struct {
	PrimaryAccessKey   string `pulumi:"primaryAccessKey"`
	SecondaryAccessKey string `pulumi:"secondaryAccessKey"`
}
