// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getBuildTask(args: GetBuildTaskArgs, opts?: pulumi.InvokeOptions): Promise<GetBuildTaskResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:containerregistry/v20180201preview:getBuildTask", {
        "buildTaskName": args.buildTaskName,
        "registryName": args.registryName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetBuildTaskArgs {
    /**
     * The name of the container registry build task.
     */
    readonly buildTaskName: string;
    /**
     * The name of the container registry.
     */
    readonly registryName: string;
    /**
     * The name of the resource group to which the container registry belongs.
     */
    readonly resourceGroupName: string;
}

/**
 * The build task that has the resource properties and all build items. The build task will have all information to schedule a build against it.
 */
export interface GetBuildTaskResult {
    /**
     * The alternative updatable name for a build task.
     */
    readonly alias: string;
    /**
     * The creation date of build task.
     */
    readonly creationDate: string;
    /**
     * The location of the resource. This cannot be changed after the resource is created.
     */
    readonly location: string;
    /**
     * The name of the resource.
     */
    readonly name: string;
    /**
     * The platform properties against which the build has to happen.
     */
    readonly platform: outputs.containerregistry.v20180201preview.PlatformPropertiesResponse;
    /**
     * The provisioning state of the build task.
     */
    readonly provisioningState: string;
    /**
     * The properties that describes the source(code) for the build task.
     */
    readonly sourceRepository: outputs.containerregistry.v20180201preview.SourceRepositoryPropertiesResponse;
    /**
     * The current status of build task.
     */
    readonly status?: string;
    /**
     * The tags of the resource.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Build timeout in seconds.
     */
    readonly timeout?: number;
    /**
     * The type of the resource.
     */
    readonly type: string;
}
