// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getWorkspace(args: GetWorkspaceArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkspaceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:machinelearningservices/v20200515preview:getWorkspace", {
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetWorkspaceArgs {
    /**
     * Name of the resource group in which workspace is located.
     */
    readonly resourceGroupName: string;
    /**
     * Name of Azure Machine Learning workspace.
     */
    readonly workspaceName: string;
}

/**
 * An object that represents a machine learning workspace.
 */
export interface GetWorkspaceResult {
    /**
     * The flag to indicate whether to allow public access when behind VNet.
     */
    readonly allowPublicAccessWhenBehindVnet?: boolean;
    /**
     * ARM id of the application insights associated with this workspace. This cannot be changed once the workspace has been created
     */
    readonly applicationInsights?: string;
    /**
     * ARM id of the container registry associated with this workspace. This cannot be changed once the workspace has been created
     */
    readonly containerRegistry?: string;
    /**
     * The creation time of the machine learning workspace in ISO8601 format.
     */
    readonly creationTime: string;
    /**
     * The description of this workspace.
     */
    readonly description?: string;
    /**
     * Url for the discovery service to identify regional endpoints for machine learning experimentation services
     */
    readonly discoveryUrl?: string;
    /**
     * The encryption settings of Azure ML workspace.
     */
    readonly encryption?: outputs.machinelearningservices.v20200515preview.EncryptionPropertyResponse;
    /**
     * The friendly name for this workspace. This name in mutable
     */
    readonly friendlyName?: string;
    /**
     * The flag to signal HBI data in the workspace and reduce diagnostic data collected by the service
     */
    readonly hbiWorkspace?: boolean;
    /**
     * The identity of the resource.
     */
    readonly identity?: outputs.machinelearningservices.v20200515preview.IdentityResponse;
    /**
     * The compute name for image build
     */
    readonly imageBuildCompute?: string;
    /**
     * ARM id of the key vault associated with this workspace. This cannot be changed once the workspace has been created
     */
    readonly keyVault?: string;
    /**
     * Specifies the location of the resource.
     */
    readonly location?: string;
    /**
     * Specifies the name of the resource.
     */
    readonly name: string;
    /**
     * The list of private endpoint connections in the workspace.
     */
    readonly privateEndpointConnections: outputs.machinelearningservices.v20200515preview.PrivateEndpointConnectionResponse[];
    /**
     * Count of private connections in the workspace
     */
    readonly privateLinkCount: number;
    /**
     * The current deployment state of workspace resource. The provisioningState is to indicate states for resource provisioning.
     */
    readonly provisioningState: string;
    /**
     * The name of the managed resource group created by workspace RP in customer subscription if the workspace is CMK workspace
     */
    readonly serviceProvisionedResourceGroup: string;
    /**
     * The list of shared private link resources in this workspace.
     */
    readonly sharedPrivateLinkResources?: outputs.machinelearningservices.v20200515preview.SharedPrivateLinkResourceResponse[];
    /**
     * The sku of the workspace.
     */
    readonly sku?: outputs.machinelearningservices.v20200515preview.SkuResponse;
    /**
     * ARM id of the storage account associated with this workspace. This cannot be changed once the workspace has been created
     */
    readonly storageAccount?: string;
    /**
     * Contains resource tags defined as key/value pairs.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Specifies the type of the resource.
     */
    readonly type: string;
    /**
     * The immutable id associated with this workspace.
     */
    readonly workspaceId: string;
}
