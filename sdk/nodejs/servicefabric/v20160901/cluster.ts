// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The cluster resource
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:servicefabric/v20160901:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The cluster resource properties
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.servicefabric.v20160901.ClusterPropertiesResponse>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ClusterArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.managementEndpoint === undefined) {
                throw new Error("Missing required property 'managementEndpoint'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.nodeTypes === undefined) {
                throw new Error("Missing required property 'nodeTypes'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["azureActiveDirectory"] = args ? args.azureActiveDirectory : undefined;
            inputs["certificate"] = args ? args.certificate : undefined;
            inputs["clientCertificateCommonNames"] = args ? args.clientCertificateCommonNames : undefined;
            inputs["clientCertificateThumbprints"] = args ? args.clientCertificateThumbprints : undefined;
            inputs["clusterCodeVersion"] = args ? args.clusterCodeVersion : undefined;
            inputs["diagnosticsStorageAccountConfig"] = args ? args.diagnosticsStorageAccountConfig : undefined;
            inputs["fabricSettings"] = args ? args.fabricSettings : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["managementEndpoint"] = args ? args.managementEndpoint : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nodeTypes"] = args ? args.nodeTypes : undefined;
            inputs["reliabilityLevel"] = args ? args.reliabilityLevel : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["reverseProxyCertificate"] = args ? args.reverseProxyCertificate : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["upgradeDescription"] = args ? args.upgradeDescription : undefined;
            inputs["upgradeMode"] = args ? args.upgradeMode : undefined;
            inputs["vmImage"] = args ? args.vmImage : undefined;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Cluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * The settings to enable AAD authentication on the cluster
     */
    readonly azureActiveDirectory?: pulumi.Input<inputs.servicefabric.v20160901.AzureActiveDirectory>;
    /**
     * This primary certificate will be used as cluster node to node security, SSL certificate for cluster management endpoint and default admin client
     */
    readonly certificate?: pulumi.Input<inputs.servicefabric.v20160901.CertificateDescription>;
    /**
     *  List of client certificates to whitelist based on common names
     */
    readonly clientCertificateCommonNames?: pulumi.Input<pulumi.Input<inputs.servicefabric.v20160901.ClientCertificateCommonName>[]>;
    /**
     * The client thumbprint details ,it is used for client access for cluster operation
     */
    readonly clientCertificateThumbprints?: pulumi.Input<pulumi.Input<inputs.servicefabric.v20160901.ClientCertificateThumbprint>[]>;
    /**
     * The ServiceFabric code version running in your cluster
     */
    readonly clusterCodeVersion?: pulumi.Input<string>;
    /**
     * The storage diagnostics account configuration details
     */
    readonly diagnosticsStorageAccountConfig?: pulumi.Input<inputs.servicefabric.v20160901.DiagnosticsStorageAccountConfig>;
    /**
     * List of custom fabric settings to configure the cluster.
     */
    readonly fabricSettings?: pulumi.Input<pulumi.Input<inputs.servicefabric.v20160901.SettingsSectionDescription>[]>;
    /**
     * Resource location.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The http management endpoint of the cluster
     */
    readonly managementEndpoint: pulumi.Input<string>;
    /**
     * The name of the cluster resource
     */
    readonly name: pulumi.Input<string>;
    /**
     * The list of node types that make up the cluster
     */
    readonly nodeTypes: pulumi.Input<pulumi.Input<inputs.servicefabric.v20160901.NodeTypeDescription>[]>;
    /**
     * Cluster reliability level indicates replica set size of system service
     */
    readonly reliabilityLevel?: pulumi.Input<string>;
    /**
     * The name of the resource group to which the resource belongs or get created
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The server certificate used by reverse proxy
     */
    readonly reverseProxyCertificate?: pulumi.Input<inputs.servicefabric.v20160901.CertificateDescription>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The policy to use when upgrading the cluster.
     */
    readonly upgradeDescription?: pulumi.Input<inputs.servicefabric.v20160901.ClusterUpgradePolicy>;
    /**
     * Cluster upgrade mode indicates if fabric upgrade is initiated automatically by the system or not
     */
    readonly upgradeMode?: pulumi.Input<string>;
    /**
     * The name of VM image VMSS has been configured with. Generic names such as Windows or Linux can be used.
     */
    readonly vmImage?: pulumi.Input<string>;
}
