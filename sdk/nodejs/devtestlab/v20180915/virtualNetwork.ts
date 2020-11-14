// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * A virtual network.
 */
export class VirtualNetwork extends pulumi.CustomResource {
    /**
     * Get an existing VirtualNetwork resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VirtualNetwork {
        return new VirtualNetwork(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:devtestlab/v20180915:VirtualNetwork';

    /**
     * Returns true if the given object is an instance of VirtualNetwork.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualNetwork {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualNetwork.__pulumiType;
    }

    /**
     * The allowed subnets of the virtual network.
     */
    public readonly allowedSubnets!: pulumi.Output<outputs.devtestlab.v20180915.SubnetResponse[] | undefined>;
    /**
     * The creation date of the virtual network.
     */
    public /*out*/ readonly createdDate!: pulumi.Output<string>;
    /**
     * The description of the virtual network.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The Microsoft.Network resource identifier of the virtual network.
     */
    public readonly externalProviderResourceId!: pulumi.Output<string | undefined>;
    /**
     * The external subnet properties.
     */
    public /*out*/ readonly externalSubnets!: pulumi.Output<outputs.devtestlab.v20180915.ExternalSubnetResponse[]>;
    /**
     * The location of the resource.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The provisioning status of the resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The subnet overrides of the virtual network.
     */
    public readonly subnetOverrides!: pulumi.Output<outputs.devtestlab.v20180915.SubnetOverrideResponse[] | undefined>;
    /**
     * The tags of the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The unique immutable identifier of a resource (Guid).
     */
    public /*out*/ readonly uniqueIdentifier!: pulumi.Output<string>;

    /**
     * Create a VirtualNetwork resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualNetworkArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.labName === undefined) {
                throw new Error("Missing required property 'labName'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["allowedSubnets"] = args ? args.allowedSubnets : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["externalProviderResourceId"] = args ? args.externalProviderResourceId : undefined;
            inputs["labName"] = args ? args.labName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["subnetOverrides"] = args ? args.subnetOverrides : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["createdDate"] = undefined /*out*/;
            inputs["externalSubnets"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["uniqueIdentifier"] = undefined /*out*/;
        } else {
            inputs["allowedSubnets"] = undefined /*out*/;
            inputs["createdDate"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["externalProviderResourceId"] = undefined /*out*/;
            inputs["externalSubnets"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["subnetOverrides"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["uniqueIdentifier"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:devtestlab/latest:VirtualNetwork" }, { type: "azure-nextgen:devtestlab/v20150521preview:VirtualNetwork" }, { type: "azure-nextgen:devtestlab/v20160515:VirtualNetwork" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(VirtualNetwork.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a VirtualNetwork resource.
 */
export interface VirtualNetworkArgs {
    /**
     * The allowed subnets of the virtual network.
     */
    readonly allowedSubnets?: pulumi.Input<pulumi.Input<inputs.devtestlab.v20180915.Subnet>[]>;
    /**
     * The description of the virtual network.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The Microsoft.Network resource identifier of the virtual network.
     */
    readonly externalProviderResourceId?: pulumi.Input<string>;
    /**
     * The name of the lab.
     */
    readonly labName: pulumi.Input<string>;
    /**
     * The location of the resource.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the virtual network.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The subnet overrides of the virtual network.
     */
    readonly subnetOverrides?: pulumi.Input<pulumi.Input<inputs.devtestlab.v20180915.SubnetOverride>[]>;
    /**
     * The tags of the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
