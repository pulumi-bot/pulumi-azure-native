// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package documentdb

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListDatabaseAccountNotebookWorkspaceConnectionInfo(ctx *pulumi.Context, args *ListDatabaseAccountNotebookWorkspaceConnectionInfoArgs, opts ...pulumi.InvokeOption) (*ListDatabaseAccountNotebookWorkspaceConnectionInfoResult, error) {
	var rv ListDatabaseAccountNotebookWorkspaceConnectionInfoResult
	err := ctx.Invoke("azurerm:documentdb:listDatabaseAccountNotebookWorkspaceConnectionInfo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDatabaseAccountNotebookWorkspaceConnectionInfoArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// The name of the notebook workspace resource.
	NotebookWorkspaceName string `pulumi:"notebookWorkspaceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The connection info for the given notebook workspace
type ListDatabaseAccountNotebookWorkspaceConnectionInfoResult struct {
	// Specifies auth token used for connecting to Notebook server (uses token-based auth).
	AuthToken string `pulumi:"authToken"`
	// Specifies the endpoint of Notebook server.
	NotebookServerEndpoint string `pulumi:"notebookServerEndpoint"`
}