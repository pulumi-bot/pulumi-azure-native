// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The manged cluster resource
 */
export class ManagedCluster extends pulumi.CustomResource {
    /**
     * Get an existing ManagedCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ManagedCluster {
        return new ManagedCluster(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:servicefabric/v20200101preview:ManagedCluster';

    /**
     * Returns true if the given object is an instance of ManagedCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ManagedCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ManagedCluster.__pulumiType;
    }

    /**
     * vm admin user password.
     */
    public readonly adminPassword!: pulumi.Output<string | undefined>;
    /**
     * vm admin user name.
     */
    public readonly adminUserName!: pulumi.Output<string>;
    /**
     * Azure active directory.
     */
    public readonly azureActiveDirectory!: pulumi.Output<outputs.servicefabric.v20200101preview.AzureActiveDirectoryResponse | undefined>;
    /**
     * The port used for client connections to the cluster.
     */
    public readonly clientConnectionPort!: pulumi.Output<number | undefined>;
    /**
     * client certificates for the cluster.
     */
    public readonly clients!: pulumi.Output<outputs.servicefabric.v20200101preview.ClientCertificateResponse[] | undefined>;
    /**
     * The cluster certificate thumbprint used node to node communication.
     */
    public /*out*/ readonly clusterCertificateThumbprint!: pulumi.Output<string>;
    /**
     * The Service Fabric runtime version of the cluster. This property can only by set the user when **upgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
     */
    public readonly clusterCodeVersion!: pulumi.Output<string | undefined>;
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
     * Describes the policy used when upgrading the cluster.
     */
    public readonly clusterUpgradeDescription!: pulumi.Output<outputs.servicefabric.v20200101preview.ClusterUpgradePolicyResponse | undefined>;
    /**
     * The upgrade mode of the cluster when new Service Fabric runtime version is available.
     *
     *   - Automatic - The cluster will be automatically upgraded to the latest Service Fabric runtime version as soon as it is available.
     *   - Manual - The cluster will not be automatically upgraded to the latest Service Fabric runtime version. The cluster is upgraded by setting the **clusterCodeVersion** property in the cluster resource.
     */
    public readonly clusterUpgradeMode!: pulumi.Output<string | undefined>;
    /**
     * The cluster dns name.
     */
    public readonly dnsName!: pulumi.Output<string>;
    /**
     * Azure resource etag.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The list of custom fabric settings to configure the cluster.
     */
    public readonly fabricSettings!: pulumi.Output<outputs.servicefabric.v20200101preview.SettingsSectionDescriptionResponse[] | undefined>;
    /**
     * the cluster Fully qualified domain name.
     */
    public /*out*/ readonly fqdn!: pulumi.Output<string>;
    /**
     * The port used for http connections to the cluster.
     */
    public readonly httpGatewayConnectionPort!: pulumi.Output<number | undefined>;
    /**
     * Describes load balancing rules.
     */
    public readonly loadBalancingRules!: pulumi.Output<outputs.servicefabric.v20200101preview.LoadBalancingRuleResponse[] | undefined>;
    /**
     * Azure resource location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Azure resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state of the managed cluster resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The sku of the managed cluster
     */
    public readonly sku!: pulumi.Output<outputs.servicefabric.v20200101preview.SkuResponse | undefined>;
    /**
     * Azure resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Azure resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ManagedCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ManagedClusterArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.adminUserName === undefined) {
                throw new Error("Missing required property 'adminUserName'");
            }
            if (!args || args.clusterName === undefined) {
                throw new Error("Missing required property 'clusterName'");
            }
            if (!args || args.dnsName === undefined) {
                throw new Error("Missing required property 'dnsName'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["adminPassword"] = args ? args.adminPassword : undefined;
            inputs["adminUserName"] = args ? args.adminUserName : undefined;
            inputs["azureActiveDirectory"] = args ? args.azureActiveDirectory : undefined;
            inputs["clientConnectionPort"] = args ? args.clientConnectionPort : undefined;
            inputs["clients"] = args ? args.clients : undefined;
            inputs["clusterCodeVersion"] = args ? args.clusterCodeVersion : undefined;
            inputs["clusterName"] = args ? args.clusterName : undefined;
            inputs["clusterUpgradeDescription"] = args ? args.clusterUpgradeDescription : undefined;
            inputs["clusterUpgradeMode"] = args ? args.clusterUpgradeMode : undefined;
            inputs["dnsName"] = args ? args.dnsName : undefined;
            inputs["fabricSettings"] = args ? args.fabricSettings : undefined;
            inputs["httpGatewayConnectionPort"] = args ? args.httpGatewayConnectionPort : undefined;
            inputs["loadBalancingRules"] = args ? args.loadBalancingRules : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["clusterCertificateThumbprint"] = undefined /*out*/;
            inputs["clusterId"] = undefined /*out*/;
            inputs["clusterState"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["fqdn"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["adminPassword"] = undefined /*out*/;
            inputs["adminUserName"] = undefined /*out*/;
            inputs["azureActiveDirectory"] = undefined /*out*/;
            inputs["clientConnectionPort"] = undefined /*out*/;
            inputs["clients"] = undefined /*out*/;
            inputs["clusterCertificateThumbprint"] = undefined /*out*/;
            inputs["clusterCodeVersion"] = undefined /*out*/;
            inputs["clusterId"] = undefined /*out*/;
            inputs["clusterState"] = undefined /*out*/;
            inputs["clusterUpgradeDescription"] = undefined /*out*/;
            inputs["clusterUpgradeMode"] = undefined /*out*/;
            inputs["dnsName"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["fabricSettings"] = undefined /*out*/;
            inputs["fqdn"] = undefined /*out*/;
            inputs["httpGatewayConnectionPort"] = undefined /*out*/;
            inputs["loadBalancingRules"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:servicefabric/latest:ManagedCluster" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ManagedCluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ManagedCluster resource.
 */
export interface ManagedClusterArgs {
    /**
     * vm admin user password.
     */
    readonly adminPassword?: pulumi.Input<string>;
    /**
     * vm admin user name.
     */
    readonly adminUserName: pulumi.Input<string>;
    /**
     * Azure active directory.
     */
    readonly azureActiveDirectory?: pulumi.Input<inputs.servicefabric.v20200101preview.AzureActiveDirectory>;
    /**
     * The port used for client connections to the cluster.
     */
    readonly clientConnectionPort?: pulumi.Input<number>;
    /**
     * client certificates for the cluster.
     */
    readonly clients?: pulumi.Input<pulumi.Input<inputs.servicefabric.v20200101preview.ClientCertificate>[]>;
    /**
     * The Service Fabric runtime version of the cluster. This property can only by set the user when **upgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
     */
    readonly clusterCodeVersion?: pulumi.Input<string>;
    /**
     * The name of the cluster resource.
     */
    readonly clusterName: pulumi.Input<string>;
    /**
     * Describes the policy used when upgrading the cluster.
     */
    readonly clusterUpgradeDescription?: pulumi.Input<inputs.servicefabric.v20200101preview.ClusterUpgradePolicy>;
    /**
     * The upgrade mode of the cluster when new Service Fabric runtime version is available.
     *
     *   - Automatic - The cluster will be automatically upgraded to the latest Service Fabric runtime version as soon as it is available.
     *   - Manual - The cluster will not be automatically upgraded to the latest Service Fabric runtime version. The cluster is upgraded by setting the **clusterCodeVersion** property in the cluster resource.
     */
    readonly clusterUpgradeMode?: pulumi.Input<string>;
    /**
     * The cluster dns name.
     */
    readonly dnsName: pulumi.Input<string>;
    /**
     * The list of custom fabric settings to configure the cluster.
     */
    readonly fabricSettings?: pulumi.Input<pulumi.Input<inputs.servicefabric.v20200101preview.SettingsSectionDescription>[]>;
    /**
     * The port used for http connections to the cluster.
     */
    readonly httpGatewayConnectionPort?: pulumi.Input<number>;
    /**
     * Describes load balancing rules.
     */
    readonly loadBalancingRules?: pulumi.Input<pulumi.Input<inputs.servicefabric.v20200101preview.LoadBalancingRule>[]>;
    /**
     * Azure resource location.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The sku of the managed cluster
     */
    readonly sku?: pulumi.Input<inputs.servicefabric.v20200101preview.Sku>;
    /**
     * Azure resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
