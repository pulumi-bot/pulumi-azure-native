// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Domain Topic.
 *
 * ## Example Usage
 * ### DomainTopics_CreateOrUpdate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const domainTopic = new azurerm.eventgrid.v20200601.DomainTopic("domainTopic", {
 *     domainName: "exampledomain1",
 *     domainTopicName: "exampledomaintopic1",
 *     resourceGroupName: "examplerg",
 * });
 *
 * ```
 */
export class DomainTopic extends pulumi.CustomResource {
    /**
     * Get an existing DomainTopic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DomainTopic {
        return new DomainTopic(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:eventgrid/v20200601:DomainTopic';

    /**
     * Returns true if the given object is an instance of DomainTopic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DomainTopic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DomainTopic.__pulumiType;
    }

    /**
     * Name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Provisioning state of the domain topic.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * Type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a DomainTopic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainTopicArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.domainName === undefined) {
                throw new Error("Missing required property 'domainName'");
            }
            if (!args || args.domainTopicName === undefined) {
                throw new Error("Missing required property 'domainTopicName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["domainName"] = args ? args.domainName : undefined;
            inputs["domainTopicName"] = args ? args.domainTopicName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:eventgrid/latest:DomainTopic" }, { type: "azurerm:eventgrid/v20190201preview:DomainTopic" }, { type: "azurerm:eventgrid/v20190601:DomainTopic" }, { type: "azurerm:eventgrid/v20200101preview:DomainTopic" }, { type: "azurerm:eventgrid/v20200401preview:DomainTopic" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(DomainTopic.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DomainTopic resource.
 */
export interface DomainTopicArgs {
    /**
     * Name of the domain.
     */
    readonly domainName: pulumi.Input<string>;
    /**
     * Name of the domain topic.
     */
    readonly domainTopicName: pulumi.Input<string>;
    /**
     * The name of the resource group within the user's subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
