// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The SourceControl Configuration object.
 */
export class SourceControlConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing SourceControlConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SourceControlConfiguration {
        return new SourceControlConfiguration(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:kubernetesconfiguration/v20191101preview:SourceControlConfiguration';

    /**
     * Returns true if the given object is an instance of SourceControlConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SourceControlConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SourceControlConfiguration.__pulumiType;
    }

    /**
     * Compliance Status of the Configuration
     */
    public /*out*/ readonly complianceStatus!: pulumi.Output<outputs.kubernetesconfiguration.v20191101preview.ComplianceStatusResponse>;
    /**
     * Option to enable Helm Operator for this git configuration.
     */
    public readonly enableHelmOperator!: pulumi.Output<string | undefined>;
    /**
     * Properties for Helm operator.
     */
    public readonly helmOperatorProperties!: pulumi.Output<outputs.kubernetesconfiguration.v20191101preview.HelmOperatorPropertiesResponse | undefined>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Instance name of the operator - identifying the specific configuration.
     */
    public readonly operatorInstanceName!: pulumi.Output<string | undefined>;
    /**
     * The namespace to which this operator is installed to. Maximum of 253 lower case alphanumeric characters, hyphen and period only.
     */
    public readonly operatorNamespace!: pulumi.Output<string | undefined>;
    /**
     * Any Parameters for the Operator instance in string format.
     */
    public readonly operatorParams!: pulumi.Output<string | undefined>;
    /**
     * Scope at which the operator will be installed.
     */
    public readonly operatorScope!: pulumi.Output<string | undefined>;
    /**
     * Type of the operator
     */
    public readonly operatorType!: pulumi.Output<string | undefined>;
    /**
     * The provisioning state of the resource provider.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Public Key associated with this SourceControl configuration (either generated within the cluster or provided by the user).
     */
    public /*out*/ readonly repositoryPublicKey!: pulumi.Output<string>;
    /**
     * Url of the SourceControl Repository.
     */
    public readonly repositoryUrl!: pulumi.Output<string | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a SourceControlConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SourceControlConfigurationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.clusterName === undefined) {
                throw new Error("Missing required property 'clusterName'");
            }
            if (!args || args.clusterResourceName === undefined) {
                throw new Error("Missing required property 'clusterResourceName'");
            }
            if (!args || args.clusterRp === undefined) {
                throw new Error("Missing required property 'clusterRp'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sourceControlConfigurationName === undefined) {
                throw new Error("Missing required property 'sourceControlConfigurationName'");
            }
            inputs["clusterName"] = args ? args.clusterName : undefined;
            inputs["clusterResourceName"] = args ? args.clusterResourceName : undefined;
            inputs["clusterRp"] = args ? args.clusterRp : undefined;
            inputs["enableHelmOperator"] = args ? args.enableHelmOperator : undefined;
            inputs["helmOperatorProperties"] = args ? args.helmOperatorProperties : undefined;
            inputs["operatorInstanceName"] = args ? args.operatorInstanceName : undefined;
            inputs["operatorNamespace"] = args ? args.operatorNamespace : undefined;
            inputs["operatorParams"] = args ? args.operatorParams : undefined;
            inputs["operatorScope"] = args ? args.operatorScope : undefined;
            inputs["operatorType"] = args ? args.operatorType : undefined;
            inputs["repositoryUrl"] = args ? args.repositoryUrl : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sourceControlConfigurationName"] = args ? args.sourceControlConfigurationName : undefined;
            inputs["complianceStatus"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["repositoryPublicKey"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["complianceStatus"] = undefined /*out*/;
            inputs["enableHelmOperator"] = undefined /*out*/;
            inputs["helmOperatorProperties"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["operatorInstanceName"] = undefined /*out*/;
            inputs["operatorNamespace"] = undefined /*out*/;
            inputs["operatorParams"] = undefined /*out*/;
            inputs["operatorScope"] = undefined /*out*/;
            inputs["operatorType"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["repositoryPublicKey"] = undefined /*out*/;
            inputs["repositoryUrl"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:kubernetesconfiguration/v20201001preview:SourceControlConfiguration" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(SourceControlConfiguration.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a SourceControlConfiguration resource.
 */
export interface SourceControlConfigurationArgs {
    /**
     * The name of the kubernetes cluster.
     */
    readonly clusterName: pulumi.Input<string>;
    /**
     * The Kubernetes cluster resource name - either managedClusters (for AKS clusters) or connectedClusters (for OnPrem K8S clusters).
     */
    readonly clusterResourceName: pulumi.Input<string>;
    /**
     * The Kubernetes cluster RP - either Microsoft.ContainerService (for AKS clusters) or Microsoft.Kubernetes (for OnPrem K8S clusters).
     */
    readonly clusterRp: pulumi.Input<string>;
    /**
     * Option to enable Helm Operator for this git configuration.
     */
    readonly enableHelmOperator?: pulumi.Input<string>;
    /**
     * Properties for Helm operator.
     */
    readonly helmOperatorProperties?: pulumi.Input<inputs.kubernetesconfiguration.v20191101preview.HelmOperatorProperties>;
    /**
     * Instance name of the operator - identifying the specific configuration.
     */
    readonly operatorInstanceName?: pulumi.Input<string>;
    /**
     * The namespace to which this operator is installed to. Maximum of 253 lower case alphanumeric characters, hyphen and period only.
     */
    readonly operatorNamespace?: pulumi.Input<string>;
    /**
     * Any Parameters for the Operator instance in string format.
     */
    readonly operatorParams?: pulumi.Input<string>;
    /**
     * Scope at which the operator will be installed.
     */
    readonly operatorScope?: pulumi.Input<string>;
    /**
     * Type of the operator
     */
    readonly operatorType?: pulumi.Input<string>;
    /**
     * Url of the SourceControl Repository.
     */
    readonly repositoryUrl?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Name of the Source Control Configuration.
     */
    readonly sourceControlConfigurationName: pulumi.Input<string>;
}
