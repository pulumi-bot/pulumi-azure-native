// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package devtestlab

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupLabVirtualmachineSchedule(ctx *pulumi.Context, args *LookupLabVirtualmachineScheduleArgs, opts ...pulumi.InvokeOption) (*LookupLabVirtualmachineScheduleResult, error) {
	var rv LookupLabVirtualmachineScheduleResult
	err := ctx.Invoke("azurerm:devtestlab:getLabVirtualmachineSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLabVirtualmachineScheduleArgs struct {
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The name of the schedule.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the virtual machine.
	VirtualMachineName string `pulumi:"virtualMachineName"`
}

// A schedule.
type LookupLabVirtualmachineScheduleResult struct {
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The properties of the resource.
	Properties SchedulePropertiesResponse `pulumi:"properties"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
}