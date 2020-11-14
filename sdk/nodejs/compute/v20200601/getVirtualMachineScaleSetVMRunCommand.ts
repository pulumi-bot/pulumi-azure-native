// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getVirtualMachineScaleSetVMRunCommand(args: GetVirtualMachineScaleSetVMRunCommandArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualMachineScaleSetVMRunCommandResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:compute/v20200601:getVirtualMachineScaleSetVMRunCommand", {
        "expand": args.expand,
        "instanceId": args.instanceId,
        "resourceGroupName": args.resourceGroupName,
        "runCommandName": args.runCommandName,
        "vmScaleSetName": args.vmScaleSetName,
    }, opts);
}

export interface GetVirtualMachineScaleSetVMRunCommandArgs {
    /**
     * The expand expression to apply on the operation.
     */
    readonly expand?: string;
    /**
     * The instance ID of the virtual machine.
     */
    readonly instanceId: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the virtual machine run command.
     */
    readonly runCommandName: string;
    /**
     * The name of the VM scale set.
     */
    readonly vmScaleSetName: string;
}

/**
 * Describes a Virtual Machine run command.
 */
export interface GetVirtualMachineScaleSetVMRunCommandResult {
    /**
     * Optional. If set to true, provisioning will complete as soon as the script starts and will not wait for script to complete.
     */
    readonly asyncExecution?: boolean;
    /**
     * Specifies the Azure storage blob where script error stream will be uploaded.
     */
    readonly errorBlobUri?: string;
    /**
     * The virtual machine run command instance view.
     */
    readonly instanceView: outputs.compute.v20200601.VirtualMachineRunCommandInstanceViewResponse;
    /**
     * Resource location
     */
    readonly location: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * Specifies the Azure storage blob where script output stream will be uploaded.
     */
    readonly outputBlobUri?: string;
    /**
     * The parameters used by the script.
     */
    readonly parameters?: outputs.compute.v20200601.RunCommandInputParameterResponse[];
    /**
     * The parameters used by the script.
     */
    readonly protectedParameters?: outputs.compute.v20200601.RunCommandInputParameterResponse[];
    /**
     * The provisioning state, which only appears in the response.
     */
    readonly provisioningState: string;
    /**
     * Specifies the user account password on the VM when executing the run command.
     */
    readonly runAsPassword?: string;
    /**
     * Specifies the user account on the VM when executing the run command.
     */
    readonly runAsUser?: string;
    /**
     * The source of the run command script.
     */
    readonly source?: outputs.compute.v20200601.VirtualMachineRunCommandScriptSourceResponse;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * The timeout in seconds to execute the run command.
     */
    readonly timeoutInSeconds?: number;
    /**
     * Resource type
     */
    readonly type: string;
}
