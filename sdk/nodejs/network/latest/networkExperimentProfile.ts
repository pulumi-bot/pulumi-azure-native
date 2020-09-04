// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Defines an Network Experiment Profile and lists of Experiments
 *
 * ## Example Usage
 * ### Creates an NetworkExperiment Profile in a Resource Group
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const networkExperimentProfile = new azurerm.network.latest.NetworkExperimentProfile("networkExperimentProfile", {
 *     enabledState: "Enabled",
 *     location: "WestUs",
 *     profileName: "MyProfile",
 *     resourceGroupName: "MyResourceGroup",
 * });
 *
 * ```
 */
export class NetworkExperimentProfile extends pulumi.CustomResource {
    /**
     * Get an existing NetworkExperimentProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): NetworkExperimentProfile {
        return new NetworkExperimentProfile(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/latest:NetworkExperimentProfile';

    /**
     * Returns true if the given object is an instance of NetworkExperimentProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkExperimentProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkExperimentProfile.__pulumiType;
    }

    /**
     * The state of the Experiment
     */
    public readonly enabledState!: pulumi.Output<string | undefined>;
    /**
     * Gets a unique read-only string that changes whenever the resource is updated.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Resource status.
     */
    public readonly resourceState!: pulumi.Output<string | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a NetworkExperimentProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkExperimentProfileArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.profileName === undefined) {
                throw new Error("Missing required property 'profileName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["enabledState"] = args ? args.enabledState : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["profileName"] = args ? args.profileName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceState"] = args ? args.resourceState : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["enabledState"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["resourceState"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/v20191101:NetworkExperimentProfile" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(NetworkExperimentProfile.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a NetworkExperimentProfile resource.
 */
export interface NetworkExperimentProfileArgs {
    /**
     * The state of the Experiment
     */
    readonly enabledState?: pulumi.Input<string>;
    /**
     * Gets a unique read-only string that changes whenever the resource is updated.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the Profile
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Profile identifier associated with the Tenant and Partner
     */
    readonly profileName: pulumi.Input<string>;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource status.
     */
    readonly resourceState?: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
