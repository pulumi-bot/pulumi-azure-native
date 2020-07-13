// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupAutomationAccountNodeConfiguration(ctx *pulumi.Context, args *LookupAutomationAccountNodeConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupAutomationAccountNodeConfigurationResult, error) {
	var rv LookupAutomationAccountNodeConfigurationResult
	err := ctx.Invoke("azurerm:automation:getAutomationAccountNodeConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAutomationAccountNodeConfigurationArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The Dsc node configuration name.
	Name string `pulumi:"name"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Definition of the dsc node configuration.
type LookupAutomationAccountNodeConfigurationResult struct {
	// The name of the resource
	Name string `pulumi:"name"`
	// Gets or sets the configuration properties.
	Properties DscNodeConfigurationPropertiesResponse `pulumi:"properties"`
	// The type of the resource.
	Type string `pulumi:"type"`
}