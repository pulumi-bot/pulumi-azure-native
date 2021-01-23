// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20201201

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListIntegrationRuntimeAuthKey(ctx *pulumi.Context, args *ListIntegrationRuntimeAuthKeyArgs, opts ...pulumi.InvokeOption) (*ListIntegrationRuntimeAuthKeyResult, error) {
	var rv ListIntegrationRuntimeAuthKeyResult
	err := ctx.Invoke("azure-nextgen:synapse/v20201201:listIntegrationRuntimeAuthKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIntegrationRuntimeAuthKeyArgs struct {
	// Integration runtime name
	IntegrationRuntimeName string `pulumi:"integrationRuntimeName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The integration runtime authentication keys.
type ListIntegrationRuntimeAuthKeyResult struct {
	// The primary integration runtime authentication key.
	AuthKey1 *string `pulumi:"authKey1"`
	// The secondary integration runtime authentication key.
	AuthKey2 *string `pulumi:"authKey2"`
}
