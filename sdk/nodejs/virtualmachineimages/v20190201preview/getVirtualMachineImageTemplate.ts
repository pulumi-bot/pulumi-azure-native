// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getVirtualMachineImageTemplate(args: GetVirtualMachineImageTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualMachineImageTemplateResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:virtualmachineimages/v20190201preview:getVirtualMachineImageTemplate", {
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

export interface GetVirtualMachineImageTemplateResult {
    /**
     * Specifies the properties used to describe the customization steps of the image, like Image source etc
     */
    readonly customize?: outputs.virtualmachineimages.v20190201preview.ImageTemplatePowerShellCustomizerResponse | outputs.virtualmachineimages.v20190201preview.ImageTemplateRestartCustomizerResponse | outputs.virtualmachineimages.v20190201preview.ImageTemplateShellCustomizerResponse[];
    /**
     * The distribution targets where the image output needs to go to.
     */
    readonly distribute: outputs.virtualmachineimages.v20190201preview.ImageTemplateManagedImageDistributorResponse | outputs.virtualmachineimages.v20190201preview.ImageTemplateSharedImageDistributorResponse | outputs.virtualmachineimages.v20190201preview.ImageTemplateVhdDistributorResponse[];
    /**
     * State of 'run' that is currently executing or was last executed.
     */
    readonly lastRunStatus: outputs.virtualmachineimages.v20190201preview.ImageTemplateLastRunStatusResponse;
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
    readonly provisioningError: outputs.virtualmachineimages.v20190201preview.ProvisioningErrorResponse;
    /**
     * Provisioning state of the resource
     */
    readonly provisioningState: string;
    /**
     * Specifies the properties used to describe the source image.
     */
    readonly source: outputs.virtualmachineimages.v20190201preview.ImageTemplateIsoSourceResponse | outputs.virtualmachineimages.v20190201preview.ImageTemplateManagedImageSourceResponse | outputs.virtualmachineimages.v20190201preview.ImageTemplatePlatformImageSourceResponse;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type: string;
}
