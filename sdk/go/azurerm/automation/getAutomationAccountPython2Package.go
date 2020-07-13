// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupAutomationAccountPython2Package(ctx *pulumi.Context, args *LookupAutomationAccountPython2PackageArgs, opts ...pulumi.InvokeOption) (*LookupAutomationAccountPython2PackageResult, error) {
	var rv LookupAutomationAccountPython2PackageResult
	err := ctx.Invoke("azurerm:automation:getAutomationAccountPython2Package", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAutomationAccountPython2PackageArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The python package name.
	Name string `pulumi:"name"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Definition of the module type.
type LookupAutomationAccountPython2PackageResult struct {
	// Gets or sets the etag of the resource.
	Etag *string `pulumi:"etag"`
	// The Azure Region where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Gets or sets the module properties.
	Properties ModulePropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
}