// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getOperationalizationCluster(args: GetOperationalizationClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetOperationalizationClusterResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:machinelearningcompute/v20170801preview:getOperationalizationCluster", {
        "clusterName": args.clusterName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetOperationalizationClusterArgs {
    /**
     * The name of the cluster.
     */
    readonly clusterName: string;
    /**
     * Name of the resource group in which the cluster is located.
     */
    readonly resourceGroupName: string;
}

/**
 * Instance of an Azure ML Operationalization Cluster resource.
 */
export interface GetOperationalizationClusterResult {
    /**
     * AppInsights configuration.
     */
    readonly appInsights?: outputs.machinelearningcompute.v20170801preview.AppInsightsPropertiesResponse;
    /**
     * The cluster type.
     */
    readonly clusterType: string;
    /**
     * Container Registry properties.
     */
    readonly containerRegistry?: outputs.machinelearningcompute.v20170801preview.ContainerRegistryPropertiesResponse;
    /**
     * Parameters for the Azure Container Service cluster.
     */
    readonly containerService?: outputs.machinelearningcompute.v20170801preview.AcsClusterPropertiesResponse;
    /**
     * The date and time when the cluster was created.
     */
    readonly createdOn: string;
    /**
     * The description of the cluster.
     */
    readonly description?: string;
    /**
     * Contains global configuration for the web services in the cluster.
     */
    readonly globalServiceConfiguration?: outputs.machinelearningcompute.v20170801preview.GlobalServiceConfigurationResponse;
    /**
     * Specifies the location of the resource.
     */
    readonly location: string;
    /**
     * The date and time when the cluster was last modified.
     */
    readonly modifiedOn: string;
    /**
     * Specifies the name of the resource.
     */
    readonly name: string;
    /**
     * List of provisioning errors reported by the resource provider.
     */
    readonly provisioningErrors: outputs.machinelearningcompute.v20170801preview.ErrorResponseWrapperResponse[];
    /**
     * The provision state of the cluster. Valid values are Unknown, Updating, Provisioning, Succeeded, and Failed.
     */
    readonly provisioningState: string;
    /**
     * Storage Account properties.
     */
    readonly storageAccount?: outputs.machinelearningcompute.v20170801preview.StorageAccountPropertiesResponse;
    /**
     * Contains resource tags defined as key/value pairs.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Specifies the type of the resource.
     */
    readonly type: string;
}
