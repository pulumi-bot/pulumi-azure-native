// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getApplication(args: GetApplicationArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:desktopvirtualization/v20191210preview:getApplication", {
        "applicationGroupName": args.applicationGroupName,
        "applicationName": args.applicationName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetApplicationArgs {
    /**
     * The name of the application group
     */
    readonly applicationGroupName: string;
    /**
     * The name of the application within the specified application group
     */
    readonly applicationName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
}

/**
 * Schema for Application properties.
 */
export interface GetApplicationResult {
    /**
     * Command Line Arguments for Application.
     */
    readonly commandLineArguments?: string;
    /**
     * Specifies whether this published application can be launched with command line arguments provided by the client, command line arguments specified at publish time, or no command line arguments at all.
     */
    readonly commandLineSetting: string;
    /**
     * Description of Application.
     */
    readonly description?: string;
    /**
     * Specifies a path for the executable file for the application.
     */
    readonly filePath?: string;
    /**
     * Friendly name of Application.
     */
    readonly friendlyName?: string;
    /**
     * the icon a 64 bit string as a byte array.
     */
    readonly iconContent: string;
    /**
     * Hash of the icon.
     */
    readonly iconHash: string;
    /**
     * Index of the icon.
     */
    readonly iconIndex?: number;
    /**
     * Path to icon.
     */
    readonly iconPath?: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Specifies whether to show the RemoteApp program in the RD Web Access server.
     */
    readonly showInPortal?: boolean;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    readonly type: string;
}
