// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * The top level Log Analytics cluster resource container.
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
    public static readonly __pulumiType = 'azure-nextgen:operationalinsights/latest:Cluster';

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
     * Configures whether billing will be only on the cluster or each workspace will be billed by its proportional use. This does not change the overall billing, only how it will be distributed. Default value is 'Cluster'
     */
    public readonly billingType!: pulumi.Output<string | undefined>;
    /**
     * The ID associated with the cluster.
     */
    public /*out*/ readonly clusterId!: pulumi.Output<string>;
    /**
     * The identity of the resource.
     */
    public readonly identity!: pulumi.Output<outputs.operationalinsights.latest.IdentityResponse | undefined>;
    /**
     * Sets whether the cluster will support availability zones. This can be set as true only in regions where Azure Data Explorer support Availability Zones. This Property can not be modified after cluster creation. Default value is 'true' if region supports Availability Zones.
     */
    public readonly isAvailabilityZonesEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Configures whether cluster will use double encryption. This Property can not be modified after cluster creation. Default value is 'true'
     */
    public readonly isDoubleEncryptionEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The associated key properties.
     */
    public readonly keyVaultProperties!: pulumi.Output<outputs.operationalinsights.latest.KeyVaultPropertiesResponse | undefined>;
    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state of the cluster.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The sku properties.
     */
    public readonly sku!: pulumi.Output<outputs.operationalinsights.latest.ClusterSkuResponse | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

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
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["billingType"] = args ? args.billingType : undefined;
            inputs["clusterName"] = args ? args.clusterName : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["isAvailabilityZonesEnabled"] = args ? args.isAvailabilityZonesEnabled : undefined;
            inputs["isDoubleEncryptionEnabled"] = args ? args.isDoubleEncryptionEnabled : undefined;
            inputs["keyVaultProperties"] = args ? args.keyVaultProperties : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["clusterId"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["billingType"] = undefined /*out*/;
            inputs["clusterId"] = undefined /*out*/;
            inputs["identity"] = undefined /*out*/;
            inputs["isAvailabilityZonesEnabled"] = undefined /*out*/;
            inputs["isDoubleEncryptionEnabled"] = undefined /*out*/;
            inputs["keyVaultProperties"] = undefined /*out*/;
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
        const aliasOpts = { aliases: [{ type: "azure-nextgen:operationalinsights/v20190801preview:Cluster" }, { type: "azure-nextgen:operationalinsights/v20200301preview:Cluster" }, { type: "azure-nextgen:operationalinsights/v20200801:Cluster" }, { type: "azure-nextgen:operationalinsights/v20201001:Cluster" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Cluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * Configures whether billing will be only on the cluster or each workspace will be billed by its proportional use. This does not change the overall billing, only how it will be distributed. Default value is 'Cluster'
     */
    readonly billingType?: pulumi.Input<string>;
    /**
     * The name of the Log Analytics cluster.
     */
    readonly clusterName: pulumi.Input<string>;
    /**
     * The identity of the resource.
     */
    readonly identity?: pulumi.Input<inputs.operationalinsights.latest.Identity>;
    /**
     * Sets whether the cluster will support availability zones. This can be set as true only in regions where Azure Data Explorer support Availability Zones. This Property can not be modified after cluster creation. Default value is 'true' if region supports Availability Zones.
     */
    readonly isAvailabilityZonesEnabled?: pulumi.Input<boolean>;
    /**
     * Configures whether cluster will use double encryption. This Property can not be modified after cluster creation. Default value is 'true'
     */
    readonly isDoubleEncryptionEnabled?: pulumi.Input<boolean>;
    /**
     * The associated key properties.
     */
    readonly keyVaultProperties?: pulumi.Input<inputs.operationalinsights.latest.KeyVaultProperties>;
    /**
     * The geo-location where the resource lives
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The sku properties.
     */
    readonly sku?: pulumi.Input<inputs.operationalinsights.latest.ClusterSku>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
