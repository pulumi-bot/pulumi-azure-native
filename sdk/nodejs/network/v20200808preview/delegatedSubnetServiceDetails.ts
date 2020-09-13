// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Delegated subnet details
 */
export class DelegatedSubnetServiceDetails extends pulumi.CustomResource {
    /**
     * Get an existing DelegatedSubnetServiceDetails resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DelegatedSubnetServiceDetails {
        return new DelegatedSubnetServiceDetails(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20200808preview:DelegatedSubnetServiceDetails';

    /**
     * Returns true if the given object is an instance of DelegatedSubnetServiceDetails.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DelegatedSubnetServiceDetails {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DelegatedSubnetServiceDetails.__pulumiType;
    }

    /**
     * Location of the DelegatedSubnet resource.
     */
    public /*out*/ readonly location!: pulumi.Output<string | undefined>;
    /**
     * The name of the DelegatedSubnet resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The current state of delegated subnet resource.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The type of the DelegatedSubnet  resource.(Microsoft.DelegatedNetwork/delegatedSubnet)
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a DelegatedSubnetServiceDetails resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DelegatedSubnetServiceDetailsArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.resourceName === undefined) {
                throw new Error("Missing required property 'resourceName'");
            }
            if (!args || args.subnetName === undefined) {
                throw new Error("Missing required property 'subnetName'");
            }
            if (!args || args.vnetName === undefined) {
                throw new Error("Missing required property 'vnetName'");
            }
            inputs["controllerID"] = args ? args.controllerID : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceName"] = args ? args.resourceName : undefined;
            inputs["subnetName"] = args ? args.subnetName : undefined;
            inputs["vnetName"] = args ? args.vnetName : undefined;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DelegatedSubnetServiceDetails.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DelegatedSubnetServiceDetails resource.
 */
export interface DelegatedSubnetServiceDetailsArgs {
    /**
     * Delegated Network Controller ID
     */
    readonly controllerID?: pulumi.Input<string>;
    /**
     * The name of the Azure Resource group of which a given DelegatedNetwork resource is part. This name must be at least 1 character in length, and no more than 90.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the resource. It must be a minimum of 3 characters, and a maximum of 63.
     */
    readonly resourceName: pulumi.Input<string>;
    /**
     * The name of the delegated subnet. This name must be at least 1 character in length, and no more than 90.
     */
    readonly subnetName: pulumi.Input<string>;
    /**
     * The name of the virtual network. This name must be at least 1 character in length, and no more than 90.
     */
    readonly vnetName: pulumi.Input<string>;
}
