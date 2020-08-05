// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Represents user credentials used for publishing activity
 */
export class SiteDeploymentSlot extends pulumi.CustomResource {
    /**
     * Get an existing SiteDeploymentSlot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SiteDeploymentSlot {
        return new SiteDeploymentSlot(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:web/v20150801:SiteDeploymentSlot';

    /**
     * Returns true if the given object is an instance of SiteDeploymentSlot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SiteDeploymentSlot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SiteDeploymentSlot.__pulumiType;
    }

    /**
     * Kind of resource
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * Resource Location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource Name
     */
    public readonly name!: pulumi.Output<string | undefined>;
    public /*out*/ readonly properties!: pulumi.Output<outputs.web.v20150801.DeploymentResponseProperties>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a SiteDeploymentSlot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SiteDeploymentSlotArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SiteDeploymentSlotArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as SiteDeploymentSlotArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.slot === undefined) {
                throw new Error("Missing required property 'slot'");
            }
            inputs["active"] = args ? args.active : undefined;
            inputs["author"] = args ? args.author : undefined;
            inputs["author_email"] = args ? args.author_email : undefined;
            inputs["deployer"] = args ? args.deployer : undefined;
            inputs["details"] = args ? args.details : undefined;
            inputs["end_time"] = args ? args.end_time : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["message"] = args ? args.message : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["slot"] = args ? args.slot : undefined;
            inputs["start_time"] = args ? args.start_time : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["properties"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SiteDeploymentSlot.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a SiteDeploymentSlot resource.
 */
export interface SiteDeploymentSlotArgs {
    /**
     * Active
     */
    readonly active?: pulumi.Input<boolean>;
    /**
     * Author
     */
    readonly author?: pulumi.Input<string>;
    /**
     * AuthorEmail
     */
    readonly author_email?: pulumi.Input<string>;
    /**
     * Deployer
     */
    readonly deployer?: pulumi.Input<string>;
    /**
     * Detail
     */
    readonly details?: pulumi.Input<string>;
    /**
     * EndTime
     */
    readonly end_time?: pulumi.Input<string>;
    /**
     * Kind of resource
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Resource Location
     */
    readonly location: pulumi.Input<string>;
    /**
     * Message
     */
    readonly message?: pulumi.Input<string>;
    /**
     * Resource Id
     */
    readonly name: pulumi.Input<string>;
    /**
     * Name of resource group
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Name of web app slot. If not specified then will default to production slot.
     */
    readonly slot: pulumi.Input<string>;
    /**
     * StartTime
     */
    readonly start_time?: pulumi.Input<string>;
    /**
     * Status
     */
    readonly status?: pulumi.Input<number>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Resource type
     */
    readonly type?: pulumi.Input<string>;
}
