// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export class JitNetworkAccessPolicy extends pulumi.CustomResource {
    /**
     * Get an existing JitNetworkAccessPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): JitNetworkAccessPolicy {
        return new JitNetworkAccessPolicy(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:security/v20200101:JitNetworkAccessPolicy';

    /**
     * Returns true if the given object is an instance of JitNetworkAccessPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is JitNetworkAccessPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === JitNetworkAccessPolicy.__pulumiType;
    }

    /**
     * Kind of the resource
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * Location where the resource is stored
     */
    public /*out*/ readonly location!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public readonly name!: pulumi.Output<string>;
    public readonly properties!: pulumi.Output<outputs.security.v20200101.JitNetworkAccessPolicyPropertiesResponse>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a JitNetworkAccessPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JitNetworkAccessPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: JitNetworkAccessPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as JitNetworkAccessPolicyArgs | undefined;
            if (!args || args.ascLocation === undefined) {
                throw new Error("Missing required property 'ascLocation'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.properties === undefined) {
                throw new Error("Missing required property 'properties'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["ascLocation"] = args ? args.ascLocation : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["location"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(JitNetworkAccessPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a JitNetworkAccessPolicy resource.
 */
export interface JitNetworkAccessPolicyArgs {
    /**
     * The location where ASC stores the data of the subscription. can be retrieved from Get locations
     */
    readonly ascLocation: pulumi.Input<string>;
    /**
     * Kind of the resource
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Name of a Just-in-Time access configuration policy.
     */
    readonly name: pulumi.Input<string>;
    readonly properties: pulumi.Input<inputs.security.v20200101.JitNetworkAccessPolicyProperties>;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
