// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A DDoS protection plan in a resource group.
 */
export class DdosProtectionPlan extends pulumi.CustomResource {
    /**
     * Get an existing DdosProtectionPlan resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DdosProtectionPlan {
        return new DdosProtectionPlan(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network:DdosProtectionPlan';

    /**
     * Returns true if the given object is an instance of DdosProtectionPlan.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DdosProtectionPlan {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DdosProtectionPlan.__pulumiType;
    }

    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Properties of the DDoS protection plan.
     */
    public readonly properties!: pulumi.Output<outputs.network.DdosProtectionPlanPropertiesFormatResponse>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a DdosProtectionPlan resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DdosProtectionPlanArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
        inputs["location"] = args ? args.location : undefined;
        inputs["name"] = args ? args.name : undefined;
        inputs["properties"] = args ? args.properties : undefined;
        inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        inputs["tags"] = args ? args.tags : undefined;
        inputs["etag"] = undefined /*out*/;
        inputs["type"] = undefined /*out*/;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DdosProtectionPlan.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DdosProtectionPlan resource.
 */
export interface DdosProtectionPlanArgs {
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the DDoS protection plan.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Properties of the DDoS protection plan.
     */
    readonly properties?: pulumi.Input<inputs.network.DdosProtectionPlanPropertiesFormat>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}