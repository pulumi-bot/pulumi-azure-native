// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200515preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupLinkedWorkspace(ctx *pulumi.Context, args *LookupLinkedWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupLinkedWorkspaceResult, error) {
	var rv LookupLinkedWorkspaceResult
	err := ctx.Invoke("azurerm:machinelearningservices/v20200515preview:getLinkedWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLinkedWorkspaceArgs struct {
	// Friendly name of the linked workspace
	LinkName string `pulumi:"linkName"`
	// Name of the resource group in which workspace is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Linked workspace.
type LookupLinkedWorkspaceResult struct {
	// Friendly name of the linked workspace.
	Name string `pulumi:"name"`
	// LinkedWorkspace specific properties.
	Properties LinkedWorkspacePropsResponse `pulumi:"properties"`
	// Resource type of linked workspace.
	Type string `pulumi:"type"`
}
