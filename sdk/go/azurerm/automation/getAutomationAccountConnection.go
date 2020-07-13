// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupAutomationAccountConnection(ctx *pulumi.Context, args *LookupAutomationAccountConnectionArgs, opts ...pulumi.InvokeOption) (*LookupAutomationAccountConnectionResult, error) {
	var rv LookupAutomationAccountConnectionResult
	err := ctx.Invoke("azurerm:automation:getAutomationAccountConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAutomationAccountConnectionArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The name of connection.
	Name string `pulumi:"name"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Definition of the connection.
type LookupAutomationAccountConnectionResult struct {
	// The name of the resource
	Name string `pulumi:"name"`
	// Gets or sets the properties of the connection.
	Properties ConnectionPropertiesResponse `pulumi:"properties"`
	// The type of the resource.
	Type string `pulumi:"type"`
}