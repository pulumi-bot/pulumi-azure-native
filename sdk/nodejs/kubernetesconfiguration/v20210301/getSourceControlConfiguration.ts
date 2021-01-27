// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

export function getSourceControlConfiguration(args: GetSourceControlConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceControlConfigurationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:kubernetesconfiguration/v20210301:getSourceControlConfiguration", {
        "clusterName": args.clusterName,
        "clusterResourceName": args.clusterResourceName,
        "clusterRp": args.clusterRp,
        "resourceGroupName": args.resourceGroupName,
        "sourceControlConfigurationName": args.sourceControlConfigurationName,
    }, opts);
}

export interface GetSourceControlConfigurationArgs {
    /**
     * The name of the kubernetes cluster.
     */
    readonly clusterName: string;
    /**
     * The Kubernetes cluster resource name - either managedClusters (for AKS clusters) or connectedClusters (for OnPrem K8S clusters).
     */
    readonly clusterResourceName: string;
    /**
     * The Kubernetes cluster RP - either Microsoft.ContainerService (for AKS clusters) or Microsoft.Kubernetes (for OnPrem K8S clusters).
     */
    readonly clusterRp: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * Name of the Source Control Configuration.
     */
    readonly sourceControlConfigurationName: string;
}

/**
 * The SourceControl Configuration object returned in Get & Put response.
 */
export interface GetSourceControlConfigurationResult {
    /**
     * Compliance Status of the Configuration
     */
    readonly complianceStatus: outputs.kubernetesconfiguration.v20210301.ComplianceStatusResponse;
    /**
     * Name-value pairs of protected configuration settings for the configuration
     */
    readonly configurationProtectedSettings?: {[key: string]: string};
    /**
     * Option to enable Helm Operator for this git configuration.
     */
    readonly enableHelmOperator?: boolean;
    /**
     * Properties for Helm operator.
     */
    readonly helmOperatorProperties?: outputs.kubernetesconfiguration.v20210301.HelmOperatorPropertiesResponse;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Instance name of the operator - identifying the specific configuration.
     */
    readonly operatorInstanceName?: string;
    /**
     * The namespace to which this operator is installed to. Maximum of 253 lower case alphanumeric characters, hyphen and period only.
     */
    readonly operatorNamespace?: string;
    /**
     * Any Parameters for the Operator instance in string format.
     */
    readonly operatorParams?: string;
    /**
     * Scope at which the operator will be installed.
     */
    readonly operatorScope?: string;
    /**
     * Type of the operator
     */
    readonly operatorType?: string;
    /**
     * The provisioning state of the resource provider.
     */
    readonly provisioningState: string;
    /**
     * Public Key associated with this SourceControl configuration (either generated within the cluster or provided by the user).
     */
    readonly repositoryPublicKey: string;
    /**
     * Url of the SourceControl Repository.
     */
    readonly repositoryUrl?: string;
    /**
     * Base64-encoded known_hosts contents containing public SSH keys required to access private Git instances
     */
    readonly sshKnownHostsContents?: string;
    /**
     * Top level metadata https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/common-api-contracts.md#system-metadata-for-all-azure-resources
     */
    readonly systemData: outputs.kubernetesconfiguration.v20210301.SystemDataResponse;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
