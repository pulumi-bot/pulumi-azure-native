// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
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
        return new Cluster(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:servicefabric/v20190301:Cluster';

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
     * The list of add-on features to enable in the cluster.
     */
    public readonly addOnFeatures!: pulumi.Output<string[] | undefined>;
    /**
     * The Service Fabric runtime versions available for this cluster.
     */
    public /*out*/ readonly availableClusterVersions!: pulumi.Output<outputs.servicefabric.v20190301.ClusterVersionDetailsResponse[]>;
    /**
     * The AAD authentication settings of the cluster.
     */
    public readonly azureActiveDirectory!: pulumi.Output<outputs.servicefabric.v20190301.AzureActiveDirectoryResponse | undefined>;
    /**
     * The certificate to use for securing the cluster. The certificate provided will be used for node to node security within the cluster, SSL certificate for cluster management endpoint and default admin client.
     */
    public readonly certificate!: pulumi.Output<outputs.servicefabric.v20190301.CertificateDescriptionResponse | undefined>;
    /**
     * Describes a list of server certificates referenced by common name that are used to secure the cluster.
     */
    public readonly certificateCommonNames!: pulumi.Output<outputs.servicefabric.v20190301.ServerCertificateCommonNamesResponse | undefined>;
    /**
     * The list of client certificates referenced by common name that are allowed to manage the cluster.
     */
    public readonly clientCertificateCommonNames!: pulumi.Output<outputs.servicefabric.v20190301.ClientCertificateCommonNameResponse[] | undefined>;
    /**
     * The list of client certificates referenced by thumbprint that are allowed to manage the cluster.
     */
    public readonly clientCertificateThumbprints!: pulumi.Output<outputs.servicefabric.v20190301.ClientCertificateThumbprintResponse[] | undefined>;
    /**
     * The Service Fabric runtime version of the cluster. This property can only by set the user when **upgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
     */
    public readonly clusterCodeVersion!: pulumi.Output<string | undefined>;
    /**
     * The Azure Resource Provider endpoint. A system service in the cluster connects to this  endpoint.
     */
    public /*out*/ readonly clusterEndpoint!: pulumi.Output<string>;
    /**
     * A service generated unique identifier for the cluster resource.
     */
    public /*out*/ readonly clusterId!: pulumi.Output<string>;
    /**
     * The current state of the cluster.
     *
     *   - WaitingForNodes - Indicates that the cluster resource is created and the resource provider is waiting for Service Fabric VM extension to boot up and report to it.
     *   - Deploying - Indicates that the Service Fabric runtime is being installed on the VMs. Cluster resource will be in this state until the cluster boots up and system services are up.
     *   - BaselineUpgrade - Indicates that the cluster is upgrading to establishes the cluster version. This upgrade is automatically initiated when the cluster boots up for the first time.
     *   - UpdatingUserConfiguration - Indicates that the cluster is being upgraded with the user provided configuration.
     *   - UpdatingUserCertificate - Indicates that the cluster is being upgraded with the user provided certificate.
     *   - UpdatingInfrastructure - Indicates that the cluster is being upgraded with the latest Service Fabric runtime version. This happens only when the **upgradeMode** is set to 'Automatic'.
     *   - EnforcingClusterVersion - Indicates that cluster is on a different version than expected and the cluster is being upgraded to the expected version.
     *   - UpgradeServiceUnreachable - Indicates that the system service in the cluster is no longer polling the Resource Provider. Clusters in this state cannot be managed by the Resource Provider.
     *   - AutoScale - Indicates that the ReliabilityLevel of the cluster is being adjusted.
     *   - Ready - Indicates that the cluster is in a stable state.
     */
    public /*out*/ readonly clusterState!: pulumi.Output<string>;
    /**
     * The storage account information for storing Service Fabric diagnostic logs.
     */
    public readonly diagnosticsStorageAccountConfig!: pulumi.Output<outputs.servicefabric.v20190301.DiagnosticsStorageAccountConfigResponse | undefined>;
    /**
     * Azure resource etag.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Indicates if the event store service is enabled.
     */
    public readonly eventStoreServiceEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The list of custom fabric settings to configure the cluster.
     */
    public readonly fabricSettings!: pulumi.Output<outputs.servicefabric.v20190301.SettingsSectionDescriptionResponse[] | undefined>;
    /**
     * Azure resource location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The http management endpoint of the cluster.
     */
    public readonly managementEndpoint!: pulumi.Output<string>;
    /**
     * Azure resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The list of node types in the cluster.
     */
    public readonly nodeTypes!: pulumi.Output<outputs.servicefabric.v20190301.NodeTypeDescriptionResponse[]>;
    /**
     * The provisioning state of the cluster resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The reliability level sets the replica set size of system services. Learn about [ReliabilityLevel](https://docs.microsoft.com/azure/service-fabric/service-fabric-cluster-capacity).
     *
     *   - None - Run the System services with a target replica set count of 1. This should only be used for test clusters.
     *   - Bronze - Run the System services with a target replica set count of 3. This should only be used for test clusters.
     *   - Silver - Run the System services with a target replica set count of 5.
     *   - Gold - Run the System services with a target replica set count of 7.
     *   - Platinum - Run the System services with a target replica set count of 9.
     */
    public readonly reliabilityLevel!: pulumi.Output<string | undefined>;
    /**
     * The server certificate used by reverse proxy.
     */
    public readonly reverseProxyCertificate!: pulumi.Output<outputs.servicefabric.v20190301.CertificateDescriptionResponse | undefined>;
    /**
     * Describes a list of server certificates referenced by common name that are used to secure the cluster.
     */
    public readonly reverseProxyCertificateCommonNames!: pulumi.Output<outputs.servicefabric.v20190301.ServerCertificateCommonNamesResponse | undefined>;
    /**
     * Azure resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Azure resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The policy to use when upgrading the cluster.
     */
    public readonly upgradeDescription!: pulumi.Output<outputs.servicefabric.v20190301.ClusterUpgradePolicyResponse | undefined>;
    /**
     * The upgrade mode of the cluster when new Service Fabric runtime version is available.
     *
     *   - Automatic - The cluster will be automatically upgraded to the latest Service Fabric runtime version as soon as it is available.
     *   - Manual - The cluster will not be automatically upgraded to the latest Service Fabric runtime version. The cluster is upgraded by setting the **clusterCodeVersion** property in the cluster resource.
     */
    public readonly upgradeMode!: pulumi.Output<string | undefined>;
    /**
     * The VM image VMSS has been configured with. Generic names such as Windows or Linux can be used.
     */
    public readonly vmImage!: pulumi.Output<string | undefined>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.clusterName === undefined) {
                throw new Error("Missing required property 'clusterName'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.managementEndpoint === undefined) {
                throw new Error("Missing required property 'managementEndpoint'");
            }
            if (!args || args.nodeTypes === undefined) {
                throw new Error("Missing required property 'nodeTypes'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["addOnFeatures"] = args ? args.addOnFeatures : undefined;
            inputs["azureActiveDirectory"] = args ? args.azureActiveDirectory : undefined;
            inputs["certificate"] = args ? args.certificate : undefined;
            inputs["certificateCommonNames"] = args ? args.certificateCommonNames : undefined;
            inputs["clientCertificateCommonNames"] = args ? args.clientCertificateCommonNames : undefined;
            inputs["clientCertificateThumbprints"] = args ? args.clientCertificateThumbprints : undefined;
            inputs["clusterCodeVersion"] = args ? args.clusterCodeVersion : undefined;
            inputs["clusterName"] = args ? args.clusterName : undefined;
            inputs["diagnosticsStorageAccountConfig"] = args ? args.diagnosticsStorageAccountConfig : undefined;
            inputs["eventStoreServiceEnabled"] = args ? args.eventStoreServiceEnabled : undefined;
            inputs["fabricSettings"] = args ? args.fabricSettings : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["managementEndpoint"] = args ? args.managementEndpoint : undefined;
            inputs["nodeTypes"] = args ? args.nodeTypes : undefined;
            inputs["reliabilityLevel"] = args ? args.reliabilityLevel : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["reverseProxyCertificate"] = args ? args.reverseProxyCertificate : undefined;
            inputs["reverseProxyCertificateCommonNames"] = args ? args.reverseProxyCertificateCommonNames : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["upgradeDescription"] = args ? args.upgradeDescription : undefined;
            inputs["upgradeMode"] = args ? args.upgradeMode : undefined;
            inputs["vmImage"] = args ? args.vmImage : undefined;
            inputs["availableClusterVersions"] = undefined /*out*/;
            inputs["clusterEndpoint"] = undefined /*out*/;
            inputs["clusterId"] = undefined /*out*/;
            inputs["clusterState"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["addOnFeatures"] = undefined /*out*/;
            inputs["availableClusterVersions"] = undefined /*out*/;
            inputs["azureActiveDirectory"] = undefined /*out*/;
            inputs["certificate"] = undefined /*out*/;
            inputs["certificateCommonNames"] = undefined /*out*/;
            inputs["clientCertificateCommonNames"] = undefined /*out*/;
            inputs["clientCertificateThumbprints"] = undefined /*out*/;
            inputs["clusterCodeVersion"] = undefined /*out*/;
            inputs["clusterEndpoint"] = undefined /*out*/;
            inputs["clusterId"] = undefined /*out*/;
            inputs["clusterState"] = undefined /*out*/;
            inputs["diagnosticsStorageAccountConfig"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["eventStoreServiceEnabled"] = undefined /*out*/;
            inputs["fabricSettings"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["managementEndpoint"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["nodeTypes"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["reliabilityLevel"] = undefined /*out*/;
            inputs["reverseProxyCertificate"] = undefined /*out*/;
            inputs["reverseProxyCertificateCommonNames"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["upgradeDescription"] = undefined /*out*/;
            inputs["upgradeMode"] = undefined /*out*/;
            inputs["vmImage"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:servicefabric/latest:Cluster" }, { type: "azure-nextgen:servicefabric/v20160901:Cluster" }, { type: "azure-nextgen:servicefabric/v20170701preview:Cluster" }, { type: "azure-nextgen:servicefabric/v20180201:Cluster" }, { type: "azure-nextgen:servicefabric/v20190301preview:Cluster" }, { type: "azure-nextgen:servicefabric/v20190601preview:Cluster" }, { type: "azure-nextgen:servicefabric/v20191101preview:Cluster" }, { type: "azure-nextgen:servicefabric/v20200301:Cluster" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Cluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * The list of add-on features to enable in the cluster.
     */
    readonly addOnFeatures?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The AAD authentication settings of the cluster.
     */
    readonly azureActiveDirectory?: pulumi.Input<inputs.servicefabric.v20190301.AzureActiveDirectory>;
    /**
     * The certificate to use for securing the cluster. The certificate provided will be used for node to node security within the cluster, SSL certificate for cluster management endpoint and default admin client.
     */
    readonly certificate?: pulumi.Input<inputs.servicefabric.v20190301.CertificateDescription>;
    /**
     * Describes a list of server certificates referenced by common name that are used to secure the cluster.
     */
    readonly certificateCommonNames?: pulumi.Input<inputs.servicefabric.v20190301.ServerCertificateCommonNames>;
    /**
     * The list of client certificates referenced by common name that are allowed to manage the cluster.
     */
    readonly clientCertificateCommonNames?: pulumi.Input<pulumi.Input<inputs.servicefabric.v20190301.ClientCertificateCommonName>[]>;
    /**
     * The list of client certificates referenced by thumbprint that are allowed to manage the cluster.
     */
    readonly clientCertificateThumbprints?: pulumi.Input<pulumi.Input<inputs.servicefabric.v20190301.ClientCertificateThumbprint>[]>;
    /**
     * The Service Fabric runtime version of the cluster. This property can only by set the user when **upgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
     */
    readonly clusterCodeVersion?: pulumi.Input<string>;
    /**
     * The name of the cluster resource.
     */
    readonly clusterName: pulumi.Input<string>;
    /**
     * The storage account information for storing Service Fabric diagnostic logs.
     */
    readonly diagnosticsStorageAccountConfig?: pulumi.Input<inputs.servicefabric.v20190301.DiagnosticsStorageAccountConfig>;
    /**
     * Indicates if the event store service is enabled.
     */
    readonly eventStoreServiceEnabled?: pulumi.Input<boolean>;
    /**
     * The list of custom fabric settings to configure the cluster.
     */
    readonly fabricSettings?: pulumi.Input<pulumi.Input<inputs.servicefabric.v20190301.SettingsSectionDescription>[]>;
    /**
     * Azure resource location.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The http management endpoint of the cluster.
     */
    readonly managementEndpoint: pulumi.Input<string>;
    /**
     * The list of node types in the cluster.
     */
    readonly nodeTypes: pulumi.Input<pulumi.Input<inputs.servicefabric.v20190301.NodeTypeDescription>[]>;
    /**
     * The reliability level sets the replica set size of system services. Learn about [ReliabilityLevel](https://docs.microsoft.com/azure/service-fabric/service-fabric-cluster-capacity).
     *
     *   - None - Run the System services with a target replica set count of 1. This should only be used for test clusters.
     *   - Bronze - Run the System services with a target replica set count of 3. This should only be used for test clusters.
     *   - Silver - Run the System services with a target replica set count of 5.
     *   - Gold - Run the System services with a target replica set count of 7.
     *   - Platinum - Run the System services with a target replica set count of 9.
     */
    readonly reliabilityLevel?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The server certificate used by reverse proxy.
     */
    readonly reverseProxyCertificate?: pulumi.Input<inputs.servicefabric.v20190301.CertificateDescription>;
    /**
     * Describes a list of server certificates referenced by common name that are used to secure the cluster.
     */
    readonly reverseProxyCertificateCommonNames?: pulumi.Input<inputs.servicefabric.v20190301.ServerCertificateCommonNames>;
    /**
     * Azure resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The policy to use when upgrading the cluster.
     */
    readonly upgradeDescription?: pulumi.Input<inputs.servicefabric.v20190301.ClusterUpgradePolicy>;
    /**
     * The upgrade mode of the cluster when new Service Fabric runtime version is available.
     *
     *   - Automatic - The cluster will be automatically upgraded to the latest Service Fabric runtime version as soon as it is available.
     *   - Manual - The cluster will not be automatically upgraded to the latest Service Fabric runtime version. The cluster is upgraded by setting the **clusterCodeVersion** property in the cluster resource.
     */
    readonly upgradeMode?: pulumi.Input<string>;
    /**
     * The VM image VMSS has been configured with. Generic names such as Windows or Linux can be used.
     */
    readonly vmImage?: pulumi.Input<string>;
}
