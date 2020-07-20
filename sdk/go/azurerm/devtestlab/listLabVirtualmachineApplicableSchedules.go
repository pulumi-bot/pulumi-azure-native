// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package devtestlab

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListLabVirtualmachineApplicableSchedules(ctx *pulumi.Context, args *ListLabVirtualmachineApplicableSchedulesArgs, opts ...pulumi.InvokeOption) (*ListLabVirtualmachineApplicableSchedulesResult, error) {
	var rv ListLabVirtualmachineApplicableSchedulesResult
	err := ctx.Invoke("azurerm:devtestlab:listLabVirtualmachineApplicableSchedules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListLabVirtualmachineApplicableSchedulesArgs struct {
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The name of the virtual machine.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Schedules applicable to a virtual machine. The schedules may have been defined on a VM or on lab level.
type ListLabVirtualmachineApplicableSchedulesResult struct {
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The properties of the resource.
	Properties ApplicableSchedulePropertiesResponse `pulumi:"properties"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
}