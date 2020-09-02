// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getVirtualMachineImageTemplate(args: GetVirtualMachineImageTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualMachineImageTemplateResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:virtualmachineimages/v20190501preview:getVirtualMachineImageTemplate", {
        "imageTemplateName": args.imageTemplateName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetVirtualMachineImageTemplateArgs {
    /**
     * The name of the image Template
     */
    readonly imageTemplateName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * Image template is an ARM resource managed by Microsoft.VirtualMachineImages provider
 */
export interface GetVirtualMachineImageTemplateResult {
    /**
     * Maximum duration to wait while building the image template. Omit or specify 0 to use the default (4 hours).
     */
    readonly buildTimeoutInMinutes?: number;
    /**
     * Specifies the properties used to describe the customization steps of the image, like Image source etc
     */
    readonly customize?: outputs.virtualmachineimages.v20190501preview.ImageTemplateCustomizerResponse[];
    /**
     * The distribution targets where the image output needs to go to.
     */
    readonly distribute: outputs.virtualmachineimages.v20190501preview.ImageTemplateDistributorResponse[];
    /**
     * The identity of the image template, if configured.
     */
    readonly identity?: outputs.virtualmachineimages.v20190501preview.ImageTemplateIdentityResponse;
    /**
     * State of 'run' that is currently executing or was last executed.
     */
    readonly lastRunStatus: outputs.virtualmachineimages.v20190501preview.ImageTemplateLastRunStatusResponse;
    /**
     * Resource location
     */
    readonly location: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * Provisioning error, if any
     */
    readonly provisioningError: outputs.virtualmachineimages.v20190501preview.ProvisioningErrorResponse;
    /**
     * Provisioning state of the resource
     */
    readonly provisioningState: string;
    /**
     * Specifies the properties used to describe the source image.
     */
    readonly source: outputs.virtualmachineimages.v20190501preview.ImageTemplateSourceResponse;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type: string;
    /**
     * Describes how virtual machine is set up to build images
     */
    readonly vmProfile?: outputs.virtualmachineimages.v20190501preview.ImageTemplateVmProfileResponse;
}
